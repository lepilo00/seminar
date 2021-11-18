package uporabnik

type User struct {
	Username string
	Password string
	Email    string
	IsAdmin  bool
}
type Uporabnik struct {
	repo repozitorijInterface
}
type repozitorijInterface interface {
	Ustvari(usr User) error
}

func Nov(repo repozitorijInterface) *Uporabnik {
	return &Uporabnik{
		repo: repo,
	}
}

func (u *Uporabnik) Ustvari(usr User) error {
	return u.repo.Ustvari(usr)
}
