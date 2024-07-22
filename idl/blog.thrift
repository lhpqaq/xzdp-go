// idl/blog.thrift
namespace go blog
struct Empty {}


struct Blog {
    1: i64 id
    2: i64 shopId
    3: i64 userId
    4: string icon
    5: string name
    6: string title
    7: string images
    8: string content
    9: i64 liked
    10: i64 comments
    11: string createTime
    12: string updateTime
}

struct BlogReq {
    1: i64 current (api.query="current");
}
service BlogSerivice {
    list<Blog> GetHotBlog(1: BlogReq request) (api.get="/blog/hot");
}