package srv

import (
	"fmt"
	"os"
)

func TestGetWd() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("TestGetWd  here is err %+v\n", err)
	}
	fmt.Printf("TestGetWd here is path, %s\n", wd)
}

func OsT() {
	allFIles := make([]*os.File, 0)
	allFIles = append(allFIles, os.Stdin)
	allFIles = append(allFIles, os.Stdout)
	allFIles = append(allFIles, os.Stderr)

}
