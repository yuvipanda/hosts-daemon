package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func collectHosts(dirPath string, targetPath string) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	tmpfile, err := ioutil.TempFile("", "hostCollector")
	if err != nil {
		panic(err)
	}
	defer tmpfile.Close()

	for _, file := range files {
		if !file.IsDir() {
			fullPath := filepath.Join(dirPath, file.Name())
			contents, err := ioutil.ReadFile(fullPath)
			if err != nil {
				panic(err)
			}
			tmpfile.WriteString("# " + fullPath + "\n")
			tmpfile.Write(contents)
			tmpfile.Write([]byte("\n"))
		}
	}

	err = tmpfile.Chmod(0644)
	if err != nil {
		panic(err)
	}
	os.Rename(tmpfile.Name(), targetPath)
}

func main() {
	dirPath := flag.String("dirpath", "/etc/hosts.d", "Path containing the individual hosts files")
	targetPath := flag.String("targetpath", "/etc/hosts", "Path to target /etc/hosts file to be written")
	scanInterval := flag.Int("scaninterval", 5, "Number of seconds between each scan of dirPath")

	flag.Parse()

	// FIXME: Use inotify!
	ticker := time.NewTicker(time.Second * time.Duration(*scanInterval))

	for _ = range ticker.C {
		collectHosts(*dirPath, *targetPath)
	}
}
