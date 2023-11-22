package scenes

import (
    "estacionamiento/models"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "gonum.org/v1/gonum/stat/distuv"
    "sync"
    "time"
)

type MainScene struct {
    window fyne.Window
}

func NewMainScene(window fyne.Window) *MainScene {
    return &MainScene{
        window: window,
    }
}

var contenedor = container.NewWithoutLayout()

func (s *MainScene) Show() {
    imagenContorno := canvas.NewImageFromFile("./assets/estacionamiento.jpg")
    imagenContorno.FillMode = canvas.ImageFillContain
    imagenContorno.Resize(fyne.NewSize(1080, 720))
    imagenContorno.Move(fyne.NewPos(0, -80))

    contenedor.Add(imagenContorno)
    s.window.SetContent(contenedor)
}

func (s *MainScene) Run() {
    p := models.NewEstacionamiento(make(chan int, 20), &sync.Mutex{})

    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            auto := models.NewVehicle(id)
            imagen := auto.GetImagenEntrada()
            imagen.Resize(fyne.NewSize(60, 100))
            imagen.Move(fyne.NewPos(90, -10))

            contenedor.Add(imagen)
            contenedor.Refresh()

            auto.Iniciar(p, contenedor, &wg)
        }(i)

        var poisson = poisson(float64(2))
        time.Sleep(time.Second * time.Duration(poisson))
    }

    wg.Wait()
}

func poisson(lambda float64) float64 {
    poisson := distuv.Poisson{Lambda: lambda, Src: nil}
    return poisson.Rand()
}
