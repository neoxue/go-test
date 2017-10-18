package main


import (
	"log"
	"log/syslog"
)

func main() {
	logwriter, e := syslog.Dial("tcp", "123.123.123.123:12345", syslog.LOG_DEBUG, "your.software.identifier")
	if e != nil {
		log.Fatal(e)
	}

	// normal logs
	logwriter.Info("This is a test!")

	// you can even send JSON like messages like this:
	logwriter.Info(`@cee:{"key1":"value1", "key2":"value2"}`)

	// more details about JSON CEE format in rsyslog:
	// http://www.rsyslog.com/json-elasticsearch/
}
