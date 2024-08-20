package image

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	image "xzdp/biz/model/image"
	service "xzdp/biz/service/image"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Upload .
// @router /upload [POST]
func Upload(ctx context.Context, c *app.RequestContext) {
	var _ image.Empty
	var req []byte
	_, err := service.NewUploadService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// 以上代码无用,因为每次更新idl都会引入service和model导致编译不通过
	file, err := c.FormFile("file")
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	ext := filepath.Ext(file.Filename)
	// 检查是否为常见图片格式
	if ext != ".jpeg" && ext != ".png" && ext != ".jpg" {
		utils.SendErrResponse(ctx, c, consts.StatusOK, errors.New("不允许的类型"))
		return
	}
	uuid, err := utils.RandomUUID()
	fp := fmt.Sprintf("%s%s", uuid, ext)
	// 获取项目的根目录
	rootDir, err := os.Getwd()
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	relativeDir := "upload/imgs/blogs"
	dir := filepath.Join(rootDir, relativeDir)
	err = utils.CreateDir(dir)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	err = c.SaveUploadedFile(file, filepath.Join(dir, fp))
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, fmt.Sprintf("/blogs/%s", fp))
}

//
//// Upload .
//// @router /upload [POST]
//func Upload(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var _ image.Empty
//	var req []byte
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
//		return
//	}
//	// 将req存为图片文件
//	_, format, err := image2.DecodeConfig(bytes.NewReader(req))
//	if err != nil {
//		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
//	}
//	// 检查是否为常见图片格式
//	if format != "jpeg" && format != "png" && format != "gif" {
//		utils.SendErrResponse(ctx, c, consts.StatusOK, errors.New("unsupported image format"))
//	}
//	//生成文件路径 /uplaod/uuid.ext
//	uuid, err := utils.RandomUUID()
//	filePath := fmt.Sprintf("/upload/img/blogs/%s.%s", uuid, format)
//	// 存储图片
//	file, err := os.Create(filePath)
//	if err != nil {
//		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
//	}
//	defer file.Close()
//	_, err = io.Copy(file, bytes.NewReader(req))
//	if err != nil {
//		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
//	}
//	_, err = service.NewUploadService(ctx, c).Run(&req)
//	if err != nil {
//		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
//		return
//	}
//
//	utils.SendSuccessResponse(ctx, c, consts.StatusOK, fmt.Sprintf("/img/blogs/%s.%s", uuid, format))
//}
