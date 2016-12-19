package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Nudua/pocketbomberman/buttonswap"
	"github.com/Nudua/pocketbomberman/gbutils"
)

func main() {
	http.HandleFunc("/buttonswap", handler)

	fmt.Println("Pocket Bomberman A+B Button Swap Webserver running...")
	fmt.Println("Hit Ctrl-C to stop...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//Simple web server that will patch any rom posted to /buttonswap and send back the patched rom
func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		err := r.ParseMultipartForm(1024 * 1024 * 2) //2mb max

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, fileInfo, err := r.FormFile("uploadfile")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		romData, err := ioutil.ReadAll(file)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Check if file is valid, don't do this for web version?
		//gameVersion := pocketbomberman.IsKnownRom(romData)

		if buttonswap.Patch(romData) {
			gbutils.FixGlobalChecksum(romData)
		} else {
			http.Error(w, "Unable to patch rom :(", http.StatusInternalServerError)
			return
		}

		extension := filepath.Ext(fileInfo.Filename)

		fileName := fileInfo.Filename[:strings.Index(fileInfo.Filename, extension)]

		fileName = fileName + "(Patched)" + extension

		//Setup headers so we can send the patched rom back with a new file name
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)

		//Send the patched rom back
		w.Write(romData)
	}
}
