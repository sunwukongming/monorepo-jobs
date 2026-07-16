package bolejiang

import (
	"time"
)

type Account struct {
	Id              int    `json:"id" xorm:"not null pk autoincr comment('主键') INT(10)"`
	Openid          string `json:"openid" xorm:"not null default '' VARCHAR(45)"`
	Unionid         string `json:"unionid" xorm:"not null default '' VARCHAR(45)"`
	Name            string `json:"name" xorm:"not null default '' comment('姓名') VARCHAR(20)"`
	Mobile          string `json:"mobile" xorm:"not null comment('手机号') index VARCHAR(11)"`
	Email           string `json:"email" xorm:"not null default '' comment('邮箱') index VARCHAR(64)"`
	Gender          string `json:"gender" xorm:"not null default '' comment('性别') VARCHAR(2)"`
	Birthday        string `json:"birthday" xorm:"not null default '' comment('生日') VARCHAR(16)"`
	Workday         string `json:"workday" xorm:"not null default '' comment('工作时间') VARCHAR(16)"`
	University      string `json:"university" xorm:"not null default '' comment('学校') VARCHAR(64)"`
	Wechat          string `json:"wechat" xorm:"not null default '' comment('微信号') VARCHAR(64)"`
	Company         string `json:"company" xorm:"not null default '' comment('公司') VARCHAR(64)"`
	Position        string `json:"position" xorm:"not null default '' comment('职位') VARCHAR(64)"`
	Industry        string `json:"industry" xorm:"not null default '' comment('行业') VARCHAR(32)"`
	Description     string `json:"description" xorm:"not null default '' comment('个人简介') VARCHAR(10240)"`
	Tags            string `json:"tags" xorm:"not null default '' comment('标签') VARCHAR(128)"`
	IsResumeWatcher int    `json:"is_resume_watcher" xorm:"not null default 0 comment('是否查看所有简历') TINYINT(4)"`
	IsAllies        int    `json:"is_allies" xorm:"not null default 0 comment('是否合作伙伴') TINYINT(4)"`
	Reward          int    `json:"reward" xorm:"not null default 0 comment('奖金') INT(11)"`
	SelfRecommend   int    `json:"self_recommend" xorm:"not null default 0 comment('自荐') INT(11)"`
	Recommend       int    `json:"recommend" xorm:"not null default 0 comment('推荐') INT(11)"`
	Collect         int    `json:"collect" xorm:"not null default 0 comment('收藏') INT(11)"`
	CurrentState    string `json:"current_state" xorm:"not null default '' comment('在职状态') VARCHAR(45)"`
	ResumeUrl       string `json:"resume_url" xorm:"not null default '' comment('简历地址') VARCHAR(255)"`
	CreatedTime     int64  `json:"created_time" xorm:"not null default 0 comment('创建时间') BIGINT(20)"`
	UpdatedTime     int64  `json:"updated_time" xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
	Status          int    `json:"status" xorm:"not null default 0 comment('用户状态') TINYINT(4)"`
}

type AccountApply struct {
	Id                  int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	AccountId           int    `json:"account_id" xorm:"not null default 0 comment('用户id') index INT(10)"`
	IsPublic            int    `json:"is_public" xorm:"not null default 0 comment('是否公开') TINYINT(4)"`
	IsFirst             int    `json:"is_first" xorm:"not null default 0 comment('是否优先') TINYINT(4)"`
	CurrentState        string `json:"current_state" xorm:"not null VARCHAR(45)"`
	CurrentCity         string `json:"current_city" xorm:"not null VARCHAR(45)"`
	CurrentIndustry     string `json:"current_industry" xorm:"not null VARCHAR(45)"`
	CurrentCompany      string `json:"current_company" xorm:"not null VARCHAR(45)"`
	CurrentPositionTag  string `json:"current_position_tag" xorm:"not null comment('当前职位类别') VARCHAR(45)"`
	CurrentPosition     string `json:"current_position" xorm:"not null VARCHAR(45)"`
	DestCity            string `json:"dest_city" xorm:"not null VARCHAR(45)"`
	DestIndustry        string `json:"dest_industry" xorm:"not null comment('期望行业') VARCHAR(45)"`
	DestCompany         string `json:"dest_company" xorm:"not null VARCHAR(45)"`
	DestPositionTag     string `json:"dest_position_tag" xorm:"not null comment('期望职位类别') VARCHAR(45)"`
	DestPosition        string `json:"dest_position" xorm:"not null VARCHAR(45)"`
	DestSalary          string `json:"dest_salary" xorm:"not null comment('目标薪资') VARCHAR(45)"`
	Education           string `json:"education" xorm:"not null VARCHAR(45)"`
	University          string `json:"university" xorm:"not null VARCHAR(45)"`
	Description         string `json:"description" xorm:"not null TEXT"`
	Contact             string `json:"contact" xorm:"not null default '' comment('联系方式') VARCHAR(200)"`
	HelpReward          int    `json:"help_reward" xorm:"not null default 0 comment('协助奖金') INT(10)"`
	HelpRewardToC       int    `json:"help_reward_to_c" xorm:"not null default 0 comment('协助奖金显示') INT(10)"`
	IsHelpRewardVisible int    `json:"is_help_reward_visible" xorm:"not null default 0 comment('协助奖金是否显示') TINYINT(3)"`
	CreatedTime         int64  `json:"created_time" xorm:"not null default 0 comment('创建时间') BIGINT(20)"`
	UpdatedTime         int64  `json:"updated_time" xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
	Status              int    `json:"status" xorm:"not null default 0 comment('求职状态') TINYINT(4)"`
}

type AccountApplyLike struct {
	Id             int `json:"id" xorm:"not null pk autoincr INT(10)"`
	AccountId      int `json:"account_id" xorm:"not null unique(uiq_account_apply) INT(10)"`
	AccountApplyId int `json:"account_apply_id" xorm:"not null unique(uiq_account_apply) INT(10)"`
}

type AccountApplyResumeAuth struct {
	Id               int   `json:"id" xorm:"not null pk autoincr INT(10)"`
	AccountApplyId   int   `json:"account_apply_id" xorm:"not null unique(uiq_account_apply_id_request_account_id) INT(10)"`
	RequestAccountId int   `json:"request_account_id" xorm:"not null unique(uiq_account_apply_id_request_account_id) INT(10)"`
	AuthState        int   `json:"auth_state" xorm:"not null default 0 comment('0 未处理 1 通过 其他拒绝') TINYINT(3)"`
	CreatedTime      int64 `json:"created_time" xorm:"not null default 0 comment('创建时间') BIGINT(20)"`
	UpdatedTime      int64 `json:"updated_time" xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
	Status           int   `json:"status" xorm:"not null default 0 TINYINT(3)"`
}

type AccountEducation struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	AccountId   int    `json:"account_id" xorm:"not null index INT(10)"`
	Name        string `json:"name" xorm:"not null comment('学校名称') VARCHAR(64)"`
	Degree      string `json:"degree" xorm:"not null comment('学历') VARCHAR(32)"`
	Profession  string `json:"profession" xorm:"not null comment('专业') VARCHAR(64)"`
	StartTime   int64  `json:"start_time" xorm:"not null comment('在职时间 - 开始') BIGINT(20)"`
	EndTime     int64  `json:"end_time" xorm:"not null comment('在职时间-结束') BIGINT(20)"`
	Experience  string `json:"experience" xorm:"not null default '' comment('经历') VARCHAR(4000)"`
	CreatedTime int64  `json:"created_time" xorm:"not null default 0 comment('创建时间') BIGINT(20)"`
	UpdatedTime int64  `json:"updated_time" xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
}

type AccountHelp struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	AccountApplyId  int    `json:"account_apply_id" xorm:"not null comment('被协助者求职id') unique(uiq_account_id_help_account_id) INT(10)"`
	HelperAccountId int    `json:"helper_account_id" xorm:"not null comment('协助人ID') index unique(uiq_account_id_help_account_id) INT(10)"`
	Company         string `json:"company" xorm:"not null default '' comment('公司') VARCHAR(255)"`
	Position        string `json:"position" xorm:"not null default '' comment('职位') VARCHAR(255)"`
	HelpPlan        string `json:"help_plan" xorm:"not null TEXT"`
	CreatedTime     int64  `json:"created_time" xorm:"not null default 0 comment('创建时间') BIGINT(20)"`
	UpdatedTime     int64  `json:"updated_time" xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
}

type AccountProject struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	AccountId   int    `json:"account_id" xorm:"not null index INT(10)"`
	Name        string `json:"name" xorm:"not null comment('项目名称') VARCHAR(64)"`
	Role        string `json:"role" xorm:"not null comment('角色') VARCHAR(32)"`
	StartTime   int64  `json:"start_time" xorm:"not null comment('项目时间 - 开始') BIGINT(20)"`
	EndTime     int64  `json:"end_time" xorm:"not null comment('项目时间-结束') BIGINT(20)"`
	Description string `json:"description" xorm:"not null comment('描述') TEXT"`
	Performance string `json:"performance" xorm:"not null comment('业绩') TEXT"`
	Link        string `json:"link" xorm:"not null default '' comment('项目连接') VARCHAR(512)"`
	CreatedTime int64  `json:"created_time" xorm:"not null default 0 comment('创建时间') BIGINT(20)"`
	UpdatedTime int64  `json:"updated_time" xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
}

type AccountWork struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	AccountId   int    `json:"account_id" xorm:"not null index INT(10)"`
	Company     string `json:"company" xorm:"not null comment('公司') VARCHAR(64)"`
	Industry    string `json:"industry" xorm:"not null comment('行业') VARCHAR(32)"`
	StartTime   int64  `json:"start_time" xorm:"not null comment('在职时间 - 开始') BIGINT(20)"`
	EndTime     int64  `json:"end_time" xorm:"not null comment('在职时间-结束') BIGINT(20)"`
	Position    string `json:"position" xorm:"not null VARCHAR(64)"`
	Content     string `json:"content" xorm:"not null comment('工作内容') TEXT"`
	Performance string `json:"performance" xorm:"not null default '' comment('工作业绩') VARCHAR(2000)"`
	Skills      string `json:"skills" xorm:"not null default '' comment('技能') VARCHAR(200)"`
	CreatedTime int64  `json:"created_time" xorm:"not null default 0 comment('创建时间') BIGINT(20)"`
	UpdatedTime int64  `json:"updated_time" xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
}

type Admin struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name          string `json:"name" xorm:"not null unique VARCHAR(255)"`
	Password      string `json:"password" xorm:"not null VARCHAR(255)"`
	Type          int    `json:"type" xorm:"comment('0=总账户、1=企业账户、2=HR账户、3=特别企业账户') index TINYINT(4)"`
	FatherAdminId int    `json:"father_admin_id" xorm:"index INT(11)"`
	CountPassage  int    `json:"count_passage" xorm:"not null default 0 INT(11)"`
	CountApply    int    `json:"count_apply" xorm:"not null default 0 INT(11)"`
	CountRelay    int    `json:"count_relay" xorm:"not null default 0 INT(11)"`
	Status        int    `json:"status" xorm:"not null default 0 comment('0、正常，1、审核中，2、停用') INT(11)"`
	CompanyId     int    `json:"company_id" xorm:"INT(11)"`
	Score         int    `json:"score" xorm:"not null default 0 comment('积分') INT(11)"`
}

type Apply struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	ConnectionId int    `json:"connection_id" xorm:"index INT(10)"`
	Name         string `json:"name" xorm:"VARCHAR(30)"`
	Phone        string `json:"phone" xorm:"index(phone) VARCHAR(30)"`
	Email        string `json:"email" xorm:"VARCHAR(30)"`
	WechatId     string `json:"wechat_id" xorm:"default '' comment('微信号') VARCHAR(50)"`
	Position     string `json:"position" xorm:"VARCHAR(30)"`
	PassageId    int    `json:"passage_id" xorm:"index(phone) INT(11)"`
	Resume       string `json:"resume" xorm:"VARCHAR(255)"`
	ResumeName   string `json:"resume_name" xorm:"VARCHAR(255)"`
	Remarks      string `json:"remarks" xorm:"VARCHAR(255)"`
	Time         string `json:"time" xorm:"VARCHAR(255)"`
	Birth        string `json:"birth" xorm:"VARCHAR(255)"`
	HomeUid      int    `json:"home_uid" xorm:"index INT(11)"`
	Status       int    `json:"status" xorm:"default 0 INT(11)"`
	NowCompany   string `json:"now_company" xorm:"VARCHAR(255)"`
	NowPosition  string `json:"now_position" xorm:"VARCHAR(255)"`
	Type         int    `json:"type" xorm:"default 1 comment('1=自己应聘、2=推荐他人') TINYINT(4)"`
}

type ApplyNoDirect struct {
	UserId  int `json:"user_id" xorm:"not null pk default 0 INT(11)"`
	ApplyId int `json:"apply_id" xorm:"not null pk default 0 INT(11)"`
}

type ArticleLike struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	ArticleTable string `json:"article_table" xorm:"not null unique(uiq_article_table_article_id_account_id) VARCHAR(20)"`
	ArticleId    int    `json:"article_id" xorm:"not null unique(uiq_article_table_article_id_account_id) INT(10)"`
	AccountId    int    `json:"account_id" xorm:"not null unique(uiq_article_table_article_id_account_id) INT(10)"`
}

type Banner struct {
	Id      int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Url     string `json:"url" xorm:"not null default '' VARCHAR(255)"`
	Type    string `json:"type" xorm:"not null default '' VARCHAR(16)"`
	JumpUrl string `json:"jump_url" xorm:"not null default '' VARCHAR(255)"`
	Sort    int    `json:"sort" xorm:"not null default 0 INT(10)"`
}

type CompanyAddress struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	CompanyId  int    `json:"company_id" xorm:"default 0 comment('公司Id') index INT(10)"`
	Address    string `json:"address" xorm:"default '' comment('地址') VARCHAR(2048)"`
	TimeInsert int    `json:"time_insert" xorm:"default 0 comment('添加时间') INT(10)"`
	Status     int    `json:"status" xorm:"default 2 comment('0-删除，1-暂不使用，2-正常') TINYINT(4)"`
}

type CompanyInfo struct {
	Id                  int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	AdminId             int    `json:"admin_id" xorm:"not null default 0 INT(10)"`
	Name                string `json:"name" xorm:"not null default '' VARCHAR(255)"`
	FakeName            string `json:"fake_name" xorm:"not null default '' comment('假名字') VARCHAR(128)"`
	Linkman             string `json:"linkman" xorm:"not null default '' VARCHAR(255)"`
	LinkmanPosition     string `json:"linkman_position" xorm:"not null default '' VARCHAR(255)"`
	LinkmanPhone        string `json:"linkman_phone" xorm:"not null default '' VARCHAR(255)"`
	Email               string `json:"email" xorm:"not null default '' VARCHAR(255)"`
	CityId              int    `json:"city_id" xorm:"not null default 0 INT(10)"`
	DistrictId          int    `json:"district_id" xorm:"not null default 0 INT(10)"`
	IndustryPath        string `json:"industry_path" xorm:"not null default '' VARCHAR(128)"`
	IndustryId          int    `json:"industry_id" xorm:"not null default 0 INT(10)"`
	CompanyscaleId      int    `json:"companyscale_id" xorm:"not null default 0 INT(10)"`
	CompanystageId      int    `json:"companystage_id" xorm:"not null default 0 INT(10)"`
	Companysynopsis     string `json:"companysynopsis" xorm:"not null TEXT"`
	Icon                string `json:"icon" xorm:"not null VARCHAR(255)"`
	SimpleName          string `json:"simple_name" xorm:"not null default '' VARCHAR(255)"`
	EditCompanysynopsis string `json:"edit_companysynopsis" xorm:"not null default '' VARCHAR(4096)"`
	Address             string `json:"address" xorm:"not null default '' VARCHAR(255)"`
	Website             string `json:"website" xorm:"not null default '' VARCHAR(255)"`
	Summary             string `json:"summary" xorm:"not null default '' comment('用于微信分享的摘要') VARCHAR(1024)"`
	IsTop               int    `json:"is_top" xorm:"not null default 0 TINYINT(3)"`
	Tags                string `json:"tags" xorm:"not null default '' comment('标签') VARCHAR(1024)"`
	IsOpen              int    `json:"is_open" xorm:"not null default 0 TINYINT(3)"`
	Articles            string `json:"articles" xorm:"not null default '' VARCHAR(2048)"`
	Sort                int    `json:"sort" xorm:"not null default 0 INT(10)"`
}

type Connection struct {
	Id                 int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	PassageId          int    `json:"passage_id" xorm:"not null unique(passage_id&user_id) INT(11)"`
	UserId             int    `json:"user_id" xorm:"not null unique(passage_id&user_id) index INT(11)"`
	FatherId           int    `json:"father_id" xorm:"not null default 0 index INT(11)"`
	RootId             int    `json:"root_id" xorm:"not null INT(11)"`
	Time               string `json:"time" xorm:"VARCHAR(255)"`
	Isrelay            int    `json:"isrelay" xorm:"not null default 0 index TINYINT(4)"`
	CountRelay         int    `json:"count_relay" xorm:"not null default 0 INT(11)"`
	CountApply         int    `json:"count_apply" xorm:"not null default 0 INT(11)"`
	CountApplyNoDirect int    `json:"count_apply_no_direct" xorm:"INT(11)"`
	NowRelayCount      int    `json:"now_relay_count" xorm:"not null default 0 INT(11)"`
}

type Cooperation struct {
	Id               int    `json:"id" xorm:"not null pk INT(10)"`
	RootCompanyId    int    `json:"root_company_id" xorm:"not null default 0 index INT(10)"`
	PassageCompanyId int    `json:"passage_company_id" xorm:"not null default 0 comment('公司id') index INT(10)"`
	Type             string `json:"type" xorm:"not null default 'cooperation' VARCHAR(32)"`
	Title            string `json:"title" xorm:"not null default '' index VARCHAR(300)"`
	Digest           string `json:"digest" xorm:"not null default '' VARCHAR(300)"`
	ImageUrl         string `json:"image_url" xorm:"not null default '' VARCHAR(1024)"`
	TimeBegin        int    `json:"time_begin" xorm:"not null default 0 INT(10)"`
	TimeEnd          int    `json:"time_end" xorm:"not null default 0 INT(10)"`
	PeopleDemand     string `json:"people_demand" xorm:"not null default '' VARCHAR(100)"`
	Cost             string `json:"cost" xorm:"not null default '' VARCHAR(100)"`
	Address          string `json:"address" xorm:"not null default '' VARCHAR(200)"`
	Content          string `json:"content" xorm:"not null MEDIUMTEXT"`
	Requirement      string `json:"requirement" xorm:"not null comment('参与条件') MEDIUMTEXT"`
	Attention        string `json:"attention" xorm:"not null comment('注意事项') MEDIUMTEXT"`
	AdminId          int    `json:"admin_id" xorm:"not null default 0 INT(11)"`
	CompanyId        int    `json:"company_id" xorm:"not null default 0 INT(11)"`
	City             string `json:"city" xorm:"not null default '' VARCHAR(100)"`
	MeetingCategory  string `json:"meeting_category" xorm:"not null default '' VARCHAR(200)"`
	PositionTag      string `json:"position_tag" xorm:"not null default '' VARCHAR(200)"`
	Experience       string `json:"experience" xorm:"not null default '' VARCHAR(100)"`
	Salary           string `json:"salary" xorm:"not null default '' VARCHAR(100)"`
	PassageIds       string `json:"passage_ids" xorm:"not null default '' comment('关联的职位id列表') VARCHAR(128)"`
	AccountId        int    `json:"account_id" xorm:"not null default 0 comment('文章account_id') INT(10)"`
	SourceUrl        string `json:"source_url" xorm:"not null default '' comment('原文地址') VARCHAR(255)"`
	IsTop            int    `json:"is_top" xorm:"not null default 0 comment('是否置顶') TINYINT(4)"`
	Pulisher         string `json:"pulisher" xorm:"not null default '' comment('发布者') VARCHAR(32)"`
	ViewCount        int    `json:"view_count" xorm:"not null default 0 comment('阅读数量') INT(10)"`
	LikeCount        int    `json:"like_count" xorm:"not null default 0 comment('点赞量') INT(10)"`
	TimeInsert       int    `json:"time_insert" xorm:"not null default 0 INT(10)"`
	TimeUpdate       int    `json:"time_update" xorm:"not null default 0 INT(10)"`
	Status           int    `json:"status" xorm:"not null default 0 TINYINT(3)"`
}

type DataCity struct {
	Id         int    `json:"id" xorm:"not null pk default 0 INT(10)"`
	CityIndex  int    `json:"city_index" xorm:"INT(11)"`
	ProvinceId int    `json:"province_id" xorm:"INT(11)"`
	Sort       int    `json:"sort" xorm:"default 0 comment('排序') INT(11)"`
	Name       string `json:"name" xorm:"VARCHAR(255)"`
}

type DataCityOld struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	CityIndex  int    `json:"city_index" xorm:"not null INT(11)"`
	ProvinceId int    `json:"province_id" xorm:"not null INT(11)"`
	Name       string `json:"name" xorm:"not null default '' VARCHAR(100)"`
}

type DataCityOld2 struct {
	Id         int    `json:"id" xorm:"not null pk default 0 INT(11)"`
	CityIndex  int    `json:"city_index" xorm:"INT(11)"`
	ProvinceId int    `json:"province_id" xorm:"INT(11)"`
	Sort       int    `json:"sort" xorm:"default 0 comment('排序') INT(11)"`
	Name       string `json:"name" xorm:"VARCHAR(255)"`
}

type DataDictionary struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Fid    int    `json:"fid" xorm:"not null unique(uiq_fid_eid) INT(10)"`
	Eid    int    `json:"eid" xorm:"not null default 0 comment('枚举值') unique(uiq_fid_eid) INT(10)"`
	Sort   int    `json:"sort" xorm:"not null default 0 INT(11)"`
	Remark string `json:"remark" xorm:"not null default '' VARCHAR(255)"`
}

type DataDictionaryConfig struct {
	Id   int    `json:"id" xorm:"not null pk autoincr comment('主键') INT(10)"`
	Fid  int    `json:"fid" xorm:"not null comment('类型') unique INT(10)"`
	Name string `json:"name" xorm:"not null comment('类型名称') VARCHAR(45)"`
	Key  string `json:"key" xorm:"not null unique VARCHAR(45)"`
}

type DataDistrict struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	CityId int    `json:"city_id" xorm:"not null INT(10)"`
	Sort   int    `json:"sort" xorm:"not null default 0 INT(11)"`
	Name   string `json:"name" xorm:"not null default '' VARCHAR(45)"`
}

type DataIndustry struct {
	Id    int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Pid   int    `json:"pid" xorm:"not null default 0 INT(10)"`
	Level int    `json:"level" xorm:"not null default 1 TINYINT(3)"`
	Name  string `json:"name" xorm:"not null default '' VARCHAR(20)"`
	Path  string `json:"path" xorm:"not null default '' VARCHAR(128)"`
	Sort  int    `json:"sort" xorm:"not null default 0 INT(10)"`
}

type DataPositionTag struct {
	Id    int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Pid   int    `json:"pid" xorm:"not null default 0 INT(10)"`
	Level int    `json:"level" xorm:"not null default 0 TINYINT(3)"`
	Name  string `json:"name" xorm:"not null default '' VARCHAR(20)"`
	Path  string `json:"path" xorm:"not null default '' VARCHAR(128)"`
	Sort  int    `json:"sort" xorm:"not null default 0 INT(10)"`
}

type DataProvince struct {
	Id   int    `json:"id" xorm:"not null pk INT(10)"`
	Name string `json:"name" xorm:"VARCHAR(50)"`
}

type Deliver struct {
	Id                       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Type                     int    `json:"type" xorm:"not null default 0 comment('投递类型: 1. 自己投递，2. 手动推荐') TINYINT(3)"`
	AccountId                int    `json:"account_id" xorm:"not null index(idx_account_id_passage_id) INT(10)"`
	AccountName              string `json:"account_name" xorm:"not null default '' comment('账号名称') VARCHAR(20)"`
	AccountMobile            string `json:"account_mobile" xorm:"not null default '' comment('账号手机号') VARCHAR(11)"`
	Name                     string `json:"name" xorm:"not null default '' comment('候选人姓名') VARCHAR(20)"`
	Mobile                   string `json:"mobile" xorm:"not null default '' comment('候选人手机号') index(idx_passage_id_mobile) VARCHAR(11)"`
	Email                    string `json:"email" xorm:"not null default '' comment('候选人邮箱') VARCHAR(64)"`
	PassageId                int    `json:"passage_id" xorm:"not null index(idx_account_id_passage_id) index(idx_passage_id_mobile) index(idx_recommend_account_id_passage_id) INT(10)"`
	PassageRecommendId       int    `json:"passage_recommend_id" xorm:"not null default 0 index INT(10)"`
	PassageRecommendPath     string `json:"passage_recommend_path" xorm:"not null default '0' comment('职位推荐路径') VARCHAR(1024)"`
	PassageRecommendPathFull string `json:"passage_recommend_path_full" xorm:"not null default '' comment('职位推荐路径，带上自己') VARCHAR(1024)"`
	RecommendAccountId       int    `json:"recommend_account_id" xorm:"not null default 0 index(idx_recommend_account_id_passage_id) INT(10)"`
	RecommendComment         string `json:"recommend_comment" xorm:"not null default '' comment('推荐评价') VARCHAR(2000)"`
	ResumeUrl                string `json:"resume_url" xorm:"not null default '' comment('简历地址') VARCHAR(255)"`
	ProgressEid              int    `json:"progress_eid" xorm:"not null default 1 comment('投递状态') INT(10)"`
	Progress                 string `json:"progress" xorm:"not null default '' comment('投递状态') VARCHAR(500)"`
	IsReal                   int    `json:"is_real" xorm:"not null default 0 comment('是否已经投递') TINYINT(3)"`
	DeliverTime              int64  `json:"deliver_time" xorm:"not null default 0 comment('投递时间') BIGINT(20)"`
	CreatedTime              int64  `json:"created_time" xorm:"not null default 0 comment('创建时间') BIGINT(20)"`
	UpdatedTime              int64  `json:"updated_time" xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
}

type Employ struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name         string `json:"name" xorm:"VARCHAR(255)"`
	Phone        string `json:"phone" xorm:"VARCHAR(255)"`
	Email        string `json:"email" xorm:"VARCHAR(255)"`
	Campany      string `json:"campany" xorm:"VARCHAR(255)"`
	ConnectionId int    `json:"connection_id" xorm:"INT(11)"`
	PassageId    int    `json:"passage_id" xorm:"INT(11)"`
	Need         string `json:"need" xorm:"VARCHAR(255)"`
}

type FinancingDemand struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Type            string `json:"type" xorm:"not null default 'financing_demand' VARCHAR(32)"`
	Title           string `json:"title" xorm:"not null default '' index VARCHAR(300)"`
	Digest          string `json:"digest" xorm:"not null default '' VARCHAR(300)"`
	ImageUrl        string `json:"image_url" xorm:"not null default '' VARCHAR(1024)"`
	TimeBegin       int    `json:"time_begin" xorm:"not null default 0 INT(10)"`
	TimeEnd         int    `json:"time_end" xorm:"not null default 0 INT(10)"`
	PeopleDemand    string `json:"people_demand" xorm:"not null default '' VARCHAR(100)"`
	Cost            string `json:"cost" xorm:"not null default '' VARCHAR(100)"`
	Address         string `json:"address" xorm:"not null default '' VARCHAR(200)"`
	Content         string `json:"content" xorm:"not null MEDIUMTEXT"`
	Requirement     string `json:"requirement" xorm:"not null comment('参与条件') MEDIUMTEXT"`
	Attention       string `json:"attention" xorm:"not null comment('注意事项') MEDIUMTEXT"`
	AdminId         int    `json:"admin_id" xorm:"not null default 0 INT(11)"`
	CompanyId       int    `json:"company_id" xorm:"not null default 0 INT(11)"`
	City            string `json:"city" xorm:"not null default '' VARCHAR(100)"`
	MeetingCategory string `json:"meeting_category" xorm:"not null default '' VARCHAR(200)"`
	PositionTag     string `json:"position_tag" xorm:"not null default '' VARCHAR(200)"`
	Experience      string `json:"experience" xorm:"not null default '' VARCHAR(100)"`
	Salary          string `json:"salary" xorm:"not null default '' VARCHAR(100)"`
	PassageIds      string `json:"passage_ids" xorm:"not null default '' comment('关联的职位id列表') VARCHAR(128)"`
	AccountId       int    `json:"account_id" xorm:"not null default 0 comment('文章account_id') INT(10)"`
	SourceUrl       string `json:"source_url" xorm:"not null default '' comment('原文地址') VARCHAR(255)"`
	IsTop           int    `json:"is_top" xorm:"not null default 0 comment('是否置顶') TINYINT(4)"`
	Pulisher        string `json:"pulisher" xorm:"not null default '' comment('发布者') VARCHAR(32)"`
	ViewCount       int    `json:"view_count" xorm:"not null default 0 comment('阅读数量') INT(11)"`
	LikeCount       int    `json:"like_count" xorm:"not null default 0 comment('点赞量') INT(10)"`
	TimeInsert      int    `json:"time_insert" xorm:"not null default 0 INT(10)"`
	TimeUpdate      int    `json:"time_update" xorm:"not null default 0 INT(10)"`
	Status          int    `json:"status" xorm:"not null default 0 TINYINT(4)"`
}

type HomeUser struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Unionid       string `json:"unionid" xorm:"unique VARCHAR(255)"`
	Password      string `json:"password" xorm:"VARCHAR(255)"`
	Phone         int64  `json:"phone" xorm:"index BIGINT(15)"`
	Nickname      string `json:"nickname" xorm:"VARCHAR(255)"`
	Realname      string `json:"realname" xorm:"not null default '' comment('真实姓名') VARCHAR(64)"`
	Sex           string `json:"sex" xorm:"VARCHAR(255)"`
	Headimgurl    string `json:"headimgurl" xorm:"VARCHAR(255)"`
	Myreward      int    `json:"myreward" xorm:"default 0 INT(11)"`
	Email         string `json:"email" xorm:"VARCHAR(255)"`
	Birthday      string `json:"birthday" xorm:"comment('出生年月日') VARCHAR(255)"`
	UserId        int    `json:"user_id" xorm:"index INT(11)"`
	Company       string `json:"company" xorm:"not null default '' comment('当前公司') VARCHAR(128)"`
	Position      string `json:"position" xorm:"not null default '' comment('当前职位') VARCHAR(64)"`
	WechatId      string `json:"wechat_id" xorm:"not null default '' comment('微信号') VARCHAR(128)"`
	SearchHistory string `json:"search_history" xorm:"not null default '' VARCHAR(10240)"`
}

type IndustryAssociation struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Type            string `json:"type" xorm:"not null default 'industry_association' VARCHAR(32)"`
	Title           string `json:"title" xorm:"not null default '' index VARCHAR(300)"`
	Digest          string `json:"digest" xorm:"not null default '' VARCHAR(300)"`
	ImageUrl        string `json:"image_url" xorm:"not null default '' VARCHAR(1024)"`
	TimeBegin       int    `json:"time_begin" xorm:"not null default 0 INT(10)"`
	TimeEnd         int    `json:"time_end" xorm:"not null default 0 INT(10)"`
	PeopleDemand    string `json:"people_demand" xorm:"not null default '' VARCHAR(100)"`
	Cost            string `json:"cost" xorm:"not null default '' VARCHAR(100)"`
	Address         string `json:"address" xorm:"not null default '' VARCHAR(200)"`
	Content         string `json:"content" xorm:"not null MEDIUMTEXT"`
	Requirement     string `json:"requirement" xorm:"not null comment('参与条件') MEDIUMTEXT"`
	Attention       string `json:"attention" xorm:"not null comment('注意事项') MEDIUMTEXT"`
	AdminId         int    `json:"admin_id" xorm:"not null default 0 INT(11)"`
	CompanyId       int    `json:"company_id" xorm:"not null default 0 INT(11)"`
	City            string `json:"city" xorm:"not null default '' VARCHAR(100)"`
	MeetingCategory string `json:"meeting_category" xorm:"not null default '' VARCHAR(200)"`
	PositionTag     string `json:"position_tag" xorm:"not null default '' VARCHAR(200)"`
	Experience      string `json:"experience" xorm:"not null default '' VARCHAR(100)"`
	Salary          string `json:"salary" xorm:"not null default '' VARCHAR(100)"`
	PassageIds      string `json:"passage_ids" xorm:"not null default '' comment('关联的职位id列表') VARCHAR(128)"`
	AccountId       int    `json:"account_id" xorm:"not null default 0 comment('文章account_id') INT(10)"`
	SourceUrl       string `json:"source_url" xorm:"not null default '' comment('原文地址') VARCHAR(255)"`
	IsTop           int    `json:"is_top" xorm:"not null default 0 comment('是否置顶') TINYINT(4)"`
	Pulisher        string `json:"pulisher" xorm:"not null default '' comment('发布者') VARCHAR(32)"`
	ViewCount       int    `json:"view_count" xorm:"not null default 0 comment('阅读数量') INT(11)"`
	LikeCount       int    `json:"like_count" xorm:"not null default 0 comment('点赞量') INT(10)"`
	TimeInsert      int    `json:"time_insert" xorm:"not null default 0 INT(10)"`
	TimeUpdate      int    `json:"time_update" xorm:"not null default 0 INT(10)"`
	Status          int    `json:"status" xorm:"not null default 0 TINYINT(4)"`
}

type IndustryInfo struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Type            string `json:"type" xorm:"not null default 'industry_info' VARCHAR(32)"`
	Title           string `json:"title" xorm:"not null default '' index VARCHAR(300)"`
	Digest          string `json:"digest" xorm:"not null default '' VARCHAR(300)"`
	ImageUrl        string `json:"image_url" xorm:"not null default '' VARCHAR(1024)"`
	TimeBegin       int    `json:"time_begin" xorm:"not null default 0 INT(10)"`
	TimeEnd         int    `json:"time_end" xorm:"not null default 0 INT(10)"`
	PeopleDemand    string `json:"people_demand" xorm:"not null default '' VARCHAR(100)"`
	Cost            string `json:"cost" xorm:"not null default '' VARCHAR(100)"`
	Address         string `json:"address" xorm:"not null default '' VARCHAR(200)"`
	Content         string `json:"content" xorm:"not null MEDIUMTEXT"`
	Requirement     string `json:"requirement" xorm:"not null comment('参与条件') MEDIUMTEXT"`
	Attention       string `json:"attention" xorm:"not null comment('注意事项') MEDIUMTEXT"`
	AdminId         int    `json:"admin_id" xorm:"not null default 0 INT(11)"`
	CompanyId       int    `json:"company_id" xorm:"not null default 0 INT(11)"`
	City            string `json:"city" xorm:"not null default '' VARCHAR(100)"`
	MeetingCategory string `json:"meeting_category" xorm:"not null default '' VARCHAR(200)"`
	PositionTag     string `json:"position_tag" xorm:"not null default '' VARCHAR(200)"`
	Experience      string `json:"experience" xorm:"not null default '' VARCHAR(100)"`
	Salary          string `json:"salary" xorm:"not null default '' VARCHAR(100)"`
	PassageIds      string `json:"passage_ids" xorm:"not null default '' comment('关联的职位id列表') VARCHAR(128)"`
	AccountId       int    `json:"account_id" xorm:"not null default 0 comment('文章account_id') INT(10)"`
	SourceUrl       string `json:"source_url" xorm:"not null default '' comment('原文地址') VARCHAR(255)"`
	IsTop           int    `json:"is_top" xorm:"not null default 0 comment('是否置顶') TINYINT(4)"`
	Pulisher        string `json:"pulisher" xorm:"not null default '' comment('发布者') VARCHAR(32)"`
	ViewCount       int    `json:"view_count" xorm:"not null default 0 comment('阅读数量') INT(11)"`
	LikeCount       int    `json:"like_count" xorm:"not null default 0 comment('点赞量') INT(10)"`
	TimeInsert      int    `json:"time_insert" xorm:"not null default 0 INT(10)"`
	TimeUpdate      int    `json:"time_update" xorm:"not null default 0 INT(10)"`
	Status          int    `json:"status" xorm:"not null default 0 TINYINT(4)"`
}

type InvestmentDemand struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Type            string `json:"type" xorm:"not null default 'investment_demand' VARCHAR(32)"`
	Title           string `json:"title" xorm:"not null default '' index VARCHAR(300)"`
	Digest          string `json:"digest" xorm:"not null default '' VARCHAR(300)"`
	ImageUrl        string `json:"image_url" xorm:"not null default '' VARCHAR(1024)"`
	TimeBegin       int    `json:"time_begin" xorm:"not null default 0 INT(10)"`
	TimeEnd         int    `json:"time_end" xorm:"not null default 0 INT(10)"`
	PeopleDemand    string `json:"people_demand" xorm:"not null default '' VARCHAR(100)"`
	Cost            string `json:"cost" xorm:"not null default '' VARCHAR(100)"`
	Address         string `json:"address" xorm:"not null default '' VARCHAR(200)"`
	Content         string `json:"content" xorm:"not null MEDIUMTEXT"`
	Requirement     string `json:"requirement" xorm:"not null comment('参与条件') MEDIUMTEXT"`
	Attention       string `json:"attention" xorm:"not null comment('注意事项') MEDIUMTEXT"`
	AdminId         int    `json:"admin_id" xorm:"not null default 0 INT(11)"`
	CompanyId       int    `json:"company_id" xorm:"not null default 0 INT(11)"`
	City            string `json:"city" xorm:"not null default '' VARCHAR(100)"`
	MeetingCategory string `json:"meeting_category" xorm:"not null default '' VARCHAR(200)"`
	PositionTag     string `json:"position_tag" xorm:"not null default '' VARCHAR(200)"`
	Experience      string `json:"experience" xorm:"not null default '' VARCHAR(100)"`
	Salary          string `json:"salary" xorm:"not null default '' VARCHAR(100)"`
	PassageIds      string `json:"passage_ids" xorm:"not null default '' comment('关联的职位id列表') VARCHAR(128)"`
	AccountId       int    `json:"account_id" xorm:"not null default 0 comment('文章account_id') INT(10)"`
	SourceUrl       string `json:"source_url" xorm:"not null default '' comment('原文地址') VARCHAR(255)"`
	IsTop           int    `json:"is_top" xorm:"not null default 0 comment('是否置顶') TINYINT(4)"`
	Pulisher        string `json:"pulisher" xorm:"not null default '' comment('发布者') VARCHAR(32)"`
	ViewCount       int    `json:"view_count" xorm:"not null default 0 comment('阅读数量') INT(11)"`
	LikeCount       int    `json:"like_count" xorm:"not null default 0 comment('点赞量') INT(10)"`
	TimeInsert      int    `json:"time_insert" xorm:"not null default 0 INT(10)"`
	TimeUpdate      int    `json:"time_update" xorm:"not null default 0 INT(10)"`
	Status          int    `json:"status" xorm:"not null default 0 TINYINT(3)"`
}

type Meeting struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Type            string `json:"type" xorm:"not null default 'meeting' VARCHAR(32)"`
	Title           string `json:"title" xorm:"not null default '' index VARCHAR(300)"`
	Digest          string `json:"digest" xorm:"not null default '' VARCHAR(300)"`
	ImageUrl        string `json:"image_url" xorm:"not null default '' VARCHAR(1024)"`
	TimeBegin       int    `json:"time_begin" xorm:"not null default 0 INT(10)"`
	TimeEnd         int    `json:"time_end" xorm:"not null default 0 INT(10)"`
	PeopleDemand    string `json:"people_demand" xorm:"not null default '' VARCHAR(100)"`
	Cost            string `json:"cost" xorm:"not null default '' VARCHAR(100)"`
	Address         string `json:"address" xorm:"not null default '' VARCHAR(200)"`
	Content         string `json:"content" xorm:"not null MEDIUMTEXT"`
	Requirement     string `json:"requirement" xorm:"not null comment('参与条件') MEDIUMTEXT"`
	Attention       string `json:"attention" xorm:"not null comment('注意事项') MEDIUMTEXT"`
	AdminId         int    `json:"admin_id" xorm:"not null default 0 INT(11)"`
	CompanyId       int    `json:"company_id" xorm:"not null default 0 INT(11)"`
	City            string `json:"city" xorm:"not null default '' VARCHAR(100)"`
	MeetingCategory string `json:"meeting_category" xorm:"not null default '' VARCHAR(200)"`
	PositionTag     string `json:"position_tag" xorm:"not null default '' VARCHAR(200)"`
	Experience      string `json:"experience" xorm:"not null default '' VARCHAR(100)"`
	Salary          string `json:"salary" xorm:"not null default '' VARCHAR(100)"`
	PassageIds      string `json:"passage_ids" xorm:"not null default '' comment('关联的职位id列表') VARCHAR(128)"`
	AccountId       int    `json:"account_id" xorm:"not null default 0 comment('文章account_id') INT(10)"`
	SourceUrl       string `json:"source_url" xorm:"not null default '' comment('原文地址') VARCHAR(255)"`
	IsTop           int    `json:"is_top" xorm:"not null default 0 comment('是否置顶') TINYINT(4)"`
	Pulisher        string `json:"pulisher" xorm:"not null default '' comment('发布者') VARCHAR(32)"`
	ViewCount       int    `json:"view_count" xorm:"not null default 0 comment('阅读数量') INT(11)"`
	LikeCount       int    `json:"like_count" xorm:"not null default 0 comment('点赞量') INT(10)"`
	TimeInsert      int    `json:"time_insert" xorm:"not null default 0 INT(10)"`
	TimeUpdate      int    `json:"time_update" xorm:"not null default 0 INT(10)"`
	Status          int    `json:"status" xorm:"not null default 0 TINYINT(3)"`
}

type MeetingJoin struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId     int    `json:"user_id" xorm:"not null default 0 INT(11)"`
	MeetingId  int    `json:"meeting_id" xorm:"not null default 0 INT(11)"`
	Name       string `json:"name" xorm:"not null default '' VARCHAR(64)"`
	Company    string `json:"company" xorm:"not null default '' VARCHAR(128)"`
	Position   string `json:"position" xorm:"not null default '' VARCHAR(64)"`
	WechatId   string `json:"wechat_id" xorm:"not null default '' comment('微信号') VARCHAR(128)"`
	Birthday   string `json:"birthday" xorm:"not null default '' VARCHAR(64)"`
	Message    string `json:"message" xorm:"not null TEXT"`
	TimeInsert int    `json:"time_insert" xorm:"not null default 0 INT(10)"`
	Status     int    `json:"status" xorm:"not null default 0 TINYINT(3)"`
}

type Passage struct {
	Id                      int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Content                 string `json:"content" xorm:"not null TEXT"`
	Title                   string `json:"title" xorm:"not null default '' VARCHAR(255)"`
	Time                    string `json:"time" xorm:"not null default '' VARCHAR(255)"`
	Summary                 string `json:"summary" xorm:"not null default '' VARCHAR(255)"`
	Icon                    string `json:"icon" xorm:"not null default '' VARCHAR(255)"`
	AdminId                 int    `json:"admin_id" xorm:"not null default 0 index INT(10)"`
	Type                    int    `json:"type" xorm:"not null default 0 comment('0=招聘、1=求职') TINYINT(4)"`
	CityId                  int    `json:"city_id" xorm:"not null default 0 index INT(10)"`
	DistrictId              int    `json:"district_id" xorm:"not null default 0 INT(10)"`
	SalaryId                int    `json:"salary_id" xorm:"not null index INT(10)"`
	ExperienceId            int    `json:"experience_id" xorm:"not null index INT(10)"`
	EducationId             int    `json:"education_id" xorm:"not null index INT(10)"`
	IndustryId              int    `json:"industry_id" xorm:"not null default 0 index INT(10)"`
	IndustryPath            string `json:"industry_path" xorm:"not null default '' VARCHAR(128)"`
	CompanyscaleId          int    `json:"companyscale_id" xorm:"not null index INT(10)"`
	CompanystageId          int    `json:"companystage_id" xorm:"not null index INT(10)"`
	CompanyAddressId        int    `json:"company_address_id" xorm:"not null default 0 comment('公司地址Id（支持多地址，为0时取默认地址）') INT(10)"`
	Companyname             string `json:"companyname" xorm:"not null index VARCHAR(255)"`
	Reward                  string `json:"reward" xorm:"not null default '' VARCHAR(255)"`
	Issue                   int    `json:"issue" xorm:"not null default 0 TINYINT(4)"`
	PositionTagId           int    `json:"position_tag_id" xorm:"not null default 0 INT(10)"`
	PositionTagPath         string `json:"position_tag_path" xorm:"not null default '' VARCHAR(128)"`
	Liangdian               string `json:"liangdian" xorm:"not null default '' VARCHAR(255)"`
	EditContent             string `json:"edit_content" xorm:"not null TEXT"`
	PositionQualification   string `json:"position_qualification" xorm:"not null comment('任职资格') TEXT"`
	PositionResearch        string `json:"position_research" xorm:"not null comment('职位调研') TEXT"`
	Mtime                   string `json:"mtime" xorm:"not null default '' VARCHAR(255)"`
	Status                  int    `json:"status" xorm:"not null default 0 comment('0=正常、1=下架') TINYINT(3)"`
	SalaryMin               int    `json:"salary_min" xorm:"INT(10)"`
	SalaryMax               int    `json:"salary_max" xorm:"INT(10)"`
	HeadCount               int    `json:"head_count" xorm:"not null default 0 comment('招聘所需人数') INT(10)"`
	CompanyRemark           string `json:"company_remark" xorm:"not null default '' VARCHAR(255)"`
	IsTop                   int    `json:"is_top" xorm:"not null default 0 TINYINT(4)"`
	IsRefined               int    `json:"is_refined" xorm:"not null default 0 comment('是否精选') TINYINT(4)"`
	IsAnonymous             int    `json:"is_anonymous" xorm:"not null default 1 comment('是否匿名') TINYINT(4)"`
	SuccessReward           string `json:"success_reward" xorm:"not null default '' comment('推荐奖金') VARCHAR(255)"`
	SuccessReward2          string `json:"success_reward2" xorm:"not null default '' VARCHAR(255)"`
	SuccessRewardRemark     string `json:"success_reward_remark" xorm:"not null default '' VARCHAR(255)"`
	InterviewReward         string `json:"interview_reward" xorm:"not null default '' comment('自荐奖金') VARCHAR(255)"`
	InterviewReward2        string `json:"interview_reward2" xorm:"not null default '' VARCHAR(255)"`
	InterviewRewardRemark   string `json:"interview_reward_remark" xorm:"not null default '' VARCHAR(255)"`
	ScholarshipReward       string `json:"scholarship_reward" xorm:"not null default '' VARCHAR(255)"`
	ScholarshipRewardRemark string `json:"scholarship_reward_remark" xorm:"not null default '' VARCHAR(255)"`
	McAdmin                 int    `json:"mc_admin" xorm:"not null default 0 comment('mc管理员ID') INT(11)"`
	RecAdmin                string `json:"rec_admin" xorm:"not null default '' comment('分配顾问') VARCHAR(255)"`
	PsgCompany              int    `json:"psg_company" xorm:"not null default 0 comment('顾客公司') index INT(10)"`
	CompanyId               int    `json:"company_id" xorm:"not null default 0 INT(10)"`
	RecommendRule           string `json:"recommend_rule" xorm:"not null comment('推荐规则') TEXT"`
	Contact                 string `json:"contact" xorm:"not null default '' comment('联系方式') VARCHAR(128)"`
}

type PassageCollection struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	HomeUid      int    `json:"home_uid" xorm:"index index(home_uid_2) index(home_uid_3) INT(11)"`
	ConnectionId int    `json:"connection_id" xorm:"index(home_uid_2) INT(11)"`
	PassageId    int    `json:"passage_id" xorm:"index(home_uid_3) INT(11)"`
	Type         string `json:"type" xorm:"default 'zhao' comment('收藏类型，复用此表') VARCHAR(20)"`
}

type PassageCompany struct {
	Id           int    `json:"id" xorm:"not null pk autoincr comment('主键') INT(10)"`
	Icon         string `json:"icon" xorm:"not null default '' VARCHAR(255)"`
	Name         string `json:"name" xorm:"not null comment('公司名称') VARCHAR(255)"`
	OutName      string `json:"out_name" xorm:"not null comment('对外简称') VARCHAR(255)"`
	CityId       int    `json:"city_id" xorm:"not null default 0 INT(10)"`
	DistrictId   int    `json:"district_id" xorm:"not null default 0 INT(10)"`
	IndustryId   int    `json:"industry_id" xorm:"not null default 0 comment('行业ID') INT(10)"`
	IndustryPath string `json:"industry_path" xorm:"not null default '' VARCHAR(128)"`
	Scale        int    `json:"scale" xorm:"not null default 0 comment('规模') INT(10)"`
	Stage        int    `json:"stage" xorm:"not null default 0 comment('阶段') INT(10)"`
	Address      string `json:"address" xorm:"not null default '' comment('地址') VARCHAR(255)"`
	AddressId    int    `json:"address_id" xorm:"not null default 0 comment('地址ID') index INT(10)"`
	Remark       string `json:"remark" xorm:"not null comment('公司简介') TEXT"`
	CompanyId    int    `json:"company_id" xorm:"not null default 0 comment('管理员企业') INT(10)"`
	AdminId      int    `json:"admin_id" xorm:"not null default 0 comment('更新adminID') INT(10)"`
	ModTime      int    `json:"mod_time" xorm:"not null default 0 comment('修改时间') INT(10)"`
	IsTop        int    `json:"is_top" xorm:"not null default 0 TINYINT(3)"`
	Tags         string `json:"tags" xorm:"not null default '' comment('标签') VARCHAR(256)"`
	Sort         int    `json:"sort" xorm:"not null default 0 INT(10)"`
}

type PassageCompanyEvent struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	RootCompanyId    int       `json:"root_company_id" xorm:"not null default 0 index INT(10)"`
	PassageCompanyId int       `json:"passage_company_id" xorm:"not null index INT(10)"`
	EventDate        time.Time `json:"event_date" xorm:"not null comment('日期') DATE"`
	Title            string    `json:"title" xorm:"not null comment('事件标题') VARCHAR(128)"`
	Description      string    `json:"description" xorm:"not null comment('事件描述') VARCHAR(1024)"`
	CreatedTime      int64     `json:"created_time" xorm:"not null default 0 BIGINT(20)"`
	UpdatedTime      int64     `json:"updated_time" xorm:"not null default 0 BIGINT(20)"`
}

type PassageCompanyMember struct {
	Id               int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	RootCompanyId    int    `json:"root_company_id" xorm:"not null default 0 index INT(10)"`
	PassageCompanyId int    `json:"passage_company_id" xorm:"not null comment('公司id') index INT(10)"`
	Name             string `json:"name" xorm:"not null comment('姓名') VARCHAR(64)"`
	AvatarUrl        string `json:"avatar_url" xorm:"not null default '' comment('头像地址') VARCHAR(255)"`
	Description      string `json:"description" xorm:"not null default '' comment('描述') VARCHAR(4096)"`
	CreatedTime      int64  `json:"created_time" xorm:"not null default 0 comment('创建时间') BIGINT(20)"`
	UpdatedTime      int64  `json:"updated_time" xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
	Status           int    `json:"status" xorm:"not null default 0 comment('用户状态') TINYINT(3)"`
}

type PassageLike struct {
	Id        int `json:"id" xorm:"not null pk autoincr INT(10)"`
	AccountId int `json:"account_id" xorm:"not null unique(uiq_home_uid_passage_id) INT(10)"`
	PassageId int `json:"passage_id" xorm:"not null unique(uiq_home_uid_passage_id) INT(10)"`
}

type PassageProject struct {
	Id        int    `json:"id" xorm:"not null pk INT(10)"`
	PassageId int    `json:"passage_id" xorm:"not null comment('职位ID') unique(idx_psgRsm) INT(10)"`
	AdminId   int    `json:"admin_id" xorm:"not null comment('hr_id') INT(10)"`
	CompanyId int    `json:"company_id" xorm:"not null comment('公司ID') INT(10)"`
	ResumeId  int    `json:"resume_id" xorm:"not null default 0 comment('简历ID') unique(idx_psgRsm) INT(10)"`
	State     int    `json:"state" xorm:"not null default 1 comment('状态') TINYINT(3)"`
	Data      string `json:"data" xorm:"not null comment('状态json') TEXT"`
	AddTime   int    `json:"add_time" xorm:"not null default 0 comment('添加时间') INT(10)"`
	ModTime   int    `json:"mod_time" xorm:"not null default 0 comment('修改时间') INT(10)"`
}

type PassageRecommend struct {
	Id                       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	PassageId                int    `json:"passage_id" xorm:"not null unique(uiq_account_id_passage_id) INT(10)"`
	AccountId                int    `json:"account_id" xorm:"not null unique(uiq_account_id_passage_id) INT(10)"`
	Path                     string `json:"path" xorm:"not null default '0' comment('传播路径') index VARCHAR(1024)"`
	PathFull                 string `json:"path_full" xorm:"not null default '0' comment('传播路径(带自己)') VARCHAR(1024)"`
	ParentPassageRecommendId int    `json:"parent_passage_recommend_id" xorm:"not null default 0 INT(10)"`
	RecommendCount           int    `json:"recommend_count" xorm:"not null default 0 comment('推荐投递次数') INT(10)"`
	RecommendCountL2         int    `json:"recommend_count_l2" xorm:"not null default 0 comment('推荐投递次数二级') INT(10)"`
	ShareCount               int    `json:"share_count" xorm:"not null default 0 comment('分享数量') INT(10)"`
	ShareCountL2             int    `json:"share_count_l2" xorm:"not null default 0 comment('分享数量2级') INT(10)"`
	CreatedTime              int64  `json:"created_time" xorm:"not null default 0 comment('创建时间') BIGINT(20)"`
	UpdatedTime              int64  `json:"updated_time" xorm:"not null default 0 comment('更新时间') BIGINT(20)"`
}

type ProfService struct {
	Id            int    `json:"id" xorm:"not null pk INT(10)"`
	Type          string `json:"type" xorm:"not null default 'prof_service' VARCHAR(32)"`
	Title         string `json:"title" xorm:"not null default '' index VARCHAR(300)"`
	Name          string `json:"name" xorm:"not null default '' comment('姓名') VARCHAR(50)"`
	Digest        string `json:"digest" xorm:"not null default '' VARCHAR(300)"`
	RecommendText string `json:"recommend_text" xorm:"not null default '' comment('推荐语') VARCHAR(300)"`
	ImageUrl      string `json:"image_url" xorm:"not null default '' VARCHAR(1024)"`
	Gender        string `json:"gender" xorm:"not null default '' comment('性别') VARCHAR(10)"`
	AdminId       int    `json:"admin_id" xorm:"not null default 0 INT(10)"`
	CompanyId     int    `json:"company_id" xorm:"not null default 0 INT(10)"`
	City          string `json:"city" xorm:"not null default '' VARCHAR(100)"`
	Experience    int    `json:"experience" xorm:"not null default 0 INT(10)"`
	Education     string `json:"education" xorm:"not null default '' VARCHAR(50)"`
	PositionTag   string `json:"position_tag" xorm:"not null comment('职位标签') VARCHAR(50)"`
	WorkingStatus string `json:"working_status" xorm:"not null default '' VARCHAR(50)"`
	MaritalStatus string `json:"marital_status" xorm:"not null default '' VARCHAR(50)"`
	CurrentSalary string `json:"current_salary" xorm:"not null default '' VARCHAR(100)"`
	ExpectSalary  int    `json:"expect_salary" xorm:"not null default 0 INT(10)"`
	Introduction  string `json:"introduction" xorm:"not null comment('基本简介') MEDIUMTEXT"`
	Evaluation    string `json:"evaluation" xorm:"not null comment('猎头评估') MEDIUMTEXT"`
	TimeInsert    int    `json:"time_insert" xorm:"not null default 0 INT(10)"`
	TimeUpdate    int    `json:"time_update" xorm:"not null default 0 INT(10)"`
	Status        int    `json:"status" xorm:"not null default 0 TINYINT(3)"`
	Service       string `json:"service" xorm:"not null comment('服务项目') TEXT"`
	Pulisher      string `json:"pulisher" xorm:"not null default '' comment('发布者') VARCHAR(200)"`
}

type ProfServiceInvite struct {
	Id         int    `json:"id" xorm:"not null pk INT(10)"`
	UserId     int    `json:"user_id" xorm:"not null default 0 INT(10)"`
	WantedId   int    `json:"wanted_id" xorm:"not null default 0 INT(10)"`
	Company    string `json:"company" xorm:"not null default '' VARCHAR(128)"`
	Name       string `json:"name" xorm:"not null default '' VARCHAR(64)"`
	Position   string `json:"position" xorm:"not null default '' VARCHAR(64)"`
	Phone      string `json:"phone" xorm:"not null default '' VARCHAR(20)"`
	WechatId   string `json:"wechat_id" xorm:"not null default '' comment('微信号') VARCHAR(128)"`
	Message    string `json:"message" xorm:"not null comment('备注') TEXT"`
	TimeInsert int    `json:"time_insert" xorm:"not null default 0 INT(10)"`
	TimeUpdate int    `json:"time_update" xorm:"not null default 0 INT(10)"`
	Status     int    `json:"status" xorm:"not null default 0 TINYINT(4)"`
}

type Resume struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserId        int    `json:"user_id" xorm:"not null default 0 comment('前台用户Id') INT(11)"`
	AdminId       int    `json:"admin_id" xorm:"not null default 0 comment('后台上传者Id') index INT(11)"`
	CompanyId     int    `json:"company_id" xorm:"not null default 0 comment('公司Id') index INT(11)"`
	Gender        string `json:"gender" xorm:"not null default '' comment('性别') VARCHAR(10)"`
	Name          string `json:"name" xorm:"not null default '' comment('姓名') VARCHAR(50)"`
	Mobile        string `json:"mobile" xorm:"not null default '' comment('手机') VARCHAR(50)"`
	Email         string `json:"email" xorm:"not null default '' comment('邮箱') VARCHAR(200)"`
	CityId        int    `json:"city_id" xorm:"not null default 0 comment('城市') INT(11)"`
	RecentCompany string `json:"recent_company" xorm:"not null default '' comment('最近公司') VARCHAR(300)"`
	Position      string `json:"position" xorm:"not null default '' comment('职位') VARCHAR(200)"`
	BirthYear     string `json:"birth_year" xorm:"not null default '' comment('出生年份') VARCHAR(10)"`
	University    string `json:"university" xorm:"not null default '' comment('毕业院校') VARCHAR(200)"`
	Major         string `json:"major" xorm:"not null default '' comment('专业') VARCHAR(100)"`
	EducationId   int    `json:"education_id" xorm:"not null default 0 comment('学历') INT(11)"`
	Source        string `json:"source" xorm:"not null default '' comment('简历来源') VARCHAR(100)"`
	OriginText    string `json:"origin_text" xorm:"not null comment('简历原文') LONGTEXT"`
	FilePath      string `json:"file_path" xorm:"not null default '' comment('简历文件位置') VARCHAR(1024)"`
	TimeInsert    int    `json:"time_insert" xorm:"not null default 0 comment('添加时间') INT(10)"`
	TimeUpdate    int    `json:"time_update" xorm:"not null default 0 comment('编辑时间') INT(10)"`
	Status        int    `json:"status" xorm:"not null default 2 comment('0-删除，1-暂不使用，2-正常') TINYINT(4)"`
	Tags          string `json:"tags" xorm:"not null default '' comment('标签') VARCHAR(255)"`
	LastOp        string `json:"last_op" xorm:"not null default '' comment('更新人') VARCHAR(45)"`
}

type ResumeRemind struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	AdminId    int    `json:"admin_id" xorm:"not null default 0 comment('管理员ID') INT(11)"`
	ResumeId   int    `json:"resume_id" xorm:"not null comment('简历ID') INT(11)"`
	RemindTime string `json:"remind_time" xorm:"not null comment('yymmdd') VARCHAR(45)"`
	AddTime    int    `json:"add_time" xorm:"not null default 0 comment('添加时间') INT(11)"`
}

type ResumeVisitHistory struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	ResumeId   int    `json:"resume_id" xorm:"not null default 0 comment('简历Id') index INT(11)"`
	AdminId    int    `json:"admin_id" xorm:"not null default 0 comment('管理者Id') INT(11)"`
	Content    string `json:"content" xorm:"not null comment('寻访记录') LONGTEXT"`
	TimeInsert int    `json:"time_insert" xorm:"not null default 0 comment('添加时间') INT(10)"`
	TimeUpdate int    `json:"time_update" xorm:"not null default 0 comment('编辑时间') INT(10)"`
	Status     int    `json:"status" xorm:"not null default 2 comment('0-删除，1-暂不使用，2-正常') TINYINT(3)"`
}

type Tags struct {
	Id  int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Tag string `json:"tag" xorm:"not null comment('标签') VARCHAR(45)"`
}

type Test struct {
	Id   int `json:"id" xorm:"not null pk autoincr INT(11)"`
	Test int `json:"test" xorm:"INT(11)"`
}

type TopicQa struct {
	Id              int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Type            string `json:"type" xorm:"not null default 'topic_qa' VARCHAR(32)"`
	Title           string `json:"title" xorm:"not null default '' index VARCHAR(300)"`
	Digest          string `json:"digest" xorm:"not null default '' VARCHAR(300)"`
	ImageUrl        string `json:"image_url" xorm:"not null default '' VARCHAR(1024)"`
	TimeBegin       int    `json:"time_begin" xorm:"not null default 0 INT(10)"`
	TimeEnd         int    `json:"time_end" xorm:"not null default 0 INT(10)"`
	PeopleDemand    string `json:"people_demand" xorm:"not null default '' VARCHAR(100)"`
	Cost            string `json:"cost" xorm:"not null default '' VARCHAR(100)"`
	Address         string `json:"address" xorm:"not null default '' VARCHAR(200)"`
	Content         string `json:"content" xorm:"not null MEDIUMTEXT"`
	Requirement     string `json:"requirement" xorm:"not null comment('参与条件') MEDIUMTEXT"`
	Attention       string `json:"attention" xorm:"not null comment('注意事项') MEDIUMTEXT"`
	AdminId         int    `json:"admin_id" xorm:"not null default 0 INT(11)"`
	CompanyId       int    `json:"company_id" xorm:"not null default 0 INT(11)"`
	City            string `json:"city" xorm:"not null default '' VARCHAR(100)"`
	MeetingCategory string `json:"meeting_category" xorm:"not null default '' VARCHAR(200)"`
	PositionTag     string `json:"position_tag" xorm:"not null default '' VARCHAR(200)"`
	Experience      string `json:"experience" xorm:"not null default '' VARCHAR(100)"`
	Salary          string `json:"salary" xorm:"not null default '' VARCHAR(100)"`
	PassageIds      string `json:"passage_ids" xorm:"not null default '' comment('关联的职位id列表') VARCHAR(128)"`
	AccountId       int    `json:"account_id" xorm:"not null default 0 comment('文章account_id') INT(10)"`
	SourceUrl       string `json:"source_url" xorm:"not null default '' comment('原文地址') VARCHAR(255)"`
	IsTop           int    `json:"is_top" xorm:"not null default 0 comment('是否置顶') TINYINT(4)"`
	Pulisher        string `json:"pulisher" xorm:"not null default '' comment('发布者') VARCHAR(32)"`
	ViewCount       int    `json:"view_count" xorm:"not null default 0 comment('阅读数量') INT(11)"`
	LikeCount       int    `json:"like_count" xorm:"not null default 0 comment('点赞量') INT(10)"`
	TimeInsert      int    `json:"time_insert" xorm:"not null default 0 INT(10)"`
	TimeUpdate      int    `json:"time_update" xorm:"not null default 0 INT(10)"`
	Status          int    `json:"status" xorm:"not null default 0 TINYINT(3)"`
}

type User struct {
	Id      int    `json:"id" xorm:"not null pk autoincr comment('0=用户，1=管理员') INT(10)"`
	Email   string `json:"email" xorm:"index VARCHAR(255)"`
	AdminId int    `json:"admin_id" xorm:"default 0 comment('1=管理员，0=用户') index INT(11)"`
	Phone   int64  `json:"phone" xorm:"unique BIGINT(15)"`
}

type UserCenter struct {
	UserId           int `json:"user_id" xorm:"not null pk default 0 INT(11)"`
	PassageId        int `json:"passage_id" xorm:"not null pk default 0 INT(11)"`
	CountRelay       int `json:"count_relay" xorm:"not null default 0 INT(11)"`
	CountApply       int `json:"count_apply" xorm:"not null default 0 INT(11)"`
	CountSecondApply int `json:"count_second_apply" xorm:"not null default 0 INT(11)"`
}
