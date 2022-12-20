package accountstorage

import (
	"Ecommerce_Golang/common"
	accountmodel "Ecommerce_Golang/module/account/model"
	"context"
)

func (s *SqlStore) ListDataWithCondition(ctx context.Context, filter *accountmodel.Filter,
	paging *common.Paging, moreKeys ...string) ([]accountmodel.Account, error) {

	var result []accountmodel.Account
	db := s.db.Table(accountmodel.Account{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.AccountId > 0 {
			db = db.Where("account_id = ?", f.AccountId)
		}
		//if len(f.Status) > 0 {
		//	db = db.Where("status in (?)", f.Status)
		//}
	}
	//fmt.Println("aaaa")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	if err := db.Offset(offset).Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//if err := s.db.Find(&result).Error; err != nil {
	//	return nil, err
	//}
	return result, nil
}
