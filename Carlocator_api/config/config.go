package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Host            string `mapstructure:"HOST"`
	Port            int    `mapstructure:"DBPORT"`
	DbName          string `mapstructure:"POSTGRES_DB"`
	DbUser          string `mapstructure:"POSTGRES_USER"`
	DbPassword      string `mapstructure:"POSTGRES_PASSWORD"`
	SslMode         string `mapstructure:"SSL_MODE"`
	AccesId         string `mapstructure:"ACCESS_ID"`
	AccessKey       string `mapstructure:"ACCESS_KEY"`
	Region          string `mapstructure:"REGION"`
	CognitoPoolId   string `mapstructure:"COGNITO_POOL_ID"`
	AppClientId     string `mapstructure:"COGNITO_APP_CLIENT_ID"`
	AppClientSecret string `mapstructure:"COGNITO_APP_CLIENT_SECRET"`
	PrivateBucket   string `mapstructure:"S3_BUCKET"`
}

var Cfg Config

func LoadConfig() error {
	viper.AddConfigPath("./")
	// export ENV=production // To load the production env
	if os.Getenv("ENV") == "production" {
		viper.SetConfigName(".env.production")
	} else if os.Getenv("ENV") == "staging" {
		viper.SetConfigName(".env.staging")
	} else if os.Getenv("ENV") == "dev" {
		viper.SetConfigName(".env.dev")
	} else {
		viper.SetConfigName(".env.local")
	}
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	viper.AutomaticEnv()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Cfg)
	return err
}
