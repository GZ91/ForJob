package initializing

import (
	"github.com/GZ91/linkreduct/internal/app/config"
	"github.com/GZ91/linkreduct/internal/app/initializing/envs"
	"github.com/GZ91/linkreduct/internal/app/logger"
)

func Configuration() (*config.Config, error) {
	envs, err := envs.ReadEnv()
	if err != nil {
		return nil, err
	}
	logger.Initializing(envs.LvlLogs)
	conf := config.New(false, envs.AddressServer, envs.AddressServerForURL, 100, 4, envs.SecretKey, envs.RootToken)
	conf.ConfigureDBPostgresql(envs.GetAddressDSN())
	return conf, nil
}
