package system

//serverModel get server information from config.yml
type serverModel struct {
	Mode                 string `yaml:"mode"`                    // run mode
	Host                 string `yaml:"host"`                    // server host
	Port                 string `yaml:"port"`                    // server port
	SystemStaticFilePath string `yaml:"system_static_file_path"` // system static file path
}

//databaseModel get database information from config.yml
type databaseModel struct {
	DBType  string `yaml:"type"`     // db type
	Connect string `yaml:"connect"`  // db connect information
	MaxIdle int    `yaml:"max_idle"` // db max idle connections
	MaxOpen int    `yaml:"max_open"` // db max open connections
}

type logModel struct {
	Level      string `yaml:"level"`
	ConfigFile string `yaml:"config_file"`
}
