package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/smallgamekit"
	"github.com/seanoneillcode/sudoku/common"
)

const boardPosX = 16
const boardPosY = 16
const charOffsetX = 2
const charOffsetY = 3
const tileSize = 16
const numTiles = 9

type Board struct {
	background   *ebiten.Image
	textRenderer smallgamekit.TextRenderer
	numbers      []string
}

func NewBoard(game *Game) *Board {
	return &Board{
		background:   game.imageRes.GetImage("background"),
		textRenderer: game.textRenderer,
		numbers: []string{
			"2", "", "", "1", "", "", "3", "", "",
			"", "", "", "", "2", "", "4", "", "",
			"", "", "", "", "", "7", "", "", "",

			"1", "2", "", "3", "", "", "2", "", "",
			"", "", "", "", "", "5", "8", "9", "",
			"", "", "", "", "2", "", "", "3", "",

			"4", "7", "", "", "", "", "3", "", "",
			"5", "6", "", "", "", "", "4", "7", "",
			"", "", "", "", "", "", "5", "1", ""},
	}
}

func (r *Board) Draw(screen *ebiten.Image) {
	// draw back panels
	// draw grid
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(r.background, op)

	// draw numbers
	for i, v := range r.numbers {
		xi := i % numTiles
		yi := i / numTiles
		xpos := (xi * tileSize) + charOffsetX + boardPosX
		ypos := (yi * tileSize) + charOffsetY + boardPosY
		r.textRenderer.DrawText(screen, v, xpos, ypos, common.TextScale)
	}

	// draw selections
}
