package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func executeYosys(path string) {
	app := "./yosys"
	// app := "cat"
	// arg0 := []string{"-p", "\"prep -auto-top; aigmap; write_json larstestOutAigmap.json\"", path}
	arg0 := []string{"-p", "prep -auto-top; aigmap; write_json larstestOutAigmap.json", path}
	// app := "ls"
	// arg0 := "-la"
	// arg1 := "\"prep -auto-top; aigmap; write_json larstestOutAigmap.json\""
	// arg2 := path
	// app := "tr"
	// arg1 := "a-z"
	// arg2 := "A-Z"
	yosysCmd := exec.Command(app, arg0...) //, arg1, arg2)
	fmt.Println(yosysCmd)

	// out, err := yosysCmd.Output()
	// if err != nil {
	// 	log.Fatalf("cmd.Output() failed with '%s'\n", err)
	// }

	// print(string(out))

	// -p "prep -auto-top; aigmap; write_json larstestOutAigmap.json" ../yosys/tests/simple/larserik/counter.v

	yosysOut, _ := yosysCmd.StdoutPipe() //Making a pipe we can reseive responses

	err := yosysCmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	yosysBytes, _ := ioutil.ReadAll(yosysOut)
	// fmt.Println("> Yosys Svarer")
	fmt.Println(string(yosysBytes))
	yosysCmd.Wait()
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
	executeYosys("/yosys/tests/simple/larserik/counter.v")
	executeNetlistsvg()
	displaySvg()

}

// /Users/larserikskraastad/Documents/ntnu/master/program/yosys/./yosys -p "prep -auto-top; aigmap; write_json testOutJson.json" /Users/larserikskraastad/Documents/ntnu/master/program/yosys/tests/simple/larserik/counter.v
