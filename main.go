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

type (
	Handler struct {
		uporabnik *uporabnik.Uporabnik
		oglas     *oglas.Neki
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

func NovHandler1(oglas *oglas.Neki) *Handler {
	return &Handler{
		oglas: oglas,
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
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
	http.HandleFunc("/login", h.Registracija)  //login.go
	http.HandleFunc("/oglas", hh.ObjavaOglasa) //objavaOglasa.go
	http.ListenAndServe(":9090", nil)

}
