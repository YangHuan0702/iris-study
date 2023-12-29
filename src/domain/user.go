package domain

type User struct {
	Id      int32  `gorm:"column:id" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Address string `gorm:"column:address" json:"address"`
	Age     int32  `gorm:"column:age" json:"age"`
	Email   string `gorm:"column:email" json:"email"`
}

func (_ *User) TableName() string {
	return "user_test"
}
