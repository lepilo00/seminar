package oglas

type Oglas struct {
	IDoglasa  int
	AvtoOglas Avto
	//Uporabnik User
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
	repoOglas repozitorijInterface
}
type repozitorijInterface interface {
	UstvariOglas(ogs Oglas) error
}

func Nov(repo repozitorijInterface) *Neki {
	return &Neki{
		repoOglas: repo,
	}
}

func (og *Neki) Ustvari(oglas Oglas) error {
	return og.repoOglas.UstvariOglas(oglas)
}
