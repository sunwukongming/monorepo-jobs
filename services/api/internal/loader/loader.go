package loader

import (
	"app/config"
	"app/internal/db/mysql"

	"gopkg.in/yaml.v2"
)

func LoadConfig(configString string) (*config.Config, error) {
	con := new(config.Config)
	con.TimeZone = "Asia/Shanghai"
	err := yaml.Unmarshal([]byte(configString), con)
	if err != nil {
		return con, err
	}
	config.Set(con)
	//重置相关的服务配置（GORM 连接，db.Default() 亦复用该连接）
	mysql.Reload()
	return con, nil
}
