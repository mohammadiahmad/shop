package storage

type Config struct {
	Host     string `default:"127.0.0.1"`
	Port     int    `default:"3306"`
	Username string `default:"admin"`
	Password string `default:"admin"`
	Database string `default:"shop"`
}
