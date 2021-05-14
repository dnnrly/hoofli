package hoofli_test

import (
	"bytes"
	_ "embed"
	"os"
	"strings"
	"testing"

	"github.com/dnnrly/hoofli"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed test/reference/plantuml/simple-example.puml
	simpleExample string
	//go:embed test/reference/plantuml/multipage-example.puml
	multipageExample string
)

func TestCreatesHarFromGooglePage(t *testing.T) {
	file, err := os.Open("test/reference/har/google-frontpage.har")
	require.NoError(t, err)

	har, err := hoofli.NewHar(file)
	require.NoError(t, err)
	require.Equal(t, 1, len(har.Log.Pages))
	require.Equal(t, 22, len(har.Log.Entries))
}

func TestCreatesHarFromHackernewsPages(t *testing.T) {
	file, err := os.Open("test/reference/har/hackernews.har")
	require.NoError(t, err)

	har, err := hoofli.NewHar(file)
	require.NoError(t, err)
	require.Equal(t, 2, len(har.Log.Pages))
	require.Equal(t, 12, len(har.Log.Entries))
}

func TestNewHar_ParsingFailure(t *testing.T) {
	_, err := hoofli.NewHar(strings.NewReader("{xyz"))
	require.Error(t, err)
}

func TestDrawHar_SinglePage(t *testing.T) {
	har := hoofli.Har{
		Log: hoofli.Log{
			Pages: []hoofli.Page{{
				ID:    "page-1",
				Title: "Example",
			}},
			Entries: []hoofli.Entry{{
				Pageref: "page-1",
				Request: hoofli.Request{
					Method: "GET",
					URL:    "https://example.com/page-1",
				},
				Response: hoofli.Response{
					Status: 200,
				},
			}},
		},
	}

	var output bytes.Buffer
	err := har.Draw(&output)

	require.NoError(t, err)
	require.Equal(t, output.String(), simpleExample)
}

func TestDrawHar_MultiPage(t *testing.T) {
	har := hoofli.Har{
		Log: hoofli.Log{
			Pages: []hoofli.Page{
				{
					ID:    "page-1",
					Title: "Example",
				},
				{
					ID:    "page-2",
					Title: "Another Example",
				},
			},
			Entries: []hoofli.Entry{
				{
					Pageref: "page-1",
					Request: hoofli.Request{
						Method: "GET",
						URL:    "https://example.com/page-1",
					},
					Response: hoofli.Response{
						Status: 200,
					},
				},
				{
					Pageref: "page-2",
					Request: hoofli.Request{
						Method: "GET",
						URL:    "https://example.com/page-2",
					},
					Response: hoofli.Response{
						Status: 200,
					},
				},
			},
		},
	}

	var output bytes.Buffer
	err := har.Draw(&output)

	require.NoError(t, err)
	require.Equal(t, output.String(), multipageExample)
}

func TestEntriesURLFilter_FixedPattern(t *testing.T) {
	entries := hoofli.Entries{
		{Request: hoofli.Request{URL: "https://example.com/page-1"}},
		{Request: hoofli.Request{URL: "https://example.com/page-2"}},
		{Request: hoofli.Request{URL: "https://another.com/"}},
	}

	filtered := entries.ExcludeByURL("another")

	require.Len(t, filtered, 2)
}

func TestEntriesURLFilter_Regex(t *testing.T) {
	entries := hoofli.Entries{
		{Request: hoofli.Request{URL: "https://example.com/page-1"}},
		{Request: hoofli.Request{URL: "https://example.com/page-2"}},
		{Request: hoofli.Request{URL: "https://another.com/"}},
	}

	filtered := entries.ExcludeByURL("(another|2)")

	require.Len(t, filtered, 1)
}

func TestEntriesResponseHeaderFilter_RegexOnHeaderValue(t *testing.T) {
	entries := hoofli.Entries{
		{Response: hoofli.Response{Headers: hoofli.Properties{
			{Name: "some-header", Value: "application/json"},
			{Name: "content-type", Value: "application/text"},
		}}},
		{Response: hoofli.Response{Headers: hoofli.Properties{{Name: "content-type", Value: "application/json"}}}},
		{Response: hoofli.Response{Headers: hoofli.Properties{{Name: "content-type", Value: "text/json"}}}},
		{Response: hoofli.Response{Headers: hoofli.Properties{{Name: "content-type", Value: "application/yaml"}}}},
		{Response: hoofli.Response{Headers: hoofli.Properties{{Name: "content-type-extra", Value: "application/yaml"}}}},
	}

	require.Len(t, entries.ExcludeByResponseHeader("content-type", ".*"), 1)
	require.Len(t, entries.ExcludeByResponseHeader("some-header", ".*"), 4)
	require.Len(t, entries.ExcludeByResponseHeader("content-type", ".+json"), 3)
	require.Len(t, entries.ExcludeByResponseHeader("content-type", "json"), 3)
}
