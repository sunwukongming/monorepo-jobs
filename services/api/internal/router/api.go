package router

import (
	"app/config"
	"app/data"
	"app/db"
	"app/internal/controllers/api/account"
	"app/internal/controllers/api/apply"
	"app/internal/controllers/api/article"
	"app/internal/controllers/api/banner"
	"app/internal/controllers/api/company"
	"app/internal/controllers/api/deliver"
	"app/internal/controllers/api/dictionary"
	"app/internal/controllers/api/help"
	"app/internal/controllers/api/passage"
	"app/internal/controllers/api/resume"
	"app/internal/controllers/api/root_company"
	"app/internal/controllers/api/user"
	"app/internal/controllers/api/wanted"
	"app/internal/controllers/api/wechat"
	"app/internal/middleware"
	"app/models/bolejiang"
	"app/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApi(routerGroup *gin.RouterGroup) {
	routerGroup.GET("cache/refresh", func(ctx *gin.Context) {
		err := data.Load()
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
		} else {
			ctx.String(http.StatusOK, "ok")
		}
	})
	if config.Get().Mode == "debug" {
		routerGroup.POST("account/login", account.LoginAction)
	}
	routerGroup.POST("account/bindWechatMobile", middleware.Auth(), account.BindWechatMobileAction)
	routerGroup.POST("account/loginWechat", account.LoginWechatAction)

	routerGroup.GET("banner/list", banner.ListAction)

	routerGroup.Any("user/info", middleware.Auth(), user.InfoAction)
	routerGroup.POST("user/wechatMobile", middleware.Auth(), account.BindWechatMobileAction)
	routerGroup.POST("user/update", middleware.Auth(), user.UpdateAction)
	routerGroup.POST("user/createApply", middleware.Auth(), user.CreateApplyAction)
	routerGroup.POST("user/updateApply", middleware.Auth(), user.UpdateApplyAction)
	routerGroup.POST("user/deleteApply", middleware.Auth(), user.DeleteApplyAction)
	routerGroup.POST("user/createEducation", middleware.Auth(), user.CreateEducationAction)
	routerGroup.POST("user/updateEducation", middleware.Auth(), user.UpdateEducationAction)
	routerGroup.POST("user/deleteEducation", middleware.Auth(), user.DeleteEducationAction)
	routerGroup.POST("user/createProject", middleware.Auth(), user.CreateProjectAction)
	routerGroup.POST("user/updateProject", middleware.Auth(), user.UpdateProjectAction)
	routerGroup.POST("user/deleteProject", middleware.Auth(), user.DeleteProjectAction)

	routerGroup.POST("user/createWork", middleware.Auth(), user.CreateWorkAction)
	routerGroup.POST("user/updateWork", middleware.Auth(), user.UpdateWorkAction)
	routerGroup.POST("user/deleteWork", middleware.Auth(), user.DeleteWorkAction)

	routerGroup.Any("dictionary/cities", dictionary.CitiesAction)
	routerGroup.Any("dictionary/industries", dictionary.IndustriesAction)
	routerGroup.Any("dictionary/positionTags", dictionary.PositionTagsAction)
	routerGroup.Any("dictionary/data", dictionary.DataAction)

	routerGroup.POST("passage/list", passage.ListAction)
	routerGroup.POST("passage/listLike", middleware.Auth(), passage.ListLikeAction)
	routerGroup.POST("passage/listRecommend", middleware.Auth(), deliver.ListAction)
	routerGroup.POST("passage/listSelfRecommend", middleware.Auth(), passage.ListSelfRecommendAction)
	routerGroup.POST("passage/get", middleware.Auth(), passage.GetAction)
	routerGroup.POST("passage/getOrigin", passage.GetOriginAction)
	routerGroup.POST("passage/like", middleware.Auth(), passage.LikeAction)
	routerGroup.POST("passage/unlike", middleware.Auth(), passage.UnlikeAction)
	routerGroup.POST("passage/deliver", middleware.Auth(), passage.DeliverAction)
	routerGroup.POST("passage/recommend", middleware.Auth(), passage.RecommendAction)
	routerGroup.POST("passage/delivers", middleware.Auth(), passage.DeliversAction)

	routerGroup.POST("company/get", company.GetAction)
	routerGroup.POST("company/list", company.ListAction)
	routerGroup.POST("company/listMember", company.ListMemberAction)
	routerGroup.POST("company/listEvent", company.ListEventAction)
	routerGroup.POST("company/listArticle", company.ListArticleAction)

	routerGroup.POST("rootCompany/get", root_company.GetAction)
	routerGroup.POST("rootCompany/list", root_company.ListAction)
	routerGroup.POST("rootCompany/listMember", root_company.ListMemberAction)
	routerGroup.POST("rootCompany/listEvent", root_company.ListEventAction)
	routerGroup.POST("rootCompany/listArticle", root_company.ListArticleAction)

	routerGroup.POST("deliver/list", middleware.Auth(), deliver.ListAction)
	routerGroup.POST("deliver/detail", middleware.Auth(), deliver.DetailAction)
	routerGroup.POST("deliver/createManual", middleware.Auth(), deliver.CreateManualAction)
	routerGroup.POST("deliver/updateManual", middleware.Auth(), deliver.UpdateManualAction)
	//routerGroup.POST("deliver/deleteManual", middleware.Auth(), deliver.DeleteManualAction)

	routerGroup.Any("wechat/barcode", wechat.BarcodeAction)
	routerGroup.Any("wechat/urlSchema", wechat.URLSchemaAction)
	routerGroup.Any("wechat/urlLink", wechat.UrlLinkAction)

	routerGroup.POST("apply/list", apply.ListAction)
	routerGroup.POST("apply/detail", apply.DetailAction)
	routerGroup.POST("apply/like", apply.LikeAction)
	routerGroup.POST("apply/unlike", apply.UnlikeAction)
	routerGroup.POST("apply/listLike", apply.ListLikeAction)
	routerGroup.POST("apply/resumeAuth", apply.ResumeAuthAction)

	for _, model := range []interface{}{bolejiang.Meeting{}, bolejiang.FinancingDemand{}, bolejiang.IndustryAssociation{}, bolejiang.IndustryInfo{}, bolejiang.InvestmentDemand{}, bolejiang.TopicQa{}, bolejiang.Cooperation{}} {
		tablename := db.Default().TableName(model)
		apiname := utils.CamelCase(tablename)
		routerGroup.POST("article/"+apiname+"s", article.GetArticlesAction(model))
		routerGroup.POST("article/"+apiname+"Detail", article.GetArticleDetailAction(model))
		routerGroup.POST("article/"+apiname+"Like", middleware.Auth(), article.GetArticleLikeAction(model))
		routerGroup.POST("article/"+apiname+"Unlike", middleware.Auth(), article.GetArticleUnlikeAction(model))
	}

	routerGroup.POST("resume/upload", resume.UploadAction)
	routerGroup.POST("wanted/get", wanted.GetAction)
	routerGroup.POST("wanted/list", wanted.ListAction)
	routerGroup.POST("wanted/like", wanted.LikeAction)
	routerGroup.POST("wanted/Unlike", wanted.UnlikeAction)
	routerGroup.POST("profService/get", wanted.GetAction)
	routerGroup.POST("profService/list", wanted.ListAction)
	routerGroup.POST("profService/lisLike", wanted.LikeAction)
	routerGroup.POST("profService/lisUnlike", wanted.UnlikeAction)

	routerGroup.POST("help/info", middleware.Auth(), help.InfoAction)
	routerGroup.POST("help/apply", middleware.Auth(), help.ApplyAction)
}
