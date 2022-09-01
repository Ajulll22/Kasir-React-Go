package models

type Transaction struct {
	Id_trx      uint `json:"id_trx" gorm:"primaryKey;type:int(10)"`
	Total_bayar uint `json:"total_bayar" gorm:"type:int(50)"`
	Id_cart     uint `json:"id_cart"`
}
