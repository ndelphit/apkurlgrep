package apktool

import (
	"os/exec"
)

func RunApktool(apk string, tempDir string) string {
	cmd := exec.Command("apktool","d", apk, "-o", tempDir, "-fq")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	return string(output)
}