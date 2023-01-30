package config

import "github.com/spf13/viper"

type Config struct {
	Port               string `mapstructure:"PORT"`
	DbUrl              string `mapstructure:"DB_URL"`
	AwsRegion          string `mapstructure:"AWS_REGION"`
	AwsAccessKeyId     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AwsBucket          string `mapstructure:"AWS_BUCKET"`
}

func LoadConfig() (c Config, err error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./internal/config")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	if err != nil {
		return
	}

	return
}
