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
