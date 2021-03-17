package console

import (
	"github.com/creack/pty"
	"io"
	"os"
	"os/exec"
)

func Run(name string, arg ...string)  {

	c := exec.Command(name, arg...)
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, f)
}
