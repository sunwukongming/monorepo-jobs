package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateWorkAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		type Request struct {
			Company     string `json:"company"`
			Industry    string `json:"industry"`
			StartTime   int64  `json:"startTime"`
			EndTime     int64  `json:"endTime"`
			Position    string `json:"position"`
			Content     string `json:"content"`
			Performance string `json:"performance"`
			Skills      string `json:"skills"`
		}
		var request Request
		var accountWork bolejiang.AccountWork
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		id := services.AuthGetAccountID(c)
		var user bolejiang.Account
		ok, err := db.Get(db.Default().Where("id = ?", id), &user)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("用户不存在")
		}
		accountWork.AccountId = user.Id
		accountWork.Company = request.Company
		accountWork.Industry = request.Industry
		accountWork.StartTime = request.StartTime
		accountWork.EndTime = request.EndTime
		accountWork.Position = request.Position
		accountWork.Content = request.Content
		accountWork.Performance = request.Performance
		accountWork.Skills = request.Skills
		accountWork.CreatedTime = time.Now().Unix()
		accountWork.UpdatedTime = time.Now().Unix()
		err = db.Default().Create(&accountWork).Error
		if err != nil {
			return nil, err
		}
		return accountWork, nil
	})
}
