package handler

import (
	"context"
	pb "github.com/i-coder-robot/go-micro-action-goods/proto/category"
	"github.com/i-coder-robot/go-micro-action-goods/service/repository"
)

type Category struct {
	Repo repository.Category
}

func (c *Category) GetCategoryByPage(ctx context.Context, req *pb.Request, res *pb.Response) (err error){
	result,e := c.Repo.GetCategoryByPage(req.ListQuery)
	if e != nil {
		return e
	}
	res.CategoryList = result
	return nil
}

func (c *Category) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error){
	result := c.Repo.Exist(req.Category)
	res.Valid = result
	return err
}

func (c *Category) Save(ctx context.Context, req *pb.Request, res *pb.Response) (err error){
	r, e := c.Repo.Save(req.Category)
	if e != nil {
		return e
	}
	res.Category=r
	return nil
}

func (c *Category) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error){
	r, e := c.Repo.Update(req.Category)
	if e != nil {
		return e
	}
	res.Valid=r
	return nil
}

func (c *Category) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error){
	b, e := c.Repo.Delete(req.Category)
	if e != nil {
		return e
	}
	res.Valid=b
	return nil
}

func (c *Category) GetCategoryById(ctx context.Context, req *pb.Request, res *pb.Response) (err error){
	res.Category = c.Repo.GetCategoryById(req.Category)
	return nil
}

func (c *Category) GetCategories4Search(ctx context.Context, req *pb.Request, res *pb.Response) (err error){
	r, e := c.Repo.GetCategories4Search(req.Category)
	if e != nil {
		return e
	}
	res.CategoryList=r
	return nil
}

//func (c *Category) SelectByLevelAndParentIdsAndNumber(ctx context.Context, req *pb.Request, res *pb.Response) (err error){
//	categories := c.Repo.SelectByLevelAndParentIdsAndNumber(req.ListQuery)
//	res.Categories=categories
//	return nil
//}