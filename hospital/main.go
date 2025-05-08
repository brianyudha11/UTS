package main

import (
	"fmt"
	"net/http"
	"project-login-mvc/Controllers"
)

func main() {
	http.HandleFunc("/", Controllers.ShowPageAuthAdmin)

	http.HandleFunc("/admin/auth", Controllers.AuthAdmin)

	http.HandleFunc("/admin/dashboard", Controllers.ShowAdminDashboard)
	
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}