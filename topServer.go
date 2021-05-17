package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
)

type templData struct {
	Hostname       string
	SelectedOption string
	TopOutput      string
}

var tmpl *template.Template
var tmplData = new(templData)
var cmdArguments string

// Is executed automatically on package load
// Reads in and parses HTML templates
func init() {
	if runtime.GOOS == "linux" {
		cmdArguments = "top -b -n 1 -o "
	} else if runtime.GOOS == "darwin" {
		cmdArguments = "top -l 1 -o "
	} else {
		println("Error: not a suitable operting system...")
		os.Exit(-1)
	}

	files := []string{
		"top.html",
	}
	tmpl = template.Must(template.ParseFiles(files...))

	tmplData.Hostname, _ = os.Hostname()
	tmplData.SelectedOption = "pid"
}

func main() {
	r := mux.NewRouter()

	// URL endpoints for HTML frontend
	r.HandleFunc("/", top).Methods("GET")

	//server := http.Server{
	//	Addr:    "localhost:8443",
	//	Handler: r,
	//}

	err := http.ListenAndServeTLS("localhost:8443",
		"/etc/letsencrypt/live/ecs-80-158-58-79.reverse.open-telekom-cloud.com/fullchain.pem",
		"/etc/letsencrypt/live/ecs-80-158-58-79.reverse.open-telekom-cloud.com/privkey.pem", r)
	log.Fatal(err)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func top(w http.ResponseWriter, r *http.Request) {
	options, ok := r.URL.Query()["option"]
	whitelist := []string{"PID", "%MEM", "%CPU", "TIME"}

	var option string
	if !ok || len(options[0]) < 1 && !contains(whitelist, options[0]) {
		option = "pid"
	} else {
		// Query()["option"] will return an array of items,
		// we only want the single item.
		option = options[0]
	}

	cmd := exec.Command("sh", "-c", cmdArguments+option)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()

	var output string
	if err != nil {
		output = "Error"
	} else {
		output = strings.TrimSpace(string(stdout.Bytes()))
	}

	tmplData.SelectedOption = option
	tmplData.TopOutput = output
	tmpl.ExecuteTemplate(w, "top.html", tmplData)
}
