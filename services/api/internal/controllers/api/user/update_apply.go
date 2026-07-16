package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateApplyAction(c *gin.Context) {
	type Request struct {
		Id                  int    `json:"id"`
		IsPublic            int    `json:"isPublic"`
		IsFirst             int    `json:"isFirst"`
		CurrentCompany      string `json:"currentCompany"`
		CurrentPositionTag  string `json:"currentPositionTag"`
		CurrentPosition     string `json:"currentPosition"`
		CurrentIndustry     string `json:"currentIndustry"`
		CurrentCity         string `json:"currentCity"`
		DestCity            string `json:"destCity"`
		DestPositionTag     string `json:"destPositionTag"`
		DestPosition        string `json:"destPosition"`
		DestCompany         string `json:"destCompany"`
		DestIndustry        string `json:"destIndustry"`
		DestSalary          string `json:"destSalary"`
		Education           string `json:"education"`
		University          string `json:"university"`
		Description         string `json:"description"`
		HelpReward          int    `json:"helpReward"`
		IsHelpRewardVisible int    `json:"isHelpRewardVisible"`
	}
	var request Request
	var accountApply bolejiang.AccountApply
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
		ok, err = db.Default().Where("id = ? and account_id = ?", request.Id, user.Id).Get(&accountApply)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("求职目标不存在")
		}

		var applies []bolejiang.AccountApply
		err = db.Default().Where("account_id = ? and id != ?", user.Id, accountApply.Id).Find(&applies)
		if err != nil {
			return err
		}

		accountApply.IsPublic = request.IsPublic
		accountApply.IsFirst = request.IsFirst
		if len(applies) == 0 {
			accountApply.IsFirst = 1
		}
		accountApply.CurrentCompany = request.CurrentCompany
		accountApply.CurrentPosition = request.CurrentPosition
		accountApply.CurrentIndustry = request.CurrentIndustry
		accountApply.CurrentPositionTag = request.CurrentPositionTag
		accountApply.CurrentCity = request.CurrentCity
		accountApply.DestCity = request.DestCity
		accountApply.DestCompany = request.DestCompany
		accountApply.DestPositionTag = request.DestPositionTag
		accountApply.DestPosition = request.DestPosition
		accountApply.DestIndustry = request.DestIndustry
		accountApply.DestSalary = request.DestSalary
		accountApply.Education = request.Education
		accountApply.University = request.University
		accountApply.Description = request.Description
		accountApply.CurrentState = user.CurrentState
		accountApply.HelpReward = request.HelpReward
		accountApply.HelpRewardToC = request.HelpReward * 7 / 10
		accountApply.IsHelpRewardVisible = request.IsHelpRewardVisible
		accountApply.UpdatedTime = time.Now().Unix()
		_, err = db.Default().Where("id = ?", accountApply.Id).AllCols().Update(accountApply)
		if err != nil {
			return err
		}
		if accountApply.IsFirst == 1 {
			db.Default().Table(bolejiang.AccountApply{}).Where("account_id = ? and id != ?", accountApply.AccountId, accountApply.Id).Update(map[string]interface{}{
				"is_first": 0,
			})
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, accountApply)
}
