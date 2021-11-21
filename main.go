package main

import (
	"fmt"
	"net/http"
	"text/template"

	//"prijava"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocolly/colly"
	"github.com/lepilo00/seminar/shramba"
	"github.com/lepilo00/seminar/storitev/oglas"
	"github.com/lepilo00/seminar/storitev/uporabnik"
	"golang.org/x/crypto/bcrypt"
)

type (
	Handler struct {
		uporabnik *uporabnik.Uporabnik
		oglas     *oglas.Oglas
	}
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func NovHandler(uporabnik *uporabnik.Uporabnik) *Handler {
	return &Handler{
		uporabnik: uporabnik,
	}
}

func NovHandler1(oglas *oglas.Oglas) *Handler {
	return &Handler{
		oglas: oglas,
	}
}

func HashPass(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func CheckPassHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
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
	hashP, err := HashPass(usr1.Password)
	if err != nil {
		fmt.Println("Napaka pri hashu passworda!")
	}

	checkPass := CheckPassHash(usr1.Password, hashP)

	fmt.Println("Password: ", usr1.Password, "\nHash: ", hashP, "\nMatch: ", checkPass)

	tpl.ExecuteTemplate(w, "login.html", nil)

	h.uporabnik.Ustvari(usr1)
}

func (hh *Handler) ObjavaOglasa(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Println("Napaka1")
	}

	oglas1 := oglas.Oglas{
		IDoglasa: 1,
		AvtoOglas: Avto{
			ZnamkaAvta:   r.FormValue("username"),
			ModelAvta:    r.FormValue("username"),
			Cena:         r.FormValue("username"),
			Letnik:       r.FormValue("username"),
			PrevozenihKM: r.FormValue("username"),
			Gorivo:       r.FormValue("username"),
			PrvaReg:      r.FormValue("username"),
			LetoProiz:    r.FormValue("username"),
			Menjalnik:    r.FormValue("username"),
			VIMstev:      r.FormValue("username"),
			KrajOgleda:   r.FormValue("username"),
		},
	}
}

func main() {
	repo, err := shramba.NovUpoabnikRepozitory()
	if err != nil {
		fmt.Println(err)
	}
	defer repo.DB.Close()

	u := uporabnik.Nov(repo)
	ogs := oglas.Nov(repoU)
	h := NovHandler(u)
	hh := NovHandler1(ogs)

	http.HandleFunc("/", index)
	http.HandleFunc("/login", h.Registracija)
	http.HandleFunc("/oglas", hh.ObjavaOglasa)
	http.ListenAndServe(":9090", nil)

}
