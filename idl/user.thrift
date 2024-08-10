// idl/user.thrift
namespace go user

struct Empty {}

struct Session {
    1: string Token;
}
struct UserLoginFrom {
    1: string Phone (api.query="phone"  go.tag='gorm:"phone"');
    2: string code (api.query="code");
    3: string password (api.query="password");
    4: Session session (api.query="session");
}

struct User {
    1: string Phone (go.tag='gorm:"phone"');
    2: string code (go.tag='gorm:"-"');
    3: string password (go.tag='gorm:"password"');
    4: i64 id (go.tag='gorm:"id"');
    5: string NickName (go.tag='gorm:"nick_name"');
    6: string icon (go.tag='gorm:"icon"');
    7: string createTime (go.tag='gorm:"create_time"');
    8: string updateTime (go.tag='gorm:"update_time"');
}

struct Result {
    1: bool success,
    2: optional string errorMsg,
    3: optional string data, // 使用 string 类型来表示泛型对象。你可以根据需要选择合适的数据类型。
    4: optional i64 total,
}

struct UserResp {
    1: User RespBody;
}

struct UserDTO {
    1: i64 id (go.tag = 'redis:"id"'),
    2: string NickName (go.tag = 'redis:"nick_name" form:"nickName" json:"nickName" query:"nickName"')
    3: string icon (go.tag = 'redis:"icon"')
}

struct UserInfo {
    1: i64 userId (go.tag='gorm:"user_id"');
    2: string city (go.tag='gorm:"city"');
    3: string introduce (go.tag='gorm:"introduce"');
    4: i64 fans (go.tag='gorm:"fans"');
    5: i64 followee (go.tag='gorm:"followee"');
    6: i64 gender (go.tag='gorm:"gender"');
    7: string birthday (go.tag='gorm:"birthday"');
    8: i64 credits (go.tag='gorm:"credits"');
    9: i64 level (go.tag='gorm:"level"');
    10: string createTime (go.tag='gorm:"create_time"');
    11: string updateTime (go.tag='gorm:"update_time"');
}

service UserService {
    UserDTO UserMe(1: Empty request) (api.get="/user/me");
    UserResp SendCode(1: UserLoginFrom request) (api.post="/user/code");
    UserResp UserLogin(1: UserLoginFrom request) (api.post="/user/login");
<<<<<<< HEAD
    UserInfo UserInfo(1: UserLoginFrom request) (api.get="/user/info/:id");
=======
    UserResp UserInfo(1: string request) (api.get="/user/:id");
>>>>>>> dev
}

