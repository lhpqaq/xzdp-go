// idl/blog_comment.thrift
namespace go blog_comment
struct Empty {}


struct BlogComment {
    1: i64 id (go.tag='gorm:"id"');
    2: i64 blogId (go.tag='gorm:"blog_id"')
    3: i64 userId (go.tag='gorm:"user_id"')
    4: i64 parentId (go.tag='gorm:"parent_id"')
    5: i64 answerId (go.tag='gorm:"answer_id"')
    6: string content (go.tag='gorm:"content"')
    7: i64 liked (go.tag='gorm:"liked"')
    8: i64 comments (go.tag='gorm:"comments"')
    9: string createTime (go.tag='gorm:"create_time"');
    10: string updateTime (go.tag='gorm:"update_time"');
}

struct CommentReq {
    1: i64 current (api.query="current");
}
struct LikeResp {
    1: bool isLiked;
}
service CommentService {
    // 获取热评
    list<BlogComment> GetHotComment(1: CommentReq request) (api.get="/comment/hot");
    // 获取评论
    list<BlogComment> GetComment(1: CommentReq request) (api.get="/comment/:blogID");
    // 发布评论
    BlogComment PostComment(1: BlogComment request) (api.post="/comment/post");
    // 点赞评论
    LikeResp LikeComment(1: string request) (api.put="/comment/like/:id");
    // 删除评论
    Empty DeleteComment(1: string request) (api.delete="/comment/:id");
}
