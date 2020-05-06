/*
 with love by @ndelphit 5/2020
*/

package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/ndelphit/apkurlgrep/command/apktool"
	dependency "github.com/ndelphit/apkurlgrep/command/dependency"
	"github.com/ndelphit/apkurlgrep/directory"
	"github.com/ndelphit/apkurlgrep/extractor"
	"os"
)



func main() {

	parser := argparse.NewParser("apkurlgrep", "ApkUrlGrep")
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