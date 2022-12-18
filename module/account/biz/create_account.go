package accountbiz

import (
	"Ecommerce_Golang/common"
	accountmodel "Ecommerce_Golang/module/account/model"
	"context"
)

type CreateAccountStore interface {
	Create(ctx context.Context, data *accountmodel.Account) error
}

type createAccountBiz struct {
	store CreateAccountStore
}

func NewCreateAccountBiz(store CreateAccountStore) *createAccountBiz {
	return &createAccountBiz{store: store}
}

func (biz *createAccountBiz) CreateAccount(ctx context.Context, account *accountmodel.Account) error {

	accountId := account.AccountId
	err := common.ValidateString(accountId)
	if err != nil {
		return err
	}

	if err = biz.store.Create(ctx, account); err != nil {
		return err
	}
	return nil
}
