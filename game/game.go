package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/smallgamekit"
	"github.com/seanoneillcode/sudoku/common"
)

type Game struct {
	textRenderer smallgamekit.TextRenderer
	imageRes     *smallgamekit.ImageResources

	board *Board
}

func NewGame() *Game {
	g := &Game{
		textRenderer: *smallgamekit.NewTextRenderer(),
		imageRes:     smallgamekit.NewImageResources("assets/images/"),
	}
	g.board = NewBoard(g)
	return g
}

func (r *Game) Update() error {
	r.board.Update()
	return nil
}

func (r *Game) Draw(screen *ebiten.Image) {
	r.board.Draw(screen)
}

func (r *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth * common.Scale, common.ScreenHeight * common.Scale
}
