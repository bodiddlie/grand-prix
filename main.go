package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	// Update the game state here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the screen with a blue color
	screen.Fill(color.RGBA{0, 0, 255, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}