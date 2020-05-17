/*
 made with love by @ndelphit 5/2020
*/

package main

import (
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/ndelphit/apkurlgrep/command/apktool"
	dependency "github.com/ndelphit/apkurlgrep/command/dependency"
	"github.com/ndelphit/apkurlgrep/directory"
	"github.com/ndelphit/apkurlgrep/extractor"
	"os"
	"path/filepath"
)

func main() {

	parser := argparse.NewParser("apkurlgrep", "ApkUrlGrep - Extract endpoints from APK files")
	apk := parser.String("a", "apk", &argparse.Options{
		Required: true,
		Help:     "Input a path to APK file.",
		Validate: func(args []string) error {
			filename := args[0]
			_, err := os.Stat(filename)
			if os.IsNotExist(err) {
				return errors.New("file does not exist")
			} else if filepath.Ext(filename) == ".apks" {
				return errors.New("this program does not support files with APKS extension. Please unpack this file first and provide the APK file (like com.application.apk) that will be located in the unpacked directory")
			} else if filepath.Ext(filename) != ".apk" {
				return errors.New("provided file have the unsupported extension. Please provide file with APK extension")
			} else if err != nil {
				return errors.New("unable to read file")
			}
			return nil
		},
	})

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
