package passage

import (
	"app/db"
	"app/models/bolejiang"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// accountShareInfo 组装「账号 + 其在某职位下的推荐/分享统计」响应体，
// GetAction 中推荐人、当前用户等多处复用。
func accountShareInfo(account bolejiang.Account, recommend bolejiang.PassageRecommend) gin.H {
	return gin.H{
		"id":                 account.Id,
		"name":               account.Name,
		"mobile":             account.Mobile,
		"recommendCount":     recommend.RecommendCount,
		"recommendCountL2":   recommend.RecommendCountL2,
		"shareCount":         recommend.ShareCount,
		"shareCountL2":       recommend.ShareCountL2,
		"passageRecommendId": recommend.Id,
	}
}

// applySimilarFilter 按「相似职位」追加行业/职位类别前缀过滤，list 与 listLike 共用。
func applySimilarFilter(query *gorm.DB, similarPassageId int) (*gorm.DB, error) {
	if similarPassageId == 0 {
		return query, nil
	}
	var similarPassage bolejiang.Passage
	ok, err := db.Get(db.Default().Where("id = ?", similarPassageId), &similarPassage)
	if err != nil {
		return query, err
	}
	if ok {
		query = query.Where("passage.industry_path like ?", similarPassage.IndustryPath+"%")
		query = query.Where("passage.position_tag_path like ?", similarPassage.PositionTagPath+"%")
	}
	return query, nil
}

// applyPassageGeoFilters 追加职位列表的「城市/区县/行业/职位类别」过滤，list 与 listLike 共用。
func applyPassageGeoFilters(query *gorm.DB, cityId, districtId int, industryPath, positionTagPath string) *gorm.DB {
	if cityId != 0 {
		query = query.Where("passage.city_id = ?", cityId)
	}
	if districtId != 0 {
		query = query.Where("passage.district_id = ?", districtId)
	}
	if industryPath != "" {
		query = query.Where("(passage.industry_path like ? or passage.industry_path = ?)", industryPath+"-%", industryPath)
	}
	if positionTagPath != "" {
		query = query.Where("(passage.position_tag_path like ? or passage.position_tag_path = ?)", positionTagPath+"-%", positionTagPath)
	}
	return query
}
