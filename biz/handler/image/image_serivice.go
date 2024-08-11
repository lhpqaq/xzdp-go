package image

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	image2 "image"
	"io"
	"os"
	"xzdp/biz/model/image"
	"xzdp/biz/service"
	"xzdp/biz/utils"
)

// Upload .
// @router /upload [POST]
func Upload(ctx context.Context, c *app.RequestContext) {
	var err error
	var _ image.Empty
	var req []byte
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// 将req存为图片文件
	_, format, err := image2.DecodeConfig(bytes.NewReader(req))
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
	}
	// 检查是否为常见图片格式
	if format != "jpeg" && format != "png" && format != "gif" {
		utils.SendErrResponse(ctx, c, consts.StatusOK, errors.New("unsupported image format"))
	}
	//生成文件路径 /uplaod/uuid.ext
	uuid, err := utils.RandomUUID()
	filePath := fmt.Sprintf("/upload/img/%s.%s", uuid, format)
	// 存储图片
	file, err := os.Create(filePath)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
	}
	defer file.Close()
	_, err = io.Copy(file, bytes.NewReader(req))
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
	}
	_, err = service.NewUploadService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, fmt.Sprintf("/img/%s.%s", uuid, format))
}
