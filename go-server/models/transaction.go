package models

type Transaction struct {
	Id_trx      int `json:"id_trx" gorm:"primaryKey;type:int(10)"`
	Total_bayar int `json:"total_bayar" gorm:"type:int(50)"`
}
