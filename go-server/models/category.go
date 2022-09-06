package models

type Category struct {
	Id_category   int    `json:"id_category,omitempty" gorm:"primaryKey;type:int(10)"`
	Nama_category string `json:"nama_category,omitempty" gorm:"type:varchar(25)"`
}
