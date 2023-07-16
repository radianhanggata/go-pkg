package initializer

import "github.com/spf13/viper"

type EV struct {
	DBHost        string `mapstructure:"POSTGRES_HOST"`
	DBUserName    string `mapstructure:"POSTGRES_USER"`
	DBSecret      string `mapstructure:"POSTGRES_SECRET"`
	DBName        string `mapstructure:"POSTGRES_DB"`
	DBPort        string `mapstructure:"POSTGRES_PORT"`
	DBSSLMode     string `mapstructure:"POSTGRES_SSLMODE"`
	DBTZ          string `mapstructure:"POSTGRES_TZ"`
	ServerPort    string `mapstructure:"PORT"`
	RedisDSN      string `mapstructure:"REDIS_DSN"`
	AccessSecret  string `mapstructure:"ACCESS_SECRET"`
	RefreshSecret string `mapstructure:"REFRESH_SECRET"`
}

func Load(path, filename string) (ev *EV, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(filename)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	ev = &EV{}
	err = viper.Unmarshal(ev)
	return
}
