package ginaccount

import (
	"Ecommerce_Golang/component/appctx"
	accountbiz "Ecommerce_Golang/module/account/biz"
	accountmodel "Ecommerce_Golang/module/account/model"
	accountstorage "Ecommerce_Golang/module/account/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAccount(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMaiDBConnection()
		var data accountmodel.Account

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := accountstorage.NewSQLStore(db)
		biz := accountbiz.CreateAccountStore(store)

		if err := biz.Create(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
