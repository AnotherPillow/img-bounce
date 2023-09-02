package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lxn/win"
)

var img *ebiten.Image
var scr_width int = int(win.GetSystemMetrics(win.SM_CXSCREEN))
var scr_height int = int(win.GetSystemMetrics(win.SM_CYSCREEN))
var step = 5

func randint(min, max int) int {
	return min + rand.Intn(max-min)
}

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	posX        int
	posY        int
	increasingX bool
	increasingY bool
}

func (g *Game) Update() error {
	rand.Seed(time.Now().UnixNano())

	//move the window
	if g.increasingX {
		g.posX += step
	} else {
		g.posX -= step
	}

	if g.increasingY {
		g.posY += step
	} else {
		g.posY -= step
	}

	//get the *screen* width and height

	//if the window is out of the screen, reset the position
	if g.posX > scr_width {
		g.increasingX = false
		g.posX = scr_width - 10
	} else if g.posX <= 0 {
		g.increasingX = true
		g.posX = 1
	}

	if g.posY > scr_height {
		g.increasingY = false
		g.posY = scr_height - 10
	} else if g.posY <= 0 {
		g.increasingY = true
		g.posY = 1
	}

	// fmt.Println("X: ", g.posX, " Y: ", g.posY, " scr_width: ", scr_width, " scr_height: ", scr_height)

	ebiten.SetWindowPosition(g.posX, g.posY)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// screen.DrawImage(img, nil)
	//fill transparent
	// screen.Fill(color.RGBA{1, 1, 1, 128})
	screen.Fill(color.NRGBA{1, 1, 1, 128})

	//draw the image, transparent bg
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(0, 0)
	// op.ColorM.Scale(1.0, 1.0, 1.0, 0.5)

	screen.DrawImage(img, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	//return 640, 480
	return 400, 400
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	//transparency
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowPosition(0, 0)

	if err := ebiten.RunGameWithOptions(&Game{}, &ebiten.RunGameOptions{ScreenTransparent: true}); err != nil {
		log.Fatal(err)
	}
}
