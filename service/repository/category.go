package repository

import (
	//pb "github.com/i-coder-robot/go-micro-action-category/proto/category"
	"github.com/jinzhu/gorm"
)

type Category interface {
	//GetCategoriesPage(ctx context.Context, in *Request, out *Response) error
	//SaveCategory(category *pb.Category) (*pb.Category, error)
	//UpdateGoodsCategory(ctx context.Context, in *Request, out *Response) error
	//GetGoodsCategoryById(ctx context.Context, in *Request, out *Response) error
	//DeleteBatch(ctx context.Context, in *Request, out *Response) error
	//GetCategoriesForSearch(ctx context.Context, in *Request, out *Response) error
	//SelectByLevelAndParentIdsAndNumber(ctx context.Context, in *Request, out *Response) error
}

type CategoryRepository struct {
	DB *gorm.DB
}
