package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"path"
)

const (
	envKey            = "ENV"
	configFile        = "system"
	baseConfigPathKey = "BASE_CONFIG_PATH"
	svcNameKey        = "SERVICE_NAME"
)

var (
	validEnvs = map[string]bool{"local": true, "stage": true, "prod": true, "test": true}
	V         *viper.Viper
)

func Init() (err error) {
	err = setupViper()
	if err != nil {
		return errors.Wrap(err, "setupViper.error")
	}

	return
}

func setupViper() (err error) {
	V = viper.New()
	if err = bindEnvVars(); err != nil {
		return errors.Wrap(err, "bindEnvVars.error")
	}

	if err = validateEnv(V.GetString(envKey)); err != nil {
		return errors.Wrap(err, "validateEnv.error")
	}

	if err = setupViperConfig(); err != nil {
		return errors.Wrap(err, "setupViperConfig.error")
	}

	return
}

func bindEnvVars() (err error) {
	if err = V.BindEnv(envKey); err != nil {
		return errors.Wrap(err, "V.BindEnv.envKey.error")
	}

	if err = V.BindEnv(baseConfigPathKey); err != nil {
		return errors.Wrap(err, "V.BindEnv.baseConfigPathKey.error")
	}

	if err = V.BindEnv(svcNameKey); err != nil {
		return errors.Wrap(err, "V.BindEnv.svcNameKey.error")
	}

	return
}

func validateEnv(env string) error {
	if valid, found := validEnvs[env]; !found || !valid {
		return errors.New(" environment " + env + " not found or invalid")
	}

	return nil
}

func setupViperConfig() error {
	V.SetConfigName(configFile)
	baseConfigPath := path.Join(V.GetString(baseConfigPathKey), V.GetString(svcNameKey), V.GetString(envKey))
	V.AddConfigPath(baseConfigPath)
	if err := V.ReadInConfig(); err != nil {
		return errors.Wrap(err, "V.ReadInConfig.error path:")
	}

	return nil
}
