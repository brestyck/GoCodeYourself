package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type MenuEntry struct {
	label string
	Next  *menu
	// звездочка нужна чтобы реализовать
	// выход обратно
	Action func()
}

/*
Инициализатор менюшки
*/
func (m *menu) AddEntry(label string, next *menu, action func()) {
	if next == nil {
		next = m
	}
	m.entries = append(m.entries, MenuEntry{
		label:  label,
		Next:   next,
		Action: action,
	})
}

type menu struct {
	entries []MenuEntry
}

func (e MenuEntry) Use() *menu {
	if e.Action != nil {
		e.Action()
	}
	return e.Next
}

var consoleActive bool

func StopConsole() {
	consoleActive = false
}

func (m menu) StartConsole() {
	consoleActive = true
	for consoleActive {
		for i, entry := range m.entries {
			fmt.Printf("%v) %v\n", i+1, entry.label)
		}

		fmt.Println()
		fmt.Print("Choose entry pls: ")

		var x int
		for {
			if n, _ := fmt.Scan(&x); n == 1 && x >= 1 && x <= len(m.entries) {
				break
			}
			fmt.Println("Say again pls")
			fmt.Print("Choose entry pls: ")
		}

		nextMenu := m.entries[x-1].Use()
		m = *nextMenu
	}
}

func main() {
	var MainMenu, newGameMenu, optionsMenu menu

	MainMenu.AddEntry("New game", &newGameMenu, nil)
	MainMenu.AddEntry("Options", &optionsMenu, nil)
	MainMenu.AddEntry("Quit", &MainMenu, StopConsole)

	newGameMenu.AddEntry("Back", &MainMenu, nil)
	optionsMenu.AddEntry("Back", &MainMenu, nil)

	MainMenu.StartConsole()
}

func gtk_startup() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	if err != nil {
		log.Fatal(err)
	}

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.ShowAll()

	gtk.Main()
}
