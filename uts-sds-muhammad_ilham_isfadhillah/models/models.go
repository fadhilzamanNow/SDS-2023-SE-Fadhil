package models

type Pembeli struct {
	Id_pembeli   uint `json:"id_pembeli" gorm:"primaryKey;autoIncrement:true"`
	Nama_pembeli string
	Email        string
	Username     string
	Password     string
}

type Penjual struct {
	Id_penjual   uint `json:"id_penjual" gorm:"primaryKey;autoIncrement:true"`
	Nama_penjual string
	Email        string
	Username     string
	Password     string
}

type Product struct {
	Id_produk   uint `json:"id_produk" gorm:"primaryKey;autoIncrement:true"`
	Nama_produk string
	harga       int
	kadaluarsa  string
}

//Tulis jawab code nya di sini
