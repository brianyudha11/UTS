package Controllers

import (
	"html/template"
	"net/http"
	"project-login-mvc/Models"
)

func ShowPageAuthAdmin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/login.html"))
    tmpl.Execute(w, nil)
}

func AuthAdmin(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        action := r.FormValue("actionAdmin")

        switch action {
			case "adminRegister":
				RegisterAdmin(w, r)
			case "adminLogin":
				LoginAdmin(w, r)
            default :
                http.Error(w, "Invalid action", http.StatusBadRequest)
        }
    } else {
        tmpl, err := template.ParseFiles("views/login.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
    }
}

func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	noTelpAdmin := r.FormValue("noTelpAdmin")
	emailAdmin := r.FormValue("emailAdmin")
	passwordAdmin := r.FormValue("passwordAdmin")

	success, errMsg := Models.CreateAdminProfile(noTelpAdmin, emailAdmin, passwordAdmin)

    tmpl, err := template.ParseFiles("views/login.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	if !success {
		tmpl.Execute(w, map[string]string{
			"Error": errMsg,
		})
		return
	}

    tmpl.Execute(w, map[string]string{
		"Success": "Akun berhasil dibuat <br> Silakan login",
	})
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	identifierAdmin := r.FormValue("identifierAdmin")
	passwordAdmin := r.FormValue("passwordAdmin")

	user := Models.LoginAdminProfile(identifierAdmin, passwordAdmin)

    if user.ID == 0 || user.Role != "ADMIN" {
        tmpl, _ := template.ParseFiles("views/login.html")
        tmpl.Execute(w, map[string]string{
            "ErrorAdmin": "Email - Username atau kata sandi salah",
        })
        return
    }

    http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func ShowAdminDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}