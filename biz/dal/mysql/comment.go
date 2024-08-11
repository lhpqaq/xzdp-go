package mysql

import (
	"context"
	"xzdp/biz/model/blog_comment"
)

func DeleteBlogComment(ctx context.Context, blogID int64) error {
	err := DB.WithContext(ctx).Where("blog_id = ?", blogID).Delete(&blog_comment.BlogComment{}).Error
	return err
}
