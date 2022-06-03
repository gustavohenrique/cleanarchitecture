package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"generator/src/fileutils"
	"generator/src/models"
)

func main() {
	http.HandleFunc("/download", download)
	http.HandleFunc("/", generate)
	port := getPortOrDefault("8003")
	log.Println("Listening", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func download(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	_, distDir := getTemplateDirs()
	file := filepath.Join(distDir, "..", filename)
	http.ServeFile(w, r, file)
}

func generate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(getMessage("Method now allowed"))
		return
	}
	sourceDir, distDir := getTemplateDirs()
	decoder := json.NewDecoder(r.Body)
	var project models.Project
	err := decoder.Decode(&project)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getMessage(err))
		return
	}
	extensions := project.GetFileExtensionsToBeReplaced()
	placeholders := project.GetPlaceholders()
	outputDir, err := fileutils.NewSed().From(sourceDir).To(distDir).Only(extensions).Replace(placeholders).Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getMessage(err))
		return
	}
	compressedFile, err := fileutils.NewCompress().Target(outputDir).Exclude(project.GetSkipDirs()).Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getMessage(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	downloadUrl := "/download?file=" + filepath.Base(compressedFile)
	w.Write(getMessage(downloadUrl))
}

func getMessage(i interface{}) []byte {
	return []byte(fmt.Sprintf(`{"message": "%s"}`, i))
}

func getTemplateDirs() (string, string) {
	sourceDir, _ := filepath.Abs(filepath.Dir(os.Getenv("SOURCE_DIR")))
	distDir, _ := filepath.Abs(filepath.Dir(os.Getenv("DIST_DIR")))
	return sourceDir, distDir
}

func getPortOrDefault(def string) string {
	port := os.Getenv("GENERATOR_PORT")
	if port == "" {
		port = def
	}
	return ":" + port
}
