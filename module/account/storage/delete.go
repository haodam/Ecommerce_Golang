package accountstorage

import (
	"context"
)

func (s *SqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table("accounts").
		Where("id = ?", id).Update("status", "0").Error; err != nil {
		return err
	}
	return nil
}
