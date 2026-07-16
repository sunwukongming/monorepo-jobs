package bootstrap

import (
	"app/config"
	"app/data"

	"app/internal/loader"
	"app/pkg/utils"
	"flag"

	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
	_ "time/tzdata"
	"unsafe"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/sirupsen/logrus"
)

type timeCodec struct {
}

func (codec *timeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	timeString := iter.ReadString()
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeString, utils.TimeLocation)
	*((*time.Time)(ptr)) = t
}

func (codec *timeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	ts := *((*time.Time)(ptr))
	return ts.UnixNano() == 0
}

func (codec *timeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ts := *((*time.Time)(ptr))
	stream.WriteString(ts.In(utils.TimeLocation).Format("2006-01-02 15:04:05"))
}

func init() {
	extra.RegisterFuzzyDecoders()
	jsoniter.RegisterTypeEncoder("time.Time", &timeCodec{})
	jsoniter.RegisterTypeDecoder("time.Time", &timeCodec{})
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "datetime",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	})

	var configPath string
	flag.StringVar(&configPath, "c", path.Join(utils.FileExecuteDir(), "config.yaml"), "配置文件")
	flag.Parse()
	configString, err := os.ReadFile(configPath)
	if err != nil {
		logrus.Fatalf("配置文件错误 %s %v", configPath, err)
	}
	loader.LoadConfig(string(configString))
	Bootstrap(config.Get())
}

func FileDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logrus.Fatal(err)
	}
	return dir
}

func Bootstrap(con *config.Config) {
	gin.SetMode(con.Mode)
	if con.Mode == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(os.Stdout)
	}
	err := os.Setenv("TZ", con.TimeZone)
	if err != nil {
		logrus.Fatal("环境变量无法设置")
	}
	dir := FileDir()
	logDir := path.Join(dir, "logs", config.Get().Name)
	func() {
		_, err := os.Stat(logDir)
		if err != nil {
			if !os.IsExist(err) { //文件不存在
				os.Mkdir(logDir, 0777)
			}
		}
	}()
	logPath := path.Join(logDir, fmt.Sprintf("api-%s.log", time.Now().Format("20060102")))
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetOutput(file)
	go func() {
		for {
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 1, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			logPath := path.Join(logDir, fmt.Sprintf("api-%s.log", time.Now().Format("20060102")))
			lastFile := file
			file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
			if err != nil {
				logrus.Fatal(err)
			}
			logrus.SetOutput(file)
			lastFile.Close()
		}
	}()
	err = data.Load()
	if err != nil {
		logrus.Fatal("数据载入失败" + err.Error())
	}
}
