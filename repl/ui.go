package repl

import (
	"fmt"
	"strings"
)

const (
	LINELOGO = "G-o-[-L-@-ß-$-] # "

	CLINELOGO = "\033[36;1;4;9mG\033[0;36;4;9m-\033[0;35;4;1mo\033[0;36;2;4;9m-\033[0;36;4;9m[\033[0;36;1;4;9m-\033[37;9;1mL\033[0;36;4;9m-\033[0;35;1;9m@\033[0;36;1;4;9m-\033[37;4;1mß\033[0;36;4;9m-\033[37;1;4m$\033[0;36;2;4;9m-\033[0;36;4;9;m]\033[0m \033[35;1;5m#\033[0m "

	LOGO = "\033[36;4;9m  __|  _ \\ |      \\   _ )  __|\033[0m\n\r\033[35;4;9m (_ | (   ||     _ \\  _ \\__ \\ \033[0m\n\r\033[36;1;4;9m\\___|\\___/____|_/  _\\___/____/\033[0m\n\r"
)

func printLineLogo(i *InOut) {
	fmt.Print(CLINELOGO)
	i.term.Cursor.AddX(len(LINELOGO))
}

func namePrompt(i *InOut) {
	fmt.Print("Enter UserName :  ")
	i.term.Cursor.AddX(len("Enter UserName :  "))
}

func logo(i *InOut) {
	fmt.Println(LOGO)

	parts := strings.Split(LOGO, "\n")

	i.term.Cursor.AddY(len(parts) + 1)
	i.AddLines(len(parts) + 1)
}

func PrintBuffer(i *InOut) {
	fmt.Print(
		string([]byte(i.lines[i.term.Cursor.Y])[i.term.Cursor.X-len(LINELOGO):]),
	)
}

func Refresh(i *InOut) {
	i.term.Cursor.CutRest()
	PrintBuffer(i)

	i.term.Cursor.MoveTo(
		i.term.Cursor.X+1,
		i.term.Cursor.Y,
	)
}

func Backspace() {
	print("\b")
}
