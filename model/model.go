package model

type DBConnect struct {
	UserName string
	PassWord string
	Addr string
	Port int
	DbName string
}

type FindAndDownload struct {
	TableName string
	Appid int64
	StartTime string
	EndTime string
	DownloadPath string
}