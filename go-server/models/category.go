package models

type Category struct {
	Id_category   uint      `json:"id_category" gorm:"primaryKey;type:int(10)"`
	Nama_category string    `json:"nama_category" gorm:"type:varchar(25)"`
	Product       []Product `json:"-" gorm:"foreignKey:Id_category;references:Id_category;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
