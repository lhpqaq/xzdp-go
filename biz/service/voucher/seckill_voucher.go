package voucher

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"sync"
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

var mu sync.Mutex

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
	user := utils.GetUser(h.Context)
	uuid, _ := utils.RandomUUID()
	sec := time.Now().Unix()
	lockValue := uuid + strconv.FormatInt(sec, 10) //由于value的全局唯一性，这里用uuid+时间戳，如需要更高精度应考虑雪花算法活其他方法生成
	lock := redis.NewLock(h.Context, user.NickName, lockValue, 10)
	ok := lock.TryLock()
	if !ok {
		return nil, errors.New("重复下单")
	}
	defer lock.UnLock(lockValue)
	return h.createOrder(*req)
}

func (h *SeckillVoucherService) createOrder(voucherId int64) (resp *int64, err error) {
	//3.判断是否已经购买
	userId := utils.GetUser(h.Context).GetID()
	err = mysql.QueryVoucherOrderByVoucherID(h.Context, userId, voucherId)
	if err != nil {
		return nil, err
	}
	//4.扣减库存
	err = mysql.UpdateVoucherStock(h.Context, voucherId)
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
		VoucherId:  voucherId,
		OrderId:    orderId,
		PayTime:    time.Now().Format("2006-01-02T15:04:05+08:00"),
		UseTime:    "0000-00-00 00:00:00",
		RefundTime: "0000-00-00 00:00:00",
	}
	err = mysql.CreateVoucherOrder(h.Context, voucherOrder)
	if err != nil {
		return nil, err
	}
	return &orderId, nil
}
