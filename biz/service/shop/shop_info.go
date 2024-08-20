package shop

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
	"xzdp/biz/dal/redis"
	"xzdp/biz/pkg/constants"

	shop "xzdp/biz/model/shop"

	"xzdp/biz/dal/mysql"

	"github.com/cloudwego/hertz/pkg/app"
)

type ShopInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewShopInfoService(Context context.Context, RequestContext *app.RequestContext) *ShopInfoService {
	return &ShopInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *ShopInfoService) Run(id int64) (resp *shop.Shop, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	key := fmt.Sprintf("%s%d", constants.CACHE_SHOP_KEY, id)
	lockKey := fmt.Sprintf("%s%d", constants.LOCK_SHOP_KEY, id)

	// 1. 从 Redis 获取数据
	result, err := redis.GetShopFromCache(h.Context, key)
	if err == nil {
		return result, nil
	}

	// 2. 缓存未命中，尝试获取锁
	isLocked := redis.TryLock(h.Context, lockKey)
	if !isLocked {
		// 锁获取失败，等待后重试
		time.Sleep(50 * time.Millisecond)
		return h.Run(id)
	}

	// 2.2 获取锁成功，再次检查缓存
	result, err = redis.GetShopFromCache(h.Context, key)
	if err == nil {
		redis.UnLock(h.Context, lockKey)
		return result, nil
	}

	// 3. 从数据库查询数据
	var shop *shop.Shop
	shop, err = mysql.QueryByID(h.Context, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			redis.RedisClient.Set(h.Context, key, "", constants.CACHE_NULL_TTL).Err()
			redis.UnLock(h.Context, lockKey)
			return nil, err
		}
		redis.UnLock(h.Context, lockKey)
		return nil, err
	}

	// 4. 数据库中存在，缓存数据
	shopJson, err := json.Marshal(*shop)
	if err != nil {
		redis.UnLock(h.Context, lockKey)
		return nil, err
	}

	err = redis.RedisClient.Set(h.Context, key, string(shopJson), constants.CACHE_SHOP_TTL).Err()
	if err != nil {
		return nil, err
	}
	redis.UnLock(h.Context, lockKey)

	return shop, nil
}
