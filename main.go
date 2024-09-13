package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/sudoku/common"
	"github.com/seanoneillcode/sudoku/game"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(common.ScreenWidth*common.Scale, common.ScreenHeight*common.Scale)
	ebiten.SetWindowTitle("Sudoku")
	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
