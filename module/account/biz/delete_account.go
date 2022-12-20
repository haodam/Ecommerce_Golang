package accountbiz

import (
	"Ecommerce_Golang/common"
	accountmodel "Ecommerce_Golang/module/account/model"
	"context"
)

type DeleteAccountStore interface {
	Delete(ctx context.Context, id int) error
	FindAccountWithCondition(ctx context.Context, condition map[string]interface{},
		morekeys ...string) (*accountmodel.Account, error)
}

type deleteAccountBiz struct {
	store DeleteAccountStore
}

func NewDeleteAccountBiz(store DeleteAccountStore) *deleteAccountBiz {
	return &deleteAccountBiz{store: store}
}

func (biz *deleteAccountBiz) DeleteAccount(ctx context.Context, id int) error {

	oldData, err := biz.store.FindAccountWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(common.EntityName, err)
	}
	if oldData.Status == 0 {
		return common.ErrEntityDeleted(common.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
