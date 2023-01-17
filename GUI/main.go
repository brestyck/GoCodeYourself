package main

import (
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// func csser() {
// 	mRefProvider, _ := gtk.CssProviderNew()
// 	if err := mRefProvider.LoadFromPath("kurwa.css"); err != nil {
// 		log.Fatal(err)
// 	}
// 	screen, _ := gdk.ScreenGetDefault()
// 	gtk.AddProviderForScreen(screen, mRefProvider, 1)

// }

func main() {
	gtk.Init(nil)

	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 50)

	cssData, _ := gtk.CssProviderNew()
	outer_b, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 50)
	scr, _ := gdk.ScreenGetDefault()
	btn, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	b1, _ := gtk.LabelNew("awidjwaiodjwioad")
	btn.Add(b1)
	sc_win, _ := gtk.ScrolledWindowNew(nil, nil)
	entry, _ := gtk.TextViewNew()
	entry.SetSizeRequest(150, 50)
	sc_win.Add(entry)
	outer_b.SetMarginBottom(24)
	outer_b.SetMarginTop(24)
	outer_b.SetMarginStart(24)
	outer_b.SetMarginEnd(24)
	sc_win.SetMinContentHeight(100)
	sc_win.SetMinContentWidth(200)

	box.Add(btn)
	box.Add(sc_win)
	outer_b.Add(box)

	err_lbl, _ := gtk.LabelNew("Watch your CSS")

	outer_b.Add(err_lbl)
	gtk.AddProviderForScreen(scr, cssData, 0)
	go func() {
		for {
			gtk.RemoveProviderForScreen(scr, cssData)
			buf, _ := entry.GetBuffer()
			txt, _ := buf.GetText(buf.GetIterAtLine(0), buf.GetEndIter(), true)
			err := cssData.LoadFromData(txt)
			if err == nil {
				gtk.AddProviderForScreen(scr, cssData, 0)
			} else {
				err_lbl.SetText(err.Error())
			}
			time.Sleep(time.Second * 3)
		}
	}()

	win.Add(outer_b)
	win.Connect("destroy", gtk.MainQuit)

	win.SetDefaultSize(400, 400)
	win.ShowAll()

	gtk.Main()
}
