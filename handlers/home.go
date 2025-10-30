package handlers

import (
	"encoding/json"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

// ViteManifest represents the structure of Vite's manifest.json
type ViteManifest map[string]struct {
	File string   `json:"file"`
	CSS  []string `json:"css,omitempty"`
}

// PageData represents data passed to templates
type PageData struct {
	Title    string
	IsDev    bool
	CSSFiles []string
	JSFiles  []string
}

// HomeHandler handles the root route
func HomeHandler(isDev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse templates - need to recursively find all .gohtml files
		tmpl := template.New("")
		err := filepath.WalkDir("ui", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && filepath.Ext(path) == ".gohtml" {
				_, err := tmpl.ParseFiles(path)
				return err
			}
			return nil
		})
		if err != nil {
			http.Error(w, "Error loading templates: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Prepare page data
		data := PageData{
			Title:    "GoVite - Home",
			IsDev:    isDev,
			CSSFiles: []string{},
			JSFiles:  []string{},
		}

		// In production, read the manifest to get asset paths
		if !isDev {
			manifestPath := filepath.Join("dist", "manifest.json")
			manifestData, err := os.ReadFile(manifestPath)
			if err != nil {
				http.Error(w, "Error reading manifest: "+err.Error(), http.StatusInternalServerError)
				return
			}

			var manifest ViteManifest
			if err := json.Unmarshal(manifestData, &manifest); err != nil {
				http.Error(w, "Error parsing manifest: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// Get the main entry point assets
			if entry, ok := manifest["main.ts"]; ok {
				data.JSFiles = append(data.JSFiles, entry.File)
				data.CSSFiles = append(data.CSSFiles, entry.CSS...)
			}
		}

		// Render the template
		if err := tmpl.ExecuteTemplate(w, "index.gohtml", data); err != nil {
			http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
