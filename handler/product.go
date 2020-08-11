package handler

import (
	"context"
	pb "github.com/i-coder-robot/go-micro-action-goods/proto/product"
	"github.com/i-coder-robot/go-micro-action-goods/service/repository"
)

type Product struct {
	Repo repository.Product
}

func (p *Product) GetProductsByPage(ctx context.Context, req *pb.Request, res *pb.Response) error{
	products, err := p.Repo.GetProductByPage(req.ListQuery)
	if err != nil {
		return err
	}
	res.ProductList=products
	return nil
}

func (p *Product) Exist(ctx context.Context, req *pb.Request, res *pb.Response) error{
	exist := p.Repo.Exist(req.Product)
	res.Valid=exist
	return nil
}

func (p *Product) Create(ctx context.Context, req *pb.Request, res *pb.Response) error{
	r, err := p.Repo.Save(req.Product)
	if err != nil {
		return err
	}
	res.Product=r
	return nil
}

func (p *Product) Update(ctx context.Context, req *pb.Request, res *pb.Response) error{
	b, err := p.Repo.Update(req.Product)
	if err != nil {
		return err
	}
	res.Valid=b
	return nil
}

func (p *Product) Delete(ctx context.Context, req *pb.Request, res *pb.Response) error{
	b, err := p.Repo.Delete(req.Product)
	if err != nil {
		return err
	}
	res.Valid=b
	return nil
}

func (p *Product) Info(ctx context.Context, req *pb.Request, res *pb.Response) error{
	r := p.Repo.GetProductById(req.Product)
	res.Product=r
	return nil
}
