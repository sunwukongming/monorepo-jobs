package services

import (
	"app/data"
	"app/internal/db/mysql/model"
	"app/internal/db/mysql/query"

	"gorm.io/gorm"
)

type PassageFull struct {
	model.Passage
	OutName string `json:"outName"`
	Address string `json:"address"`

	CompanyName    string `json:"companyName"`
	CompanyOutName string `json:"companyOutName"`
	CompanyRemark  string `json:"companyRemark"`
	CompanyAddress string `json:"companyAddress"`

	PassageCompany PassageCompanyFull `json:"passageCompany"`
	RootCompany    RootCompanyFull    `json:"rootCompany"`

	CityName        string `json:"cityName"`
	DistrictName    string `json:"districtName"`
	IndustryName    string `json:"industryName"`
	PositionTagName string `json:"positionTagName"`
	IsLike          bool   `json:"isLike"`

	Amounts struct {
		CompanyPassageAmount int `json:"companyPassageAmount"`
		DeliverAmount        int `json:"deliverAmount"`
		InterviewAmount      int `json:"interviewAmount"`
	} `json:"amounts"`
}

func PassageGetFullByID(id uint32, accountID uint32) (PassageFull, error) {
	passages, err := PassageListFullByIDs([]uint32{id}, accountID)
	if err != nil {
		return PassageFull{}, err
	}
	if len(passages) == 0 {
		return PassageFull{}, gorm.ErrRecordNotFound
	}
	return passages[0], nil
}

func PassageListFullByIDs(ids []uint32, accountID uint32) ([]PassageFull, error) {
	passages, err := query.Passage.Where(query.Passage.ID.In(ids...)).Order(query.Passage.Mtime.Desc(), query.Passage.ID.Desc()).Find()
	if err != nil {
		return nil, err
	}
	passageIds := make([]uint32, 0, len(passages))
	psgCompanies := make([]uint32, 0, len(passages))
	for _, passage := range passages {
		passageIds = append(passageIds, passage.ID)
		psgCompanies = append(psgCompanies, passage.PsgCompany)
	}

	passageCompanies, err := query.PassageCompany.Where(query.PassageCompany.ID.In(psgCompanies...)).Find()
	if err != nil {
		return nil, err
	}

	passageCompanyFulls, err := CompanyFullFind(passageCompanies)
	if err != nil {
		return nil, err
	}

	rootCompanyIDs := make([]uint32, 0, len(passageCompanies))
	for i := range passageCompanies {
		rootCompanyIDs = append(rootCompanyIDs, passageCompanies[i].CompanyID)
	}

	rootCompanies, err := query.CompanyInfo.Where(query.CompanyInfo.ID.In(rootCompanyIDs...)).Find()
	if err != nil {
		return nil, err
	}

	rootCompanyFulls, err := RootCompanyFullFind(rootCompanies)
	if err != nil {
		return nil, err
	}

	var passageLikes []*model.PassageLike
	var delivers []struct {
		PassageID   uint32
		ProgressEid uint32
		Count       int
	}
	if accountID > 0 {
		passageLikes, err = query.PassageLike.Where(query.PassageLike.AccountID.Eq(accountID), query.PassageLike.PassageID.In(passageIds...)).Find()
		if err != nil {
			return nil, err
		}

		err := query.Deliver.Select(query.Deliver.PassageID, query.Deliver.ProgressEid, query.Deliver.ID.Count().As("count")).
			Where(query.Deliver.PassageID.In(passageIds...), query.Deliver.RecommendAccountID.Eq(accountID), query.Deliver.IsReal.Eq(1)).
			Group(query.Deliver.PassageID, query.Deliver.ProgressEid).Scan(&delivers)
		if err != nil {
			return nil, err
		}
	}

	var passageResponses []PassageFull
	for _, passage := range passages {
		item := PassageFull{
			Passage:         *passage,
			CityName:        data.CityMap[int(passage.CityID)].Name,
			DistrictName:    data.DistrictMap[int(passage.DistrictID)].Name,
			IndustryName:    data.IndustryMap[passage.IndustryPath].Name,
			PositionTagName: data.PositionTagMap[passage.PositionTagPath].Name,
		}
		for _, passageCompany := range passageCompanyFulls {
			if item.PsgCompany == passageCompany.ID {
				item.Address = passageCompany.Address
				item.OutName = passageCompany.OutName

				item.CompanyName = passageCompany.Name
				item.CompanyOutName = passageCompany.OutName
				item.CompanyRemark = passageCompany.Remark
				item.CompanyAddress = passageCompany.Address
				item.Amounts.CompanyPassageAmount = passageCompany.PassageAmount

				item.PassageCompany = *passageCompany
			}
		}

		for i := range rootCompanies {
			if item.PassageCompany.CompanyID == rootCompanyFulls[i].ID {
				item.RootCompany = *(rootCompanyFulls[i])
			}
		}

		for _, passageLike := range passageLikes {
			if item.Passage.ID == passageLike.AccountID {
				item.IsLike = true
			}
		}
		for _, deliver := range delivers {
			if deliver.PassageID == item.Passage.ID {
				item.Amounts.DeliverAmount += deliver.Count
				if deliver.ProgressEid > 3 {
					item.Amounts.InterviewAmount += deliver.Count
				}
			}
		}
		passageResponses = append(passageResponses, item)
	}
	return passageResponses, nil
}
