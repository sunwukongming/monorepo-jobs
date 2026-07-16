package user

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	Gender        string `json:"gender"`
	Birthday      string `json:"birthday"`
	Workday       string `json:"workday"`
	Company       string `json:"company"`
	Position      string `json:"position"`
	Wechat        string `json:"wechat"`
	University    string `json:"university"`
	Industry      string `json:"industry"`
	Description   string `json:"description"`
	Tags          string `json:"tags"`
	Reward        int    `json:"reward"`
	SelfRecommend int    `json:"selfRecommend"`
	Recommend     int    `json:"recommend"`
	Collect       int    `json:"collect"`
	CurrentState  string `json:"currentState"`
	Apply         Apply  `json:"apply"`
}

type Apply struct {
	IsPublic           int    `json:"isPublic"`
	CurrentState       string `json:"currentState"`
	CurrentCompany     string `json:"currentCompany"`
	CurrentPositionTag string `json:"currentPositionTag"`
	CurrentPosition    string `json:"currentPosition"`
	CurrentIndustry    string `json:"currentIndustry"`
	CurrentCity        string `json:"currentCity"`
	DestCity           string `json:"destCity"`
	DestPositionTag    string `json:"destPositionTag"`
	DestPosition       string `json:"destPosition"`
	DestCompany        string `json:"destCompany"`
	DestIndustry       string `json:"destIndustry"`
	DestSalary         string `json:"destSalary"`
	Education          string `json:"education"`
	University         string `json:"university"`
	Description        string `json:"description"`
}

func UpdateAction(c *gin.Context) {
	var request UpdateRequest
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		request.Name = utils.StringTrim(request.Name)
		//request.Mobile = utils.StringTrim(request.Mobile)
		request.Email = utils.StringTrim(request.Email)
		toUpdateRelevant := false
		accountId := services.AuthGetAccountID(c)
		var account bolejiang.Account
		ok, err := db.Default().Where("id = ?", accountId).Get(&account)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("用户不存在")
		}
		cols := []string{}
		if utils.StringTrim(request.Name) != "" && account.Name != utils.StringTrim(request.Name) {
			account.Name = utils.StringTrim(request.Name)
			cols = append(cols, "name")
			toUpdateRelevant = true
		}
		if request.Company != "" {
			account.Company = request.Company
			cols = append(cols, "company")
		}
		if utils.StringTrim(request.Email) != "" && account.Email != utils.StringTrim(request.Email) {
			account.Email = request.Email
			cols = append(cols, "email")
			toUpdateRelevant = true
		}
		/*
			if utils.StringTrim(request.Mobile) != "" && account.Mobile != utils.StringTrim(request.Mobile) {
				account.Mobile = utils.StringTrim(request.Mobile)
				cols = append(cols, "mobile")
				toUpdateRelevant = true
			}
		*/
		if request.University != "" {
			account.University = request.University
			cols = append(cols, "university")
		}
		if request.Position != "" {
			account.Position = request.Position
			cols = append(cols, "position")
		}
		if request.Wechat != "" {
			account.Wechat = request.Wechat
			cols = append(cols, "wechat")
		}
		if request.Industry != "" {
			account.Industry = request.Industry
			cols = append(cols, "industry")
		}
		if request.Description != "" {
			account.Description = request.Description
			cols = append(cols, "description")
		}
		if request.Tags != "" {
			account.Tags = request.Tags
			cols = append(cols, "tags")
		}
		if request.CurrentState != "" && account.CurrentState != request.CurrentState {
			account.CurrentState = request.CurrentState
			toUpdateRelevant = true
			cols = append(cols, "current_state")
		}
		if request.Birthday != "" {
			account.Birthday = request.Birthday
			cols = append(cols, "birthday")
		}
		if request.Workday != "" {
			account.Workday = request.Workday
			cols = append(cols, "workday")
		}

		if request.Gender != "" {
			account.Gender = request.Gender
			cols = append(cols, "gender")
		}

		_, err = db.Default().ID(account.Id).Cols(cols...).Update(account)
		if err != nil {
			return err
		}
		if toUpdateRelevant {
			services.AccountUpdateRelevant(account)
		}
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, nil)
}
