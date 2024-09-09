package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
)

// A map to hold file suffix to script mapping
var suffixToScriptMap = make(map[string]string)
var mutex sync.Mutex

func main() {
	// Target directory and script folder as params
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <target-directory> <script-directory>\n", os.Args[0])
	}

	targetDir := os.Args[1]
	scriptDir := os.Args[2]

	// Load .sh scripts and map them to file suffixes
	err := loadScripts(scriptDir)
	if err != nil {
		log.Fatalf("Failed to load scripts: %s", err)
	}

	// Monitor the target directory for new files
	err = monitorDirectory(targetDir)
	if err != nil {
		log.Fatalf("Failed to monitor directory: %s", err)
	}
}

// Load .sh scripts from the script directory and create suffix mappings
func loadScripts(scriptDir string) error {
	files, err := os.ReadDir(scriptDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sh") {
			suffix := strings.TrimSuffix(file.Name(), ".sh")
			mutex.Lock()
			suffixToScriptMap["."+suffix] = filepath.Join(scriptDir, file.Name())
			mutex.Unlock()
			fmt.Printf("Mapped %s to script %s\n", "."+suffix, file.Name())
		}
	}

	return nil
}

// Monitor the target directory for new files
func monitorDirectory(targetDir string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	err = watcher.Add(targetDir)
	if err != nil {
		return err
	}

	log.Println("Watching directory:", targetDir)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			if event.Op&fsnotify.Create == fsnotify.Create {

				fileInfo, err := os.Stat(event.Name)
				if err == nil && !fileInfo.IsDir() && (fileInfo.Size() > 0) {
                    log.Println("------------------------------")
					log.Printf("> %d  %s\n", fileInfo.Size(), event.Name)
					//<M-C-S-F5>|| event.Op&fsnotify.Rename == fsnotify.Rename {
					// When a new file is created, check if the suffix matches any of the scripts
					fileExtension := filepath.Ext(event.Name)
					mutex.Lock()
					scriptPath, exists := suffixToScriptMap[strings.ToLower(fileExtension)]
					mutex.Unlock()

					if exists {
						// Execute the script with the file name as an argument
						go runScript(scriptPath, event.Name)
					}
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			log.Println("Error:", err)
		}
	}
}

// Run the .sh script passing the new file name as an argument
func runScript(scriptPath, fileName string) {
	cmd := exec.Command("bash", scriptPath, fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to execute script %s for file %s: %s\n", scriptPath, fileName, err)
	}
}
