package hoofli

import (
	"fmt"
	"io"
	"net/url"
)

func (h *Har) Draw(w io.Writer) error {
	fmt.Fprintln(w, "@startuml")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "participant Browser")
	fmt.Fprintln(w, "")
	for _, p := range h.Log.Pages {
		fmt.Fprintf(w, "->Browser : %s\n", p.Title)
		fmt.Fprintln(w, "activate Browser")
		for _, e := range h.Log.Entries {
			dest := e.Request.URL
			if url, err := url.Parse(dest); err == nil {
				dest = url.Host
			}
			fmt.Fprintf(w, "Browser->\"%s\" ++ : %s %s\n", dest, e.Request.Method, e.Request.URL)
			fmt.Fprintf(w, "return %d\n", e.Response.Status)
		}
		fmt.Fprintln(w, "deactivate Browser")
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "@enduml")
	return nil
}
