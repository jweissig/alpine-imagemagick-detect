package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// parse and execute the template
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.5.html
// https://github.com/astaxie/build-web-application-with-golang/blob/master/en/04.4.md
// http://stackoverflow.com/questions/15202448/go-formfile-for-multiple-files
// upload logic
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// generate semi-random filename
	rand.Seed(time.Now().UnixNano())
	name := randSeq(10)

	r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// log some data to console
	fmt.Printf("Processing: %v\n", handler.Header)

	// create the directory path if needed
	os.MkdirAll("./tmp/", 0755)

	// save the file
	f, err := os.OpenFile("./tmp/"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)

	// probe file
	probedata := probe(name)

	fmt.Fprintf(w, "%v", probedata)

}

func probe(filename string) string {

	// convert example.png json:
	cmd := exec.Command("convert", "/app/tmp/"+filename, "json:")
	//fmt.Println(cmd)
	response, err := cmd.CombinedOutput()

	if err != nil {
		return "ERROR in cmd.CombinedOutput()"
	}

	// dump contents of command
	//fmt.Println(response)

	// hopefully return json
	return string(response)

}

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)

	log.Println("Listening 5000...")
	http.ListenAndServe(":5000", nil)

}
