package main

import (
	"github.com/lepilo00/seminar/storitev/oglas"
	"github.com/lepilo00/seminar/storitev/uporabnik"
)

type (
	Handler struct {
		uporabnik *uporabnik.Uporabnik
		oglas     *oglas.Neki
	}
)

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
