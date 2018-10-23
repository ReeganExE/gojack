package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/generaltso/vibrant"

	"github.com/PuerkitoBio/goquery"
)

var (
	sourceChars     = strings.Split("UWJHDGMAYIXNRLBPK", "")
	actual          = strings.Split("0123456789cufr112", "")
	userAgent       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
	regDecodeUrl    = regexp.MustCompile(`decode_download_url\("([^"]+)", "([^"]+)", "([^"]+)"`)
	regExtractAlbum = regexp.MustCompile(`<u>Album(.*?)href="(https?:\/\/chiasenhac.vn\/nghe-album[^"]+)"`)
	regAlbum        = regexp.MustCompile(`chiasenhac.vn\/nghe-album`)
	regImage        = regexp.MustCompile(`rel="image_src"\s+href="([^"]+)"`)
	regTitleArist   = regexp.MustCompile(`<meta name="title"\s+content="([^"]+)"`)
)

// HttpClient HTTP client interface
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// ChiaSeNhac client
type ChiaSeNhac struct {
	client HttpClient
}

// Colors color
type Colors struct {
	Background string `json:"background"`
	Text       string `json:"text"`
}

// Palette Palette
type Palette map[string]*Colors

// Track a track detail
type Track struct {
	Album   string   `json:"album,omitempty"`
	Artist  string   `json:"artist"`
	Cover   string   `json:"cover,omitempty"`
	ID      string   `json:"id,omitempty"`
	Title   string   `json:"title"`
	Uid     string   `json:"uid,omitempty"`
	URL     string   `json:"url"`
	Media   string   `json:"media,omitempty"`
	Palette *Palette `json:"palette,omitempty"`
}

// AsJSON marshall current object as JSON string
func (t *Track) AsJSON() []byte {
	bytes, e := json.Marshal(t)
	if e != nil {
		log.Fatal(e)
	}

	return bytes
}

func (c *Colors) String() string {
	bytes, e := json.Marshal(c)
	if e != nil {
		log.Fatal(e)
	}

	return string(bytes)
}

func newColorFromSwatch(v *vibrant.Swatch) *Colors {
	if v == nil {
		return nil
	}

	return &Colors{v.Color.RGBHex(), v.Color.BodyTextColor().String()}
}

// NewCsn creates new CSN instance
func NewCsn(client HttpClient) *ChiaSeNhac {
	return &ChiaSeNhac{client}
}

// GetList gets all the tracks in an album or playlist
func (c *ChiaSeNhac) GetList(link string) ([]Track, error) {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)

	resp, e := c.client.Do(req)
	if e != nil {
		return nil, e
	}

	defer resp.Body.Close()
	// Load the HTML document
	doc, e := goquery.NewDocumentFromReader(resp.Body)
	if e != nil {
		return nil, e
	}

	var tracks []Track

	doc.Find("#playlist .tbtable tr[class]").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".musictitle")
		band := strings.TrimSpace(title.Closest("span.gen").Text())
		titleAndArtist := strings.SplitN(band, " - ", 2)
		val, _ := title.Attr("href")
		tracks = append(tracks, Track{
			Title:  titleAndArtist[0],
			Artist: titleAndArtist[1],
			URL:    val,
		})
	})

	return tracks, nil
}

// GetTrack gets the detail of a track. Included streamUrl (media), cover art, album
func (c *ChiaSeNhac) GetTrack(link string) (track *Track, e error) {
	html, e := c.getAsHtml(link)
	if e != nil {
		return
	}

	track = &Track{URL: link}

	if regTitleArist.Match(html) {
		submatch := regTitleArist.FindSubmatch(html)
		matcheds := strings.Split(string(submatch[1]), " ~ ")
		track.Title = matcheds[0]
		track.Artist = matcheds[1]
	}

	if regImage.Match(html) {
		submatch := regImage.FindSubmatch(html)
		track.Cover = string(submatch[1])
		palette, _ := c.GetPalette(track.Cover)
		track.Palette = palette
	}

	if regAlbum.MatchString(link) {
		track.Album = link
	} else if regExtractAlbum.Match(html) {
		submatch := regExtractAlbum.FindSubmatch(html)
		track.Album = string(submatch[2])
	}

	track.Media = parseParams(html)

	return
}

// GetPreset gets the recently shared tracks
// Ex: http://chiasenhac.vn/mp3/beat-playback/us-instrumental/
func (c *ChiaSeNhac) GetPreset(link string) ([]Track, error) {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)

	resp, e := c.client.Do(req)
	if e != nil {
		return nil, e
	}

	doc, e := goquery.NewDocumentFromReader(resp.Body)
	if e != nil {
		return nil, e
	}

	var tracks []Track

	doc.Find(".bod table .genmed a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		list, _ := c.GetList(href)
		tracks = append(tracks, list...)
	})
	return tracks, nil
}

// GetMedia returns a stream URL
func (c *ChiaSeNhac) GetMedia(link string) (string, error) {
	html, _ := c.getAsHtml(link)
	mediaLink := parseParams(html)

	if mediaLink == "" {
		return "", fmt.Errorf("Link is invalid or moved: %s", link)
	}
	return mediaLink, nil
}

func (c *ChiaSeNhac) getAsHtml(link string) ([]byte, error) {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "mq", Value: "i%3A320%3B"})
	req.Header.Set("User-Agent", userAgent)

	resp, e := c.client.Do(req)
	if e != nil {
		return nil, e
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *ChiaSeNhac) GetPalette(imgSrc string) (*Palette, error) {
	req, err := http.NewRequest("GET", imgSrc, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)

	resp, e := c.client.Do(req)
	if e != nil {
		return nil, e
	}

	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("unable to decode image %s", imgSrc)
	}

	palette, err := vibrant.NewPaletteFromImage(img)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("unable to decode image %s", imgSrc)
	}

	awesome := palette.ExtractAwesome()

	rs := Palette{
		"muted":      newColorFromSwatch(awesome["Muted"]),
		"lightMuted": newColorFromSwatch(awesome["LightMuted"]),
		"darkMuted":  newColorFromSwatch(awesome["DarkMuted"]),
		"color":      newColorFromSwatch(awesome["Vibrant"]),
		"light":      newColorFromSwatch(awesome["LightVibrant"]),
		"dark":       newColorFromSwatch(awesome["DarkVibrant"]),
	}

	return &rs, nil
}

func parseParams(html []byte) (link string) {
	if regDecodeUrl.Match(html) {
		matches := regDecodeUrl.FindAllSubmatch(html, -1)
		matched := matches[0]
		link = decodeDownloadUrl(string(matched[1]), string(matched[2]), string(matched[3]))
	}

	return
}

func decodeDownloadUrl(http, code, tail string) string {
	for i := range sourceChars {
		code = strings.Replace(code, sourceChars[i], actual[i], -1)
	}

	return http + code + tail
}
