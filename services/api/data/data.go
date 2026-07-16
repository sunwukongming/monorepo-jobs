package data

import (
	"app/db"
	"app/models/bolejiang"
	"fmt"
)

const FidIndustry = 2
const FidPositionTag = 48

var Cities []bolejiang.DataCity
var Districts []bolejiang.DataDistrict
var Dictionaries []bolejiang.DataDictionary
var Industries []bolejiang.DataIndustry
var PositionTags []bolejiang.DataPositionTag
var CityMap map[int]bolejiang.DataCity = make(map[int]bolejiang.DataCity, 0)
var DistrictMap map[int]bolejiang.DataDistrict = make(map[int]bolejiang.DataDistrict, 0)
var IndustryMap map[string]bolejiang.DataIndustry = make(map[string]bolejiang.DataIndustry, 0)
var PositionTagMap map[string]bolejiang.DataPositionTag = make(map[string]bolejiang.DataPositionTag, 0)
var ComposedCities []bolejiang.ComposedCity
var ComposedIndustries []bolejiang.ComposedIndustry
var ComposedPositionTags []bolejiang.ComposedPositionTag

var DictionaryMap map[string]DictionaryBox

type DictionaryBox struct {
	Key  string                     `json:"key"`
	Name string                     `json:"name"`
	List []bolejiang.DataDictionary `json:"list"`
}

func (box DictionaryBox) GetItem(id int) bolejiang.DataDictionary {
	for i := range box.List {
		if box.List[i].Id == id {
			return box.List[i]
		}
	}
	return bolejiang.DataDictionary{}
}

func Load() error {
	cities := make([]bolejiang.DataCity, 0)
	err := db.Default().OrderBy("sort, id").Find(&cities)
	if err != nil {
		return err
	}
	districts := make([]bolejiang.DataDistrict, 0)
	err = db.Default().OrderBy("sort, id").Find(&districts)
	if err != nil {
		return err
	}

	industries := make([]bolejiang.DataIndustry, 0)
	err = db.Default().OrderBy("level, sort, id").Find(&industries)
	if err != nil {
		return err
	}

	positionTags := []bolejiang.DataPositionTag{}
	err = db.Default().OrderBy("level, sort, id").Find(&positionTags)
	if err != nil {
		return err
	}

	dictionaries := []bolejiang.DataDictionary{}
	err = db.Default().OrderBy("fid, sort, id").Find(&dictionaries)
	if err != nil {
		return err
	}

	for _, city := range cities {
		CityMap[city.Id] = city
	}
	for _, district := range districts {
		DistrictMap[district.Id] = district
	}

	composedCities := make([]bolejiang.ComposedCity, 0)
	for _, city := range cities {
		composedCity := bolejiang.ComposedCity{
			DataCity: city,
			Children: []bolejiang.DataDistrict{},
		}
		for _, district := range districts {
			if district.CityId == composedCity.Id {
				composedCity.Children = append(composedCity.Children, district)
			}
		}
		composedCities = append(composedCities, composedCity)
	}

	composedIndustries := make([]bolejiang.ComposedIndustry, 0)
	for _, industry := range industries {
		dbPath := industry.Path
		if industry.Level == 1 {
			industry.Path = fmt.Sprintf("%d", industry.Id)
			composedIndustries = append(composedIndustries, bolejiang.ComposedIndustry{
				DataIndustry: industry,
				Children:     []bolejiang.ComposedIndustry{},
			})
		}
		if industry.Level == 2 {
			for i := range composedIndustries {
				if composedIndustries[i].Id == industry.Pid {
					industry.Path = fmt.Sprintf("%d-%d", composedIndustries[i].Id, industry.Id)
					composedIndustries[i].Children = append(composedIndustries[i].Children, bolejiang.ComposedIndustry{
						DataIndustry: industry,
						Children:     []bolejiang.ComposedIndustry{},
					})
				}
			}
		}
		if industry.Level == 3 {
			for i := range composedIndustries {
				for j := range composedIndustries[i].Children {
					if composedIndustries[i].Children[j].Id == industry.Pid {
						industry.Path = fmt.Sprintf("%d-%d-%d", composedIndustries[i].Id, composedIndustries[i].Children[j].Id, industry.Id)
						composedIndustries[i].Children[j].Children = append(composedIndustries[i].Children[j].Children, bolejiang.ComposedIndustry{
							DataIndustry: industry,
							Children:     []bolejiang.ComposedIndustry{},
						})
					}
				}
			}
		}
		if dbPath != industry.Path {
			_, _ = db.Default().ID(industry.Id).Cols("path").Update(industry)
		}
		IndustryMap[industry.Path] = industry
	}

	composedPositionTags := make([]bolejiang.ComposedPositionTag, 0)

	for _, positionTag := range positionTags {

		dbPath := positionTag.Path
		if positionTag.Level == 1 {
			positionTag.Path = fmt.Sprintf("%d", positionTag.Id)
			composedPositionTags = append(composedPositionTags, bolejiang.ComposedPositionTag{
				DataPositionTag: positionTag,
				Children:        []bolejiang.ComposedPositionTag{},
			})
		}
		if positionTag.Level == 2 {
			for i := range composedPositionTags {
				if composedPositionTags[i].Id == positionTag.Pid {
					positionTag.Path = fmt.Sprintf("%d-%d", composedPositionTags[i].Id, positionTag.Id)
					composedPositionTags[i].Children = append(composedPositionTags[i].Children, bolejiang.ComposedPositionTag{
						DataPositionTag: positionTag,
						Children:        []bolejiang.ComposedPositionTag{},
					})
				}
			}
		}
		if positionTag.Level == 3 {
			for i := range composedPositionTags {
				for j := range composedPositionTags[i].Children {
					if composedPositionTags[i].Children[j].Id == positionTag.Pid {
						positionTag.Path = fmt.Sprintf("%d-%d-%d", composedPositionTags[i].Id, composedPositionTags[i].Children[j].Id, positionTag.Id)
						composedPositionTags[i].Children[j].Children = append(composedPositionTags[i].Children[j].Children, bolejiang.ComposedPositionTag{
							DataPositionTag: positionTag,
							Children:        []bolejiang.ComposedPositionTag{},
						})
					}
				}
			}
		}
		if dbPath != positionTag.Path {
			_, _ = db.Default().ID(positionTag.Id).Cols("path").Update(positionTag)
		}
		PositionTagMap[positionTag.Path] = positionTag
	}

	Cities = cities
	Districts = districts
	Industries = industries
	PositionTags = positionTags
	ComposedCities = composedCities
	ComposedIndustries = composedIndustries
	ComposedPositionTags = composedPositionTags

	//获取字典配置
	var ddcs []bolejiang.DataDictionaryConfig
	err = db.Default().Find(&ddcs)
	if err != nil {
		return err
	}
	DictionaryMap = make(map[string]DictionaryBox)
	for _, ddc := range ddcs {
		dictBox := DictionaryBox{}
		dictBox.Key = ddc.Key
		dictBox.Name = ddc.Name
		for _, item := range dictionaries {
			if item.Fid == ddc.Fid {
				dictBox.List = append(dictBox.List, item)
			}
		}
		DictionaryMap[ddc.Key] = dictBox
	}
	return nil
}

func GetCities() []bolejiang.DataCity {
	return Cities
}

func GetComposedCities() []bolejiang.ComposedCity {
	return ComposedCities
}

func GetIndustries() []bolejiang.DataIndustry {
	return Industries
}

func GetComposedIndustries() []bolejiang.ComposedIndustry {
	return ComposedIndustries
}

func GetPositionTags() []bolejiang.DataPositionTag {
	return PositionTags
}

func GetComposedPositionTags() []bolejiang.ComposedPositionTag {
	return ComposedPositionTags
}
