package route

import (
	"belajar-backend/database"
	"belajar-backend/models"

	"github.com/gofiber/fiber/v2"
)

///Membuat CRUD Dasar dan meyimpan nya di database

func CreateMahasiswa(c *fiber.Ctx) error {

	var data map[string]string

	mahasiswa := models.Mahasiswa{
		Phone:          data["phone"],
		Email:          data["email"],
		Nama_mahasiswa: data["nama_mahasiswa"],
		Password:       data["password"],
		Jurusan:        data["jurusan"],
	}

	database.DB.Create(mahasiswa)

	return c.JSON(fiber.Map{})
}

func CreateDosen(c *fiber.Ctx) error {

	var data map[string]string

	dosen := models.Dosen{
		Phone:      data["phone"],
		Email:      data["email"],
		Nama_dosen: data["nama_dosen"],
		Password:   data["password"],
		Jurusan:    data["jurusan"],
	}

	database.DB.Create(dosen)

	return c.JSON(fiber.Map{})
}

func CreateTugas(c *fiber.Ctx) error {

	var data map[string]string

	tugas := models.Tugas{
		Matakuliah: data["mata_kuliah"],
		Deadline:   data["deadline"],
	}

	database.DB.Create(tugas)

	return c.JSON(fiber.Map{})
}

func UpdateMahasiswa(c *fiber.Ctx) error {

	id_mahasiswa := c.Params("id_mahasiswa")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var mahasiswa models.Mahasiswa
	database.DB.Find(&mahasiswa)

	update := models.Mahasiswa{
		Phone:          data["phone"],
		Email:          data["email"],
		Nama_mahasiswa: data["nama_mahasiswa"],
		Password:       data["password"],
		Jurusan:        data["jurusan"],
	}
	//mengambil database untuk di update

	database.DB.Model(&mahasiswa).Where("id_mahasiswa = ?", id_mahasiswa).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "Data Mahasiswa telah di Update",
	})
}

func UpdateDosen(c *fiber.Ctx) error {

	id_dosen := c.Params("id_dosen")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var dosen models.Dosen
	database.DB.Find(&dosen)

	update := models.Dosen{
		Phone:      data["phone"],
		Email:      data["email"],
		Nama_dosen: data["nama_dosen"],
		Password:   data["password"],
		Jurusan:    data["jurusan"],
	}
	//mengambil database untuk di update

	database.DB.Model(&dosen).Where("id_dosen = ?", id_dosen).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "Data Dosen telah di Update",
	})
}

func UpdateTugas(c *fiber.Ctx) error {

	id_tugas := c.Params("id_Tugas")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var tugas models.Tugas
	database.DB.Find(&tugas)

	update := models.Tugas{
		Matakuliah: data["mata_kuliah"],
		Deadline:   data["deadline"],
	}
	//mengambil database untuk di update

	database.DB.Model(&tugas).Where("id_tugas = ?", id_tugas).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "Data Tugas telah di Update",
	})
}

func GetMahasiswa(c *fiber.Ctx) error {
	var mahasiswa []models.Mahasiswa
	database.DB.Find(&mahasiswa)
	return c.JSON(fiber.Map{
		"data": mahasiswa,
	})
}

func GetDosen(c *fiber.Ctx) error {
	var dosen []models.Dosen
	database.DB.Find(&dosen)
	return c.JSON(fiber.Map{
		"data": dosen,
	})
}

func GetTugas(c *fiber.Ctx) error {
	var tugas []models.Tugas
	database.DB.Find(&tugas)
	return c.JSON(fiber.Map{
		"data": tugas,
	})
}

func DeleteDosen(c *fiber.Ctx) error {

	var dosen models.Dosen

	id_dosen := c.Params("id_dosen")

	database.DB.Where("id_user = ?", id_dosen).Delete(dosen)

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah di hapus",
	})

}

func DeleteMahasiswa(c *fiber.Ctx) error {

	var mahasiswa models.Mahasiswa

	id_mahasiswa := c.Params("id_mahasiswa")

	database.DB.Where("id_user = ?", id_mahasiswa).Delete(mahasiswa)

	return c.JSON(fiber.Map{
		"Pesan": "Data Mahasiswa telah di hapus",
	})

}

func DeleteTugas(c *fiber.Ctx) error {

	var tugas models.Tugas

	id_tugas := c.Params("id_tugas")

	database.DB.Where("id_user = ?", id_tugas).Delete(tugas)

	return c.JSON(fiber.Map{
		"Pesan": "Data Tugas telah di hapus",
	})

}
