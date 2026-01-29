package pkgs

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"strconv"
	"tagprint/models"
	"time"

	"github.com/redis/go-redis/v9"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/spf13/viper"
)

var Log *slog.Logger
var ServerConfig models.ServerConfig

func Rec1(owner string) {
	err := recover()
	if err != nil {
		fmt.Println(owner, err)
		Log.Error(owner, err)
	}
}
func init() {
	logf := &lumberjack.Logger{
		Filename:   "logs/golog.log", // 日志文件的位置
		MaxSize:    5,                // 文件最大尺寸（以MB为单位）
		MaxBackups: 30,               // 保留的最大旧文件数量
		MaxAge:     400,              // 保留旧文件的最大天数
		Compress:   true,             // 是否压缩/归档旧文件
		LocalTime:  true,             // 使用本地时间创建时间戳
	}
	Log = slog.New(slog.NewTextHandler(logf, nil))

	workDir, _ := os.Getwd()
	v := viper.New()
	v.SetConfigFile(path.Join(workDir, "config.yaml"))
	if err := v.ReadInConfig(); err != nil {
		Log.Error("配置文件读取失败", "error", err)
		os.Exit(0)
		return
	}
	ServerConfig = models.ServerConfig{}
	if err := v.Unmarshal(&ServerConfig); err != nil {
		Log.Error("解析结构体失败")
		return
	}

}
func GetRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     *ServerConfig.RedisAddr,
		Password: "",                         // no password set
		DB:       int(*ServerConfig.RedisDB), // use default DB
	})
	return rdb
}
func GetTimeSpan() int64 {
	now := time.Now()
	// 获取毫秒级 Unix 时间戳
	millisecondTimestamp := now.UnixNano() / int64(time.Millisecond)
	return millisecondTimestamp
}

func GerRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     *ServerConfig.RedisAddr,
		Password: *ServerConfig.RedisPassword, // no password set
		DB:       int(*ServerConfig.RedisDB),  // use default DB
	})
	return rdb
}
func GetTimeStr() string {
	now := time.Now()
	// 格式化时间为指定的字符串格式
	formattedTime := now.Format("20060102150405")
	return formattedTime
}
func GetTime(now time.Time) string {
	// 格式化时间为指定的字符串格式
	formattedTime := now.Format("2006-01-02 15:04:05.000")
	return formattedTime
}
func GetTimestamp(isSec bool) string {
	currentTime := time.Now()
	if isSec {
		return fmt.Sprintf("%d", currentTime.Unix())

	} else {
		milliseconds := currentTime.UnixNano() / 1e6
		return strconv.FormatInt(milliseconds, 2)
	}
}
func Err(e error) {
	if e != nil {
		Log.Error(e.Error())
	}
}
