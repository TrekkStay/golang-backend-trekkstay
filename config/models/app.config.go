package models

type AppConfig struct {
	BuildEnv       string `mapstructure:"BUILD_ENV"`
	Version        string `mapstructure:"VERSION"`
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceHost    string `mapstructure:"SERVICE_HOST"`
	ServicePort    uint   `mapstructure:"SERVICE_PORT"`
	ServiceTimeout uint   `mapstructure:"SERVICE_TIMEOUT"`
}
