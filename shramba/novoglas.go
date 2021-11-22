package shramba

import (
	"database/sql"
	"fmt"

	"github.com/lepilo00/seminar/storitev/oglas"
	//"github.com/lepilo00/seminar/storitev/uporabnik"
)

type OglasRepozitorij struct {
	DB *sql.DB
}

func NovOglasRepozitory() (*OglasRepozitorij, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb") //ne dela
	if err != nil {
		return nil, fmt.Errorf("napaka pri kreiranju repozitorija: %w", err)
	}

	return &OglasRepozitorij{
		DB: db,
	}, nil
}

func (r *OglasRepozitorij) UstvariOglas(ogs oglas.Oglas) error {
	_, err := r.DB.Exec(`INSERT INTO oglas (IDoglasa, ZnamkaAvta,ModelAvta,Cena,Letnik,PrevozenihKM,Gorivo,PrvaReg,LetoProiz,Menjalnik,VIMstev,KrajOgleda) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, ogs.IDoglasa, ogs.AvtoOglas.ZnamkaAvta, ogs.AvtoOglas.ModelAvta, ogs.AvtoOglas.Cena, ogs.AvtoOglas.Letnik, ogs.AvtoOglas.PrevozenihKM, ogs.AvtoOglas.Gorivo, ogs.AvtoOglas.PrvaReg, ogs.AvtoOglas.LetoProiz, ogs.AvtoOglas.Menjalnik, ogs.AvtoOglas.VIMstev, ogs.AvtoOglas.KrajOgleda)
	if err != nil {
		return fmt.Errorf("napaka pri vstavljanju v bazo: %w", err)
	}

	return nil
}
