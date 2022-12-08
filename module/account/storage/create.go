package accountstorage

import (
	accountmodel "Ecommerce_Golang/module/account/model"
	"context"
)

func (s *SqlStore) Create(ctx context.Context, data *accountmodel.Account) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
