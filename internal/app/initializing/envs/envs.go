package envs

import (
	"github.com/caarlos0/env/v6"
)

type EnvVars struct {
	AddressServer       string `env:"SERVER_ADDRESS"`
	AddressServerForURL string `env:"BASE_URL"`
	LvlLogs             string `env:"LOG_LEVEL"`
	CSDBLogin           string `env:"DSN_Login"`
	CSDBPassword        string `env:"DSN_Password"`
	CSDBAddress         string `env:"DSN_Address"`
	CSDBSslmode         string `env:"DSN_Sslmode"`
	CSDBBaseName        string `env:"DSN_BaseName"`
	SecretKey           string `env:"Secret_Key"`
}

func ReadEnv() (*EnvVars, error) {

	envs := EnvVars{}

	if err := env.Parse(&envs); err != nil {
		return nil, err
	}

	return &envs, nil
}

func (r *EnvVars) GetAddressDSN() string {
	return "postgres://" + r.CSDBLogin + ":" + r.CSDBPassword + "@" + r.CSDBAddress + "/" + r.CSDBBaseName + "?sslmode=" + r.CSDBSslmode
}
