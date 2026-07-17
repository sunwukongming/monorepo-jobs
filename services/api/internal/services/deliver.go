package services

import (
	"app/db"
	"app/models/bolejiang"
)

func DeliverGetByPassageIdAndAccountId(passageId int, accountId int) (*bolejiang.Deliver, error) {
	var deliver bolejiang.Deliver
	ok, err := db.Get(db.Default().Where("account_id = ? and passage_id = ?", accountId, passageId), &deliver)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return &deliver, nil
}

func DeliverGetByPassageIdAndMobile(passageId int, mobile string) (*bolejiang.Deliver, error) {
	var deliver bolejiang.Deliver
	ok, err := db.Get(db.Default().Where("mobile = ? and passage_id = ?", mobile, passageId), &deliver)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return &deliver, nil
}
