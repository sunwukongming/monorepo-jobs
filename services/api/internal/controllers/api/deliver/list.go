package deliver

import (
	"app/db"
	"app/internal/services"
	"app/models/bolejiang"
	"app/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 我投递的人
func ListAction(c *gin.Context) {
	type Request struct {
		Keyword  string `json:"keyword"`
		Page     int    `json:"page"`
		PageSize int    `json:"pageSize"`
		IsSelf   int    `json:"isSelf"`
	}

	services.Handle(c, func() (interface{}, error) {
		var page *services.Page
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			return nil, err
		}
		accountId := services.AuthGetAccountID(c)
		//获取当前用户（Common 中间件已校验并写入 context，直接复用避免重复查询）
		accountPtr, err := services.AuthGetAccountOrError(c)
		if err != nil {
			return nil, err
		}
		currentAccount := *accountPtr
		// 仅取投递(deliver)本表数据；职位详情由下方 PassageListFullByIDs 单独补全。
		// （原 xorm `extends` 多表嵌入扫描 GORM 不支持，故改为单表查询 + 二次聚合。）
		var delivers []bolejiang.Deliver
		page = services.NewPage(request.Page, request.PageSize)
		query := db.Default().Table("deliver").Select("deliver.*").
			Joins("LEFT JOIN passage ON deliver.passage_id = passage.id").
			Joins("LEFT JOIN account ON deliver.account_id = account.id").
			Where("deliver.is_real > 0").
			Order("deliver.deliver_time desc, deliver.created_time desc")
		if request.IsSelf != 0 {
			query = query.Where("deliver.mobile = ?", currentAccount.Mobile)
		} else {
			query = query.Where("deliver.recommend_account_id = ? and deliver.mobile != ?", accountId, currentAccount.Mobile)
		}
		if request.Keyword != "" {
			query = query.Where("(passage.title like ? or passage.edit_content like ? or deliver.name like ? or deliver.mobile like ?)",
				"%"+request.Keyword+"%", "%"+request.Keyword+"%", "%"+request.Keyword+"%", "%"+request.Keyword+"%")
		}
		err = page.Execute(query, &delivers)
		if err != nil {
			return nil, err
		}

		passageIDs := make([]uint32, 0, len(delivers))
		for _, deliver := range delivers {
			passageIDs = append(passageIDs, uint32(deliver.PassageId))
		}
		passageFulls, err := services.PassageListFullByIDs(passageIDs, uint32(utils.IntVal(accountId)))
		if err != nil {
			return nil, err
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
		return page, nil
	})
}
