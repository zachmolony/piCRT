package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var dev = true 

var basePath string

func init() {
	if dev {
		basePath = "/mnt/hagrid/piCRT/"
	} else {
		basePath = "/home/pi/media/"
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	category := r.URL.Path[len("/play/"):]
	path := basePath + category

	findCmd := exec.Command("bash", "-c", fmt.Sprintf("find '%s' -type f \\( -iname '*.mp4' -o -iname '*.mkv' -o -iname '*.avi' -o -iname '*.mov' -o -iname '*.webm' -o -iname '*.flv' -o -iname '*.mpeg' \\) | shuf -n 1", path))
	filePathBytes, err := findCmd.Output()
	if err != nil {
		http.Error(w, "Failed to find media", http.StatusInternalServerError)
		return
	}

	filePath := strings.TrimSpace(string(filePathBytes)) 

	exec.Command("pkill", "-f", "mpv").Run()

	mpvCmd := exec.Command("bash", "-c", fmt.Sprintf("mpv --fs --loop-playlist=inf --shuffle '%s'/*", path))
	err = mpvCmd.Start() 

	if err != nil {
		http.Error(w, "Failed to start playback", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Now playing: %s", category)
	fmt.Printf("Playing: %s\n", filePath)
}


func stopHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cmd := exec.Command("bash", "-c", "pkill -f mpv")
	cmd.Start()

	fmt.Fprintf(w, "Stopping playback")
}


func getCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var categoriesData []string

	entries, err := os.ReadDir(basePath)
	if err != nil {
		http.Error(w, "Failed to find categories", http.StatusInternalServerError)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			categoriesData = append(categoriesData, entry.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categoriesData)
}


func main() {
	http.HandleFunc("/play/", playHandler)
	http.HandleFunc("/stop", stopHandler)
	http.HandleFunc("/categories", getCategoriesHandler)

	http.Handle("/", http.FileServer(http.Dir("/home/pi/piCRT/build")))

	fmt.Println("Listening on port 5000")
	http.ListenAndServe("0.0.0.0:5000", nil)
}
