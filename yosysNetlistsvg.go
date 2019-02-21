package main

import (
	"fmt"
	"log"
	"os/exec"
)

func executeYosys(path string) {
	app := "yosys/yosys"
	arg0 := []string{"-p", "prep -auto-top; aigmap; write_json larstestOutAigmap.json", path}
	yosysCmd := exec.Command(app, arg0...) //, arg1, arg2)

	// yosysOut, _ := yosysCmd.StdoutPipe() //Making a pipe we can reseive responses
	err := yosysCmd.Run()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	// yosysBytes, _ := ioutil.ReadAll(yosysOut)
	// fmt.Println("> Yosys Svarer")
	// fmt.Println(string(yosysBytes))
	// yosysCmd.Wait()
}

func executeNetlistsvg() {
	app := "node"
	arg := []string{"netlistsvg/bin/netlistsvg", "larstestOutAigmap.json", "-o", "larsern2Aigmap.svg"}
	netlistsvgCmd := exec.Command(app, arg...)

	err := netlistsvgCmd.Run()
	if err != nil {
		log.Fatalf("netlistsvgCmd.Run() failed with '%s'\n", err)
	}
}

func displaySvg() {
	app := "display"
	arg := "larsern2Aigmap.svg"
	displayCmd := exec.Command(app, arg)
	err := displayCmd.Run()

	if err != nil {
		log.Fatalf("displayCmd.Run() failed with '%s'\n", err)
	}
}

func main() {
	fmt.Println("Running")
	executeYosys("yosys/tests/simple/larserik/counter.v")
	executeNetlistsvg()
	displaySvg()
	fmt.Println("Done")

}

// /Users/larserikskraastad/Documents/ntnu/master/program/yosys/./yosys -p "prep -auto-top; aigmap; write_json testOutJson.json" /Users/larserikskraastad/Documents/ntnu/master/program/yosys/tests/simple/larserik/counter.v
