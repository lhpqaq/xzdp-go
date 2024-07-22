// idl/shop.thrift
namespace go shop

// import "github.com/hertz-contrib/sessions"

struct Empty {}

struct ShopType {
    1: i64 id
    2: string name
    3: string icon
    4: i32 sort
}

service ShopService {
    list<ShopType> ShopList(1: Empty request) (api.get="/shop-type/list");
}
