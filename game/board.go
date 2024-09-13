package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/smallgamekit"
	"github.com/seanoneillcode/sudoku/common"

	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const boardPosX = 16
const boardPosY = 16
const charOffsetX = 2
const charOffsetY = 3
const tileSize = 16
const tileSizeHalf = 8
const numTiles = 9

type Board struct {
	background        *ebiten.Image
	tileSelectedImage *ebiten.Image
	textRenderer      smallgamekit.TextRenderer
	numbers           []string
	selectedTileIndex int
}

func NewBoard(game *Game) *Board {
	return &Board{
		background:        game.imageRes.GetImage("background"),
		tileSelectedImage: game.imageRes.GetImage("tile-selected"),
		textRenderer:      game.textRenderer,
		numbers: []string{
			"2", "", "", "1", "", "", "3", "", "",
			"", "", "", "", "2", "", "4", "", "",
			"", "", "", "", "", "7", "", "", "",

			"1", "2", "", "3", "", "", "2", "", "",
			"", "", "", "", "", "5", "8", "9", "",
			"", "", "", "", "2", "", "", "3", "",

			"4", "7", "", "", "", "", "3", "", "",
			"5", "6", "", "", "", "", "4", "7", "",
			"", "", "", "", "", "", "5", "1", "",
		},
		selectedTileIndex: 15,
	}
}

func (r *Board) Draw(screen *ebiten.Image) {
	// draw background
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(common.Scale, common.Scale)
	screen.DrawImage(r.background, op)

	// draw selection
	if r.selectedTileIndex != -1 {
		op := &ebiten.DrawImageOptions{}
		xi := r.selectedTileIndex % numTiles
		yi := r.selectedTileIndex / numTiles
		xpos := (xi * tileSize) + boardPosX
		ypos := (yi * tileSize) + boardPosY
		op.GeoM.Translate(float64(xpos), float64(ypos))

		op.GeoM.Scale(common.Scale, common.Scale)
		screen.DrawImage(r.tileSelectedImage, op)
	}

	// draw numbers
	for i, v := range r.numbers {
		xi := i % numTiles
		yi := i / numTiles
		xpos := (xi * tileSize) + charOffsetX + boardPosX
		ypos := (yi * tileSize) + charOffsetY + boardPosY
		r.textRenderer.DrawText(screen, v, xpos, ypos, common.TextScale)
	}
}

func (r *Board) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		mx, my := MousePos()

		r.selectedTileIndex = -1
		log.Printf("mx: %v, my: %v", mx, my)
		if mx > boardPosX && mx < boardPosX+(tileSize*numTiles) {
			if my > boardPosY && my < boardPosY+(tileSize*numTiles) {
				log.Printf("clicked on board!")
				mx = mx - boardPosX
				my = my - boardPosY
				xi := int(mx / tileSize)
				yi := int(my / tileSize)
				r.selectedTileIndex = (yi * 9) + xi
			}
		}
	}
}

func MousePos() (float64, float64) {
	x, y := ebiten.CursorPosition()
	return float64(x / common.Scale), float64(y / common.Scale)
}
