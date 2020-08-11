package handler

import (
	goodsPB "github.com/i-coder-robot/go-micro-action-goods/proto/category"
	productPB "github.com/i-coder-robot/go-micro-action-goods/proto/product"
	db "github.com/i-coder-robot/go-micro-action-goods/providers/database"
	"github.com/i-coder-robot/go-micro-action-goods/service/repository"
	"github.com/micro/go-micro/v2/server"
)

type Handler struct {
	Server server.Server
}

func (srv *Handler) Register() {
	goodsPB.RegisterCategoriesHandler(srv.Server, srv.Category())
	productPB.RegisterProductsHandler(srv.Server,srv.Product())
}

func (srv *Handler) Category() *Category {
	return &Category{
		Repo: &repository.CategoryRepository{
			DB: db.DB,
		},
	}
}

func (srv *Handler) Product() *Product {
	return &Product{
		Repo: &repository.ProductRepository{
			DB: db.DB,
		},
	}
}
