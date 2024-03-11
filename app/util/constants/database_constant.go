package constants

const (
	Postgres = "postgres"
	Mysql    = "mysql"
	Sqlite   = "sqlite"
	Mariadb  = "mariadb"
)

const (
	PostgresDSN = "host=%s user=%s password=%s dbname=%s port=%v sslmode=%s search_path=%s"
	MysqlDSN    = "%s:%s@tcp(%s:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	SqliteDSN   = "file:%s?cache=shared&mode=rwc"
)
