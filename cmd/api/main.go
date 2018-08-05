package main

import (
	"attrapi/pkg/attr"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.POST("/api/v0/attr", HandleText)
	port := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, r))
}

func HandleText(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var t *attr.Text
	err := decoder.Decode(&t)
	if err != nil {
		makeResponse(fmt.Sprintf("%v", err), w, http.StatusBadRequest)
		return
	}
	log.Println(t.TimeToRead())
	//TimeToRead(len(data))

}
