package models

type Cart struct {
	Id_cart     uint          `json:"id_cart" gorm:"primaryKey;type:int(10)"`
	Jumlah      uint          `json:"jumlah" gorm:"type:int(10)"`
	Total_harga uint          `json:"total_harga" gorm:"type:int(50)"`
	Status_cart bool          `json:"status_cart" gorm:"default:true"`
	Id_product  uint          `json:"id_product"`
	Transaction []Transaction `gorm:"foreignKey:Id_cart;references:Id_cart;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
