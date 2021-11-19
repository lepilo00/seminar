package oglas

type Oglas struct {
	IDoglasa  int
	AvtoOglas Avto
}

type Avto struct {
	ZnamkaAvta   string
	ModelAvta    string
	Cena         float32
	Letnik       int
	PrevozenihKM int
	Gorivo       string
	PrvaReg      string
	LetoProiz    string
	Menjalnik    string
	VIMstev      string
	KrajOgleda   string
}

type Oglass struct {
	repo repozitorijInterface
}
type repozitorijInterface interface {
	Ustvari(og Oglas) error
}

func Nov(repo repozitorijInterface) *Oglas {
	return &Oglas{
		repo: repo,
	}
}

func (og *Oglas) Ustvari(oglas Oglas) error {
	return og.repo.Ustvari(oglas)
}
