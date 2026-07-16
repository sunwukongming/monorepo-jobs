package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"

	"github.com/gin-gonic/gin"
)

func InfoAction(c *gin.Context) {
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
	err := func() error {
		id := services.AuthGetAccountID(c)
		ok, err := db.Default().Where("id = ?", id).Get(&user)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("用户不存在")
		}
		var accountApplies []bolejiang.AccountApply
		err = db.Default().Where("account_id = ?", user.Id).OrderBy("is_first desc, created_time desc, id desc").Find(&accountApplies)
		if err != nil {
			return err
		}
		var accountWorks []bolejiang.AccountWork
		err = db.Default().Where("account_id = ?", user.Id).OrderBy("start_time desc, id desc").Find(&accountWorks)
		if err != nil {
			return err
		}
		var accountProjects []bolejiang.AccountProject
		err = db.Default().Where("account_id = ?", user.Id).OrderBy("start_time desc, id desc").Find(&accountProjects)
		if err != nil {
			return err
		}
		var accountEducations []bolejiang.AccountEducation
		err = db.Default().Where("account_id = ?", user.Id).OrderBy("start_time desc, id desc").Find(&accountEducations)
		if err != nil {
			return err
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
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, resp)
}
