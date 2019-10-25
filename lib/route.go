package lib

import (
	util "github.com/eagle7410/go_util/libs"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	// Tech
	r.HandleFunc("/ping", util.Ping)
	r.HandleFunc("/wait", util.Ping)
	r.HandleFunc("/", toIndex)

	return r
}


func toIndex(w http.ResponseWriter, r *http.Request) {
	ver := "v4"
	timeString := time.Now().Format(time.RFC3339)
	util.Logf("Handler app %v at %v, url %v",ver, timeString, r.URL)
	time.Sleep(time.Second * 8)
	util.Logf("Handler app %v at %v, url %v End",ver, timeString, r.URL)
	w.Write([]byte("Ready " + ver))
	w.WriteHeader(http.StatusOK)
}
