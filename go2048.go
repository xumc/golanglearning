package main

import (
	"github.com/tncardoso/gocurses"
)

func main() {
	gocurses.Initscr()
	defer gocurses.End()
	gocurses.Cbreak()
	gocurses.Noecho()
	gocurses.Stdscr.Keypad(true)

	gocurses.Attron(gocurses.A_BOLD)
	gocurses.Addstr("Hello World!")
	gocurses.Refresh()

	wind := gocurses.NewWindow(10,40,10,10)
	wind.Box(0,0)
	wind.Refresh()

	gocurses.Stdscr.Getch()

	wind.Mvaddstr(2, 10, "this is moved cur")
	wind.Refresh()

	gocurses.Stdscr.Getch()
	wind.Mvaddstr(3, 10, "this is moved cur")
	wind.Refresh()

	gocurses.Stdscr.Getch()

}