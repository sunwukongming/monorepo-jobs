package main

import (
	"app/db"
	_ "app/init"
	"app/internal/db/mysql/query"
	"app/models/bolejiang"
	"fmt"
	"log"
)

func main() {
	// var configPath string
	// flag.StringVar(&configPath, "c", path.Join(utils.FileExecuteDir(), "config.yaml"), "配置文件")
	// flag.Parse()
	// configString, err := os.ReadFile(configPath)
	// if err != nil {
	// 	logrus.Fatalf("配置文件错误 %s %v", configPath, err)
	// }
	// loader.LoadConfig(string(configString))
	// bootstrap.Bootstrap(config.Get())

	var recommends []bolejiang.PassageRecommend
	err := db.Default().Where("path = ''").Find(&recommends)
	if err != nil {
		log.Fatal(err.Error())
	}
	query.Account.Create()

	for _, recommend := range recommends {
		if recommend.ParentPassageRecommendId == 0 {
			recommend.Path = "0"
			recommend.PathFull = recommend.GetFullPath()
			_, err := db.Default().ID(recommend.Id).Cols("path").Update(recommend)
			if err != nil {
				log.Println(err)
			}
		} else {
			var parentRecommend bolejiang.PassageRecommend
			ok, err := db.Default().ID(recommend.ParentPassageRecommendId).Get(&parentRecommend)
			if err != nil {
				log.Println(err)
			}
			if !ok {
				log.Println("not extists")
			}
			if parentRecommend.Path != "" {
				recommend.Path = fmt.Sprintf("%s-%d", parentRecommend.Path, parentRecommend.Id)
				recommend.PathFull = recommend.GetFullPath()
				_, err := db.Default().ID(recommend.Id).Cols("path", "path_full").Update(recommend)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
