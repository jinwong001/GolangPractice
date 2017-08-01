package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
	"strconv"
	"os"
	"path/filepath"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//		templ, err := template.New("go-web").Parse(
		//`Package name:{{.Name}}
		//Package NumFunc:{{.NumFunc}}
		//Package NumVars: {{.NumVars}}`)

		templ, err := template.ParseFiles("package.tmpl")

		if err != nil {
			log.Fatalf("Parse:%v", err)
			fmt.Fprintf(w, "Parse:%v", err)
			return
		}

		err = templ.Execute(w, Package{Name:"my", NumFuncs:12, NumVars:1})

		if err != nil {
			log.Fatalf("Execute:%v", err)
			fmt.Fprintf(w, "Execute:%v", err)
			return
		}
	})

	log.Println("server is starting")

	log.Fatal(http.ListenAndServe(":1234", nil))

}

func main1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("http.tmpl")

		if err != nil {
			log.Fatalf("Parse:%v", err)
			fmt.Fprintf(w, "Parse:%v", err)
			return
		}

		err = templ.Execute(w, r)

		if err != nil {
			log.Fatalf("Execute:%v", err)
			fmt.Fprintf(w, "Execute:%v", err)
			return
		}
	})

	log.Println("server is starting")

	log.Fatal(http.ListenAndServe(":1234", nil))

}

func main3() {
	wd, err := os.Getwd();
	if err != nil {
		log.Fatalf("Getwd:%v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles(filepath.Join(wd, "score.tmpl"))

		if err != nil {
			log.Fatalf("Parse:%v", err)
			fmt.Fprintf(w, "Parse:%v", err)
			return
		}

		score := r.FormValue("score")
		num, _ := strconv.Atoi(score)
		err = templ.Execute(w, num)

		if err != nil {
			log.Fatalf("Execute:%v", err)
			fmt.Fprintf(w, "Execute:%v", err)
			return
		}
	})

	log.Println("server is starting")

	log.Fatal(http.ListenAndServe(":1234", nil))

}

func main4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("header.tmpl")
		if err != nil {
			fmt.Fprintf(w, "Parse :%v", err)
			return
		}

		err = templ.Execute(w, r)
		if err != nil {
			fmt.Fprintf(w, "Execute:%v", err)
			return
		}
	})

	log.Println("Server is Starting...")
	log.Fatal(http.ListenAndServe(":1234", nil))
}

func main5() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("map.tmpl")
		if err != nil {
			fmt.Fprintf(w, "Parse:%v", err)
		}

		err = templ.Execute(w, map[string]interface{}{
			"Requset":r,
			"Score":98,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute:%v", err)
			return
		}
	})
	log.Println("Server is staring...")
	log.Fatal(http.ListenAndServe(":1234", nil))
}

func main7() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd: %v", err)
	}
	log.Print("Work directory:", wd)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pkg := &Package{
			Name:     "go-web",
			NumFuncs: 12,
			NumVars:  1200,
		}

		tmpl, err := template.New("main_v1.tmpl").Funcs(template.FuncMap{
			"NumFuncs": func() int {
				return pkg.NumFuncs
			},
			"Str2html": func(str string) template.HTML {
				return template.HTML(str)
			},
			"Divide": func(num int) int {
				return num / 2
			},
			"Add": func(num int) int {
				return num + 100
			},
		}).ParseFiles(filepath.Join(wd, "main_v1.tmpl"))
		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
			return
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"Name":        pkg.Name,
			"NumFuncs":    pkg.NumFuncs,
			"NumVars":     pkg.NumVars,
			"NumVarsHTML": `<li>Number of functions: 1200</li>`,
			"Maps": map[string]map[string]string{
				"Level1": map[string]string{
					"Name": "go-web",
				},
			},
			"Nums": []int{1, 2, 3, 4, 5, 6, 7},
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd:%w", wd)
	}

	log.Println("work dirctory:", wd)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pkg := Package{
			Name:"go web",
			NumVars:20,
			NumFuncs:30,
		}

		templ, err := template.New("complex.tmpl").Funcs(template.FuncMap{
			"NumFunc":func() int {
				return pkg.NumFuncs
			},
			"Str2Html":func(str string) template.HTML {
				return template.HTML(str)
			},
			"Divide":func(num int) int {
				return num / 2
			},
			"Add":func(num int) int {
				return num + 100
			},


		}).ParseFiles(filepath.Join(wd, "template/complex.tmpl"))

		if err != nil {
			fmt.Fprintf(w, "PaseFiles:%v", err)
			return
		}

		err = templ.Execute(w, map[string]interface{}{
			"Name":pkg.Name,
			"NumVars":pkg.NumVars,
			"NumFuncs":pkg.NumFuncs,
			"Nums": []int{1, 2, 3, 4, 5},
			"NumVarsHTML":`<li>Number of functions: 1200</li>`,
			"Maps": map[string]map[string]string{
				"Level1": map[string]string{
					"Name": "go-web",
				},
			},
		})

		if err != nil {
			fmt.Fprintf(w, "Execute:%v", err)
			return
		}
	})

	log.Println("server is starting...")
	log.Fatal(http.ListenAndServe(":1234", nil))
}