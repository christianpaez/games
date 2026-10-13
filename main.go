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
	width  float64
	height float64
}

func (ball *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ball.x, ball.y)
	screen.DrawImage(ball.image, op)
}

type Bat struct {
	x, y  float64
	speed float64
	image *ebiten.Image
	width float64
}

func (bat *Bat) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bat.x, bat.y)
	screen.DrawImage(bat.image, op)
}

func (bat *Bat) Update(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.bat.x-float64(g.bat.speed) >= 0 {
		g.bat.x = g.bat.x - float64(g.bat.speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.bat.x+g.bat.speed <= float64(WIDTH)-g.bat.width {
		g.bat.x = g.bat.x + float64(g.bat.speed)
	}
}

func (g *Game) Update() error {
	g.bat.Update(g)
	ballPreviousBottomY := g.ball.y + g.ball.height
	g.ball.x += g.ball.dx * g.ball.speed

	g.ball.y += g.ball.dy * g.ball.speed

	if g.ball.x >= float64(WIDTH)-g.ball.width {
		g.ball.x = float64(WIDTH) - g.ball.width
		g.ball.dx *= -1
	}

	if g.ball.x <= 0 {
		g.ball.x = 0
		g.ball.dx *= -1
	}

	if ballPreviousBottomY <= g.bat.y &&
		g.ball.y+g.ball.height >= g.bat.y &&
		g.ball.x+g.ball.width > g.bat.x &&
		g.ball.x < g.bat.x+g.bat.width {
		g.ball.dy *= -1
	}

	if g.ball.y <= 0 {
		g.ball.dy *= -1
	}

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
		speed: 8,
		width: float64(batImage.Bounds().Dx()),
	}

	ball := Ball{
		image:  ballImage,
		x:      float64(WIDTH) / 2,
		y:      float64(HEIGHT) * 0.2,
		speed:  4,
		dx:     -1,
		dy:     1,
		width:  float64(ballImage.Bounds().Dx()),
		height: float64(ballImage.Bounds().Dy()),
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
