package constants

import "time"

const (
	CACHE_SHOP_TYPE_LIST_KEY = "cache:shopType"
	MAX_PAGE_SIZE            = 10
	LOGIN_CODE_KEY           = "login:code:"
	LOGIN_USER_KEY           = "login:token:"
	DEFAULT_PAGE_SIZE        = 5
	CACHE_SHOP_KEY           = "cache:shop:"
	CACHE_NULL_TTL           = time.Minute * 2
	CACHE_SHOP_TTL           = time.Minute * 30
	LOCK_SHOP_KEY            = "lock:shop:"
	LOGIN_CODE_EXPIRE        = 300
)

const (
	FOLLOW_USER_KEY = "follow:user:"
)

const (
	BLOG_LIKED_KEY = "blog:liked:"
	FEED_KEY       = "feed:"
)
