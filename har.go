package hoofli

import (
	"encoding/json"
	"io"
	"time"
)

type Property struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Properties []Property

type Content struct {
	Mimetype string `json:"mimeType"`
	Size     int    `json:"size"`
	Text     string `json:"text"`
}

type Page struct {
	Starteddatetime time.Time `json:"startedDateTime"`
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Pagetimings     struct {
		Oncontentload int `json:"onContentLoad"`
		Onload        int `json:"onLoad"`
	} `json:"pageTimings"`
}

type Pages []Page

type Response struct {
	Status      int        `json:"status"`
	Statustext  string     `json:"statusText"`
	Httpversion string     `json:"httpVersion"`
	Headers     Properties `json:"headers"`
	Cookies     Properties `json:"cookies"`
	Content     Content    `json:"content"`
	Redirecturl string     `json:"redirectURL"`
	Headerssize int        `json:"headersSize"`
	Bodysize    int        `json:"bodySize"`
}

type Timing struct {
	Blocked int `json:"blocked"`
	DNS     int `json:"dns"`
	Connect int `json:"connect"`
	Ssl     int `json:"ssl"`
	Send    int `json:"send"`
	Wait    int `json:"wait"`
	Receive int `json:"receive"`
}

type Cache struct {
	Afterrequest struct {
		Expires      string `json:"expires"`
		Lastfetched  string `json:"lastFetched"`
		Etag         string `json:"eTag"`
		Fetchcount   string `json:"fetchCount"`
		Datasize     string `json:"_dataSize"`
		Lastmodified string `json:"_lastModified"`
		Device       string `json:"_device"`
	} `json:"afterRequest"`
}

type Request struct {
	Bodysize    int        `json:"bodySize"`
	Method      string     `json:"method"`
	URL         string     `json:"url"`
	Httpversion string     `json:"httpVersion"`
	Headers     Properties `json:"headers"`
	Cookies     Properties `json:"cookies"`
	Querystring Properties `json:"queryString"`
	Headerssize int        `json:"headersSize"`
	Postdata    Content    `json:"postData"`
}

type Entry struct {
	Pageref         string    `json:"pageref"`
	Starteddatetime time.Time `json:"startedDateTime"`
	Response        Response  `json:"response,omitempty"`
	Timings         Timing    `json:"timings"`
	Time            int       `json:"time"`
	Securitystate   string    `json:"_securityState"`
	Serveripaddress string    `json:"serverIPAddress,omitempty"`
	Connection      string    `json:"connection,omitempty"`
	Cache           Cache     `json:"cache,omitempty"`
	Request         Request   `json:"request,omitempty"`
}

type Entries []Entry

// Har represents an entire HTTP Archive
// See the following for more details:
// - https://en.wikipedia.org/wiki/HAR_(file_format)
// - http://www.softwareishard.com/blog/har-12-spec/
// - https://w3c.github.io/web-performance/specs/HAR/Overview.html
type Har struct {
	Log struct {
		Version string   `json:"version"`
		Creator Property `json:"creator"`
		Browser Property `json:"browser"`
		Pages   Pages    `json:"pages"`
		Entries Entries  `json:"entries"`
	} `json:"log"`
}

func NewHar(r io.Reader) (*Har, error) {
	var har Har
	err := json.NewDecoder(r).Decode(&har)
	return &har, err
}
