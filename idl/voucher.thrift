namespace go voucher

struct Empty {}

struct Voucher {
    1: i64 id (go.tag='gorm:"id"');
    2: i64 shopId (go.tag='gorm:"shop_id"');
    3: string title (go.tag='gorm:"title"');
    4: string subTitle (go.tag='gorm:"sub_title"');
    5: string rules (go.tag='gorm:"rules"');
    6: i64 payValue (go.tag='gorm:"pay_value"');
    7: i64 actualValue (go.tag='gorm:"actual_value"');
    8: i8 type (go.tag='gorm:"type"');
    9: i8 status (go.tag='gorm:"status"');
    10: string createTime (go.tag='gorm:"create_time"');
    11: string updateTime (go.tag='gorm:"update_time"');
}

struct SeckillVoucher {
    1: i64 voucherId (go.tag='gorm:"voucher_id"');
    2: i32 stock (go.tag='gorm:"stock"');
    3: string createTime (go.tag='gorm:"create_time"');
    4: string beginTime (go.tag='gorm:"begin_time"');
    5: string endTime (go.tag='gorm:"end_time"');
    6: string updateTime (go.tag='gorm:"update_time"');
}

struct VoucherOrder {
    1: i64 id (go.tag='gorm:"column:id;primaryKey;autoIncrement:false"');
    2: i64 userId (go.tag='gorm:"column:user_id"');
    3: i64 voucherId (go.tag='gorm:"column:voucher_id"');
    4: i32 payType (go.tag='gorm:"column:pay_type;default:1"');
    5: i32 status (go.tag='gorm:"column:status;default:1"');
    6: string createTime (go.tag='gorm:"column:create_time;default:CURRENT_TIMESTAMP"');
    7: string payTime (go.tag='gorm:"column:pay_time"');
    8: string useTime (go.tag='gorm:"column:use_time"');
    9: string refundTime (go.tag='gorm:"column:refund_time"');
    10: string updateTime (go.tag='gorm:"column:update_time;default:CURRENT_TIMESTAMP;autoUpdateTime"');
}

service VoucherService {
    list<Voucher> VoucherList(1: Empty request) (api.get="/voucher/list/:id");
    i64 SeckillVoucher(1: Empty request) (pi.post="/seckill/:id");
}