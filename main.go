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
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("html/index.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "html/index.html", nil)
}

func main() {
	repo, err := shramba.NovUpoabnikRepozitory()
	if err != nil {
		fmt.Println(err)
	}
	defer repo.DB.Close()

	repoOglas, err := shramba.NovOglasRepozitory()
	if err != nil {
		fmt.Println(err)
	}

	u := uporabnik.Nov(repo)
	oglas12 := oglas.Nov(repoOglas)
	h := NovHandler(u)
	hh := NovHandler1(oglas12)

	http.HandleFunc("/", index)
	http.HandleFunc("/login", h.Login)                //login.go
	http.HandleFunc("/objavaOglasa", hh.ObjavaOglasa) //objavaOglasa.go
	http.HandleFunc("/registracija", Registracija)
	http.ListenAndServe(":9090", nil)

}
