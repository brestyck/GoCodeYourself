package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type Field struct {
	x rune
	y rune
}

var labelInfo string = "Сейчас ходит игрок X"
var winner byte = 0
var globTurn byte = 0

func turnCSS() {
	mRefProvider, _ := gtk.CssProviderNew()
	if err := mRefProvider.LoadFromPath("styles.css"); err != nil {
		log.Fatal(err)
	}
	screen, _ := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, mRefProvider, 1)
}

// Функция, определяющая ходящего игрока
func determineTurn() byte {
	switch globTurn {
	case 0:
		globTurn++
		return 1
	case 1:
		globTurn--
		return 2
	default:
		return 0
	}
}

// Функция проверки победы
func checkWinner(movesArray map[Field]byte, width rune, height rune, rowlen rune) bool {
	var yi rune
	var xi rune
	var rowcount rune
	var chain rune
	for yi = 0; yi < height; yi++ {
		for xi = 0; xi < (width-rowlen)+1; xi++ {
			if movesArray[Field{xi, yi}] != 0 {
				chain = 0
				if yi < (height-rowlen)+1 {
					// \ diagonal
					if xi < (width-rowlen)+1 {
						for rowcount = 0; rowcount <= rowlen; rowcount++ {
							if chain == rowlen {
								winner = movesArray[Field{xi, yi}]
								return true
							}
							if movesArray[Field{xi + rowcount, yi + rowcount}] == movesArray[Field{xi, yi}] {
								chain++
								continue
							}
							chain = 0
							break
						}
					}
					chain = 0
					// Vertical check
					for rowcount = 0; rowcount <= rowlen; rowcount++ {
						if chain == rowlen {
							winner = movesArray[Field{xi, yi}]
							return true
						}
						if movesArray[Field{xi, yi + rowcount}] == movesArray[Field{xi, yi}] {
							chain++
							continue
						}
						chain = 0
						break
					}
				}
				chain = 0
				// Horizontal check
				for rowcount = 0; rowcount <= rowlen; rowcount++ {
					if chain == rowlen {
						winner = movesArray[Field{xi, yi}]
						return true
					}
					if movesArray[Field{xi + rowcount, yi}] == movesArray[Field{xi, yi}] {
						chain++
						continue
						// If we found, continue
					}
					// One dismatch and we skip
					chain = 0
					break
				}
				// End of horizontal check
			}
		}
	}
	chain = 0
	// / diagonal
	for yi = 0; yi < (height-rowlen)+1; yi++ {
		for xi = width - 1; xi > rowlen-2; xi-- {
			if movesArray[Field{xi, yi}] != 0 {
				for rowcount = 0; rowcount <= rowlen; rowcount++ {
					if chain == rowlen {
						winner = movesArray[Field{xi, yi}]
						return true
					}
					if movesArray[Field{xi - rowcount, yi + rowcount}] == movesArray[Field{xi, yi}] {
						chain++
						continue
						// If we found, continue
					}
					// One dismatch and we skip
					chain = 0
					break
				}
			}
		}
	}
	return false
}

func ifTie(coordes map[Field]byte) (tie bool) {
	for _, n := range coordes {
		if n == 0 {
			tie = true
		}
	}
	if !tie {
		winner = 3
	}
	return !tie
}

// Функция вычисления ячейки хода
func calculateMoveCoords(coordes map[Field]byte, column rune, height rune) (Field, bool) {
	if winner == 0 {
		for i := height - 1; i >= 0; i-- {
			if coordes[Field{column, i}] == 0 {
				return Field{column, i}, true
			}
		}
	}
	return Field{0, 0}, false
}

// Возвращает символ игрока по номеру ходящего
func getMoveLiteral(moveNum byte) string {
	switch moveNum {
	case 1:
		return "X"
	case 2:
		return "O"
	default:
		return ""
	}
}

func labelWorker() {
	switch winner {
	case 1:
		labelInfo = "Игрок X победил"
	case 2:
		labelInfo = "Игрок O победил"
	case 3:
		labelInfo = "Ничья"
	default:
		labelInfo = fmt.Sprintf("Сейчас ходит игрок %s", getMoveLiteral(globTurn+1))
	}
}

func createArea(height rune, width rune, mWind *gtk.Window) *gtk.Box {
	fieldCoordesArray := make(map[glib.Object]Field)
	fieldMovesArray := make(map[Field]byte)
	buttonsHandlers := make(map[Field]*gtk.Button)
	mainArea, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	mainArea.SetMarginBottom(20)
	mainArea.SetMarginEnd(20)
	mainArea.SetMarginStart(20)
	mainArea.SetMarginTop(20)
	var i rune
	var j rune
	for i = 0; i < height; i++ {
		row, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 4)
		for j = 0; j < width; j++ {
			field, _ := gtk.ButtonNewWithLabel("")
			fieldCoordesArray[*field.Object] = Field{j, i}
			buttonsHandlers[Field{j, i}] = field
			fieldMovesArray[Field{j, i}] = 0

			field.Connect("clicked", func(b *gtk.Button) {
				col := fieldCoordesArray[*b.Object].x
				pos, ok := calculateMoveCoords(fieldMovesArray, col, height)
				if ok {
					turn := determineTurn()
					fieldMovesArray[pos] = turn
					buttonsHandlers[pos].SetLabel(getMoveLiteral(turn))
					buttonsHandlers[pos].SetName(fmt.Sprint(turn))
					checkWinner(fieldMovesArray, width, height, 4)
					ifTie(fieldMovesArray)
					labelWorker()
				}
			})

			field.SetSizeRequest(100, 100)
			row.Add(field)
		}
		mainArea.Add(row)
	}
	return mainArea
}

func main() {
	gtk.Init(nil)
	var InitializeMenu func() *gtk.Window
	InitializeMenu = func() *gtk.Window {
		win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
		outerBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
		outerBox.SetMarginBottom(20)
		outerBox.SetMarginEnd(20)
		outerBox.SetMarginTop(20)
		outerBox.SetMarginStart(20)
		innerUpward, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
		innerDownward, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
		goBtn, _ := gtk.ButtonNewWithLabel("GO")
		goBtn.SetName("InitMenu")
		heiDesc, _ := gtk.LabelNew("Height")
		heiSelect, _ := gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, 4, 10, 1)
		widDesc, _ := gtk.LabelNew("Width")
		widSelect, _ := gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, 4, 10, 1)
		heiSelect.SetSizeRequest(120, 40)
		widSelect.SetSizeRequest(120, 40)
		heiSelect.SetValue(6)
		widSelect.SetValue(7)
		innerUpward.Add(heiDesc)
		innerUpward.Add(heiSelect)
		innerUpward.Add(widDesc)
		innerUpward.Add(widSelect)
		innerDownward.Add(goBtn)
		outerBox.Add(innerUpward)
		outerBox.Add(innerDownward)
		win.Add(outerBox)
		goBtn.Connect("clicked", func() {
			turnCSS()
			mainWindow, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
			mainWindow.SetTitle("4 In A Row")
			upperLabel, _ := gtk.LabelNew("This is an upper label")
			upperLabel.SetName("upper")
			glib.IdleAdd(func() bool { upperLabel.SetLabel(labelInfo); upperLabel.SetName(fmt.Sprint(winner)); return true })
			outerBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
			outerBox.Add(upperLabel)
			areaH := rune(heiSelect.GetValue())
			areaW := rune(widSelect.GetValue())

			// Main Area
			area := createArea(areaH, areaW, mainWindow)

			outerBox.Add(area)
			mainWindow.Add(outerBox)
			mainWindow.Connect("destroy", gtk.MainQuit)
			mainWindow.ShowAll()
			win.Hide()
		})
		win.Connect("destroy", gtk.MainQuit)
		return win
	}
	initi := InitializeMenu()
	initi.ShowAll()

	gtk.Main()
}
