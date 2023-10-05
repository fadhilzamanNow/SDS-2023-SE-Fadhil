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

	pembeli := models.Pembeli{
		Nama_pembeli: data["nama_pembeli"],
		Email:        data["email"],
		Password:     data["password"],
	}

	database.DB.Create(pembeli)

	return c.JSON(fiber.Map{
		"Pesan": "Data telah berhasil di tambahkan",
	})
}

// Lengkapi Code Berikut untuk untuk Mengambil data untuk semua user - user [DONE]
func GetAllData(c *fiber.Ctx) error {

	var pembeli models.Pembeli

	database.DB.Find(&pembeli)

	return c.JSON(fiber.Map{
		"data": pembeli,
	})

}

//Lengkapi Code berikut Untuk Mengambil data dari id_user berdasarkan Parameter [DONE]

func GetUserByid(c *fiber.Ctx) error {

	var pembeli models.Pembeli

	id_pembeli := c.Params("id_pembeli")

	database.DB.Where("id_pembeli = ?", id_pembeli).Find(&pembeli)
	return c.JSON(fiber.Map{
		"id_pembeli": id_pembeli,
		"data":       pembeli,
	})
}

func Delete(c *fiber.Ctx) error {

	var pembeli models.Pembeli

	id_pembeli := c.Params("id_pembeli")

	database.DB.Where("id_pembeli = ?", id_pembeli).Delete(pembeli)

	return c.JSON(fiber.Map{
		"Pesan": "Data telah di hapus",
	})
}

// mengupdate data user berdasarkan id_user yang di tempatkan di parameter [DONE]
func Update(c *fiber.Ctx) error {

	id_user := c.Params("id_user")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var pembeli models.Pembeli
	database.DB.Find(&pembeli)
	//data yang di ubah
	//membuat variable user berdasarkan model user
	var user models.Pembeli

	update := models.Pembeli{
		Nama_pembeli: data["nama"],
		Email:        data["email"],
		Username:     data["username"],
		Password:     data["password"],
	}
	//mengambil database untuk di update

	database.DB.Model(&user).Where("id_user = ?", id_user).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah di Update",
	})
}
