package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateProjectAction(c *gin.Context) {
	type Request struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Role        string `json:"role"`
		StartTime   int64  `json:"startTime"`
		EndTime     int64  `json:"endTime"`
		Description string `json:"description"`
		Performance string `json:"performance"`
		Link        string `json:"link"`
	}
	var request Request
	var accountProject bolejiang.AccountProject
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

		ok, err = db.Default().Where("id = ? and account_id = ?", request.ID, user.Id).Get(&accountProject)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("项目信息不存在")
		}

		accountProject.AccountId = user.Id
		accountProject.Name = request.Name
		accountProject.Role = request.Role
		accountProject.StartTime = request.StartTime
		accountProject.EndTime = request.EndTime
		accountProject.Description = request.Description
		accountProject.Performance = request.Performance
		accountProject.Link = request.Link
		accountProject.UpdatedTime = time.Now().Unix()
		_, err = db.Default().ID(accountProject.Id).AllCols().Update(accountProject)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, accountProject)
}
