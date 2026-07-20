package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
)

func InfoAction(c *gin.Context) {
	services.Handle(c, func() (interface{}, error) {
		var user bolejiang.Account
		type Resp struct {
			bolejiang.Account
			Apply      bolejiang.AccountApply       `json:"apply"`
			Applies    []bolejiang.AccountApply     `json:"applies"`
			Works      []bolejiang.AccountWork      `json:"works"`
			Educations []bolejiang.AccountEducation `json:"educations"`
			Projects   []bolejiang.AccountProject   `json:"projects"`
		}
		var resp Resp
		id := services.AuthGetAccountID(c)
		ok, err := db.Get(db.Default().Where("id = ?", id), &user)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("用户不存在")
		}
		var accountApplies []bolejiang.AccountApply
		err = db.Default().Where("account_id = ?", user.Id).Order("is_first desc, created_time desc, id desc").Find(&accountApplies).Error
		if err != nil {
			return nil, err
		}
		var accountWorks []bolejiang.AccountWork
		err = db.Default().Where("account_id = ?", user.Id).Order("start_time desc, id desc").Find(&accountWorks).Error
		if err != nil {
			return nil, err
		}
		var accountProjects []bolejiang.AccountProject
		err = db.Default().Where("account_id = ?", user.Id).Order("start_time desc, id desc").Find(&accountProjects).Error
		if err != nil {
			return nil, err
		}
		var accountEducations []bolejiang.AccountEducation
		err = db.Default().Where("account_id = ?", user.Id).Order("start_time desc, id desc").Find(&accountEducations).Error
		if err != nil {
			return nil, err
		}
		resp.Account = user
		if len(accountApplies) > 0 {
			resp.Apply = accountApplies[0]
		} else {
			resp.Apply = bolejiang.AccountApply{}
		}
		resp.Applies = accountApplies
		resp.Works = accountWorks
		resp.Educations = accountEducations
		resp.Projects = accountProjects
		return resp, nil
	})
}
