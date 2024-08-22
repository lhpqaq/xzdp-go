package mysql

import (
	"context"
	"xzdp/biz/model/voucher"
)

func QueryVoucherByID(ctx context.Context, id int64) ([]*voucher.Voucher, error) {
	var voucherList []*voucher.Voucher

	err = DB.WithContext(ctx).Where("shop_id = ?", id).Find(&voucherList).Error
	if err != nil {
		return nil, err
	}

	return voucherList, nil
}
