package nlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "s", "on", "second"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
	// if !reflect.DeepEqual(expected, tokens) {
	// 	t.Fatalf("Expected %#v, got %#v", expected, tokens)
	// }
}
