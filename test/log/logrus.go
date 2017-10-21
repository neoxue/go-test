package main

import (
	"github.com/Sirupsen/logrus"
	"os"
	"github.com/neoxue/comos/libs"
	"fmt"
)


// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func main() {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	//log.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	logrus.SetFormatter(&logrus.JSONFormatter{})
	 file, err := os.OpenFile("/data1/ms/log/logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	logrus.SetOutput(file)
	 if err == nil {
		 log.Out = file
	 } else {
	  logrus.Fatal("Failed to log to file, using default stderr")
	 }
	 logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	test()
}

func test() {
	err := libs.NewErrors("test error ", "test", 101, "test")
	logrus.Info(fmt.Sprintf("%+v", err))
	logrus.Info("test");
}