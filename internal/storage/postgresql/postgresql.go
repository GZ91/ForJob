package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/GZ91/linkreduct/internal/app/logger"
	"github.com/GZ91/linkreduct/internal/errorsapp"
	"github.com/GZ91/linkreduct/internal/models"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"sync"
	"time"
)

type DB struct {
	conf           ConfigerStorage
	generatorRunes GeneratorRunes
	ps             string
	db             *sql.DB
	chsURLsForDel  chan []models.StructDelURLs
	chURLsForDel   chan models.StructDelURLs
}

func New(ctx context.Context, config ConfigerStorage, generatorRunes GeneratorRunes) (*DB, error) {
	db := &DB{conf: config, generatorRunes: generatorRunes}
	ConfDB := db.conf.GetConfDB()
	db.ps = ConfDB.StringServer
	err := db.openDB()
	if err != nil {
		return nil, err
	}
	err = db.createTable(ctx)
	if err != nil {
		return nil, err
	}
	db.chURLsForDel = make(chan models.StructDelURLs)
	return db, err
}

func (d *DB) openDB() error {
	db, err := sql.Open("pgx", d.ps)
	if err != nil {
		logger.Log.Error("failed to open the database", zap.Error(err))
		return err
	}
	d.db = db
	return nil
}

func (d *DB) createTable(ctx context.Context) error {
	con, err := d.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer con.Close()
	_, err = con.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS short_origin_reference 
(
	id serial PRIMARY KEY,
	token VARCHAR(45)  NOT NULL,
	ShortURL VARCHAR(250) NOT NULL,
	OriginalURL TEXT
);`)
	return err
}

func (d *DB) Ping(ctx context.Context) error {
	con, err := d.db.Conn(ctx)
	if err != nil {
		logger.Log.Error("failed to connect to the database", zap.Error(err))
		return err
	}
	defer con.Close()
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) GetLinksToken(ctx context.Context, userID string) ([]models.ReturnedStructURL, error) {
	con, err := d.db.Conn(ctx)
	if err != nil {
		logger.Log.Error("failed to connect to the database", zap.Error(err))
		return nil, err
	}
	defer con.Close()

	rows, err := con.QueryContext(ctx, `SELECT ShortURL, OriginalURL
	FROM short_origin_reference WHERE userID = $1`, userID)
	if err != nil || rows.Err() != nil {
		if err != sql.ErrNoRows || rows.Err() != sql.ErrNoRows {
			logger.Log.Error("when reading data from the database", zap.Error(err))
			return nil, err
		}
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	returnData := make([]models.ReturnedStructURL, 0)

	for rows.Next() {
		var shortURL, originalURL string
		rows.Scan(&shortURL, &originalURL)
		returnData = append(returnData, models.ReturnedStructURL{OriginalURL: originalURL, ShortURL: shortURL})

	}
	return returnData, nil
}

func (d *DB) InitializingRemovalChannel(ctx context.Context, chsURLs chan []models.StructDelURLs) error {
	d.chsURLsForDel = chsURLs
	go d.GroupingDataForDeleted(ctx)
	go d.FillBufferDelete(ctx)
	return nil
}

func (d *DB) GroupingDataForDeleted(ctx context.Context) {

	var wg sync.WaitGroup
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			close(d.chURLsForDel)
			return
		default:
			sliceVal := <-d.chsURLsForDel
			wg.Add(1)
			go func(*sync.WaitGroup) {
				for _, val := range sliceVal {
					d.chURLsForDel <- val
				}
				wg.Done()
			}(&wg)
		}
	}
}

func (d *DB) FillBufferDelete(ctx context.Context) {
	t := time.NewTicker(time.Second * 10)
	var listForDel []models.StructDelURLs
	for {
		select {
		case val := <-d.chURLsForDel:
			listForDel = append(listForDel, val)
		case <-t.C:
			if len(listForDel) > 0 {
				d.deletedURLs(listForDel)
				listForDel = nil
			}
		}

	}
}

func (d *DB) deletedURLs(listForDel []models.StructDelURLs) {
	ctx := context.Background()
	tx, err := d.db.Begin()
	defer tx.Rollback()
	if err != nil {
		logger.Log.Error("error when trying to create a connection to the database", zap.Error(err))
		return
	}
	pr, err := tx.PrepareContext(ctx, "UPDATE short_origin_reference SET deletedFlag = true WHERE ShortURL = $1 and userID=$2")
	if err != nil {
		logger.Log.Error("error when trying to create a runtime request template", zap.Error(err))
		return
	}
	for _, val := range listForDel {
		pr.Exec(val.URL, val.UserID)
	}
	tx.Commit()
}

func (d *DB) AddBatchLink(ctx context.Context, batchLinks []string) (releasedBatchURL map[string]string, errs error) {
	var UserID string
	var userIDCTX models.CtxString = "userID"
	UserIDVal := ctx.Value(userIDCTX)
	if UserIDVal != nil {
		UserID = UserIDVal.(string)
	}

	tx, err := d.db.Begin()
	defer tx.Rollback()

	if err != nil {
		return nil, err
	}
	reqShortURL, err := tx.PrepareContext(ctx, "SELECT COUNT(id) FROM short_origin_reference WHERE shorturl = $1")
	if err != nil {
		logger.Log.Error("When initializing a shortcut search pattern", zap.Error(err))
		return nil, err
	}
	reqLongLinkInBase, err := tx.PrepareContext(ctx, "SELECT shorturl FROM short_origin_reference WHERE originalurl = $1 LIMIT 1")
	if err != nil {
		logger.Log.Error("when initializing a long link search pattern", zap.Error(err))
		return nil, err
	}
	execInsertLongURLInBase, err := tx.PrepareContext(ctx, "INSERT INTO short_origin_reference(uuid, shorturl, originalurl, userID) VALUES ($1, $2, $3, $4);")
	if err != nil {
		logger.Log.Error("when initializing the add string pattern", zap.Error(err))
		return nil, err
	}

	lenShort := d.conf.GetStartLenShortLink()
	index := 0
	for _, incomingLink := range batchLinks {
		shorturl := ""
		row := reqLongLinkInBase.QueryRowContext(ctx, incomingLink.OriginalURL)
		err := row.Scan(&shorturl)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				logger.Log.Error("when searching for a scan of the result of finding a repeat of a long link", zap.Error(err))
				return nil, err
			}
		}
		if shorturl != "" {
			releasedBatchURL = append(releasedBatchURL, models.ReleasedBatchURL{
				CorrelationID: incomingLink.CorrelationID,
				ShortURL:      shorturl,
			})
			errs = errorsapp.ErrLinkAlreadyExists
			continue
		}
		for {
			shorturl = d.generatorRunes.RandStringRunes(lenShort)
			row := reqShortURL.QueryRowContext(ctx, shorturl)
			var countShorturl int
			err := row.Scan(&countShorturl)
			if err != nil {
				logger.Log.Error("when searching for a scan of the result of finding a repeat of a short link", zap.Error(err))
				return nil, err
			}
			if countShorturl == 0 {
				break
			}
			index++
			if index == d.conf.GetMaxIterLen() {
				lenShort++
				index = 0
			}
		}

		_, err = execInsertLongURLInBase.ExecContext(ctx, uuid.New().String(), shorturl, incomingLink.OriginalURL, UserID)
		if err != nil {
			logger.Log.Error("When creating a string with a long link in the database", zap.Error(err))
			tx.Rollback()
			return nil, err
		}
		releasedBatchURL = append(releasedBatchURL, models.ReleasedBatchURL{
			CorrelationID: incomingLink.CorrelationID,
			ShortURL:      shorturl,
		})
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return
}
