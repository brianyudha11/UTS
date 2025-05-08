package Databases

import "time"

type Users struct {
	ID       int    `json:"id"`
	NoTelpon string
	Email    string `json:"email"`   
	Password string `json:"password"`
	Role     string `json:"role"`    
	Active   bool   `json:"active"`  
}

type Admins struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Nama         string    `json:"nama"`
	Jabatan      string    `json:"jabatan"`
	Alamat       string    `json:"alamat"`
	TanggalMasuk time.Time `json:"tanggal_masuk"`
}