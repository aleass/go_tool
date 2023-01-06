package main

type DrinkInfo struct {
	DrinkCount int64  `gorm:"column:drink_count"  desc:""`
	Time       string `gorm:"column:time"  desc:""`
	User       string `gorm:"column:users"  desc:""`
}

func (DrinkInfo) TableName() string {
	return "drink_info"
}

type DrinkDetail struct {
	Drink int64  `gorm:"column:drink"  desc:""`
	Time  string `gorm:"column:time"  desc:""`
	User  string `gorm:"column:users"  desc:""`
}

func (DrinkDetail) TableName() string {
	return "drink_detail"
}
