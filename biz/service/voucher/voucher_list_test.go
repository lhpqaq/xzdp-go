package voucher

import (
	"context"
	"testing"
	"xzdp/biz/dal/mysql"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestVoucherListService_Run(t *testing.T) {
	ctx := context.Background()
	mysql.Init()
	c := app.NewContext(1)
	s := NewVoucherListService(ctx, c)
	// init req and assert value
	resp, err := s.Run(1)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
