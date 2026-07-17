package passage

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"

	"github.com/gin-gonic/gin"
)

func ListSelfRecommendAction(c *gin.Context) {
	type Request struct {
		Keyword  string `json:"keyword"`
		Page     int    `json:"page"`
		PageSize int    `json:"pageSize"`
	}

	var page *services.Page
	var request Request
	err := func() error {
		if err := c.ShouldBindJSON(&request); err != nil {
			return err
		}
		accountId := services.AuthGetAccountID(c)
		var delivers []bolejiang.Deliver
		page = services.NewPage(request.Page, request.PageSize)
		// 仅取投递(deliver)本表数据；对应职位详情由下方 PassageListFullByIDs 单独补全。
		// （原 xorm `extends` 多表嵌入扫描 GORM 不支持，故改为单表查询 + 二次聚合。）
		query := db.Default().Table("deliver").Select("deliver.*").
			Joins("LEFT JOIN passage ON deliver.passage_id = passage.id").
			Where("deliver.account_id = ? and deliver.is_real > 0", accountId).
			Order("deliver.deliver_time desc, deliver.created_time desc")
		if request.Keyword != "" {
			query = query.Where("(passage.title like ? or passage.edit_content like ? or account.name like ? or account.mobile like ?)",
				"%"+request.Keyword+"%", "%"+request.Keyword+"%", "%"+request.Keyword+"%", "%"+request.Keyword+"%")
		}
		err := page.Execute(query, &delivers)
		if err != nil {
			return err
		}

		passageIDs := make([]uint32, 0, len(delivers))
		for _, deliver := range delivers {
			passageIDs = append(passageIDs, uint32(deliver.PassageId))
		}
		passageFulls, err := services.PassageListFullByIDs(passageIDs, uint32(utils.IntVal(accountId)))
		if err != nil {
			return err
		}

		rows := []interface{}{}
		for _, deliver := range delivers {
			row := map[string]any{}
			for _, passageFull := range passageFulls {
				if passageFull.Passage.ID == uint32(deliver.PassageId) {
					row["passage"] = passageFull
				}
			}
			row["deliver"] = deliver
			rows = append(rows, row)
		}
		page.List = rows
		return nil
	}()
	if err != nil {
		services.ResponseError(c, -1, err.Error(), nil)
		return
	}
	services.ResponseSuccess(c, page)
}
