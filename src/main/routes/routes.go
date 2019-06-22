package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"../print"
)

type AsciiArt struct {
	Art string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/main/public/index.html")
	t.Execute(w, nil)
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/main/public/style.css")
	t.Execute(w, nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("createHandler is running")

	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("userImage")
	if err != nil {
		fmt.Println("Error in retrieving file from form data")
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Printf("uploaded file")

	userFile, err := ioutil.TempFile("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/AsciiArt/userImages", "user-*.jpeg")
	fmt.Println(userFile.Name())

	if err != nil {
		fmt.Println("Error in storing tempfile")
		fmt.Println(err)
		return
		//PLaceholder for redirect to an Error page
	}
	defer userFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error in reading file")
		fmt.Println(err)
		return
		//PLaceholder for redirect to an Error page
	}
	userFile.Write(fileBytes)

	fmt.Println("fileUploaded")

	printOut := print.Printer(userFile.Name())

	defer os.Remove(userFile.Name())

	p := AsciiArt{Art: printOut}

	t, _ := template.ParseFiles("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/main/public/create.html")
	t.Execute(w, p)

}
func SetupRoutes() {
	fmt.Println("setup is running")
	http.Handle("*", http.FileServer(http.Dir("/public")))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)

	fmt.Println("setup is running 2")

}
