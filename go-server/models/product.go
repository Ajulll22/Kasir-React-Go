package models

type Product struct {
	Id_product     uint   `json:"id_product" gorm:"primaryKey;type:int(10)"`
	Kode_product   string `json:"kode_product" gorm:"type:varchar(10)"`
	Nama_product   string `json:"nama_product" gorm:"type:varchar(25)"`
	Harga_product  uint   `json:"harga_product" gorm:"type:int(25)"`
	Status_product bool   `json:"status_product" gorm:"default:true"`
	Gambar_product string `json:"gambar_product" gorm:"type:varchar(25)"`
	Id_category    uint   `json:"id_category"`
	Cart           []Cart `gorm:"foreignKey:Id_product;references:Id_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
