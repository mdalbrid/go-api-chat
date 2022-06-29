package configs

/*
var ServerConfig server.Config
var DbConfig *configUtil.DatabasePostgresConfig


func init() {
	DbConfig = configUtil.GetDatabasePostgresConfig()
	viper := configUtil.GetViper()
	ServerConfig = server.Config{
		Addr: fmt.Sprintf(`%s:%d`,
			viper.GetString("CHAT.SERVER_HOST"),
			viper.GetInt("CHAT.SERVER_PORT"),
		),
		WriteTimeout: time.Duration(viper.GetInt("CHAT.SERVER_WRITE_TIMEOUT")) * time.Second,
		ReadTimeout:  time.Duration(viper.GetInt("CHAT.SERVER_READ_TIMEOUT")) * time.Second,
	}
}

*/
