package hoofli

import (
	"fmt"
	"io"
)

func (h *Har) Draw(w io.Writer) error {
	fmt.Fprintln(w, "@startuml")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "participant Browser")
	fmt.Fprintln(w, "")
	for _, p := range h.Log.Pages {
		fmt.Fprintf(w, "->Browser : %s\n", p.Title)
		fmt.Fprintln(w, "activate Browser")
		fmt.Fprintln(w, "deactivate Browser")
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "@enduml")
	return nil
}
