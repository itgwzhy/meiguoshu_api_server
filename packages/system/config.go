package system

import (
	"github.com/jinzhu/gorm"
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
	Servicer *serverModel
	Database *databaseModel
	LogModel *logModel
}

//LoadConfig
func LoadConfig() (err error) {
	var (
		configData []byte
	)
	configData, err = ioutil.ReadFile("packages/system/system.yml")
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
	return
}

//Init Database
func InitDatabase() {
	var (
		db *gorm.DB
		err	error
	)
	switch Config.Database.DBType {
	case "postgres":
		db, err = gorm.Open("postgres",Config.Database.Connect)
	case "mysql":
		db, err = gorm.Open("mysql", Config.Database.Connect)

	}
	if err != nil || db == nil {
		log.Errorf("数据库连接失败：%s", err)
		log.Flush()
		os.Exit(-1)
	}

}
