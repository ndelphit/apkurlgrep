package directory


import (
	"io/ioutil"
	"log"
	"os"
)

func CreateTempDir() string {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	return tmp
}

func RemoveTempDir(tempDir string)  {
	defer os.RemoveAll(tempDir)
}