package accountbiz

import (
	"context"
)

type DeleteAccountStore interface {
	Delete(ctx context.Context, id int) error
}

type deleteAccountBiz struct {
	store DeleteAccountStore
}

func NewDeleteAccount(store DeleteAccountStore) *deleteAccountBiz {
	return &deleteAccountBiz{store: store}
}

func (biz *deleteAccountBiz) DeleteAccount(ctx context.Context, id int) error {

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
