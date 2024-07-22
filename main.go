package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/mikerybka/util"
)

func main() {
	intervalMinutes := util.RequireEnvVar("INTERVAL_MINUTES")
	min, err := strconv.Atoi(intervalMinutes)
	if err != nil {
		log.Fatal("invalid env var:", "INTERVAL_MINUTES:", err)
	}
	for {
		start := time.Now()
		fmt.Println("starting backup at", start)

		backup()

		end := time.Now()
		fmt.Println("backup complete in", end.Sub(start))
		fmt.Println("sleeping for", min, "minutes")

		time.Sleep(time.Minute * time.Duration(min))
	}
}

func backup() {
	cmd := exec.Command("rsync", "-avz", "--delete", "/from/", "/to/")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
