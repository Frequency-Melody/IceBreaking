package config

type config struct {
	Mysql Mysql
	OSS   OSS
}

type Mysql struct {
	Database string
	Host     string
	User     string
	Pwd      string
	Port     string
}

type OSS struct {
	EndPoint        string //region
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}
