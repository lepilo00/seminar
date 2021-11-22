package main

import (
	"fmt"
	"net/http"

	//"prijava"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocolly/colly"
	"github.com/lepilo00/seminar/storitev/oglas"
	//"github.com/lepilo00/seminar/shramba"
	//"github.com/lepilo00/seminar/storitev/oglas"
)

// klic funkcije v main programu
func (hh *Handler) ObjavaOglasa(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Println("Napaka1")
	}

	ogs := oglas.Oglas{
		IDoglasa: 1,
		AvtoOglas: oglas.Avto{
			ZnamkaAvta:   r.FormValue("username"),
			ModelAvta:    r.FormValue("username"),
			Cena:         23,
			Letnik:       1,
			PrevozenihKM: 15,
			Gorivo:       r.FormValue("username"),
			PrvaReg:      r.FormValue("username"),
			LetoProiz:    r.FormValue("username"),
			Menjalnik:    r.FormValue("username"),
			VIMstev:      r.FormValue("username"),
			KrajOgleda:   r.FormValue("username"),
		},
	}

	hh.oglas.Ustvari(ogs)

}
