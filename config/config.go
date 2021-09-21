package config

type config struct {
	Mysql Mysql
}

type Mysql struct {
	Database string
	Host     string
	User     string
	Pwd      string
	Port     string
	Table    string
}
