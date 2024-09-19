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
	SHOP_GEO_KEY             = "shop:geo:"
	//CACHE_SHOP_TTL    = 10 * time.Second
	LOCK_SHOP_KEY     = "lock:shop:"
	LOCK_KEY          = "lock:"
	LOGIN_CODE_EXPIRE = 300
	USER_SIGN_KEY     = "user:sign:"
	HLL_UV_KEY        = "HyperLogLog:uv"
)

// cache 相关
const (
	CACHE_USERDTO_KEY    = "userdto:"
	CACHE_USERDTO_EXPIRE = time.Minute * 5
)

// follow 相关
const (
	FOLLOW_USER_KEY = "follow:user:"
)

// blog 相关
const (
	BLOG_LIKED_KEY = "blog:liked:"
	FEED_KEY       = "feed:"
)

// message 相关
const (
	MESSAGE_STREAM_KEY = "message.stream:"
	STREAM_READ_GROUP  = "stream.group:1"
	STREAM_CONSUMER    = "stream.consume:"
)

// voucher 相关
const (
	BEGIN_TIMESTAMP  = 1725120000
	ICRID_KEY        = "id:"
	COUNT_BIT        = 32
	VOUCHER_LOCK_KEY = "voucher:lock:"
)
