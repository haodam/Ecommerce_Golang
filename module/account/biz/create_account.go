package accountbiz

import (
	accountmodel "Ecommerce_Golang/module/account/model"
	"context"
	"errors"
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

func (biz *createAccountBiz) CreateAccount(ctx context.Context, data *accountmodel.Account) error {
	if data.Name == "" {
		return errors.New("name cannot be empty")
	}
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
