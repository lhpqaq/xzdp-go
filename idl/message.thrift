// idl/message.thrift
namespace go message
include "./user.thrift"
struct Empty {}
struct Message {
  1: i64 from (go.tag='gorm:"-"');
  2: i64 to (go.tag='gorm:"-"');
  3: string content (go.tag='gorm:"-"');
  4: string type (go.tag='gorm:"-"');
  5: string time (go.tag='gorm:"-"');
}
struct MessageResp {
  1: user.UserDTO user;
  2: Message message;
}
service MessageService {
   string Sse(1: string request) (api.get="/message/sse");
}
