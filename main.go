package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

//------suffixToScript------------------------------------

var suffixToScriptMap = make(map[string]string)
var mutex sync.Mutex

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

func runScript(scriptPath, fileName string) {
	cmd := exec.Command("bash", scriptPath, fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to execute script %s for file %s: %s\n", scriptPath, fileName, err)
	}
}


// ------------------------------------------
// Stack to hold files to be processed

var fileStack []string   


func checkFileStack(str string) bool {
	mutex.Lock()
	defer mutex.Unlock()
    for _, v := range fileStack {
        if v == str {
            return true
        }
    }
    return false
}       

func addFileToProcess(fileName string) {
	if exists := checkFileStack(fileName); exists {
		log.Printf("🚮 %s\n", filepath.Base(fileName))
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

    fileStack = append(fileStack, fileName)	
	log.Printf("➕ %s\n", filepath.Base(fileName))
}

func processFileStack() {
	mutex.Lock()
	defer mutex.Unlock()

	for len(fileStack) > 0 {
		fileName := fileStack[len(fileStack)-1]
		fileStack = fileStack[:len(fileStack)-1]

		log.Printf("📜 %s\n", filepath.Base(fileName))
		
        fileInfo, err := os.Stat(fileName)
        if err == nil && !fileInfo.IsDir() && (fileInfo.Size() > 0) {
		    fileExtension := filepath.Ext(fileName)
	        scriptPath, exists := suffixToScriptMap[strings.ToLower(fileExtension)]

	        if exists {
		        go runScript(scriptPath, fileName)
	        } else {
	            log.Printf("🤷 File not handled: %s\n", filepath.Base(fileName))
	        }
	       
	    } else {
	           log.Printf("🤷 File gone or missing: %s\n", filepath.Base(fileName))
	    }
				
	}
}

//------------------------------------------

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <target-directory> <script-directory>\n", os.Args[0])
	}

	targetDir := os.Args[1]
	scriptDir := os.Args[2]

	err := loadScripts(scriptDir)
	if err != nil {
		log.Fatalf("Failed to load scripts: %s", err)
	}

	err = monitorDirectory(targetDir)
	if err != nil {
		log.Fatalf("Failed to monitor directory: %s", err)
	}
}

func monitorDirectory(targetDir string) error {
    handler := NewTimeoutHandler(2 * time.Second)
    
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
					go addFileToProcess(event.Name)
					handler.Trigger(processFileStack)
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
