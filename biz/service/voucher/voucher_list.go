package voucher

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"xzdp/biz/dal/mysql"
	voucher "xzdp/biz/model/voucher"
)

type VoucherListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewVoucherListService(Context context.Context, RequestContext *app.RequestContext) *VoucherListService {
	return &VoucherListService{RequestContext: RequestContext, Context: Context}
}

func (h *VoucherListService) Run(id int64) (resp *[]*voucher.SeckillVoucher, err error) {
	result, err := mysql.QueryShopVoucherByShopID(h.Context, id)
	return &result, nil
}
