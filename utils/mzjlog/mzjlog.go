package mzjlog

import (
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/vladoatanasov/logrus_amqp"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/sohlich/elogrus.v2"
)

func main() {
	/*var log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{}) //设置日志的输出格式为json格式，还可以设置为text格式
	log.SetOutput(os.Stdout)                  //设置日志的输出为标准输出
	log.SetLevel(logrus.InfoLevel)            //设置日志的显示级别，这一级别以及更高级别的日志信息将会输出
	//用日志实例的方式使用日志
	log.Out = os.Stdout //日志标准输出
	file, err := os.OpenFile("golang.log", os.O_CREATE|os.O_WRONLY, 1)
	if err == nil {
		log.Out = file
	} else {
		logrus.Info("failed to log to file")
	}
	logrus.WithFields(logrus.Fields{
		"filename": "123.txt",
	}).Info("将日志信息输出到文件中")
	logrus.WithFields(logrus.Fields{
		"filename": "123.txt",
	}).Info("将日志信息输出到文件中asda")*/

}

var logPath = "log"
var fileName = "wfw.log"
var baseLogPaht string

func init() { //不想new,所以直接点第一次就会去做一个日志基础配置，需要说明一点是go-micro也带有logrus，所以那边就用它的插件了
	logrus.SetFormatter(&logrus.JSONFormatter{}) //设置日志的输出格式为json格式，还可以设置为text格式
	logrus.SetOutput(os.Stdout)                  //设置日志的输出为标准输出
	logrus.SetLevel(logrus.InfoLevel)            //设置日志的显示级别，这一级别以及更高级别的日志信息将会输出
	log.AddHook(GetHook(logPath, fileName, time.Hour*24*7, time.Hour))
}

//Info Info日志
func Info(msg string) {
	logrus.WithFields(logrus.Fields{
		"filename": baseLogPaht,
	}).Info(msg)
}

//Error Error
func Error(msg string) {
	logrus.WithFields(logrus.Fields{
		"filename": baseLogPaht,
	}).Error(msg)
}

//Fatal Fatal
func Fatal(msg string) {
	logrus.WithFields(logrus.Fields{
		"filename": baseLogPaht,
	}).Fatal(msg)
}

//Warn Warn
func Warn(msg string) {
	logrus.WithFields(logrus.Fields{
		"filename": baseLogPaht,
	}).Warn(msg)
}

func GetHook(logPath, logFileName string, maxAge, rotationTime time.Duration) *lfshook.LfsHook {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.Mkdir(logPath, 0777)
		// 再修改权限
		os.Chmod(logPath, 0777)
	}
	baseLogPaht = path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+".Bug%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	infoFiles, err := rotatelogs.New(
		baseLogPaht+".Msg%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	return lfshook.NewHook(lfshook.WriterMap{
		//log.DebugLevel: infoFiles, // 为不同级别设置不同的输出目的
		//log.InfoLevel:  infoFiles,
		//debug等不做记录了，浪费资源
		log.WarnLevel:  infoFiles,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.JSONFormatter{})
}

//ConfigAmqpLogger config logrus log to amqp
func ConfigAmqpLogger(server, username, password, exchange, exchangeType, virtualHost, routingKey string) {
	hook := logrus_amqp.NewAMQPHookWithType(server, username, password, exchange, exchangeType, virtualHost, routingKey)
	log.AddHook(hook)
}

//ConfigESLogger config logrus log to es
func ConfigESLogger(esURL string, esHost string, index string) {
	client, err := elastic.NewClient(elastic.SetURL(esURL))
	if err != nil {
		log.Errorf("config es logger error. %+v", errors.WithStack(err))
	}
	esHook, err := elogrus.NewElasticHook(client, esHost, log.DebugLevel, index)
	if err != nil {
		log.Errorf("config es logger error. %+v", errors.WithStack(err))
	}
	log.AddHook(esHook)
}
