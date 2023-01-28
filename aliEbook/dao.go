package main

var model = &AliEBook{}

func GetAllBooks() ([]AliEBook, error) {
	var res []AliEBook
	db := Mssql.Model(model).Find(&res)
	if db.Error != nil {
		return nil, db.Error
	}
	return res, nil
}
