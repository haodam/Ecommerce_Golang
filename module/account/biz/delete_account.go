package accountbiz

import (
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

func NewDeleteAccount(store DeleteAccountStore) *deleteAccountBiz {
	return &deleteAccountBiz{store: store}
}

func (biz *deleteAccountBiz) DeleteAccount(ctx context.Context, id int) error {

	//var oldData *accountmodel.Account := biz.store.FindAccountWithCondition(ctx, map[string]interface{}{"id": id})
	//
	//if oldData.Status  == 0{
	//	return errors.New("data has been deleted")
	//}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
