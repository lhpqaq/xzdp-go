package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	shop "xzdp/biz/model/shop"
	"xzdp/biz/pkg/constants"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ShopListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewShopListService(Context context.Context, RequestContext *app.RequestContext) *ShopListService {
	return &ShopListService{RequestContext: RequestContext, Context: Context}
}

func QueryShopType(ctx context.Context) ([]*shop.ShopType, error) {
	var shopTypeList []*shop.ShopType

	// 从 Redis 获取数据
	shopTypeJsonList, err := redis.RedisClient.LRange(ctx, constants.CACHE_SHOP_TYPE_LIST_KEY, 0, -1).Result()
	if err == nil && len(shopTypeJsonList) > 0 {
		for _, shopTypeJson := range shopTypeJsonList {
			var shopType shop.ShopType
			err := json.Unmarshal([]byte(shopTypeJson), &shopType)
			if err != nil {
				log.Printf("Error unmarshalling shop type: %v", err)
				continue
			}
			shopTypeList = append(shopTypeList, &shopType)
		}
		return shopTypeList, nil
	}

	// 如果 Redis 没有数据，从数据库获取数据
	err = mysql.DB.Order("sort asc").Find(&shopTypeList).Error
	if err != nil {
		return nil, err
	}

	if len(shopTypeList) == 0 {
		return nil, fmt.Errorf("no exist shop type")
	}

	// 将数据存储到 Redis
	for _, shopType := range shopTypeList {
		shopTypeJson, err := json.Marshal(shopType)
		if err != nil {
			log.Printf("Error marshalling shop type: %v", err)
			continue
		}
		redis.RedisClient.RPush(ctx, constants.CACHE_SHOP_TYPE_LIST_KEY, shopTypeJson)
	}

	return shopTypeList, nil
}

func (h *ShopListService) Run(req *shop.Empty) (resp *[]*shop.ShopType, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	shopTypeList, err := QueryShopType(h.Context)
	if err != nil {
		return nil, err
	}

	return &shopTypeList, nil
}
