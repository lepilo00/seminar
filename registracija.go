package main

import (
	"fmt"
	"net/http"

	"github.com/lepilo00/seminar/storitev/uporabnik"
	"golang.org/x/crypto/bcrypt"
)

func HashPassRegistracija(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func (h *Handler) Registracija(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Println("Napaka1")
	}

	usr1 := uporabnik.User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Email:    r.FormValue("email"),
		IsAdmin:  false,
	}

	fmt.Println(usr1.Password)

	//kreiranje hash oblike passworda
	_, err := HashPassRegistracija(usr1.Password)
	if err != nil {
		fmt.Println("Napaka pri hashu passworda!")
	}

	tpl.ExecuteTemplate(w, "html/login.html", nil)

	h.uporabnik.Ustvari(usr1)
}
