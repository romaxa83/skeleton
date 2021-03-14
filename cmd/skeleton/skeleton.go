package main

import "github.com/romaxa83/skeleton/internal/app"

func main() {

	config := app.InitConfig()
	app.Create(config)
}


//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"os"
//)
//
//type Color string
//
////func init()  {
////	fmt.Println("init")
////}
//
//const (
//	ColorBlack  Color = "\u001b[30m"
//	ColorRed          = "\u001b[31m"
//	ColorGreen        = "\u001b[32m"
//	ColorYellow       = "\u001b[33m"
//	ColorBlue         = "\u001b[34m"
//	ColorReset        = "\u001b[0m"
//)
//
//func colorize(color Color, message string) {
//	fmt.Println(string(color), message, string(ColorReset))
//}
//
//type Config struct {
//	Applications []Application
//}
//
//type Application struct {
//	Name string
//	Versions []Version
//	Default string
//}
//
//type Version struct {
//	Version	string
//	Number	string
//}
//
//type ConfigSkeleton struct {
//	server string
//	php string
//	db string
//}
////
//func vers(data )  {
//
//	fmt.Println(data)
//
//	//for _, val := range data {
//	//	fmt.Println("[" + val.Number + "] " + val.Version + "[]")
//	//}
//}
//
//func main()  {
//

//	}
//
//	for _, value := range config.Applications{
//		fmt.Println("--" + value.Name + " ["+ value.Default +"]")
//
//		vers(value)
//
//
//		//fmt.Printf("%T\n", value)
//
//
//		//fmt.Printf("%T\n", value.Versions)
//		//for _, val := range value.Versions {
//		//	fmt.Println(val)
//		//	fmt.Println("[" + val.Number + "] " + val.Version + "[]")
//		//}
//
//		var input string
//		fmt.Scanln(&input)
//		//skeleton := ConfigSkeleton{value.Name: input}
//		fmt.Println(input)
//	}
//
//	fmt.Println()
//	fmt.Print(len(config.Applications))

	//useColor := flag.Bool("color", false, "display colorized output")
	//flag.Parse()
	//
	//if *useColor {
	//	colorize(ColorBlue, "Hello, DigitalOcean!")
	//	return
	//}
	//fmt.Println("Hello, DigitalOcean!")
	//
	//fmt.Print("Choice: \n")
	//fmt.Print("[1]: \n")
	//fmt.Print("[2]: \n")
	//fmt.Print("[3]: \n")
	//var input string
	//var one string
	//fmt.Scanln(&input)
	//fmt.Println(input)
	//
	//fmt.Print("Choice: \n")
	//fmt.Print("[4]: \n")
	//fmt.Print("[5]: \n")
	//fmt.Print("[6]: \n")
	//fmt.Scanln(&one)
	//fmt.Println(one)

//}


//package main
//
//import (
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//	list := tview.NewList().
//		AddItem("List item 1", "Some explanatory text", 'a', nil).
//		AddItem("List item 2", "Some explanatory text", 'b', nil).
//		AddItem("List item 3", "Some explanatory text", 'c', nil).
//		AddItem("List item 4", "Some explanatory text", 'd', nil).
//		AddItem("Quit", "Press to exit", 'q', func() {
//			app.Stop()
//		})
//	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
//		panic(err)
//	}
//}