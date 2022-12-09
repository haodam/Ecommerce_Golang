package accountstorage

import (
	accountmodel "Ecommerce_Golang/module/account/model"
	"context"
)

func (s *SqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(accountmodel.Account{}.TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}
	return nil
}
