package shramba

import (
	"database/sql"
	"fmt"

	"github.com/lepilo00/seminar/storitev/uporabnik"
)

type UporabnikRepozitorij struct {
	DB *sql.DB
}

func NovUpoabnikRepozitory() (*UporabnikRepozitorij, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		return nil, fmt.Errorf("napaka pri kreiranju repozitorija: %w", err)
	}

	return &UporabnikRepozitorij{
		DB: db,
	}, nil
}

func (r *UporabnikRepozitorij) Ustvari(usr uporabnik.User) error {
	_, err := r.DB.Exec(`INSERT INTO user (username, password, email, isAdmin) VALUES (?, ?, ?, ?)`, usr.Username, usr.Password, usr.Email, usr.IsAdmin)
	if err != nil {
		return fmt.Errorf("napaka pri vstavljanju v bazo: %w", err)
	}

	return nil
}
