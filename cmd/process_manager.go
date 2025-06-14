package main

import (
	"log"
	"os/exec"
	"time"
)


func restartHyprpaper() {
	log.Println("restarting hyprpaper...")

	killCmd := exec.Command("pkill", "hyprpaper")
	killOutput, killErr := killCmd.CombinedOutput()
	if killErr != nil {
		log.Printf("warning: failed to kill hyprpaper via pkill: %v\noutput: %s", killErr, killOutput)
	} else {
		log.Println("hyprpaper killed.")
		time.Sleep(500 * time.Millisecond)
	}

	startCmd := exec.Command("hyprpaper")
	err := startCmd.Start()
	if err != nil {
		log.Printf("error starting hyprpaper: %v", err)
	} else {
		log.Println("hyprpaper started.")
	}
}

func restartSwaync() {
	log.Println("restarting swaync...")

	killCmd := exec.Command("pkill", "swaync")
	killOutput, killErr := killCmd.CombinedOutput()
	if killErr != nil {
		log.Printf("warning: failed to kill swaync via pkill: %v\noutput: %s", killErr, killOutput)
	} else {
		log.Println("swaync killed.")
		time.Sleep(500 * time.Millisecond)
	}

	startCmd := exec.Command("swaync")
	err := startCmd.Start()
	if err != nil {
		log.Printf("error starting swaync: %v", err)
	} else {
		log.Println("swaync started.")
	}
}