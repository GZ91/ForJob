package postgresql

import "github.com/GZ91/linkreduct/internal/storage/postgresql/postgresqlconfig"

type ConfigerStorage interface {
	GetMaxIterLen() int
	GetConfDB() *postgresqlconfig.ConfigDB
	GetStartLenShortLink() int
}

type GeneratorRunes interface {
	RandStringRunes(int) string
}
