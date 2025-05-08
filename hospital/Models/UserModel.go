package Models

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
	"project-login-mvc/Databases"
)

func CreateAdminProfile(noTelpAdmin, emailAdmin, passwordAdmin string) (bool, string) {
	userFile, _ := os.ReadFile("Datas/Users.json")
	var users []Databases.Users
	_ = json.Unmarshal(userFile, &users)

	for _, user := range users {
		if user.Email == emailAdmin {
			return false, "Email Sudah Terdaftar"
		}
	}

	for _, user := range users {
		if user.NoTelpon == noTelpAdmin {
			return false, "Nomor Telepon Sudah Terdaftar"
		}
	}

	rand.Seed(time.Now().UnixNano())
	userID := rand.Intn(9000) + 1000


	newUser := Databases.Users{
		ID: userID,
		NoTelpon: noTelpAdmin,
		Email: emailAdmin,
		Password: passwordAdmin,
		Role : "ADMIN",
		Active : true,
	}

	users = append(users, newUser)
	userJSON, _ := json.MarshalIndent(users, "", "	")
	_ = os.WriteFile("Datas/Users.json", userJSON, os.ModePerm)

	adminFile, _ := os.ReadFile("Datas/Admins.json")
	var admins []Databases.Admins
	_ = json.Unmarshal(adminFile, &admins)

	now := time.Now()
	newAdmin := Databases.Admins{
		ID: rand.Intn(999999),
		UserID: userID,
		Nama: "-",
		Jabatan: "-",
		Alamat: "-",
		TanggalMasuk: now,
	}

	admins = append(admins, newAdmin)
	adminJSON, _ := json.MarshalIndent(admins, "", "  ")
	_ = os.WriteFile("Datas/Admins.json", adminJSON, os.ModePerm)	

	return true, ""
}

func LoginAdminProfile(identifier, password string) Databases.Users {
	userFile, _ := os.ReadFile("Datas/Users.json")
	var users []Databases.Users
	_ = json.Unmarshal(userFile, &users)

	for _, user := range users {
		if user.NoTelpon == identifier {
			if user.Password == password {
				return user
			}
		}
	}

	for _, user := range users {
		if user.Email == identifier {
			if user.Password == password {
				return user
			}
		}
	}

	return Databases.Users{}
}