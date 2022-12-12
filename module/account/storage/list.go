package accountstorage

import (
	"Ecommerce_Golang/common"
	accountmodel "Ecommerce_Golang/module/account/model"
	"context"
)

func (s *SqlStore) ListAccountWithCondition(ctx context.Context, filter *accountmodel.Filter, paging *common.Paging,
	morekeys ...string) ([]accountmodel.Account, error) {

	var result []accountmodel.Account
	db := s.db.Where("status in (1)")

	if f := filter; f != nil {
		if f.AccountId > 0 {
			db = db.Where("account_id = ?", f.AccountId)
		}
	}

	if err := s.db.Where(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
