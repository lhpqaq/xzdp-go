package utils

import (
	"context"
	"xzdp/biz/model/xzdp"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	response := xzdp.NewFailureResponse(err.Error())
	c.JSON(code, response)
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	response := xzdp.NewSuccessResponse(data)
	c.JSON(code, response)
}

func SendRawResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}
