package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sillyhatxu/consul-client"
	"github.com/sillyhatxu/environment-config"
	"github.com/sillyhatxu/logrus-client"
	"github.com/sillyhatxu/logrus-client/filehook"
	"github.com/sillyhatxu/logrus-client/logstashhook"
	"github.com/sillyhatxu/user-backend/config"
	"github.com/sillyhatxu/user-backend/dao"
	"github.com/sillyhatxu/user-backend/grpc/grpcserver"
	"github.com/sirupsen/logrus"
	"net"
)

func init() {
	cfgFile := flag.String("c", "config-local.conf", "config file")
	flag.Parse()
	err := envconfig.ParseEnvironmentConfig(&config.Conf.EnvConfig)
	if err != nil {
		panic(err)
	}
	envconfig.ParseConfig(*cfgFile, func(content []byte) {
		err := toml.Unmarshal(content, &config.Conf)
		if err != nil {
			panic(fmt.Sprintf("unmarshal toml object error. %v", err))
		}
	})
	fields := logrus.Fields{
		"project":  config.Conf.Project,
		"module":   config.Conf.Module,
		"@version": "1",
		"type":     "project-log",
	}
	logrusconf.NewLogrusConfig(
		logrusconf.Fields(fields),
		logrusconf.FileConfig(filehook.NewFileConfig(config.Conf.Log.FilePath, filehook.Open(config.Conf.Log.OpenLogfile))),
		logrusconf.LogstashConfig(logstashhook.NewLogstashConfig(config.Conf.EnvConfig.LogstashURL, logstashhook.Open(config.Conf.Log.OpenLogstash), logstashhook.Fields(fields))),
	).Initial()
}

func main() {
	consulServer := consul.NewConsulServer(
		config.Conf.EnvConfig.ConsulAddress,
		config.Conf.Module,
		config.Conf.Host,
		config.Conf.GRPCPort,
		consul.CheckType(consul.HealthCheckGRPC),
	)
	err := consulServer.Register()
	if err != nil {
		panic(err)
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Conf.EnvConfig.UserName,
		config.Conf.EnvConfig.Password,
		config.Conf.EnvConfig.Host,
		config.Conf.EnvConfig.Port,
		config.Conf.EnvConfig.Schema,
	)
	logrus.Infof("dataSourceName : %s", dataSourceName)
	err = dao.Initial(dataSourceName, config.Conf.EnvConfig.DDLPath)
	if err != nil {
		panic(err)
	}
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Conf.GRPCPort))
	if err != nil {
		panic(err)
	}
	go grpcserver.InitialGRPC(grpcListener)
	forever := make(chan bool)
	<-forever
}
