package websever

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//跨域
	router.Use(Cors())
	//连接数据库url——http://127.0.0.1:8000/connectdb
	/** 例：
	curl -vd '{"UserName":"root","PassWord":"_gTQ699qwX","Addr":"100.121.190.3","Port": 15944,"DbName":"ReportData"}' http://127.0.0.1:8000/connectdb
	*/
	router.POST("/connectdb", ConnectDB)

	//根据时间插入指定条数的数据url——http://127.0.0.1:8000/insertdata
	/**例：
	curl -vd '{"TableName":"report_data_async_2019_08","Data":[{"AppId": 1255334997,"TaskType":"atomic-trans","Duration":10,"CreatedAt":"2019-08-01 06:00:0:"h264","UserDefined7":100,"UserDefined8":4072}],"Standard":"CreatedAt","Num":6}' http://127.0.0.1:8000/insertdata
	*/
	router.POST("/insertdata", InsertByStandard)

	//根据时间插入指定条数的数据url——http://127.0.0.1:8000/findbytimeanddownload
	/**例：
	{"TableName":"report_data_async_2019_08","Appid":1255334997,"StartTime":"2019-08-01 06:00:00","EndTime":"2019-08-06 20:00:00","DownloadPath":"reportData_month8.xlsx"}
	*/
	router.POST("/findbytimeanddownload", FindByTimeAndDownload)

	////根据指定字段查询并求和url——http://127.0.0.1:8000/findandsum
	/**例：
	curl -vd '{"TableName":"report_data_async_2019_08","Condition":[{"app_id":1255334997,"task_type":"atomic-trans","mode":1}],"GoalSum":"duration","StartTime":"2019-08-01 00:00:00","EndTime":"2019-08-06 20:00:00","DownloadPath":"reportDataSum_month8.xlsx"}' http://127.0.0.1:8000/findandsum
	*/
	router.POST("/findandsum", FindAndSum)

	return router
}



//跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "text/plain")                                                                                                                                                                    // 设置返回格式是json text/plain
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
