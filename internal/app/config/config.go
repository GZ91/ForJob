package config

import (
	"github.com/GZ91/linkreduct/internal/storage/postgresql/postgresqlconfig"
	"sync"
)

type Config struct {
	debug             bool
	addressServer     string
	addressServerURL  string
	maxIterLen        int
	startLenShortLink int
	configDB          *postgresqlconfig.ConfigDB
	mutex             sync.Mutex
	secretKey         string
	roottoken         string
}

func New(debug bool, addressServer, addressServerURL string, maxIterRuneGen int, startLenShortLink int, SecretKey string, RootToken string) *Config {
	return &Config{
		debug:             debug,
		addressServer:     addressServer,
		maxIterLen:        maxIterRuneGen,
		addressServerURL:  addressServerURL,
		startLenShortLink: startLenShortLink,
		secretKey:         SecretKey,
		roottoken:         RootToken,
	}
}

func (r *Config) ConfigureDBPostgresql(StringServer string) {
	r.configDB = postgresqlconfig.New(StringServer)
}

func (r *Config) GetConfDB() *postgresqlconfig.ConfigDB {
	return r.configDB
}

func (r *Config) GetAddressServerURL() string {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.addressServerURL
}

func (r *Config) GetDebug() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.debug
}

func (r *Config) GetAddressServer() string {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.addressServer
}

func (r *Config) GetMaxIterLen() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.maxIterLen
}

func (r *Config) GetStartLenShortLink() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.startLenShortLink
}

func (r *Config) GetSecretKey() string {
	return r.secretKey
}

func (r *Config) GetRootToken() string {
	return r.roottoken
}
