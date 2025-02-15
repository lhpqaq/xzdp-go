package mysql

import (
	"context"
	"errors"
	"xzdp/biz/model/voucher"

	"gorm.io/gorm"
)

func QueryShopVoucherByShopID(ctx context.Context, id int64) ([]*voucher.SeckillVoucher, error) {
	var voucherList []*voucher.SeckillVoucher

	err = DB.WithContext(ctx).Raw("select v.*,s.stock,s.begin_time,s.end_time from tb_voucher v left join tb_seckill_voucher s on s.voucher_id = v.id where v.shop_id = ?", id).
		Find(&voucherList).Error
	if err != nil {
		return nil, err
	}

	return voucherList, nil
}

func QueryVoucherByID(ctx context.Context, id int64) (*voucher.SeckillVoucher, error) {
	var seckillVoucher voucher.SeckillVoucher
	err = DB.WithContext(ctx).Where("voucher_id = ?", id).Order("create_time desc").Limit(1).Find(&seckillVoucher).Error
	return &seckillVoucher, err
}

func QueryVoucherOrderByVoucherID(ctx context.Context, userId int64, id int64) error {
	var voucherOrder voucher.VoucherOrder
	err = DB.WithContext(ctx).Where("voucher_id = ? and user_id=?", id, userId).Limit(1).Find(&voucherOrder).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("重复下单")
		}
	} else if voucherOrder.ID != 0 {
		return errors.New("重复下单")
	}
	return nil
}

func UpdateVoucherStock(ctx context.Context, id int64) error {
	result := DB.WithContext(ctx).Model(&voucher.SeckillVoucher{}).
		Where("voucher_id = ? and stock > 0", id).
		Update("stock", gorm.Expr("stock - ?", 1))
	if result.RowsAffected == 0 {
		return errors.New("库存不足，扣减失败")
	}
	return result.Error
}

func CreateVoucherOrder(ctx context.Context, voucherOrder *voucher.VoucherOrder) error {
	return DB.WithContext(ctx).Create(voucherOrder).Error
}
