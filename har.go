package hoofli

import (
	"encoding/json"
	"io"
	"regexp"
	"strings"
	"time"
)

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (p Property) String() string {
	return p.Name + "=" + p.Value
}

type Properties []Property

func (p Properties) String() string {
	ps := []string{}
	for _, v := range p {
		ps = append(ps, v.String())
	}
	return "[" + strings.Join(ps, ", ") + "]"
}

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
		Oncontentload float32 `json:"onContentLoad"`
		Onload        float32 `json:"onLoad"`
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
	Blocked float32 `json:"blocked"`
	DNS     float32 `json:"dns"`
	Connect float32 `json:"connect"`
	Ssl     float32 `json:"ssl"`
	Send    float32 `json:"send"`
	Wait    float32 `json:"wait"`
	Receive float32 `json:"receive"`
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
	Time            float32   `json:"time"`
	Securitystate   string    `json:"_securityState"`
	Serveripaddress string    `json:"serverIPAddress,omitempty"`
	Connection      string    `json:"connection,omitempty"`
	Cache           Cache     `json:"cache,omitempty"`
	Request         Request   `json:"request,omitempty"`
}

type Entries []Entry

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

func NewHar(r io.Reader) (*Har, error) {
	var har Har
	err := json.NewDecoder(r).Decode(&har)
	return &har, err
}
