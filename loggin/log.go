package loggin

import (
    //"bufio"
	//"fmt"
	//"io"
	//"os"
	//"time"
    "github.com/jeanphorn/log4go"
)
type Level int
const (
    FINEST Level = iota
    FINE
    DEBUG
    TRACE
    INFO
    WARNING
    ERROR
    CRITICAL
)

func Log4(Logfile, Msg string)  {
    log := log4go.NewLogger()
    log.LoadConfiguration("logging.json")
    log.AddFilter("file", log4go.INFO, log4go.NewFileLogWriter(Logfile,false,true))
    log.Info(Msg)
    log.Close()
  return
}