package services

import (
	"app/db"
	"app/models/bolejiang"
	"errors"

	"github.com/sirupsen/logrus"
)

func AccountGetByMobile(mobile string) (*bolejiang.Account, error) {
	var account bolejiang.Account
	ok, err := db.Get(db.Default().Where("mobile = ?", mobile), &account)
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
	result := db.Default().Model(&account).Where("id = ?", account.Id).Select("mobile").Updates(account)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("手机号更新失败")
	}
	var accounts []bolejiang.Account
	err := db.Default().Where("mobile = ? and id != ?", mobile, account.Id).Find(&accounts).Error
	if err != nil {
		return err
	}
	accountIds := []int{}
	for _, item := range accounts {
		accountIds = append(accountIds, item.Id)
	}
	if len(accountIds) > 0 {
		err = db.Default().Model(&bolejiang.Account{}).Where("id IN ?", accountIds).Updates(map[string]interface{}{"mobile": ""}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func AccountUpdateRelevant(account bolejiang.Account) {
	err := db.Default().Model(&bolejiang.AccountApply{}).Where("account_id = ?", account.Id).Updates(map[string]interface{}{
		"current_state": account.CurrentState,
	}).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"action":  "account_relevant_update",
			"message": err.Error(),
		})
	}

	err = db.Default().Exec(`update deliver as d set
	account_name = (select name from account as a where d.account_id = a.id),
	account_mobile = (select mobile from account as a where d.account_id = a.id),
	name = (select name from account as a where d.account_id = a.id),
	mobile = (select mobile from account as a where d.account_id = a.id),
	email = (select email from account as a where d.account_id = a.id)
	where account_id = ? `, account.Id).Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"action":  "account_relevant_update",
			"message": err.Error(),
		})
	}
}
