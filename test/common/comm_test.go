package common_test

import "github.com/Sirupsen/logrus"

func test() {
	logrus.WithFields(logrus.Fields{"testkey":"testvalue"}).Debug("value")
}
