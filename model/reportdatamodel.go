package model

type ReportDataModel struct {
	Id             int64  `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	CreatedAt      string `gorm:"column:created_at"`
	TaskType       string `gorm:"column:task_type"`
	VodTaskId      string `gorm:"column:vod_task_id"`
	AppId          int64  `gorm:"column:app_id"`
	Region         string `gorm:"column:region"`
	FileId         int64  `gorm:"column:file_id"`
	BizId          int64  `gorm:"column:biz_id"`
	TaskId         string `gorm:"column:task_id"`
	TaskHost       string `gorm:"column:task_host"`
	TaskStatus     int64  `gorm:"column:task_status"`
	TaskErrCode    int64  `gorm:"column:task_error_code"`
	ReqTime        string `gorm:"column:req_time"`
	ScheduleTime   string `gorm:"column:schedule_time"`
	RealStartTime  string `gorm:"column:real_start_time"`
	FinishTime     string `gorm:"column:finish_time"`
	FileSize       int64  `gorm:"column:file_size"`
	FileNum        int64  `gorm:"column:file_num"`
	Duration       int64  `gorm:"column:duration"`
	ReqPara        string `gorm:"column:req_para"`
	TaskResult     string `gorm:"column:task_result"`
	ReportIp       string `gorm:"column:report_ip"`
	AtomicTaskId   string `gorm:"column:atomic_task_id"`
	UserDefined1   int64  `gorm:"column:user_defined1"`
	UserDefined2   int64  `gorm:"column:user_defined2"`
	UserDefined3   int64  `gorm:"column:user_defined3"`
	UserDefined7   int64  `gorm:"column:user_defined7"`
	UserDefined8   int64  `gorm:"column:user_defined8"`
	UserDefined4   string `gorm:"column:user_defined4"`
	UserDefined5   string `gorm:"column:user_defined5"`
	UserDefined6   string `gorm:"column:user_defined6"`
	OutputFileSize int64  `gorm:"column:output_file_size"`
	NonBill        int64  `gorm:"column:non_bill"`
	NonVod         int64  `gorm:"column:non_vod"`
	Mode           int64  `gorm:"column:mode"`
}

//插入数据的方法
func (r *ReportDataModel) Create() error {
	return DB.Self.Create(&r).Error
}

////设置表名
//func (r *ReportDataModel) TableName() string {
//	return "report_data_async_2019_06"
//}
var tableName string

//设置表名
func (r *ReportDataModel) TableName() string {
	return tableName
}

func GetTableName(table string) string {
	tableName = table
	return tableName
}
