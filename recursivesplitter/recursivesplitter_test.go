package recursivesplitter

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

// Test the split text method
func TestSplitText(t *testing.T) {
	chunkSize := 60
	overlap := 20
	content, errRead := os.ReadFile("testdata/split_text_1.md")
	if errRead != nil {
		t.Fatalf("Fatal error occured when reading testdata due to %s", errRead)
	}
	document := string(content)
	splitter := NewRecursiveSplitter(
		chunkSize,
		overlap,
		[]string{"\n\n", "\n", "。", "，", " ", ",", "."},
	)
	// Timeout checking
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	result, errChunk := splitter.SplitText(
		ctx,
		document,
	)
	if errChunk != nil {
		t.Fatalf("Fatal error occured when running test due to %s", errChunk)
	}
	t.Log(result)
	t.Log(len(result))
	for _, i := range result {
		if !strings.Contains(document, i) {
			t.Errorf("%s does not exists in original document", i)
		}
	}
}
