package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"sync"
)

type Estacionamiento struct {
	espacios      chan int
	puerta        *sync.Mutex
	espace [20]bool
}

func NewEstacionamiento(espacios chan int, puertaMu *sync.Mutex) *Estacionamiento {
	return &Estacionamiento{
		espacios:      espacios,
		puerta:        puertaMu,
		espace: [20]bool{},
	}
}

func (p *Estacionamiento) GetEspacios() chan int {
	return p.espacios
}

func (p *Estacionamiento) GetPuertaMu() *sync.Mutex {
	return p.puerta
}

func (p *Estacionamiento) GetEspace() [20]bool {
	return p.espace
}

func (p *Estacionamiento) Setespace(espace [20]bool) {
	p.espace = espace
}

type CarQueue struct {
	Contenedor *fyne.Container
	Imagen     *canvas.Image
}

func (cq *CarQueue) ColaSalida() {
	cq.Imagen.Move(fyne.NewPos(80, 280))
	cq.Contenedor.Add(cq.Imagen)
	cq.Contenedor.Refresh()
}
