package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	fs := http.FileServer(http.Dir("../website/dist"))
	http.Handle("/", fs)
	http.HandleFunc("/api/rebuild", rebuildHandler)

	log.Print("Listening on :80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func rebuildHandler(w http.ResponseWriter, r *http.Request) {
	err := rebuild()
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	fmt.Fprintf(w, "Rebuild complete")
}

func rebuild() error {
	log.Println("Rebuilding...")
	curDir, err := os.Getwd()
	log.Println("curDir", curDir)
	if err != nil {
		return err
	}
	if curDir != "/website" {
		err := os.Chdir("website")
		if err != nil {
			return err
		}
	}
	cmd := exec.Command("npm", "run", "build")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
