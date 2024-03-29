package readline

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/BuprenorphineKid/golabs/pkg/cli"
)

const ()

var Term = cli.NewTerminal()

// Application Input and Output meshed into a single construct that
// holds all the state relevent to working with the input and
// output.
type Input struct {
	reader      io.Reader
	writer      io.Writer
	Rbuf        rbuf
	Wbuf        wbuf
	Spbuf       spbuf
	Mvbuf       mvbuf
	Fbuf        fbuf
	Ctrlkey     chan string
	Lines       []line
	done        chan struct{}
	CntrlCode   chan int
	ScrollCount int
}

// Constructor function for the Input Output struct.
func NewInput() *Input {
	i := Input{
		writer: os.Stdout,
		reader: os.Stdin,
	}

	i.Rbuf = rbuf("")
	i.Wbuf = wbuf("")
	i.Spbuf = spbuf("")
	i.Mvbuf = mvbuf("")
	i.Fbuf = fbuf("")
	i.Ctrlkey = make(chan string)
	i.Lines = make([]line, 1, Term.Lines)
	i.CntrlCode = make(chan int)
	i.ScrollCount = 0

	return &i
}

// InOut method for Adding n amount of lines.
func (i *Input) AddLines(n int) {
	newlines := make([]line, n, n)

	for i := range newlines {
		newlines[i] = ""
	}

	if Term.Cursor.Y < len(i.Lines) {
		return
	}

	i.Lines = append(i.Lines, newlines...)
}

// The read method is used for recieving one byte of input at a
// time and appending it to the Filter buffer(Fbuf) unless
// it recieves an escape byte, in which case 2 more bytes
// will be read and then its entirety will be sent to the
// Filter buffer (Fbuf).
func (i *Input) read() {
	if Term.IsRaw != true {
		panic("Not able to enter into raw mode :(.")
	}

	var buf [1]byte

	_, err := i.reader.(*os.File).Read(buf[:])

	if err != nil {
		panic(err)
	}

	i.Fbuf = append(i.Fbuf, buf[:]...)

	for strings.HasPrefix(string(i.Fbuf), "\x1b") && len(i.Fbuf) < 3 {
		var buf [1]byte

		_, err := i.reader.(*os.File).Read(buf[:])

		if err != nil {
			panic(err)
		}

		i.Fbuf = append(i.Fbuf, buf[:]...)
	}

	if strings.HasPrefix(string(i.Fbuf), "~") {
		i.Fbuf = fbuf(strings.Trim(string(i.Fbuf), "~"))
	}
}

// The write method is used to write to the current line
// so that it can be processed by the output system.
func (i *Input) write(buf []byte) {
	if string(buf) == "" {
		return
	}

	i.Lines[Term.Cursor.Y-1] = i.Lines[Term.Cursor.Y-1].insert(buf, Term.Cursor.X)
}

func (i *Input) Scroll() {
	i.ScrollCount++

	for j := range i.Lines[6:] {
		Term.Cursor.MoveTo(len(INPROMPT), j+5)
		Term.Cursor.CutRest()

		if j == len(i.Lines[6:])-1 {
			return
		}

		fmt.Print(i.Lines[6:][j+i.ScrollCount])
	}
}

func (i *Input) ScrollBack() {
	i.ScrollCount--

	for j := range i.Lines[6:] {
		Term.Cursor.MoveTo(len(INPROMPT), j+5)
		Term.Cursor.CutRest()

		if j == len(i.Lines[6:])-1 {
			return
		}

		fmt.Print(i.Lines[6:][j+i.ScrollCount])
	}

}

// Most users will be more inclined to use Take() instead
// because it does the dirty work for you.
//
// Start Input Loop that after its done reading input and
// filling buffers, concurrently processes each.
func ReadLine(i *Input) *ReturnLine {
	for {
		select {
		case code := <-i.CntrlCode:
			if code == 0 {
				select {
				case c := <-i.CntrlCode:
					switch c {
					case 1:
					case 2:
						return nil
					}
				}
			}
		default:
		}

		i.done = make(chan struct{}, 1)
		i.read()

		var nlwg sync.WaitGroup
		nlwg.Add(1)

		var bufs = []buffer{&i.Fbuf, &i.Rbuf, &i.Spbuf, &i.Mvbuf, &i.Wbuf}
		go ProccessBuffers(bufs, i, &nlwg)

		nlwg.Wait()

		select {
		case <-i.done:
			return &ReturnLine{
				&i.Lines[Term.Cursor.Y-2],
				Term.Cursor.Y - 5,
			}
		default:
			out.SetLine(string(i.Lines[Term.Cursor.Y-1]))
		}
		out.Devices["main"].(Display).RenderLine()

	}
}
