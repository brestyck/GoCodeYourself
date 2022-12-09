package main

import (
	"github.com/gotk3/gotk3/gtk"
)

func yesIwant() (*gtk.Window, *gtk.Button) {
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 12)
	box.SetMarginTop(24)
	box.SetMarginBottom(24)
	box.SetMarginStart(24)
	box.SetMarginEnd(24)
	win.Add(box)
	label, _ := gtk.LabelNew("К сожалению, сейчас у нас нет слона, так что мы не сможем Вам его продать.")
	ohNo, _ := gtk.ButtonNewWithLabel("Жаль!")
	box.Add(label)
	box.Add(ohNo)
	return win, ohNo
}

func noIdont() (*gtk.Window, *gtk.Button, *gtk.Button) {
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 12)
	box.SetMarginTop(24)
	box.SetMarginBottom(24)
	box.SetMarginStart(24)
	box.SetMarginEnd(24)
	win.Add(box)
	label, _ := gtk.LabelNew("Может, Вы всё-таки хотите ли купить слона?")
	yBtn, _ := gtk.ButtonNewWithLabel("Да")
	nBtn, _ := gtk.ButtonNewWithLabel("Нет")
	box.Add(label)
	box.Add(yBtn)
	box.Add(nBtn)
	return win, yBtn, nBtn
}

func main() {
	gtk.Init(nil)

	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 12)
	box.SetMarginTop(24)
	box.SetMarginBottom(24)
	box.SetMarginStart(24)
	box.SetMarginEnd(24)

	label, _ := gtk.LabelNew("Вы хотите купить слона?")

	btnYes, _ := gtk.ButtonNewWithLabel("Да!")
	btnNo, _ := gtk.ButtonNewWithLabel("Нет!")

	var goAgain func()

	noWins := []*gtk.Window{}
	noWins = append(noWins, win)
	var checkmate func()

	doCheckmate := true

	checkmate = func() {
		isAct := false
		for _, n := range noWins {
			if n.IsVisible() {
				isAct = true
			}
		}
		if !isAct && doCheckmate {
			gtk.MainQuit()
		}
	}

	yiw := func() {
		doCheckmate = false
		for _, n := range noWins {
			if n != nil {
				n.Destroy()
			}
		}
		winOhNo, ohNoBtn := yesIwant()
		winOhNo.ShowAll()
		noWins = append(noWins, winOhNo)
		ohNoBtn.Connect("clicked", func() {
			gtk.MainQuit()
		})
		winOhNo.Connect("destroy", gtk.MainQuit)
	}

	goAgain = func() {
		winGoAgain, by, bn := noIdont()
		noWins = append(noWins, winGoAgain)
		winGoAgain.ShowAll()
		winGoAgain.Connect("destroy", checkmate)
		by.Connect("clicked", yiw)
		bn.Connect("clicked", goAgain)
	}

	btnYes.Connect("clicked", yiw)
	btnNo.Connect("clicked", goAgain)
	win.Connect("destroy", checkmate)
	box.Add(label)
	box.Add(btnNo)
	box.Add(btnYes)
	win.Add(box)
	win.SetTitle("Купи слона!")

	// win.Connect("destroy", nil)
	win.ShowAll()
	gtk.Main()
}
