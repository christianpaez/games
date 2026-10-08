package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	ball *ebiten.Image
	bat  *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(100, 100)

	screen.DrawImage(g.ball, op)

	op.GeoM.Translate(500, 100)
	screen.DrawImage(g.bat, op)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return 640, 480
}

func main() {
	ball, _, err := ebitenutil.NewImageFromFile("assets/ball.png")
	bat, _, err := ebitenutil.NewImageFromFile("assets/paddle.png")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		ball: ball,
		bat:  bat,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pong")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
