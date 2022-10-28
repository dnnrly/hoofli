package hoofli

import (
	"fmt"
	"io"
	"net/url"
)

// Draw writes a plantuml formatted sequence diagram representing the Har to the writer
func (h *Har) Draw(w io.Writer) error {
	// initiatorTypesUsed stores the used initiator types for rendering the
	// correct legend at the end. Using a map over a list gives us automatic
	// deduplication.
	initiatorTypesUsed := map[string]bool{
		"script":   false,
		"renderer": false,
		"other":    false,
		"":         false,
	}
	fmt.Fprintln(w, "@startuml")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "participant Browser")
	fmt.Fprintln(w, "")
	for _, p := range h.Log.Pages {
		fmt.Fprintf(w, "->Browser : %s\n", p.Title)
		fmt.Fprintln(w, "activate Browser")
		for i := range h.Log.Entries {
			if p.ID == h.Log.Entries[i].Pageref {
				dest := h.Log.Entries[i].Request.URL
				if parsedURL, err := url.Parse(dest); err == nil {
					dest = parsedURL.Host
				}
				fmt.Fprintf(w, "Browser-[#%s]->\"%s\" ++ : %s %s\n",
					InitiatorTypeToColor(h.Log.Entries[i].Initiator.Type),
					dest,
					h.Log.Entries[i].Request.Method,
					h.Log.Entries[i].Request.URL,
				)
				fmt.Fprintf(w, "return %d\n", h.Log.Entries[i].Response.Status)
				// remember we used an initiator of this type
				initiatorTypesUsed[h.Log.Entries[i].Initiator.Type] = true
			}
		}
		// add a note explaining the colors of the connections
		if shouldRenderLegend(initiatorTypesUsed) {
			fmt.Fprintf(w, "note over Browser: Connection color represents initiator type:")
			for t, _ := range initiatorTypesUsed {
				color := InitiatorTypeToColor(t)
				fmt.Fprintf(w, "\\n<font color=%s>%s (%s)</font>", color, t, color)
			}
			fmt.Fprintf(w, "\n")
		}
		fmt.Fprintln(w, "deactivate Browser")
	}
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "@enduml")
	return nil
}

// InitiatorTypeToColor converts the value found in
// log.entries[]._initiator.type to a drawable color for plantuml. If the type
// is not known yet, it will be rendered in the color specified as defaultColor.
func InitiatorTypeToColor(strType string) string {
	const defaultColor = "black"
	colors := map[string]string{
		"script":   "red",
		"renderer": "blue",
		"other":    "green",
	}
	out, ok := colors[strType]
	if !ok {
		return defaultColor
	}
	return out
}

// shouldRenderLegend decides if the legend explaining the initiator types
// should be rendered. It will decide to render the legend if at least one
// initiator type other than the "unspecified" type was used.
func shouldRenderLegend(initiatorTypesUsed map[string]bool) bool {
	for initiatorType, wasUsed := range initiatorTypesUsed {
		if initiatorType != "" && wasUsed {
			return true
		}
	}
	return false
}
