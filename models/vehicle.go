package models

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/storage"
    "math/rand"
    "sync"
    "time"
)

// Definición de la estructura Vehicle
type Vehicle struct {
    id              int
    tiempoLim       time.Duration
    espacioAsignado int
    imagenEntrada   *canvas.Image
    imagenSalida    *canvas.Image // Usamos solo la imagen de salida
}

// Función NewVehicle crea una nueva instancia de Vehicle
func NewVehicle(id int) *Vehicle {
    // Carga imágenes desde archivos y establece campos iniciales
    imagenEntrada := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
    imagenSalida := canvas.NewImageFromURI(storage.NewFileURI("./assets/car_salida.png"))
    return &Vehicle{
        id:              id,
        tiempoLim:       time.Duration(rand.Intn(20)+20) * time.Second,
        espacioAsignado: 0,
        imagenEntrada:   imagenEntrada,
        imagenSalida:    imagenSalida,
    }
}

// Función Entrar permite que el Vehiclemóvil entre al estacionamiento
func (a *Vehicle) Entrar(p *Estacionamiento, contenedor *fyne.Container) {
    p.GetEspacios() <- a.GetId()
    p.GetPuertaMu().Lock()

    espacios := p.GetEspaciosArray()
    const (
        columnasPorGrupo   = 10
        espacioEntreFilas  = 10
        espacioHorizontal  = 70
        espacioVertical    = 100
    )

    time.Sleep(time.Millisecond * 1500)
    for i := 0; i < len(espacios); i++ {
        if !espacios[i] {
            espacios[i] = true
            a.espacioAsignado = i

            fila := i / columnasPorGrupo

            x := float32(320 + (i%columnasPorGrupo)*espacioHorizontal)
            y := float32(180 + fila*espacioEntreFilas + (i/columnasPorGrupo)*espacioVertical)

            a.imagenEntrada.Move(fyne.NewPos(x, y))
            break
        }
    }

    p.SetEspaciosArray(espacios)
    p.GetPuertaMu().Unlock()
    contenedor.Refresh()
}

// Función Salir permite que el Vehiclemóvil salga del estacionamiento
func (a *Vehicle) Salir(p *Estacionamiento, contenedor *fyne.Container) {
    <-p.GetEspacios()
    p.GetPuertaMu().Lock()

    spacesArray := p.GetEspaciosArray()
    spacesArray[a.espacioAsignado] = false
    p.SetEspaciosArray(spacesArray)

    p.GetPuertaMu().Unlock()

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

    contenedor.Remove(a.imagenSalida)
    contenedor.Refresh()
}


// Función Iniciar inicia el proceso del Vehiclemóvil en el estacionamiento
func (a *Vehicle) Iniciar(p *Estacionamiento, contenedor *fyne.Container, wg *sync.WaitGroup) {
    a.Avanzar(6) // Realiza una animación de avance

    a.Entrar(p, contenedor) // El Vehiclemóvil entra en el estacionamiento

    time.Sleep(a.tiempoLim) // Espera el tiempo límite

    contenedor.Remove(a.imagenEntrada)

    a.Salir(p, contenedor) // El Vehiclemóvil sale del estacionamiento

    wg.Done() // Indica que el Vehiclemóvil ha terminado su proceso
}

// Función Avanzar realiza una animación de avance del Vehiclemóvil
func (a *Vehicle) Avanzar(pasos int) {
    for i := 0; i < pasos; i++ {
        a.imagenEntrada.Move(fyne.NewPos(a.imagenEntrada.Position().X, a.imagenEntrada.Position().Y+20))
        time.Sleep(time.Millisecond * 200)
    }
}

// Función GetId devuelve el identificador del Vehiclemóvil
func (a *Vehicle) GetId() int {
    return a.id
}

// Función GetImagenEntrada devuelve la imagen de entrada del Vehiclemóvil
func (a *Vehicle) GetImagenEntrada() *canvas.Image {
    return a.imagenEntrada
}
