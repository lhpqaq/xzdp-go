package voucher

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	voucher "xzdp/biz/model/voucher"
)

func TestSeckillVoucherService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewSeckillVoucherService(ctx, c)
	// init req and assert value
	req := &voucher.Empty{}
	resp, err := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
