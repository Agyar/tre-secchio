package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"log"
	"os"
)

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Unable to init termbox: %s", err.Error()))
	}

	view := NewObjectListView()
	if err := view.Start(); err != nil {
		log.Fatalln(fmt.Sprintf("Unable to start view: %s", err.Error()))
	}

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	for {
		ev := <-eventQueue
		switch {
		case ev.Key == termbox.KeyArrowUp:
			view.Up()
		case ev.Key == termbox.KeyArrowDown:
			view.Down()
		case ev.Key == termbox.KeyEnter:
			view.Dive()
		case ev.Key == termbox.KeyBackspace2:
			view.Back()
		case ev.Ch == 'd':
			view.Download()
		case ev.Key == termbox.KeyEsc:
			termbox.Close()
			os.Exit(0)
		}
	}
}
