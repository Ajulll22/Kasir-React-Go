package models

type Product struct {
	Id_product     int      `json:"id_product,omitempty" gorm:"primaryKey;type:int(10)"`
	Kode_product   string   `json:"kode_product,omitempty" gorm:"type:varchar(10)"`
	Nama_product   string   `json:"nama_product,omitempty" gorm:"type:varchar(25)"`
	Harga_product  int      `json:"harga_product,omitempty" gorm:"type:int(25)"`
	Status_product bool     `json:"status_product,omitempty" gorm:"default:true"`
	Gambar_product string   `json:"gambar_product,omitempty" gorm:"type:varchar(25)"`
	Id_category    int      `json:"id_category,omitempty"`
	Category       Category `gorm:"foreignKey:Id_category;references:Id_category;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category,omitempty"`
}
