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
	Header   string
	Homepage string
	Created  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	printOut := print.Printer("../AsciiArt/t-simpsons-oral-history-august-2007.jpg")
	printOutH := print.Printer("../AsciiArt/headerHomeJPEG.jpg")
	p := AsciiArt{Homepage: printOut, Header: printOutH, Created: "nil"}
	t, _ := template.ParseFiles("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/main/public/index.html")
	t.Execute(w, p)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	printOut := print.Printer("../AsciiArt/errorJPEG.jpg")
	p := AsciiArt{Homepage: "nil", Header: "nil", Created: printOut}
	t, _ := template.ParseFiles("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/main/public/error.html")
	t.Execute(w, p)
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/main/public/style.css")
	t.Execute(w, nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
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

	printOut := print.Printer(userFile.Name())
	printOutH := print.Printer("../AsciiArt/art2JPEG.jpg")

	defer os.Remove(userFile.Name())

	p := AsciiArt{Homepage: "nil", Header: printOutH, Created: printOut}

	t, _ := template.ParseFiles("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/main/public/create.html")
	t.Execute(w, p)

}
func SetupRoutes() {
	// fmt.Println("setup is running")
	// http.Handle("*", http.FileServer(http.Dir("/public")))
	http.HandleFunc("/oops", errorHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)

	// fmt.Println("setup is running 2")

}
