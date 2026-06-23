package fixedsplitter

import (
	"os"
	"strings"
	"testing"

	"github.com/1Vewton/textsplitter"
)

// Test whether the FixedSplitter implements TextSplitter
func TestInterface(t *testing.T) {
	var splitter interface{} = &FixedSplitter{}
	_, ok := splitter.(textsplitter.TextSplitter)
	if !ok {
		t.Errorf("FixedSplitter does not implements TextSplitter")
	}
}

// Test the SplitText of the FixedSplitter
func TestSplitText(t *testing.T) {
	chunkSize := 60
	overlap := 20
	content, errRead := os.ReadFile("testdata/split_text.txt")
	if errRead != nil {
		t.Fatalf("Fatal error occured when reading testdata due to %s", errRead)
	}
	document := string(content)
	splitter := NewFixedSplitter(
		chunkSize,
		overlap,
	)
	result, errChunk := splitter.SplitText(
		t.Context(),
		document,
	)
	if errChunk != nil {
		t.Fatalf("Fatal error occured when running test due to %s", errChunk)
	}
	t.Log(result)
	for _, i := range result {
		chunkLength := len([]rune(i))
		if chunkLength > chunkSize {
			t.Errorf("The %s is longer than chunk size", i)
		}
		if !strings.Contains(document, i) {
			t.Errorf("%s does not exists", i)
		}
	}
}
