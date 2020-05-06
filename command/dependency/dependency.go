package command

import (
	"os/exec"
)

func AreAllReady() bool {
	var areAllReady = true

	areAllReady = isApktoolInstalled()
	if areAllReady != true {
		return false
	}

	return true
}


func isApktoolInstalled() bool {
	_, err := exec.LookPath("apktool")
	if err != nil {
		panic("Didn't find 'apktool' executable. Install instructions is here: https://ibotpeaches.github.io/Apktool/install/")
		return false
	}

	return true
}