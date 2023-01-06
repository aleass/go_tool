package main

import "time"

// 喝水
func Drinking() (int64, error) {
	//添加详细表
	now := time.Now()
	var db = DrinkDetail{
		User:  user,
		Drink: 1,
		Time:  now.Format("2006-01-02 15:04:05"),
	}
	row := Mssql.Create(&db)
	if row.Error != nil {
		return 0, row.Error
	}

	//每天
	date := now.Format("2006-01-02")
	drinkInfo, err := GetDrinkInfo(date)
	if err != nil {
		return 0, err
	}
	if len(drinkInfo.Time) > 0 {
		drinkInfo.DrinkCount++
		row = Mssql.Model(model).Where("time = ? and users = ?", date, user).Limit(1).
			Update("drink_count", drinkInfo.DrinkCount)
	} else {
		drinkInfo.DrinkCount = 1
		drinkInfo.Time = date
		drinkInfo.User = user
		row = Mssql.Create(&drinkInfo)
	}
	return drinkInfo.DrinkCount, row.Error
}

var model = &DrinkInfo{}

func GetDrinkInfo(now string) (*DrinkInfo, error) {
	var res = new(DrinkInfo)
	db := Mssql.Model(model).Where("time = ? and users = ?", now, user).Find(res)
	if db.Error != nil {
		return nil, db.Error
	}
	return res, nil
}
