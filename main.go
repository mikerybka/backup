package main

import (
	"fmt"
	"os"

	"github.com/mikerybka/util"
)

func main() {
	key := os.Args[1]
	src := os.Args[2]
	dst := os.Args[3]
	err := util.BackupDir(src, key, dst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
