package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProjectAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		type Request struct {
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
		accountProject.AccountId = user.Id
		accountProject.Name = request.Name
		accountProject.Role = request.Role
		accountProject.StartTime = request.StartTime
		accountProject.EndTime = request.EndTime
		accountProject.Description = request.Description
		accountProject.Performance = request.Performance
		accountProject.Link = request.Link
		accountProject.CreatedTime = time.Now().Unix()
		accountProject.UpdatedTime = time.Now().Unix()
		err = db.Default().Create(&accountProject).Error
		if err != nil {
			return nil, err
		}
		return accountProject, nil
	})
}
