package models

type Cart struct {
	Id_cart     int     `json:"id_cart" gorm:"primaryKey;type:int(10)"`
	Jumlah      int     `json:"jumlah" gorm:"type:int(10)"`
	Total_harga int     `json:"total_harga" gorm:"type:int(50)"`
	Keterangan  string  `json:"keterangan" gorm:"type:varchar(75)"`
	Status_cart bool    `json:"status_cart" gorm:"default:true"`
	Id_product  int     `json:"id_product"`
	Product     Product `gorm:"foreignKey:Id_product;references:Id_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}
