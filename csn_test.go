package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

var link = "http://chiasenhac.vn/nghe-album/so-far-away~martin-garrix-david-guetta-jamie-scott-romy-dya~tsvct563qvfhkw.html"
var mockClient = newMockClient()

type MockClient struct {
	resp []byte
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBuffer(m.resp)),
		StatusCode: 200,
		Status:     "200 OK",
	}, nil
}

func newMockClient() *MockClient {
	sampleHtml, e := ioutil.ReadFile("test.htm")
	panicIf(e)
	return &MockClient{resp: sampleHtml}
}

func TestChiaSeNhac_GetList(t *testing.T) {
	csn := NewCsn(mockClient)
	tracks, e := csn.GetList(link)

	panicIf(e)
	assertEqual(t, 1, len(tracks))
	assertEqual(t, "So Far Away", tracks[0].Title)
}

func TestChiaSeNhac_GetTrackDetail(t *testing.T) {
	csn := NewCsn(mockClient)
	track, e := csn.GetTrack(link)

	panicIf(e)
	assertEqual(t, "So Far Away", track.Title)
	assertEqual(t, "http://data00.chiasenhac.com/downloads/1848/1/1847624-e3bfedf1/m4a/So%20Far%20Away%20-%20Martin%20Garrix_%20David%20Guett.m4a", track.Media)
}

func TestChiaSeNhac_GetMedia(t *testing.T) {
	csn := NewCsn(mockClient)
	media, e := csn.GetMedia(link)

	panicIf(e)
	assertEqual(t, "http://data00.chiasenhac.com/downloads/1848/1/1847624-e3bfedf1/m4a/So%20Far%20Away%20-%20Martin%20Garrix_%20David%20Guett.m4a", media)
}
