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
	return 1280, 800
}

func main() {
	// Set the window size to match the logical size and ensure aspect ratio is preserved
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowTitle("Steam Deck Resolution Game")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}