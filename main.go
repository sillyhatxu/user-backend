package main

import (
	"github.com/sillyhatxu/logrus-client"
	"github.com/sillyhatxu/logrus-client/filehook"
	"github.com/sillyhatxu/logrus-client/logstashhook"
	"github.com/sillyhatxu/user-backend/config"
	"github.com/sirupsen/logrus"
)

func init() {
	fields := logrus.Fields{
		"project":  config.Conf.Project,
		"module":   config.Conf.Module,
		"@version": "1",
		"type":     "project-log",
	}
	logrusconf.NewLogrusConfig(
		logrusconf.Fields(fields),
		logrusconf.FileConfig(filehook.NewFileConfig(config.Conf.Log.FilePath)),
		logrusconf.LogstashConfig(logstashhook.NewLogstashConfig(config.Conf.EnvConfig.LogstashURL, logstashhook.Fields(fields))),
	).Initial()
}

func main() {

}
