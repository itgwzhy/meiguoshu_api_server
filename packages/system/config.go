package system

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/seelog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	DB     *gorm.DB
	Config *configModel
)

type configModel struct {
	Server   *serverModel
	Database *databaseModel
	Log      *logModel
}

//LoadConfig
func LoadConfig() (err error) {
	var (
		configData []byte
	)
	configData, err = ioutil.ReadFile("config.yml")
	if err != nil {
		log.Errorf("配置文件读取失败: %s", err)
		log.Flush()
		os.Exit(-1)
	}

	err = yaml.Unmarshal(configData, &Config)
	if err != nil {
		log.Errorf("配置文件解析失败：%s", err)
		log.Flush()
		os.Exit(-1)
	}

	fmt.Printf("%+v\n", Config)

	return
	LoadLog()
	InitDatabase()
	return
}

//Init log
func LoadLog() {
	//Config.LogModel.ConfigFile
	l, e := log.LoggerFromConfigAsFile("seelog.xml")
	if e != nil {
		fmt.Println("日志加载失败：", e)
		os.Exit(-1)
	}
	log.ReplaceLogger(l)
}

//Init Database
func InitDatabase() {
	var (
		db  *gorm.DB
		err error
	)
	fmt.Println(Config.Database.DBType)

	switch Config.Database.DBType {
	case "postgres":
		db, err = gorm.Open("postgres", Config.Database.Connect)
	case "mysql":
		db, err = gorm.Open("mysql", Config.Database.Connect)

	}
	if err != nil || db == nil {
		log.Errorf("数据库连接失败: %s", err)
		log.Errorf("数据库配置信息: %v", Config.Database)
		log.Flush()
		os.Exit(-1)
	}

	DB = db.Set("gorm:save_associations", false)
	//创建数据库
	DBMigrate(db)
	return
}
