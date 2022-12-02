// //////////////////
// НАЧАЛО ПРИМЕРА //
// //////////////////
package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func BtnLabelConvert(num int) string {
	return fmt.Sprintf("Btn %d", num)
}

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	if err != nil {
		log.Fatal(err)
	}

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	win.Add(box)

	btnAmount := 0

	positions := map[glib.Object]int{}

	var a func(b *gtk.Button)

	a = func(b *gtk.Button) {
		btnAmount += 1
		fmt.Println(b.GetChildren())

		newBtn, err := gtk.ButtonNewWithLabel(BtnLabelConvert(btnAmount))
		// fmt.Println("! Hit button with addr ", *b.Object)
		// fmt.Println("Set newbutton to addr ", &newBtn)
		// fmt.Println("Position of current button is ", positions[b])
		positions[*newBtn.Object] = positions[*b.Object] + 1
		// fmt.Println("New button position set to ", positions[newBtn])
		box.GetChildren().Foreach(func(btn any) {
			button, _ := btn.(*gtk.Widget)
			button_id := button.Object
			if positions[*button_id] >= positions[*newBtn.Object] {
				positions[*button_id]++
			}
		})
		if err != nil {
			log.Fatal(err)
		}
		newBtn.Connect("clicked", a)
		box.Add(newBtn)
		box.ReorderChild(newBtn, positions[*newBtn.Object])
		box.ShowAll()
	}

	b, _ := gtk.ButtonNewWithLabel("First")
	positions[*b.Object] = -1

	a(b)

	// Показываем окно
	win.ShowAll()
	// Цикл обработки событий
	gtk.Main()
}

///////////////////
// КОНЕЦ ПРИМЕРА //
///////////////////
