package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireParameter(t *testing.T) {
	server := NewServer()

	req, err := http.NewRequest("GET", "/media", nil)
	panicIf(err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	assertEqual(t, http.StatusBadRequest, rr.Code)
}

func TestChiaSeNhac_GetTrack(t *testing.T) {
	server := NewServerWithClient(mockClient)

	req, err := http.NewRequest("GET", "/detail?link="+link, nil)
	panicIf(err)

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	track := Track{}
	json.Unmarshal([]byte(`
	{
		"album": "http://chiasenhac.vn/nghe-album/so-far-away~martin-garrix-david-guetta-jamie-scott-romy-dya~tsvct563qvfhkw.html",
		"artist": "Martin Garrix; David Guetta; Jamie Scott; Romy Dya",
		"cover": "http://125.212.211.135/~csn/data/cover/80/79755.jpg",
		"id": "",
		"title": "So Far Away",
		"uid": "",
		"url": "http://chiasenhac.vn/nghe-album/so-far-away~martin-garrix-david-guetta-jamie-scott-romy-dya~tsvct563qvfhkw.html",
		"media": "http://data00.chiasenhac.com/downloads/1848/1/1847624-e3bfedf1/m4a/So%20Far%20Away%20-%20Martin%20Garrix_%20David%20Guett.m4a"
	}`), &track)

	expected, _ := json.Marshal(track)

	assertEqual(t, http.StatusOK, rr.Code)
	assertEqual(t, string(expected), rr.Body.String())
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Fatalf("expected %v, but got %v", expected, actual)
	}
}

func panicIf(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
