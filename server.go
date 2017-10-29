package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"path"
	"fmt"
	"strings"
	"os"
)

func getOptinText(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "By tapping \"Send\", you agree that Spotlight will slowly suck the battery out of your device with diagnostics logs. Here's a link: <a href=\"http://www.apple.com/legal/privacy\">http://www.apple.com/legal/privacy</a>")
}

func validateTicket(w http.ResponseWriter, r *http.Request) {

}

func postHandle(w http.ResponseWriter, r *http.Request) {
	//b, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	panic(err)
	//}
	//log.Print(hex.EncodeToString(b))
	r.ParseMultipartForm(0)
	log.Print("Values:")
	for name, value := range r.MultipartForm.Value {
		log.Printf("%s => %s", name, value)
	}
	log.Print("Files:")
	for name, fileValue := range r.MultipartForm.File {
		log.Printf("Working with %s", name)
		for _, value := range fileValue {
			filename := "files/" + escapeFileName(value.Filename)

			// Read and unzip for further use
			fileThing, _ := value.Open()
			byteThing, _ := ioutil.ReadAll(fileThing)
			unzippedThing, _ := GunzipData(byteThing)


			_ = ioutil.WriteFile(filename, unzippedThing, os.ModePerm)
			log.Printf("%s => %s", value.Filename, filename)
		}
	}
}

func escapeFileName(name string) (properName string)  {
	return strings.Replace(name, "/", "_", -1)
}

func respondWithFile(w http.ResponseWriter, r *http.Request) {
	// Serve the file needed per the path name.
	file, err := ioutil.ReadFile("files/" + path.Base(r.URL.Path))
	if err != nil {
		print(err)
	}
	w.Write(file)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s\n%s", r.Method, r.Host, r.URL, r.UserAgent())
		for name, value := range r.PostForm {
			log.Printf("%s => %s", name, value)
		}
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/configurations/retail/mobileBehaviorScan_1.1.plist", respondWithFile)
	http.HandleFunc("/MR3Server/GetOptinText", getOptinText)
	http.HandleFunc("/MR3Server/ValidateTicket", validateTicket)
	http.HandleFunc("/MR3Server/MR3Post", postHandle)
	http.HandleFunc("/server.crt", respondWithFile)
	go func() {
		log.Fatal(http.ListenAndServeTLS(":443", "files/server.crt", "files/server.key", logRequest(http.DefaultServeMux)))
	}()
	log.Fatal(http.ListenAndServe(":80", logRequest(http.DefaultServeMux)))
}
