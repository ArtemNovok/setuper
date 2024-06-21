package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const (
	mainFile             = "package main\n\nfunc main(){\n\n}"
	dirCreatedMsg        = "Directory created: "
	mainFileCreatedMsg   = "file main.go created"
	mainFileFormattedMsg = "file main.go formatted"
	goModCreatedMsg      = "go.mod created with module name: "
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
	Cyan  = "\033[36m"
)

func main() {
	foldName := flag.String("n", "undefined", "name of the project folder and go.mod if name of mode file is not defined")
	GoModName := flag.String("m", "", "name of the go module")
	flag.Parse()
	SetUpProd(*foldName, *GoModName)
}
func PrintSuccessln(msg string) {
	fmt.Println(Green + msg + Reset)
}
func PrintBlueln(msg string) {
	fmt.Println(Cyan + msg + Reset)
}
func PrintSuccess(msg string) {
	fmt.Print(Green + msg + Reset)
}
func PrintErrorMsg(err error) {
	fmt.Println(Red + fmt.Sprintf("Got error: %s", err.Error()) + Reset)
}
func SetUpProd(foldName, GoModName string) {
	err := SetUp(foldName, GoModName)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			PrintErrorMsg(err)
			return
		} else {
			PrintErrorMsg(err)
			MustClear()
		}
	}
}
func SetUp(foldName, GoModName string) error {
	fmt.Println(foldName)
	if GoModName == "" {
		GoModName = foldName
	}
	err := os.Mkdir(foldName, os.ModePerm)
	if err != nil {
		return err
	}
	PrintSuccess(dirCreatedMsg)
	PrintBlueln(foldName)
	err = os.Chdir(foldName)
	if err != nil {
		return err
	}
	f, err := os.OpenFile("main.go", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	PrintSuccessln(mainFileCreatedMsg)
	_, err = f.WriteString(mainFile)
	if err != nil {
		return err
	}
	f.Close()
	PrintSuccessln(mainFileFormattedMsg)
	goMod := exec.Command("go", "mod", "init", GoModName)
	_, err = goMod.Output()
	if err != nil {
		return err
	}
	PrintSuccess(goModCreatedMsg)
	PrintBlueln(GoModName)
	return nil
}

func MustClear() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	PrintSuccessln(dir)
	err = os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
}
