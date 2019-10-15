package lib

import (
	util "github.com/eagle7410/go_util/libs"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	//TODO: clear Do somethining ...
	//r.PathPrefix("/static/").Handler(
	//	staticAccess(
	//		http.StripPrefix(
	//			"/static/",
	//			http.FileServer(http.Dir(path.Join(ENV.WorkDir, "/front/dist"))))))

	// Tech
	r.HandleFunc("/ping", util.Ping)
	r.HandleFunc("/", toIndex)

	return r
}

//TODO: clear
//func staticAccess(handler http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		//TODO: clear Maybe check access
//		handler.ServeHTTP(w, r)
//
//	})
//}

func toIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ping", http.StatusSeeOther)
}
