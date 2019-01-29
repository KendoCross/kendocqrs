package logs

import "github.com/astaxie/beego/logs"

const (
	FILE = 1 << iota
	CONSO
)

//ConsoleLogs 控制台日志
var ConsoleLogs *logs.BeeLogger

//FileLogs 文件日志
var FileLogs *logs.BeeLogger

//初始化_日志工具
func InitLogs() {
	ConsoleLogs = logs.NewLogger(1000)
	ConsoleLogs.SetLogger(logs.AdapterConsole)
	ConsoleLogs.EnableFuncCallDepth(true)

	FileLogs = logs.NewLogger(1000)
	FileLogs.EnableFuncCallDepth(true)
	FileLogs.Async()
	logsConf := `{"filename":"./logs/beego_store.log","maxdays ":1,"level ": 6,"perm":"0777"}`

	FileLogs.SetLogger(logs.AdapterFile, logsConf)
	ConsoleLogs.Info("logsConf:", logsConf)
}

func Info(logType int, format string, v ...interface{}) {
	if (logType & FILE) == FILE {
		FileLogs.Info(format, v)
	}
	if logType&CONSO == CONSO {
		ConsoleLogs.Info(format, v)
	}
}

func Error(logType int, format string, v ...interface{}) {
	if (logType & FILE) == FILE {
		FileLogs.Error(format, v)
	}
	if logType&CONSO == CONSO {
		ConsoleLogs.Error(format, v)
	}
}
