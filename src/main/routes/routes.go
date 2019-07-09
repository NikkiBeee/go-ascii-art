package routes

import (
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
	printOutE := print.Printer("../AsciiArt/errorJPEG.jpg")
	p := AsciiArt{Homepage: "nil", Header: "nil", Created: printOutE}
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
		printOutE := print.Printer("../AsciiArt/errorJPEG.jpg")
		p := AsciiArt{Homepage: "nil", Header: "nil", Created: printOutE}
		t, _ := template.ParseFiles("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/main/public/error.html")
		t.Execute(w, p)
	}
	defer file.Close()

	userFile, err := ioutil.TempFile("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/AsciiArt/userImages", "user-*.jpeg")

	if err != nil {
		printOutE := print.Printer("../AsciiArt/errorJPEG.jpg")
		p := AsciiArt{Homepage: "nil", Header: "nil", Created: printOutE}
		t, _ := template.ParseFiles("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/main/public/error.html")
		t.Execute(w, p)
	}
	defer userFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		printOutE := print.Printer("../AsciiArt/errorJPEG.jpg")
		p := AsciiArt{Homepage: "nil", Header: "nil", Created: printOutE}
		t, _ := template.ParseFiles("/Users/nicolebent/Desktop/GHP_1094/Senior-Phase/stackathon/src/main/public/error.html")
		t.Execute(w, p)
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
	http.HandleFunc("/oops", errorHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)

}
