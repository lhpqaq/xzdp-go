package shop

import (
	"context"

	shop "xzdp/biz/model/shop"

	"xzdp/biz/dal/mysql"
	"xzdp/biz/pkg/constants"

	"github.com/cloudwego/hertz/pkg/app"
)

type ShopOfTypeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewShopOfTypeService(Context context.Context, RequestContext *app.RequestContext) *ShopOfTypeService {
	return &ShopOfTypeService{RequestContext: RequestContext, Context: Context}
}

func (h *ShopOfTypeService) Run(req *shop.ShopOfTypeReq) (resp *[]*shop.Shop, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var shops []*shop.Shop
	err = mysql.DB.Where("type_id = ?", req.TypeId).Offset(int((req.Current - 1) * constants.DEFAULT_PAGE_SIZE)).Limit(constants.DEFAULT_PAGE_SIZE).Find(&shops).Error
	if err != nil {
		return nil, err
	}
	return &shops, nil
}
