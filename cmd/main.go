package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

    "containers_packer/internal/packer"
)

type PageData struct {
	Result      map[int]int
	TestsOutput template.HTML 
	Error       string
}

var tpl *template.Template

func main() {
	// Parse templates once at startup
	tpl = template.Must(template.ParseFiles(
		"web/templates/index.html",
		"web/templates/form.html",
		"web/templates/result.html",
		"web/templates/testsResult.html",
	))

    // server static files
    http.Handle("/static/",
        http.StripPrefix("/static/",
            http.FileServer(http.Dir("web/static")),
        ),
    )

    // routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tests", testsHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			data.Error = err.Error()
			renderTemplate(w, data)

			return
		}

		var containers []int
		for _, v := range r.Form["containers"] {
			val, err := strconv.Atoi(v)
			if err != nil {
				data.Error = "Invalid container: " + v
				renderTemplate(w, data)

				return
			}

			containers = append(containers, val)
		}

		// Goods
		goodsStr := r.FormValue("goods")
		goods, err := strconv.Atoi(goodsStr)
		if err != nil {
			data.Error = "Invalid goods value"
			renderTemplate(w, data)

			return
		}

		// Call packer
		res, err := packer.Pack(containers, goods)
		if err != nil {
			data.Error = err.Error()
		} else {
			data.Result = res
		}
	}

	renderTemplate(w, data)
}

func testsHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("./packer.test", "-test.v")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	output := out.String()
	if err != nil {
		output += "\nError: " + err.Error()
	}

    output = strings.ReplaceAll(output, "\n", "<br>")
    data := PageData{
        TestsOutput: template.HTML(output), // stays as HTML
    }

	renderTemplate(w, data)
}

func renderTemplate(w http.ResponseWriter, data PageData) {
	err := tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
