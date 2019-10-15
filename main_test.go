package main

import (
	. "github.com/smartystreets/goconvey/convey"
	sw "go_frame/lib"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCircle(t *testing.T) {

	if err := sw.ENV.Init(); err != nil {
		log.Fatalf("[0;31m Error on initializing envirement: %s :[39m", err)
	}

	router := sw.GetRouter()

	ts := httptest.NewServer(sw.LogRequest(router))
	defer ts.Close()

	Convey("App run", t, func() {
		client := http.Client{}
		req, err := http.NewRequest("GET", ts.URL+"/ping", nil)
		So(err, ShouldBeNil)

		resp, err := client.Do(req)
		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, http.StatusOK)

		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		bodyString := string(body)
		So(bodyString[:4], ShouldEqual, "PONG")
	})
}
