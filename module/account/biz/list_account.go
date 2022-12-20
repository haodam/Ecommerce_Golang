package accountbiz

import (
	"Ecommerce_Golang/common"
	accountmodel "Ecommerce_Golang/module/account/model"
	"context"
)

type ListAccountStore interface {
	ListDataWithCondition(ctx context.Context, filter *accountmodel.Filter,
		paging *common.Paging, moreKeys ...string) ([]accountmodel.Account, error)
}

type listAccountBiz struct {
	store ListAccountStore
}

func NewListAccountBiz(store ListAccountStore) *listAccountBiz {
	return &listAccountBiz{store: store}
}

func (biz *listAccountBiz) ListAccount(ctx context.Context, filter *accountmodel.Filter, paging *common.Paging,
	moreKeys ...string) ([]accountmodel.Account, error) {

	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(common.EntityName, nil)
	}

	return result, nil
}
