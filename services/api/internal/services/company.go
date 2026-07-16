package services

import (
	"app/data"
	"app/internal/db/mysql/model"
	"app/internal/db/mysql/query"
)

type PassageCompanyFull struct {
	model.PassageCompany
	PassageAmount int    `json:"passageAmount"`
	CityName      string `json:"cityName"`
	DistrictName  string `json:"districtName"`
	IndustryName  string `json:"industryName"`
	ScaleName     string `json:"scaleName"`
	StageName     string `json:"stageName"`
}

func CompanyFullGet(company *model.PassageCompany) (*PassageCompanyFull, error) {
	companyFulls, err := CompanyFullFind([]*model.PassageCompany{company})
	if err != nil {
		return nil, err
	}
	return companyFulls[0], nil
}

func CompanyFullFind(companies []*model.PassageCompany) ([]*PassageCompanyFull, error) {
	ids := []uint32{}
	for _, company := range companies {
		ids = append(ids, company.ID)
	}
	type Amount struct {
		PsgCompany uint32
		Amount     int
	}
	var amounts []Amount
	// err := db.Default().Table(bolejiang.Passage{}).Select("psg_company, count(*) as amount").GroupBy("psg_company").
	// 	In("psg_company", ids).Find(&amounts)
	err := query.Passage.Select(query.Passage.PsgCompany, query.Passage.ID.Count().As("amount")).
		Where(query.Passage.PsgCompany.In(ids...)).Where(query.Passage.IsAnonymous.Eq(0)).
		Group(query.Passage.PsgCompany).Scan(&amounts)
	if err != nil {
		return nil, err
	}
	companyFulls := make([]*PassageCompanyFull, 0, len(companies))
	for i, company := range companies {
		amount := 0
		for _, item := range amounts {
			if item.PsgCompany == company.ID {
				amount = item.Amount
			}
		}
		companyFull := &PassageCompanyFull{
			PassageCompany: *companies[i],
			PassageAmount:  amount,
			CityName:       data.CityMap[int(companies[i].CityID)].Name,
			DistrictName:   data.DistrictMap[int(companies[i].DistrictID)].Name,
			IndustryName:   data.IndustryMap[companies[i].IndustryPath].Name,
			ScaleName:      data.DictionaryMap["scale"].GetItem(int(companies[i].Scale)).Remark,
			StageName:      data.DictionaryMap["stage"].GetItem(int(companies[i].Stage)).Remark,
		}
		companyFulls = append(companyFulls, companyFull)
	}
	return companyFulls, nil
}
