// idl/image.thrift
namespace go image
struct Empty {}
struct UploadResp {
    1: string url;
}
service ImageSerivice {
    UploadResp Upload(1: binary request) (api.post="/upload/blog");
}
