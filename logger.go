package GoIpBot

import "github.com/sirupsen/logrus"

func Log(module string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"module": module,
	})
}
