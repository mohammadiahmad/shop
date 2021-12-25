package readisearch

type Config struct {
	Host     	string `default:"127.0.0.1"`
	Port     	int    `default:"6379"`
	Db 			string `default:"0"`
}
