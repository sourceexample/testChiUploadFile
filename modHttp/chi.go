package modHttp

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

var g_chiMux *chi.Mux = nil

func Chi_Initialize() error {
	if g_chiMux != nil {
		return nil
	}
	g_chiMux = chi.NewRouter()

	g_chiMux.Get("/", handleHomepage)
	g_chiMux.Post("/upload", uploadFile)
	g_chiMux.Get("/test1.jpg", getjpg)

	http.ListenAndServe(":8080", g_chiMux)

	return nil
}

func handleHomepage(w http.ResponseWriter, r *http.Request) {
	data1, err := os.ReadFile("index.html")
	if err != nil {
		w.Write([]byte("sorry, no data"))
		return
	}

	w.Write(data1)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	out, err := os.Create("test1.jpg")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer out.Close()
	io.Copy(out, file)

	w.Write([]byte("done: " + handler.Filename))
	// read all of the contents of our uploaded file into a
	// byte array
	// fileBytes, err := io.ReadAll(file)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // write this byte array to our temporary file
	// tempFile.Write(fileBytes)
}

func getjpg(w http.ResponseWriter, r *http.Request) {
	img, err := os.ReadFile("test1.jpg")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("wrong"))
	}
	w.Header().Set("Content-Type", http.DetectContentType(img))
	w.Write(img)
}
