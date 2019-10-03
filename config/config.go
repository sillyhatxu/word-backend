package config

var Conf config

type config struct {
	Project   string
	Module    string
	Host      string
	GRPCPort  int               `toml:"grpc_port"`
	Log       logConf           `toml:"log_conf"`
	EnvConfig environmentConfig `toml:"env_config"`
}

type environmentConfig struct {
	UserName      string `toml:"user_name" env:"SILLYHAT.DB.WORD.USERNAME"`
	Password      string `toml:"password" env:"SILLYHAT.DB.WORD.PASSWORD"`
	Host          string `toml:"host" env:"SILLYHAT.DB.WORD.HOST"`
	Port          string `toml:"port" env:"SILLYHAT.DB.WORD.PORT"`
	Schema        string `toml:"schema" env:"SILLYHAT.DB.WORD.SCHEMA"`
	DDLPath       string `toml:"ddl_path" env:"SILLYHAT.DB.WORD.DDL.PATH"`
	LogstashURL   string `toml:"logstash_url" env:"SILLYHAT.LOGSTASH.URL"`
	ConsulAddress string `toml:"consul_address" env:"SILLYHAT.HOST.CONSUL"`
}

type logConf struct {
	OpenLogstash bool   `toml:"open_logstash"`
	OpenLogfile  bool   `toml:"open_logfile"`
	FilePath     string `toml:"file_path"`
}
