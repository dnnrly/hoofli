package hoofli

import "time"

type Property struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Properties []Property

type Content struct {
	Mimetype string    `json:"mimeType"`
	Size     int       `json:"size"`
	Text     time.Time `json:"text"`
}

type Har struct {
	Log struct {
		Version string   `json:"version"`
		Creator Property `json:"creator"`
		Browser Property `json:"browser"`
		Pages   []struct {
			Starteddatetime time.Time `json:"startedDateTime"`
			ID              string    `json:"id"`
			Title           string    `json:"title"`
			Pagetimings     struct {
				Oncontentload int `json:"onContentLoad"`
				Onload        int `json:"onLoad"`
			} `json:"pageTimings"`
		} `json:"pages"`
		Entries []struct {
			Pageref         string    `json:"pageref"`
			Starteddatetime time.Time `json:"startedDateTime"`
			Response        struct {
				Status      int        `json:"status"`
				Statustext  string     `json:"statusText"`
				Httpversion string     `json:"httpVersion"`
				Headers     Properties `json:"headers"`
				Cookies     Properties `json:"cookies"`
				Content     Content    `json:"content"`
				Redirecturl string     `json:"redirectURL"`
				Headerssize int        `json:"headersSize"`
				Bodysize    int        `json:"bodySize"`
			} `json:"response,omitempty"`
			Timings struct {
				Blocked int `json:"blocked"`
				DNS     int `json:"dns"`
				Connect int `json:"connect"`
				Ssl     int `json:"ssl"`
				Send    int `json:"send"`
				Wait    int `json:"wait"`
				Receive int `json:"receive"`
			} `json:"timings"`
			Time            int    `json:"time"`
			Securitystate   string `json:"_securityState"`
			Serveripaddress string `json:"serverIPAddress,omitempty"`
			Connection      string `json:"connection,omitempty"`
			Cache           struct {
				Afterrequest struct {
					Expires      string `json:"expires"`
					Lastfetched  string `json:"lastFetched"`
					Etag         string `json:"eTag"`
					Fetchcount   string `json:"fetchCount"`
					Datasize     string `json:"_dataSize"`
					Lastmodified string `json:"_lastModified"`
					Device       string `json:"_device"`
				} `json:"afterRequest"`
			} `json:"cache,omitempty"`
			Request struct {
				Bodysize    int        `json:"bodySize"`
				Method      string     `json:"method"`
				URL         string     `json:"url"`
				Httpversion string     `json:"httpVersion"`
				Headers     Properties `json:"headers"`
				Cookies     Properties `json:"cookies"`
				Querystring Properties `json:"queryString"`
				Headerssize int        `json:"headersSize"`
				Postdata    Content    `json:"postData"`
			} `json:"request,omitempty"`
		} `json:"entries"`
	} `json:"log"`
}
