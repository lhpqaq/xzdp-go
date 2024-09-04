package voucher

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	voucherModel "xzdp/biz/model/voucher"
	"xzdp/biz/utils"
)

type SeckillVoucherService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSeckillVoucherService(Context context.Context, RequestContext *app.RequestContext) *SeckillVoucherService {
	return &SeckillVoucherService{RequestContext: RequestContext, Context: Context}
}

func (h *SeckillVoucherService) Run(req *int64) (resp *int64, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	//0.查询秒杀券
	voucher, err := mysql.QueryVoucherByID(h.Context, *req)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	//1.判断是否开始&&结束
	layout := "2006-01-02T15:04:05+08:00"
	beginTime, _ := time.Parse(layout, voucher.GetBeginTime())
	endTime, _ := time.Parse(layout, voucher.GetEndTime())
	if beginTime.After(now) {
		return nil, errors.New("秒杀还未开始")
	}
	if endTime.Before(now) {
		return nil, errors.New("秒杀已经结束")
	}
	//2.判断库存是否充足
	if voucher.GetStock() <= 0 {
		return nil, errors.New("已抢空")
	}
	//3.判断是否已经购买
	userId := utils.GetUser(h.Context).GetID()
	order, err := mysql.QueryVoucherOrderByVoucherID(h.Context, userId, *req)
	if order != nil {
		return nil, err
	}
	//4.扣减库存
	err = mysql.UpdateVoucherStock(h.Context, *req)
	if err != nil {
		return nil, err
	}
	//5.创建订单
	orderId, err := redis.NextId(h.Context, "order")
	if err != nil {
		return nil, err
	}
	voucherOrder := &voucherModel.VoucherOrder{
		UserId:     userId,
		VoucherId:  *req,
		OrderId:    orderId,
		PayTime:    time.Now().Format(layout),
		UseTime:    "0000-00-00 00:00:00",
		RefundTime: "0000-00-00 00:00:00",
	}
	err = mysql.CreateVoucherOrder(h.Context, voucherOrder)
	if err != nil {
		return nil, err
	}
	return &orderId, nil
}
