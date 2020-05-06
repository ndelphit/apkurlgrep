/*
 with love by @delphit 5/2020
*/

package main

import (
	"apkurlgreps/command/apktool"
	dependency "apkurlgreps/command/dependency"
	"apkurlgreps/directory"
	"apkurlgreps/extractor"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)



func main() {

	parser := argparse.NewParser("apkurlgreps", "AndroidLinksExtractor")
	apk := parser.String("a", "apk", &argparse.Options{Required: true, Help: "Input a path to APK file."})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(-1)
	}

	var baseApk = *apk
	var tempDir = directory.CreateTempDir()

	dependency.AreAllReady()
	apktool.RunApktool(baseApk, tempDir)
	extractor.Extract(tempDir)
	directory.RemoveTempDir(tempDir)
}