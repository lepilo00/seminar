package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocolly/colly"
	"golang.org/x/crypto/bcrypt"

	"github.com/lepilo00/seminar/storitev/uporabnik"
)

func HashPass(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func CheckPassHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

// klic funkcije v main programu
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
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

	tpl.ExecuteTemplate(w, "html/login.html", usr1)

	//h.uporabnik.Ustvari(usr1)
}
