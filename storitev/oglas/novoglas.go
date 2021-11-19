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

type Neki struct {
	repo repozitorijInterface
}
type repozitorijInterface interface {
	Ustvari(og Oglas) error
}

func Nov(repo repozitorijInterface) *Neki {
	return &Neki{
		repo: repo,
	}
}

func (og *Neki) Ustvari(oglas Oglas) error {
	return og.repo.Ustvari(ogs)
}
