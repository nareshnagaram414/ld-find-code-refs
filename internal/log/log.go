package log

import (
	"io/ioutil"
	"log"
	"os"
)

// Global package level loggers
var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Stdout  *log.Logger
)

// Init overrides the default loggers that write to stdout
func Init(debug bool) {
	debugHandle := ioutil.Discard
	if debug {
		debugHandle = os.Stdout
	}

	Debug = log.New(debugHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Stdout = log.New(os.Stdout, "", 0)
}
