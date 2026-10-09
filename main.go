package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	WIDTH  = 640
	HEIGHT = 480
)

type Game struct {
	ball Ball
	bat  Bat
}

type Ball struct {
	x, y   float64
	dx, dy float64
	speed  float64
	image  *ebiten.Image
}

func (ball *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ball.x, ball.y)
	screen.DrawImage(ball.image, op)
}

type Bat struct {
	x, y  float64
	speed uint8
	image *ebiten.Image
}

func (bat *Bat) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bat.x, bat.y)
	screen.DrawImage(bat.image, op)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.bat.x = g.bat.x - float64(g.bat.speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.bat.x = g.bat.x + float64(g.bat.speed)
	}

	g.ball.x += g.ball.dx * g.ball.speed

	g.ball.y += g.ball.dy * float64(g.ball.speed)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ball.Draw(screen)
	g.bat.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return WIDTH, HEIGHT
}

func main() {
	ballImage, _, err := ebitenutil.NewImageFromFile("assets/ball.png")
	if err != nil {
		log.Fatal(err)
	}
	batImage, _, err := ebitenutil.NewImageFromFile("assets/paddle.png")
	if err != nil {
		log.Fatal(err)
	}

	bat := Bat{
		image: batImage,
		x:     float64(WIDTH) / 2.5,
		y:     float64(HEIGHT) * 0.8,
		speed: 4,
	}

	ball := Ball{
		image: ballImage,
		x:     float64(WIDTH) / 2,
		y:     float64(HEIGHT) * 0.2,
		speed: 2,
		dx:    1,
		dy:    1,
	}
	game := &Game{
		ball: ball,
		bat:  bat,
	}

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Pong")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
