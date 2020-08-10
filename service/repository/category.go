package repository

import (
	"github.com/jinzhu/gorm"
	pb "github.com/i-coder-robot/go-micro-action-goods/proto/category"
)

type Category interface {
	//GetCategoriesPage(ctx context.Context, in *Request, out *Response) error
	//SaveCategory(ctx context.Context, in *Request, out *Response) error
	//UpdateGoodsCategory(ctx context.Context, in *Request, out *Response) error
	//GetGoodsCategoryById(ctx context.Context, in *Request, out *Response) error
	//DeleteBatch(ctx context.Context, in *Request, out *Response) error
	//GetCategoriesForSearch(ctx context.Context, in *Request, out *Response) error
	//SelectByLevelAndParentIdsAndNumber(ctx context.Context, in *Request, out *Response) error
}

type CategoryRepository struct {
	DB *gorm.DB
}
