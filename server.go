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

var basePath string
var currentlyPlaying string

func init() {
	basePath = os.Getenv("PICRT_MEDIA_PATH")
	if basePath == "" {
		basePath = "/home/pi/media/"
	}
}

func setCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getVideosHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
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
	setCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/play/")
	parts := strings.SplitN(path, "/", 2)
	category := parts[0]
	categoryPath := filepath.Join(basePath, category)

	exec.Command("pkill", "-f", "mpv").Run()

	if len(parts) == 2 {
		// Play a specific video
		video := parts[1]
		filePath := filepath.Join(categoryPath, video)
		currentlyPlaying = fmt.Sprintf("%s/%s", category, video)
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
	currentlyPlaying = fmt.Sprintf("%s (shuffled)", category)
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

func nowPlayingHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"nowPlaying": currentlyPlaying})
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	cmd := exec.Command("bash", "-c", "pkill -f mpv")
	cmd.Start()

	fmt.Fprintf(w, "Stopping playback")
}

func getCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
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

func getCategoryInfoHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	type CatInfo struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}
	var result []CatInfo

	entries, err := os.ReadDir(basePath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]CatInfo{})
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			category := entry.Name()
			files, _ := os.ReadDir(filepath.Join(basePath, category))
			count := 0
			for _, f := range files {
				if !f.IsDir() {
					name := f.Name()
					if strings.HasSuffix(strings.ToLower(name), ".mp4") ||
						strings.HasSuffix(strings.ToLower(name), ".mkv") ||
						strings.HasSuffix(strings.ToLower(name), ".avi") ||
						strings.HasSuffix(strings.ToLower(name), ".mov") ||
						strings.HasSuffix(strings.ToLower(name), ".webm") ||
						strings.HasSuffix(strings.ToLower(name), ".flv") ||
						strings.HasSuffix(strings.ToLower(name), ".mpeg") {
						count++
					}
				}
			}
			result = append(result, CatInfo{Name: category, Count: count})
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/play/", playHandler)
	http.HandleFunc("/stop", stopHandler)
	http.HandleFunc("/categories", getCategoriesHandler)
	http.HandleFunc("/videos/", getVideosHandler)
	http.HandleFunc("/nowplaying", nowPlayingHandler)
	http.HandleFunc("/categoryinfo", getCategoryInfoHandler)

	http.Handle("/", http.FileServer(http.Dir("/home/pi/piCRT/build")))

	fmt.Println("Listening on port 5000")
	http.ListenAndServe("0.0.0.0:5000", nil)
}
