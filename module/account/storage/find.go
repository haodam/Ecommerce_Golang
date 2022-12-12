package accountstorage

import (
	accountmodel "Ecommerce_Golang/module/account/model"
	"context"
)

func (s *SqlStore) FindAccountWithCondition(ctx context.Context, condition map[string]interface{},
	morekeys ...string) (*accountmodel.Account, error) {

	var data accountmodel.Account
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
