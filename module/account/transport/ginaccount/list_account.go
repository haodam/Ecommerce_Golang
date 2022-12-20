package ginaccount

import (
	"Ecommerce_Golang/common"
	"Ecommerce_Golang/component/appctx"
	accountbiz "Ecommerce_Golang/module/account/biz"
	accountmodel "Ecommerce_Golang/module/account/model"
	accountstorage "Ecommerce_Golang/module/account/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListAccount(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//pagingData.Fulfill()

		var filter accountmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//filter.Status = []int{1}
		//var result []accountmodel.Account

		store := accountstorage.NewSQLStore(db)
		biz := accountbiz.NewListAccountBiz(store)

		result, err := biz.ListAccount(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
