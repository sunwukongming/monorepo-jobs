package apply

import (
	"strings"

	"app/data"
	"app/pkg/utils"

	"gorm.io/gorm"
)

// ListRequest 求职列表 / 求职收藏列表的公共请求参数。
type ListRequest struct {
	Keyword             string `json:"keyword"`
	DestCity            string `json:"destCity"`
	DestIndustry        string `json:"destIndustry"`
	DestPosition        string `json:"destPosition"`
	DestCityId          int    `json:"destCityId"`
	DestIndustryId      int    `json:"destIndustryId"`
	DestPositionTagId   int    `json:"destPositionTagId"`
	DestIndustryPath    string `json:"destIndustryPath"`
	DestPositionTagPath string `json:"destPositionTagPath"`
	IsHelpRewardVisible string `json:"isHelpRewardVisible"`
	Page                int    `json:"page"`
	PageSize            int    `json:"pageSize"`
}

// applyDestFilters 追加「期望城市 / 行业 / 职位类别」相关的动态过滤条件。
// list 与 listLike 完全共用（列名不带表前缀，两处 SQL 一致）。
func applyDestFilters(session *gorm.DB, request ListRequest) *gorm.DB {
	if request.DestCityId != 0 {
		for _, v := range data.Cities {
			if v.Id == request.DestCityId {
				session = session.Where("dest_city like ?", "%"+v.Name+"%")
				break
			}
		}
	}
	// 按 industryId / industryPath 过滤：两种入口最终都是「该行业结点 + 其所有子级」
	if request.DestIndustryId != 0 {
		for _, v := range data.Industries {
			if v.Id == request.DestIndustryId {
				session = whereOrLike(session, "dest_industry", industryNamePatterns(v.Path))
				break
			}
		}
	}
	if request.DestIndustryPath != "" {
		session = whereOrLike(session, "dest_industry", industryNamePatterns(request.DestIndustryPath))
	}
	if request.DestPositionTagId != 0 {
		for _, v := range data.PositionTags {
			if request.DestPositionTagId == v.Id {
				session = whereOrLike(session, "dest_position_tag", positionTagNamePatterns(v.Path))
				break
			}
		}
	}
	if request.DestPositionTagPath != "" {
		session = whereOrLike(session, "dest_position_tag", positionTagNamePatterns(request.DestPositionTagPath))
	}
	return session
}

// whereOrLike 追加形如 "col like ? or col like ? ..." 的条件；patterns 为空则原样返回。
func whereOrLike(session *gorm.DB, col string, patterns []interface{}) *gorm.DB {
	if len(patterns) == 0 {
		return session
	}
	parts := make([]string, len(patterns))
	for i := range patterns {
		parts[i] = col + " like ?"
	}
	return session.Where(strings.Join(parts, " or "), patterns...)
}

// industryNamePatterns 返回「path 对应结点及其所有子级」的名称 like 模式。
func industryNamePatterns(path string) []interface{} {
	patterns := make([]interface{}, 0)
	for _, item := range data.Industries {
		if item.Path == path || utils.StringStartsWith(item.Path, path+"-") {
			patterns = append(patterns, "%"+item.Name+"%")
		}
	}
	return patterns
}

// positionTagNamePatterns 返回「path 对应结点及其所有子级」的名称 like 模式。
func positionTagNamePatterns(path string) []interface{} {
	patterns := make([]interface{}, 0)
	for _, item := range data.PositionTags {
		if item.Path == path || utils.StringStartsWith(item.Path, path+"-") {
			patterns = append(patterns, "%"+item.Name+"%")
		}
	}
	return patterns
}
