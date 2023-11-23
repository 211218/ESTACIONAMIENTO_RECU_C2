package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"sync"
)

type Parking struct {
	espacios      chan int
	puerta        *sync.Mutex
	espace [20]bool
}

type VehicleA struct {
	Contenedor *fyne.Container
	Imagen     *canvas.Image
}

func NewParking(espacios chan int, puertaMu *sync.Mutex) *Parking {
	return &Parking{
		espacios:      espacios,
		puerta:        puertaMu,
		espace: [20]bool{},
	}
}


func (cq *VehicleA) Salida() {
	cq.Imagen.Move(fyne.NewPos(80, 280))
	cq.Contenedor.Add(cq.Imagen)
	cq.Contenedor.Refresh()
}

func (p *Parking) GetEspacios() chan int {
	return p.espacios
}

func (p *Parking) GetDoor() *sync.Mutex {
	return p.puerta
}

func (p *Parking) GetEspace() [20]bool {
	return p.espace
}

func (p *Parking) Setespace(espace [20]bool) {
	p.espace = espace
}
