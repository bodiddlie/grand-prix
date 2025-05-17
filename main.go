package main

import (
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	Curves []*CurveCard
	Hand []Renderable
}

func (g *Game) Update() error {
	// Update the game state here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255})

	for i, card := range g.Hand {
		x := 100 + i*150
		y := 100
		card.Render(screen, x, y)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 800
}

func main() {
	// Set the window size to match the logical size and ensure aspect ratio is preserved
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowTitle("Flat Out")

	curves := []*CurveCard{
		NewCurveCard("sweepers"),
		NewCurveCard("castle"),
		NewCurveCard("flatout"),
		NewCurveCard("lucky"),
		NewCurveCard("parabola"),
		NewCurveCard("troubled_water"),
		NewCurveCard("zorro"),
	}

	// Shuffle the list of curves
	rand.Shuffle(len(curves), func(i, j int) { curves[i], curves[j] = curves[j], curves[i] })

	// Take the top 3 curves and put them in the hand
	hand := make([]Renderable, 3)
	for i := 0; i < 3; i++ {
		hand[i] = curves[i]
	}
	game := &Game{
		Curves: curves,
		Hand: hand,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Renderable interface {
	Render(screen *ebiten.Image, x, y int)
}

type CurveCard struct {
	cardImg *ebiten.Image
}

func NewCurveCard(cardName string) *CurveCard {
	// Load the card image here
	cardImg, _, err := ebitenutil.NewImageFromFile("assets/" + cardName + ".png")
	if err != nil {
		log.Fatal(err)
	}

	return &CurveCard{
		cardImg: cardImg,
	}
}

func (c *CurveCard) Render(screen *ebiten.Image, x, y int) {
	// Draw the card image at the specified position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(c.cardImg, op)
}

func (c *CurveCard) Update() {
	// Update the card's position or state here if needed
	// For example, you could animate the card or handle user input
}