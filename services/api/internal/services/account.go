package services

import (
	"app/db"
	"app/models/bolejiang"
	"errors"

	"github.com/sirupsen/logrus"
)

func AccountGetByMobile(mobile string) (*bolejiang.Account, error) {
	var account bolejiang.Account
	ok, err := db.Default().Where("mobile = ?", mobile).Get(&account)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return &account, nil
}

func AccountBindMobile(account bolejiang.Account, mobile string) error {
	account.Mobile = mobile
	n, err := db.Default().ID(account.Id).Cols("mobile").Update(account)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("手机号更新失败")
	}
	var accounts []bolejiang.Account
	err = db.Default().Where("mobile = ? and id != ?", mobile, account.Id).Find(&accounts)
	if err != nil {
		return err
	}
	accountIds := []int{}
	for _, item := range accounts {
		accountIds = append(accountIds, item.Id)
	}
	if len(accountIds) > 0 {
		_, err = db.Default().Table(bolejiang.Account{}).In("id", accountIds).Update(map[string]interface{}{"mobile": ""})
		if err != nil {
			return err
		}
	}
	return nil
}

func AccountUpdateRelevant(account bolejiang.Account) {
	_, err := db.Default().Table(bolejiang.AccountApply{}).Where("account_id = ?", account.Id).Update(map[string]interface{}{
		"current_state": account.CurrentState,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"action":  "account_relevant_update",
			"message": err.Error(),
		})
	}

	_, err = db.Default().Exec(`update deliver as d set 
	account_name = (select name from account as a where d.account_id = a.id),
	account_mobile = (select mobile from account as a where d.account_id = a.id),
	name = (select name from account as a where d.account_id = a.id),
	mobile = (select mobile from account as a where d.account_id = a.id),
	email = (select email from account as a where d.account_id = a.id)
	where account_id = ? `, account.Id)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"action":  "account_relevant_update",
			"message": err.Error(),
		})
	}
}
