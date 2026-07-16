package services

import (
	"app/db"
	"strings"

	"github.com/sirupsen/logrus"
)

func CountUpdatePassageRecommendByPath(path string) error {
	ids := "(" + strings.Join(strings.Split(path, "-"), ",") + ")"
	var err error
	// 更新职位推荐数量
	_, err = db.Default().Exec(`update passage_recommend as pr set 
		recommend_count = (select count(*) from deliver where passage_id = pr.passage_id and recommend_account_id = pr.account_id and is_real = 1 and recommend_account_id != account_id), 
		recommend_count_l2 = (select count(*) from deliver where passage_id = pr.passage_id and passage_recommend_path = concat(pr.path, '-', pr.id) and is_real = 1), 
		share_count = (select count(*) from deliver where passage_id = pr.passage_id and recommend_account_id = pr.account_id and recommend_account_id != account_id),
		share_count_l2 = (select count(*) from deliver where passage_id = pr.passage_id and passage_recommend_path = concat(pr.path, '-', pr.id))
		where id in ` + ids)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"action":  "counter_update",
			"message": err.Error(),
		}).Error()
		return err
	}
	return nil
}

func CountUpdatePassageRecommend(passageRecommendId int) error {
	var err error
	// 更新职位推荐数量
	_, err = db.Default().Exec(`update passage_recommend as pr set 
		recommend_count = (select count(*) from deliver where passage_id = pr.passage_id and recommend_account_id = pr.account_id and is_real = 1 and recommend_account_id != account_id), 
		recommend_count_l2 = (select count(*) from deliver where passage_id = pr.passage_id and passage_recommend_path = concat(pr.path, '-', pr.id) and is_real = 1), 
		share_count = (select count(*) from deliver where passage_id = pr.passage_id and recommend_account_id = pr.account_id and recommend_account_id != account_id),
		share_count_l2 = (select count(*) from deliver where passage_id = pr.passage_id and passage_recommend_path = concat(pr.path, '-', pr.id))
		where id = ?`, passageRecommendId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"action":  "counter_update",
			"message": err.Error(),
		}).Error()
		return err
	}
	return nil
}

func CountUpdateAccount(accountId int) error {
	var err error
	//更新用户自荐数量
	_, err = db.Default().Exec("update account as a set self_recommend = (select count(*) from deliver where account_id = a.id and is_real = 1) where id = ?", accountId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"action":  "counter_update",
			"message": err.Error(),
		}).Error()
	}
	// 更新用户推荐数量
	_, err = db.Default().Exec("update account as a set recommend = (select count(*) from deliver where recommend_account_id = a.id and is_real = 1 and recommend_account_id != account_id) where id = ?", accountId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"action":  "counter_update",
			"message": err.Error(),
		}).Error()
	}
	return nil
}
