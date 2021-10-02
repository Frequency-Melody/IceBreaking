package config

type config struct {
	Mysql    Mysql
	OSS      OSS
	Hduhelp  Hduhelp
	FrontEnd FrontEnd
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

type Hduhelp struct {
	ClientId     string
	ClientSecret string
}

type FrontEnd struct {
	Home          string
	AuthFailedUrl string
}
