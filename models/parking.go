package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"sync"
)

type Parking struct {
	ChannelSpace chan int
	puerta   *sync.Mutex
	espace   [20]bool
}

type VehicleA struct {
	Contenedor *fyne.Container
	Imagen     *canvas.Image
}

func NewParking(ChannelSpace chan int, puertaMu *sync.Mutex) *Parking {
	return &Parking{
		ChannelSpace: ChannelSpace,
		puerta:   puertaMu,
		espace:   [20]bool{},
	}
}

func (p *Parking) GetChannelSpace() chan int {
	return p.ChannelSpace
}

func (p *Parking) GetDoor() *sync.Mutex {
	return p.puerta
}

func (p *Parking) GetEspace() [20]bool {
	return p.espace
}

func (p *Parking) SetEspace(espace [20]bool) {
	p.espace = espace
}
