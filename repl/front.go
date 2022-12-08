package repl

import (
	"labs/cli"
)

type User struct {
	CmdCount int
	Name     string
	CmdHist  []string
}

func NewUser() *User {
	h := make([]string, 1, 150)

	var u User
	u.CmdCount = 1
	u.CmdHist = h
	u.CmdHist[0] = "begin"

	return &u
}

func (u *User) setName(name string) {
	u.Name = name
}

func (u *User) addCmd(cmd string) {
	u.CmdHist = append(u.CmdHist, cmd)
	u.CmdCount++
}

func repl(i *InOut, lab *Lab, usr *User) {
	curLine := StartInputLoop(i)

	input := *curLine

	DetermineCmd(lab, string(input), usr, i)

	repl(i, lab, usr)
}

func Run() {
	cli.Ready()

	term := cli.NewTerminal()
	term.Clear()

	welcome()

	inout := newInOut(term)

	lab := NewLab()
	usr := NewUser()

	repl(inout, lab, usr)
}