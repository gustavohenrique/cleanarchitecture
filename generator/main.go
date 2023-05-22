package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"generator/assets"
	"generator/src/fileutils"
	"generator/src/models"
)

func main() {
	http.HandleFunc("/download", download)
	http.HandleFunc("/generate", generate)
	http.Handle("/", assets.New().GetFS())
	port := getPortOrDefault("8003")
	log.Println("Listening", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func download(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	engine := r.URL.Query().Get("engine")
	filesystem := models.NewFilesystem(engine)
	file := filesystem.Download(filename)
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	http.ServeFile(w, r, file)
}

func generate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(getMessage("Method not allowed"))
		return
	}
	decoder := json.NewDecoder(r.Body)
	var project models.Project
	err := decoder.Decode(&project)
	if err != nil || !project.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getMessage(err))
		return
	}
	filesystem := models.NewFilesystem(project.GetEngine())
	extensions := filesystem.GetExtensions()
	templateData := project.GetTemplateData()
	parsedDir, err := fileutils.
		NewSed().
		From(filesystem.GetRepo()).
		To(filesystem.Dist(project.GetName())).
		Exclude(filesystem.GetSkipDirs()).
		Only(extensions).
		Replace(templateData).
		Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getMessage(err))
		return
	}
	compressedFile, err := fileutils.
		NewCompress().
		Input(parsedDir).
		Output(filesystem.GetDownload()).
		Name(project.GetName()).
		Exclude(filesystem.GetSkipDirs()).
		Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getMessage(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	downloadUrl := fmt.Sprintf("/download?file=%s&engine=%s", filepath.Base(compressedFile), project.GetEngine())
	w.Write(getMessage(downloadUrl))
}

func getMessage(i interface{}) []byte {
	return []byte(fmt.Sprintf(`{"message": "%s"}`, i))
}

func getPortOrDefault(def string) string {
	port := os.Getenv("GENERATOR_PORT")
	if port == "" {
		port = def
	}
	return ":" + port
}
