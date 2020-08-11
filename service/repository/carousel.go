package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	pb "github.com/i-coder-robot/go-micro-action-goods/proto/carousel"
	"github.com/micro/go-micro/v2/util/log"
)

type Carousel interface {
	List(query *pb.ListQuery) ([]*pb.Carousel,error)
    Save(carousel *pb.Carousel) (*pb.Carousel,error)
	Update (carousel *pb.Carousel) (bool,error)
	Info (carousel *pb.Carousel) (*pb.Carousel,error)
	Delete(carousel pb.Carousel) (bool,error)
}

type CarouselRepository struct {
	DB *gorm.DB
}

func (repo *CarouselRepository) List(req *pb.ListQuery) (carousels []*pb.Carousel,err error){
	db := repo.DB
	limit := req.Limit
	offset := req.Page * 10
	sort := req.Sort
	if err = db.Order(sort).Limit(limit).Offset(offset).Find(&carousels).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return carousels, nil
}

func (repo *CarouselRepository) Save(carousel *pb.Carousel) (*pb.Carousel,error){
	err := repo.DB.Create(carousel).Error
	if err != nil {
		log.Log(err)
		return carousel, fmt.Errorf("用户注册失败")
	}
	return carousel, nil
}

func (repo *CarouselRepository) Update (carousel *pb.Carousel) (bool,error){
	if carousel.CarouselId == "" {
		return false, fmt.Errorf("请传入更新 ID")
	}
	c := &pb.Carousel{
		CarouselId: carousel.CarouselId,
	}
	err := repo.DB.Model(c).Update(c).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

func (repo *CarouselRepository) Info (carousel *pb.Carousel) (c *pb.Carousel,err error){
	err=repo.DB.Where([]string{carousel.CarouselId}).Find(&c).Error
	if err != nil {
		return nil,err
	}
	return c,nil
}

func (repo *CarouselRepository) Delete(carousel pb.Carousel) (bool,error){
	c := &pb.Carousel{CarouselId: carousel.CarouselId}
	err := repo.DB.Delete(c).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}