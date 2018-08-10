package main

import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/common"
	"github.com/MEIGUOSHU/meiguoshu_api_server/models"
	"github.com/MEIGUOSHU/meiguoshu_api_server/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
	"io"
	"os"
	"time"
)

// Get a filename based on the date, just for the sugar.
func todayFilename() string {
	today := time.Now().Format("2006-01-02")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	// Open the file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Application{},
		&models.Category{},
		&models.Consultation{},
		&models.Contact{},
		&models.Link{},
		&models.News{},
		&models.Product{},
	)
}

func main() {
	//创建日志文件
	f := newLogFile()
	defer f.Close()
	app := iris.Default()

	//输出日志信息
	//app.Logger().SetOutput(f)
	app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))

	//初始化路由
	handler := router.InitRouter(app)

	//初始化数据库
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=meiguoshu sslmode=disable password=Woxihuanni1")
	if err != nil {
		app.Logger().Infof("数据库连接失败，错误信息：%s", err)
		return
	}
	//全局禁用表名复数
	db.SingularTable(true)

	//添加表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "meiguoshu_" + defaultTableName
	}

	common.DB = db.Set("gorm:save_associations", false)

	//根据模型创建表
	DBMigrate(common.DB)

	defer db.Close()
	//启动程序
	handler.Run(
		iris.Addr(":8000"),
		iris.WithoutBanner,
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
