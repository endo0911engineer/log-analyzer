package watcher

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

// WatchFile sets up a file watcher on the given path
func WatchFile(path string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}
	defer watcher.Close()

	//親ディレクトリを監視する（ファイル変更イベントはディレクトリレベルで発生するため）
	dir := getDir(path)
	err = watcher.Add(dir)
	if err != nil {
		return fmt.Errorf("failed to add watch on directory: %w", err)
	}

	log.Printf("Watching file: %s", path)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			//イベントが対象ファイルで、かつWriteイベントなら反応
			if event.Name == path && event.Op&fsnotify.Write == fsnotify.Write {
				log.Printf("Modified file: %s", event.Name)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			log.Println("error", err)
		}
	}
}

// getDir extracts directory from full file path
func getDir(path string) string {
	lastSlash := -1
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			lastSlash = i
			break
		}
	}
	if lastSlash == -1 {
		return "."
	}
	return path[:lastSlash]
}
