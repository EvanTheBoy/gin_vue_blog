package core

//这个是用来记录程序的运行情况的。日志是一种记录程序中发生的事件或数据的方式，
//可以帮助开发者或用户分析程序的性能、错误、状态等。
//logrus是一种日志库，就是一组提供日志功能的代码，可以方便地在Go语言中使用。

import (
	"bytes"
	"fmt"
	"gin_vue_blog/gin_vue_blog_server/global"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct {
}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer //设置一个缓冲区。我们规定通过这个函数返回出去的都要放在缓冲区里
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	log := global.Config.Logger

	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		//Sprintf()函数的作用是返回一个字符串
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//设置格式。Fprintf()函数会将输出写入到一个io.Writer接口类型的变量中
		//这个Fprintf()函数的第一个变量就是io.Writer接口类型的变量，最后这个函数返回出去的也是这个的.Bytes
		_, err := fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n",
			log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n",
			log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
		if err != nil {
			return nil, err
		}
	}
	return b.Bytes(), nil
}

func InitDefaultLogger() {
	//这里设置的是全局logrus
	logrus.SetOutput(os.Stdout)                           //设置输出类型。这个Stdout就是说把日志内容输出到控制台
	logrus.SetReportCaller(global.Config.Logger.ShowLine) //开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{})                  //设置自己定义的Formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //设置最低的level
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                //新建一个实例
	mLog.SetOutput(os.Stdout)                           //设置输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) //开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{})                  //设置自己定义的Formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) //设置最低的Level
	InitDefaultLogger()
	return mLog
}
