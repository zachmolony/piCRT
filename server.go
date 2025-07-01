package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
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

func getVideosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	categoryEnc := strings.TrimPrefix(r.URL.Path, "/videos/")
	category, err := url.PathUnescape(categoryEnc)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]string{})
		return
	}
	path := filepath.Join(basePath, category)

	entries, err := os.ReadDir(path)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]string{})
		return
	}

	videos := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			name := entry.Name()
			if strings.HasSuffix(strings.ToLower(name), ".mp4") ||
				strings.HasSuffix(strings.ToLower(name), ".mkv") ||
				strings.HasSuffix(strings.ToLower(name), ".avi") ||
				strings.HasSuffix(strings.ToLower(name), ".mov") ||
				strings.HasSuffix(strings.ToLower(name), ".webm") ||
				strings.HasSuffix(strings.ToLower(name), ".flv") ||
				strings.HasSuffix(strings.ToLower(name), ".mpeg") {
				videos = append(videos, name)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videos)
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	path := strings.TrimPrefix(r.URL.Path, "/play/")
	parts := strings.SplitN(path, "/", 2)
	category := parts[0]
	categoryPath := filepath.Join(basePath, category)

	exec.Command("pkill", "-f", "mpv").Run()

	if len(parts) == 2 {
		// Play a specific video
		video := parts[1]
		filePath := filepath.Join(categoryPath, video)
		mpvCmd := exec.Command("mpv", "--fs", filePath)
		err := mpvCmd.Start()
		if err != nil {
			http.Error(w, "Failed to start playback", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Now playing: %s/%s", category, video)
		fmt.Printf("Playing: %s\n", filePath)
		return
	}

	// Shuffle mode: play all videos in the category, shuffled and looped
	mpvCmd := exec.Command("bash", "-c", fmt.Sprintf(
		"mpv --fs --loop-playlist=inf --shuffle '%s'/*",
		categoryPath,
	))
	err := mpvCmd.Start()
	if err != nil {
		http.Error(w, "Failed to start playback", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Now playing (shuffled): %s", category)
	fmt.Printf("Playing shuffled playlist: %s/*\n", categoryPath)
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
	http.HandleFunc("/videos/", getVideosHandler)

	http.Handle("/", http.FileServer(http.Dir("/home/pi/piCRT/build")))

	fmt.Println("Listening on port 5000")
	http.ListenAndServe("0.0.0.0:5000", nil)
}
