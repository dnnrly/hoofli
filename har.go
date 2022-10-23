package hoofli

import (
	"encoding/json"
	"io"
	"regexp"
	"time"
)

// Property is a key value pair
type Property struct {
	// Name of the value represented
	Name string `json:"name"`

	// Value being represented
	Value string `json:"value"`
}

// Properties is a collection of Property
type Properties []Property

// Content is the payload of a request or response
type Content struct {
	Mimetype string `json:"mimeType"`
	Size     int    `json:"size"`
	Text     string `json:"text"`
}

// Page is the data that relates to a single browser page rather the requests and responses
type Page struct {
	Starteddatetime time.Time `json:"startedDateTime"`
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Pagetimings     struct {
		Oncontentload float32 `json:"onContentLoad"`
		Onload        float32 `json:"onLoad"`
	} `json:"pageTimings"`
}

// Pages is a collection of Page
type Pages []Page

// Response represents the data about the response to a request
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

// Timing contains the timing of various different parts of the request and response for an entry
type Timing struct {
	Blocked float32 `json:"blocked"`
	DNS     float32 `json:"dns"`
	Connect float32 `json:"connect"`
	Ssl     float32 `json:"ssl"`
	Send    float32 `json:"send"`
	Wait    float32 `json:"wait"`
	Receive float32 `json:"receive"`
}

// Cache represents data cached from a request's response
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

// Request represents information about the request portion of an Entry
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

// Initiator represents information about the initiator of the request which
// caused the Entry
type Initiator struct {
	Type string `json:"type"`
}

// Entry represents a single entry in the network log
type Entry struct {
	Pageref         string    `json:"pageref"`
	Starteddatetime time.Time `json:"startedDateTime"`
	Response        Response  `json:"response,omitempty"`
	Timings         Timing    `json:"timings"`
	Time            float32   `json:"time"`
	Securitystate   string    `json:"_securityState"`
	Serveripaddress string    `json:"serverIPAddress,omitempty"`
	Connection      string    `json:"connection,omitempty"`
	Cache           Cache     `json:"cache,omitempty"`
	Request         Request   `json:"request,omitempty"`
	Initiator       Initiator `json:"_initiator,omitempty"`
}

// Entries is a group of entries
type Entries []Entry

// ExcludeByURL filters out entries whose request URL matches the supplied regex
func (e Entries) ExcludeByURL(pattern string) Entries {
	result := Entries{}
	for i := range e {
		p := regexp.MustCompile(pattern)
		if !p.MatchString(e[i].Request.URL) {
			result = append(result, e[i])
		}
	}
	return result
}

// ExcludeByResponseHeader filters out entries by named headers whose value matches the supplied regex
func (e Entries) ExcludeByResponseHeader(header, value string) Entries {
	result := Entries{}
	p := regexp.MustCompile(value)
	for i := range e {
		matched := false
		for j := range e[i].Response.Headers {
			h := e[i].Response.Headers[j]
			if h.Name == header && p.MatchString(h.Value) {
				matched = true
			}
		}
		if !matched {
			result = append(result, e[i])
		}
	}
	return result
}

// Log represents the contents of a Har log
type Log struct {
	Version string   `json:"version"`
	Creator Property `json:"creator"`
	Browser Property `json:"browser"`
	Pages   Pages    `json:"pages"`
	Entries Entries  `json:"entries"`
}

// Har represents an entire HTTP Archive
// See the following for more details:
// - https://en.wikipedia.org/wiki/HAR_(file_format)
// - http://www.softwareishard.com/blog/har-12-spec/
// - https://w3c.github.io/web-performance/specs/HAR/Overview.html
type Har struct {
	Log Log `json:"log"`
}

// NewHar creates a new Har from a reader
func NewHar(r io.Reader) (*Har, error) {
	var har Har
	err := json.NewDecoder(r).Decode(&har)
	return &har, err
}
