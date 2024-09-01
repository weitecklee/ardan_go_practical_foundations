package nlp

import (
	"os"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

// var tokenizeCases = []struct {
// 	text   string
// 	tokens []string
// }{
// 	{"Who's on first", []string{"who", "s", "on", "first"}},
// 	{"", nil},
// }

type tokenizeCase struct {
	Text   string
	Tokens []string
}

func loadTokenizeCases(t *testing.T) []tokenizeCase {
	data, err := os.ReadFile("testdata/tokenize_cases.toml")
	require.NoError(t, err, "Read file")

	var testCases struct {
		Cases []tokenizeCase
	}

	err = toml.Unmarshal(data, &testCases)
	require.NoError(t, err, "Unmarshal TOML")
	return testCases.Cases
}

func TestTokenizeTable(t *testing.T) {
	// var testCases struct{ Cases []tokenizeCase }
	// _, err := toml.DecodeFile("tokenize_cases.toml", &testCases)
	// require.NoError(t, err, "DecodeFile")
	// for _, tc := range testCases.Cases {
	for _, tc := range loadTokenizeCases(t) {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "on", "second"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
	// if !reflect.DeepEqual(expected, tokens) {
	// 	t.Fatalf("Expected %#v, got %#v", expected, tokens)
	// }
}

func FuzzTokenize(f *testing.F) {
	f.Fuzz(func(t *testing.T, text string) {
		tokens := Tokenize(text)
		lText := strings.ToLower(text)
		for _, tok := range tokens {
			if !strings.Contains(lText, tok) {
				t.Fatal(tok)
			}
		}
	})
}

// go test -fuzz . -fuzztime 10s -v
