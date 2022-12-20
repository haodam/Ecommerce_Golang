package ginaccount

import (
	"Ecommerce_Golang/common"
	"Ecommerce_Golang/component/appctx"
	accountbiz "Ecommerce_Golang/module/account/biz"
	accountstorage "Ecommerce_Golang/module/account/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindAccount(ctx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := ctx.GetMaiDBConnection()
		paramStr := c.Query("id")
		id, err := strconv.Atoi(paramStr)
		if err != nil {
			fmt.Println(common.ErrInvalidRequest(err))
			return
		}

		if err := c.ShouldBind(&id); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := accountstorage.NewSQLStore(db)
		biz := accountbiz.NewFindAccountBiz(store)

		result, err := biz.FindAccount(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}
