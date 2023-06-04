package app

type ApplicationConfig struct {
	HttpHost   string
	HttpPort   uint64
	DBHost     string
	DBPort     uint64
	DBName     string
	DBUser     string
	DBPassword string
}
