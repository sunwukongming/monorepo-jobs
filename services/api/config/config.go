package config

import (
	"github.com/sirupsen/logrus"
)

// Config 配置
type Config struct {
	Name     string `yaml:"name"`
	Mode     string `yaml:"mode"`
	Listen   string `yaml:"listen"`
	TimeZone string `yaml:"timezone"`
	Key      string `yaml:"key"`
	DBS      DBS    `yaml:"dbs"`
	Redis    Redis  `yaml:"redis"`
	OSS      OSS    `yaml:"oss"`
	Wechat   Wechat `yaml:"wechat"`
}

// Wechat 微信小程序凭据
type Wechat struct {
	AppID     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

// OSS 对象存储
type OSS struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	Bucket          string `yaml:"bucket"`
}

// DBS 数据库
type DBS struct {
	Default DB `yaml:"default"`
}

// DB 数据库
type DB struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
	PoolSize int    `yaml:"poolSize"`
	Prefix   string `yaml:"prefix"`
}

// Redis 缓存
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Index    int    `yaml:"index"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"poolSize"`
}

var con *Config

// 设置配置
func Set(config *Config) {
	con = config
}

// Get 获取配置
func Get() *Config {
	if con == nil {
		logrus.Fatal("缺乏配置文件")
	}
	return con
}
