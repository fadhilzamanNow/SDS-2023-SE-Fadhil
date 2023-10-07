package models

type Admin struct {
	Id_admin   uint `json:"id_admin" gorm:"primaryKey;autoIncrement:true"`
	Nama_admin string
	Email      string
	Username   string
	Password   string
}

type User struct {
	Id_user   uint `json:"id_user" gorm:"primaryKey;autoIncrement:true"`
	Nama_user string
	Email     string
	Username  string
	Password  string
}

//Tulis jawab code nya di sini [dDONE]
