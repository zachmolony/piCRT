package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

var categories = map[string]string{
	"anime":     "/home/pi/media/anime",
	"skate": 		 "/home/pi/media/skate",
	"jdm":			 "/home/pi/media/jdm",
	"longplays": "/home/pi/media/longplays",
	"misc":  		 "/home/pi/media/misc",	
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Path[len("/play/"):]
	path, exists := categories[category]

	if !exists {
		http.Error(w, "Invalid category", http.StatusBadRequest)
		return
	}

	findCmd := exec.Command("bash", "-c", fmt.Sprintf("find '%s' -type f \\( -iname '*.mp4' -o -iname '*.mkv' -o -iname '*.avi' -o -iname '*.mov' -o -iname '*.webm' -o -iname '*.flv' -o -iname '*.mpeg' \\) | shuf -n 1", path))
	filePathBytes, err := findCmd.Output()
	if err != nil {
		http.Error(w, "Failed to find media", http.StatusInternalServerError)
		return
	}

	filePath := strings.TrimSpace(string(filePathBytes)) 

	exec.Command("pkill", "-f", "mpv").Run()

	mpvCmd := exec.Command("mpv", "--fs", filePath)
	err = mpvCmd.Start() 

	if err != nil {
		http.Error(w, "Failed to start playback", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Now playing: %s", category)
	fmt.Printf("Playing: %s\n", filePath)
}


func stopHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("bash", "-c", "pkill -f mpv")
	cmd.Start()

	fmt.Fprintf(w, "Stopping playback")
}



func main() {
	http.HandleFunc("/play/", playHandler)
	http.HandleFunc("/stop", stopHandler)

	http.Handle("/", http.FileServer(http.Dir("/home/pi/piCRT/build")))

	fmt.Println("Listening on port 5000")
	http.ListenAndServe("0.0.0.0:5000", nil)
}
