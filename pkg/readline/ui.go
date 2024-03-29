package readline

import "strings"

const (
	INPROMPT  = "G-o-[-L-@-ß-$-] <> "
	CINPROMPT = "\033[36;1;4;9mG\033[0;36;4;9m-\033[0;35;4;1mo\033[0;36;2;4;9m-\033[0;36;4;9m[\033[0;36;1;4;9m-\033[37;9;1mL\033[0;36;4;9m-\033[0;35;1;9m@\033[0;36;1;4;9m-\033[37;4;1mß\033[0;36;4;9m-\033[37;1;4m$\033[0;36;2;4;9m-\033[0;36;4;9;m]\033[0m \033[35;1;9;5m<>\033[0m "

	LOGO  = "  __|  _ \\ |      \\   _ )  __|\n\r (_ | (   ||     _ \\  _ \\__ \\ \n\r\\___|\\___/____|_/  _\\___/____/\n\r"
	CLOGO = "\033[36;4;9m  __|  _ \\ |      \\   _ )  __|\033[0m\n\r\033[35;4;9m (_ | (   ||     _ \\  _ \\__ \\ \033[0m\n\r\033[36;1;4;9m\\___|\\___/____|_/  _\\___/____/\033[0m\n\r"

	ANDPROMPT  = "G-o-[-L-@-ß-$-] << "
	CANDPROMPT = "\033[36;1;4;9mG\033[0;36;4;9m-\033[0;35;4;1mo\033[0;36;2;4;9m-\033[0;36;4;9m[\033[0;36;1;4;9m-\033[37;9;1mL\033[0;36;4;9m-\033[0;35;1;9m@\033[0;36;1;4;9m-\033[37;4;1mß\033[0;36;4;9m-\033[37;1;4m$\033[0;36;2;4;9m-\033[0;36;4;9;m]\033[0m \033[34;1;9;5m<<\033[0m "

	OUTPROMPT  = ">>"
	COUTPROMPT = "\033[35;9;1m<<\033[0m "
)

func Logo(i *Input) {
	print(CLOGO)

	Term.Cursor.AddY(len(strings.Split(LOGO, "\n\r")) + 1)
	i.AddLines(len(strings.Split(LOGO, "\n\r")))
}

func Backspace() {
	print("\b")
}

func Tab() {
	print("    ")
}
