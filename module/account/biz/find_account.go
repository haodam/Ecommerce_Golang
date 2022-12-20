package accountbiz

import (
	"Ecommerce_Golang/common"
	accountmodel "Ecommerce_Golang/module/account/model"
	"context"
)

type FindAccountStore interface {
	FindAccountWithCondition(ctx context.Context, condition map[string]interface{},
		moreKeys ...string) (*accountmodel.Account, error)
}

type findAccountBiz struct {
	store FindAccountStore
}

func NewFindAccountBiz(store FindAccountStore) *findAccountBiz {
	return &findAccountBiz{store: store}
}

func (biz *findAccountBiz) FindAccount(ctx context.Context, id int, moreKeys ...string) (*accountmodel.Account, error) {

	//var data accountmodel.Account
	result, err := biz.store.FindAccountWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(common.EntityName, nil)
	}

	return result, nil
}
