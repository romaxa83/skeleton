package main

import (
	"github.com/romaxa83/skeleton/internal/app"
)

func main() {




	//cmd := exec.Command("touch", "test1.txt")
	//cmd.Stdin = strings.NewReader("some input")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("in all caps: %q\n", out.String())


	config := app.InitConfig()
	app.Create(config)
}