package models

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/storage"
    "math/rand"
    "sync"
    "time"
    "fmt"
)

type Vehicle struct {
    id              int
    tiempoLim       time.Duration
    espacioAsignado int
    imagenEntrada   *canvas.Image
    imagenSalida    *canvas.Image 
}

func NewVehicle(id int) *Vehicle {
    imagenEntrada := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
    imagenSalida := canvas.NewImageFromURI(storage.NewFileURI("./assets/carLeft.png"))
    return &Vehicle{
        id:              id,
        tiempoLim:       time.Duration(rand.Intn(20)+20) * time.Second,
        espacioAsignado: 0,
        imagenEntrada:   imagenEntrada,
        imagenSalida:    imagenSalida,
    }
}

func (a *Vehicle) Entrar(p *Estacionamiento, contenedor *fyne.Container) {
    p.GetEspacios() <- a.GetId()
    p.GetDoor().Lock()

    espacios := p.GetEspace()
    const (
        columna   = 10
        espaceFilas  = 10
        espaceHorizontal  = 70
        espaceVertical    = 100
    )

    time.Sleep(time.Millisecond * 1500)
    for i := 0; i < len(espacios); i++ {
        if !espacios[i] {
            espacios[i] = true
            a.espacioAsignado = i

            fila := i / columna

            x := float32(320 + (i%columna)*espaceHorizontal)
            y := float32(180 + fila*espaceFilas + (i/columna)*espaceVertical)

            a.imagenEntrada.Move(fyne.NewPos(x, y))
            break
        }
    }
    fmt.Printf("Vehículo %d Entrando al estacionamiento\n", a.GetId())
    p.Setespace(espacios)
    p.GetDoor().Unlock()
    contenedor.Refresh()
}

func (a *Vehicle) Salir(p *Estacionamiento, contenedor *fyne.Container) {
    <-p.GetEspacios()
    p.GetDoor().Lock()

    spacesArray := p.GetEspace()
    spacesArray[a.espacioAsignado] = false
    p.Setespace(spacesArray)

    p.GetDoor().Unlock()

    x := a.imagenEntrada.Position().X
    y := a.imagenEntrada.Position().Y

    contenedor.Remove(a.imagenEntrada)
    a.imagenSalida.Resize(fyne.NewSize(60, 100))
    a.imagenSalida.Move(fyne.NewPos(x, y))

    contenedor.Add(a.imagenSalida)
    contenedor.Refresh()

    for i := 0; i < 10; i++ {
        a.imagenSalida.Move(fyne.NewPos(a.imagenSalida.Position().X, a.imagenSalida.Position().Y-30))
        time.Sleep(time.Millisecond * 200)
    }
    fmt.Printf("Vehículo %d Saliendo del estacionamiento\n", a.GetId())
    contenedor.Remove(a.imagenSalida)
    contenedor.Refresh()
}


func (a *Vehicle) Iniciar(p *Estacionamiento, contenedor *fyne.Container, wg *sync.WaitGroup) {
    a.Avanzar(6) 
    a.Entrar(p, contenedor)
    time.Sleep(a.tiempoLim)

    contenedor.Remove(a.imagenEntrada)

    a.Salir(p, contenedor)

    wg.Done()
}

func (a *Vehicle) Avanzar(pasos int) {
    for i := 0; i < pasos; i++ {
        a.imagenEntrada.Move(fyne.NewPos(a.imagenEntrada.Position().X, a.imagenEntrada.Position().Y+20))
        time.Sleep(time.Millisecond * 200)
    }
}


func (a *Vehicle) GetId() int {
    return a.id
}

func (a *Vehicle) GetImagenEntrada() *canvas.Image {
    return a.imagenEntrada
}
