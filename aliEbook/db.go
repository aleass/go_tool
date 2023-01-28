package main

type AliEBook struct {
	Ebook int64  `gorm:"column:ebook"  desc:""`
	Name  string `gorm:"column:name"  desc:""`
}

func (AliEBook) TableName() string {
	return "ali_ebook"
}
