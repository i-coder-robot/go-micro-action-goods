package repository

import (
	"fmt"
	pb "github.com/i-coder-robot/go-micro-action-goods/proto/product"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
)

type Product interface {
	GetProductByPage(query *pb.ListQuery) ([]*pb.Product, error)
	Exist(user *pb.Product) bool
	Save(category *pb.Product) (*pb.Product, error)
	Update(category *pb.Product) (bool, error)
	Delete(category *pb.Product) (bool,error)
	GetProductById(query *pb.Product) (*pb.Product)
}

type ProductRepository struct {
	DB *gorm.DB
}

func (repo *ProductRepository) GetProductByPage(req *pb.ListQuery) (products []*pb.Product,err error){
	db := repo.DB
	limit := req.Limit
	offset := req.Page * 10
	sort := req.Sort
	if err = db.Order(sort).Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return products, nil
}
func (repo *ProductRepository) Exist(proudct *pb.Product) bool{
	var count int
	if proudct.ProductName != "" {
		repo.DB.Model(&proudct).Where("product_name= ?", proudct.ProductName).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

func (repo *ProductRepository) Save(product *pb.Product) (*pb.Product, error){
	if exist := repo.Exist(product); exist == true {
		return product, fmt.Errorf("商品已存在")
	}
	err := repo.DB.Create(product).Error
	if err != nil {
		log.Log(err)
		return product, fmt.Errorf("商品添加失败")
	}
	return product, nil
}
func (repo *ProductRepository) Update(product *pb.Product) (bool, error){
	if product.ProductId == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	c := &pb.Product{
		ProductId: product.ProductId,
	}
	err := repo.DB.Model(c).Update(product).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
func (repo *ProductRepository) Delete(product *pb.Product) (bool,error){
	c := &pb.Product{ProductId: product.ProductId}
	err := repo.DB.Delete(c).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
func (repo *ProductRepository) GetProductById(product *pb.Product) (*pb.Product){
	repo.DB.Where([]string{product.ProductId}).Find(&product)
	return product
}
