package services

import (
	"app/data"
	"app/internal/db/mysql/model"
	"app/internal/db/mysql/query"
)

type RootCompanyFull struct {
	model.CompanyInfo
	PassageAmount    int    `json:"passageAmount"`
	CityName         string `json:"cityName"`
	DistrictName     string `json:"districtName"`
	IndustryName     string `json:"industryName"`
	CompanyscaleName string `json:"companyscaleName"`
	CompanystageName string `json:"companystageName"`
}

func RootCompanyFullGet(company *model.CompanyInfo) (*RootCompanyFull, error) {
	companyFulls, err := RootCompanyFullFind([]*model.CompanyInfo{company})
	if err != nil {
		return nil, err
	}
	return companyFulls[0], nil
}

func RootCompanyFullFind(companies []*model.CompanyInfo) ([]*RootCompanyFull, error) {
	ids := []uint32{}
	for _, company := range companies {
		ids = append(ids, company.ID)
	}

	childCompanies, err := query.PassageCompany.Where(query.PassageCompany.CompanyID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	childCompanyFulls, err := CompanyFullFind(childCompanies)
	if err != nil {
		return nil, err
	}
	rootCompanyFulls := make([]*RootCompanyFull, 0, len(companies))
	for i, company := range companies {
		rootCompanyFull := &RootCompanyFull{
			CompanyInfo:      *companies[i],
			CityName:         data.CityMap[int(companies[i].CityID)].Name,
			DistrictName:     data.DistrictMap[int(companies[i].DistrictID)].Name,
			IndustryName:     data.IndustryMap[companies[i].IndustryPath].Name,
			CompanyscaleName: data.DictionaryMap["scale"].GetItem(int(companies[i].CompanyscaleID)).Remark,
			CompanystageName: data.DictionaryMap["stage"].GetItem(int(companies[i].CompanystageID)).Remark,
		}
		for _, child := range childCompanyFulls {
			if child.CompanyID == company.ID {
				rootCompanyFull.PassageAmount += child.PassageAmount
			}
		}
		rootCompanyFulls = append(rootCompanyFulls, rootCompanyFull)
	}
	return rootCompanyFulls, nil
}
