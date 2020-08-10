package repository

import (
	"fmt"
	pb "github.com/i-coder-robot/go-micro-action-goods/proto/category"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
)

type Category interface {
	GetCategoryByPage(query *pb.ListQuery) ([]*pb.Category, error)
	Exist(user *pb.Category) bool
	Save(category *pb.Category) (*pb.Category, error)
	Update(category *pb.Category) (bool, error)
	Delete(category *pb.Category) (bool,error)
	GetCategoryById(query *pb.Category) (*pb.Category)
	GetCategories4Search (category *pb.Category) ([]*pb.Category,error)
	//SelectByLevelAndParentIdsAndNumber(query *pb.ListQuery) (categories []*pb.Category)
}

type CategoryRepository struct {
	DB *gorm.DB
}

func (repo *CategoryRepository) GetCategoryByPage(req *pb.ListQuery) (categories []*pb.Category, err error){
	db := repo.DB
	limit := req.Limit
	offset := req.Page * 10
	sort := req.Sort
	if err = db.Order(sort).Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return categories, nil
}

func (repo *CategoryRepository) Exist(category *pb.Category) bool {
	var count int
	if category.CategoryName != "" {
		repo.DB.Model(&category).Where("category_name= ?", category.CategoryName).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

func (repo *CategoryRepository) Save(category *pb.Category) (*pb.Category, error){
	if exist := repo.Exist(category); exist == true {
		return category, fmt.Errorf("用户注册已存在")
	}
	err := repo.DB.Create(category).Error
	if err != nil {
		log.Log(err)
		return category, fmt.Errorf("用户注册失败")
	}
	return category, nil
}

func (repo *CategoryRepository) Update(category *pb.Category) (bool, error){
	if category.CategoryId == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	c := &pb.Category{
		CategoryId: category.CategoryId,
	}
	err := repo.DB.Model(c).Update(category).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

func (repo *CategoryRepository) Delete(category *pb.Category) (bool, error){
	c := &pb.Category{CategoryId: category.CategoryId}
	err := repo.DB.Delete(c).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

func (repo *CategoryRepository) GetCategoryById(category *pb.Category) (c *pb.Category){
	repo.DB.Where([]string{category.CategoryId}).Find(&c)
	return c
}

func (repo *CategoryRepository) GetCategories4Search(category *pb.Category) (categories []*pb.Category, err error){
	err=repo.DB.Where("category_name",category.CategoryName).Find(categories).Error
	return categories,err
}

//func (repo *CategoryRepository) SelectByLevelAndParentIdsAndNumber(category *pb.Category) (categories []*pb.Category){
//	ids:= category.CategoryIds
//	idArr:=strings.Split(ids,",")
//	repo.DB.Where("level=",category.CategoryLevel).Where(idArr).Find(&categories)
//	return categories
//}

