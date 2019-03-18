package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/bitly/go-simplejson"
)

const (
	configFilePath string = "../config/config.json"
)

var (
	// TestDSN 测试数据库链接
	TestDSN string
	// DSN 正式数据库链接
	DSN string
	// Port 端口
	Port int
	// Apns 是否正式环境
	Apns bool
)

func getString(j *simplejson.Json, key string) string {

	if j == nil || key == "" {
		log.Panicln("INVALID ADDRESS OR KEY IS EMPTY")
	}
	s, err := j.Get(key).String()
	if err != nil {
		log.Panicln(err, key)

	}
	return s

}

func getInt(j *simplejson.Json, key string) int {
	if j == nil || key == "" {
		log.Panicln("INVALID ADDRESS OR KEY IS EMPTY")

	}
	i, err := j.Get(key).Int()
	if err != nil {
		log.Panicln(err, key)

	}
	return i
}

func getBool(j *simplejson.Json, key string) bool {
	if j == nil || key == "" {
		log.Panicln("INVALID ADDRESS OR KEY IS EMPTY")
	}
	i, err := j.Get(key).Bool()
	if err != nil {
		log.Panicln(err, key)
	}
	return i
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conf, err := filepath.Abs("./" + configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("loading config: %s", conf)
	f, err := os.Open(conf)
	if err != nil {
		log.Panicln(err)
	}
	config, err := simplejson.NewFromReader(f)
	if err != nil {
		log.Panicln(err)
	}

	TestDSN = getString(config, "TestDSN")
	DSN = getString(config, "DSN")
	Port = getInt(config, "Port")
	Apns = getBool(config, "Apns")
}
