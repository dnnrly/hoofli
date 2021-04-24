package hoofli_test

import (
	"os"
	"strings"
	"testing"

	"github.com/dnnrly/hoofli"
	"github.com/stretchr/testify/require"
)

func TestCreatesHarFromGooglePage(t *testing.T) {
	file, err := os.Open("test/reference/har/google-frontpage.har")
	require.NoError(t, err)

	har, err := hoofli.NewHar(file)
	require.NoError(t, err)
	require.Equal(t, 1, len(har.Log.Pages))
	require.Equal(t, 22, len(har.Log.Entries))
}

func TestNewHar_ParsingFailure(t *testing.T) {
	_, err := hoofli.NewHar(strings.NewReader("{xyz"))
	require.Error(t, err)
}
