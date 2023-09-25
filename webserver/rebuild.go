package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func rebuildHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	err := rebuild()
	if err != nil {
		//return http 500
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Rebuild failed"}`))
		rebuildHistogramMetric.WithLabelValues(r.Method, r.URL.Path, "500").Observe(float64(time.Since(start).Seconds()))
		return
	}
	//return http 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": "Rebuild successful"}`))
	rebuildHistogramMetric.WithLabelValues(r.Method, r.URL.Path, "200").Observe(float64(time.Since(start).Seconds()))
}

func rebuild() error {
	log.Println("Rebuilding...")
	err := moveToCorrectFolder("/website")
	if err != nil {
		log.Println("Error moving to correct folder")
		return err
	}
	cmd := exec.Command("npm", "run", "build")
	out, err := cmd.Output()
	log.Println("Out", string(out))
	if err != nil {
		log.Println("Error running npm run build")
		log.Println(err.Error())
		return err
	}
	fmt.Println(string(out))
	return copyDistFolderToServeFolder()
}

func copyDistFolderToServeFolder() error {
	// recursively copy dist folder to serve folder replacing all files
	err := moveToCorrectFolder("/")
	if err != nil {
		log.Println("Error moving to correct folder")
		return err
	}
	// check if the serve folder exists, and delete it if it does
	if _, err := os.Stat("website/serve"); !os.IsNotExist(err) {
		log.Println("Serve folder exists, deleting...")
		err := os.RemoveAll("website/serve")
		if err != nil {
			log.Println("Error deleting serve folder")
			return err
		}
	}

	// copy dist folder to serve folder
	cmd := exec.Command("cp", "-p", "-r", "website/dist/", "website/serve")
	out, err := cmd.CombinedOutput()
	log.Println("Copying dist folder to serve folder , cp -p -r website/dist/ website/serve")
	log.Println("out", string(out))
	if err != nil {
		log.Println("Error running cp -r dist/* serve/")
		log.Println(err.Error())
		return err
	}
	fmt.Println(string(out))
	return nil
}

func moveToCorrectFolder(correctFolder string) error {
	curDir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory")
		return err
	}
	if curDir != correctFolder {
		err := os.Chdir(correctFolder)
		if err != nil {
			log.Println("Error changing directory to website")
			return err
		}
	}
	return nil
}
