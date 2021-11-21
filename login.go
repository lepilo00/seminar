package main

import (
	"fmt"
	"net/http"

	//"prijava"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocolly/colly"

	//"github.com/lepilo00/seminar/shramba"
	//"github.com/lepilo00/seminar/storitev/oglas"
	"github.com/lepilo00/seminar/storitev/uporabnik"
)

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
	hashP, err := HashPass(usr1.Password)
	if err != nil {
		fmt.Println("Napaka pri hashu passworda!")
	}

	checkPass := CheckPassHash(usr1.Password, hashP)

	fmt.Println("Password: ", usr1.Password, "\nHash: ", hashP, "\nMatch: ", checkPass)

	tpl.ExecuteTemplate(w, "login.html", nil)

	h.uporabnik.Ustvari(usr1)
}
