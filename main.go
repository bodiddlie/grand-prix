package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	Curves []*CurveCard
}

func (g *Game) Update() error {
	// Update the game state here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})

	for _, card := range g.Curves {
		card.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 800
}

func main() {
	// Set the window size to match the logical size and ensure aspect ratio is preserved
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowTitle("Steam Deck Resolution Game")
	game := &Game{}

	// Initialize cards with their respective images
	game.Curves = []*CurveCard{
		NewCurveCard("sweepers", 100, 100),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Sprite interface {
	Draw(screen *ebiten.Image)
	Update()
}

type CurveCard struct {
	cardImg *ebiten.Image
	x, y, w, h int
}

func NewCurveCard(cardName string, x, y int) *CurveCard {
	// Load the card image here
	cardImg, _, err := ebitenutil.NewImageFromFile("assets/" + cardName + ".png")
	if err != nil {
		log.Fatal(err)
	}

	return &CurveCard{
		cardImg: cardImg,
		x:       x,
		y:       y,
		w:       cardImg.Bounds().Dx(),
		h:       cardImg.Bounds().Dy(),
	}
}

func (c *CurveCard) Draw(screen *ebiten.Image) {
	// Draw the card image at the specified position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x), float64(c.y))
	screen.DrawImage(c.cardImg, op)
}

func (c *CurveCard) Update() {
	// Update the card's position or state here if needed
	// For example, you could animate the card or handle user input
}