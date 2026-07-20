package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateEducationAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		type Request struct {
			Name       string `json:"name"`
			Degree     string `json:"degree"`
			Profession string `json:"profession"`
			StartTime  int64  `json:"startTime"`
			EndTime    int64  `json:"endTime"`
			Experience string `json:"experience"`
		}
		var request Request
		var accountEducation bolejiang.AccountEducation
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
		accountEducation.AccountId = user.Id
		accountEducation.Name = request.Name
		accountEducation.Degree = request.Degree
		accountEducation.Profession = request.Profession
		accountEducation.StartTime = request.StartTime
		accountEducation.EndTime = request.EndTime
		accountEducation.Experience = request.Experience
		accountEducation.CreatedTime = time.Now().Unix()
		accountEducation.UpdatedTime = time.Now().Unix()
		err = db.Default().Create(&accountEducation).Error
		if err != nil {
			return nil, err
		}
		return accountEducation, nil
	})
}
