package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

type Space struct {
	gorm.Model
	Name          string           `json:"name"` // 桌球、麻将...
	Num           int64            `json:"num"`  // 1、2、3...
	Price         float64          `json:"price"`
	PriceRuleType pb.PriceRuleType `json:"price_rule_type"` // 定时、记时
}

func (*Space) TableName() string {
	return "space"
}

func (space *Space) ToPb() *pb.Space {
	if space == nil {
		return nil
	}
	return &pb.Space{
		Name:          space.Name,
		Num:           space.Num,
		Price:         space.Price,
		PriceRuleType: space.PriceRuleType,
	}
}
