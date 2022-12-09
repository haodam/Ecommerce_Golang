package ginaccount

import "C"
import (
	"Ecommerce_Golang/component/appctx"
	accountbiz "Ecommerce_Golang/module/account/biz"
	accountstorage "Ecommerce_Golang/module/account/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteAccount(ctx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		DB := ctx.GetMaiDBConnection()
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := accountstorage.NewSQLStore(DB)
		biz := accountbiz.NewDeleteAccount(store)

		if err := biz.DeleteAccount(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": 1,
		})
	}
}
