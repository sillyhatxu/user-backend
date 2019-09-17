package config

var Conf config

type config struct {
	Project   string
	Module    string
	Host      string
	HttpPort  int               `toml:"http_port"`
	GRPCPort  int               `toml:"grpc_port"`
	Log       logConf           `toml:"log_conf"`
	DB        database          `toml:"database"`
	EnvConfig environmentConfig `toml:"env_config"`
}

type environmentConfig struct {
	LogstashURL   string `toml:"logstash_url" env:"SILLYHAT.LOGSTASH.URL"`
	ConsulAddress string `toml:"consul_address" env:"SILLYHAT.HOST.CONSUL"`
}

type database struct {
	DataSourceName string `toml:"data_source_name"`
	DDLPath        string `toml:"ddl_path"`
}

type logConf struct {
	OpenLogstash bool   `toml:"open_logstash"`
	OpenLogfile  bool   `toml:"open_logfile"`
	FilePath     string `toml:"file_path"`
}
