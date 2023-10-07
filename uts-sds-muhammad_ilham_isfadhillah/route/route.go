package route

import (
	"uts/database"
	"uts/models"

	"github.com/gofiber/fiber/v2"
)

// buatlah endpoint Insert Data sesuai dengan Colection Postman [DONE]
func InsertData(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(data); err != nil {
		return err
	}

	user := models.User{
		Nama_user: data["nama_pembeli"],
		Email:     data["email"],
		Username:  data["username"],
		Password:  data["password"],
	}

	database.DB.Create(user)

	return c.JSON(fiber.Map{
		"Pesan": "Data telah berhasil di tambahkan",
	})
}

// Lengkapi Code Berikut untuk untuk Mengambil data untuk semua user - user [DONE]
func GetAllData(c *fiber.Ctx) error {

	var user models.User

	database.DB.Find(&user)

	return c.JSON(fiber.Map{
		"data": user,
	})

}

//Lengkapi Code berikut Untuk Mengambil data dari id_user berdasarkan Parameter [DONE]

func GetUserByid(c *fiber.Ctx) error {

	var user models.User

	id_user := c.Params("id_user")

	database.DB.Where("id_user = ?", id_user).Find(&user)
	return c.JSON(fiber.Map{
		"id_pembeli": id_user,
		"data":       user,
	})
}

func Delete(c *fiber.Ctx) error {

	var user models.User

	id_user := c.Params("id_user")

	database.DB.Where("id_user = ?", id_user).Delete(user)

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah di hapus",
	})
}

// mengupdate data user berdasarkan id_user yang di tempatkan di parameter [DONE]
func Update(c *fiber.Ctx) error {

	id_user := c.Params("id_user")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User
	database.DB.Find(&user)
	//data yang di ubah
	//membuat variable user berdasarkan model user
	// var user models.User

	update := models.User{
		Nama_user: data["nama"],
		Email:     data["email"],
		Username:  data["username"],
		Password:  data["password"],
	}
	//mengambil database untuk di update

	database.DB.Model(&user).Where("id_user = ?", id_user).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah di Update",
	})
}
