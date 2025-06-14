package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
)

func expandTilde(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}
	usr, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr, path[1:]), nil
}


func execPython(scriptPath string) {
	if !fileExist(scriptPath) {
		log.Fatalln("one of .py scripts doesn't exist")
	}
	cmd := exec.Command("python", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			log.Fatalln("python script exited with error")
		}
		log.Fatalln("failed to execute Python script")
	}

}

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatalln("image doesn't exist through this path")
	}
	if err != nil {
		log.Fatalln("internal error")
	}
	return !info.IsDir()
}

func isImage(filetype string) bool {
	return slices.Contains(mimetypes, filetype)
}