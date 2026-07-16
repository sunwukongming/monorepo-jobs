package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateWorkAction(c *gin.Context) {
	type Request struct {
		ID          string `json:"id"`
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
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		id := services.AuthGetAccountID(c)
		var user bolejiang.Account
		ok, err := db.Default().Where("id = ?", id).Get(&user)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("用户不存在")
		}

		ok, err = db.Default().Where("id = ? and account_id = ?", request.ID, user.Id).Get(&accountWork)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("工作信息不存在")
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
		accountWork.UpdatedTime = time.Now().Unix()
		_, err = db.Default().ID(accountWork.Id).Update(accountWork)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, accountWork)
}
