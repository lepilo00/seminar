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
	repoOglas repozitorijOglasInterface
}
type repozitorijOglasInterface interface {
	UstvariOglas(ogs Oglas) error
}

func Nov(repoOglas repozitorijOglasInterface) *Neki {
	return &Neki{
		repoOglas: repoOglas,
	}
}

func (oglas12 *Neki) Ustvari(ogs Oglas) error {
	return oglas12.repoOglas.UstvariOglas(ogs)
}
