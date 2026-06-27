package recursivesplitter

import (
	"context"
	"strings"

	"github.com/1Vewton/textsplitter/fixedsplitter"
)

// RecursiveSplitter splits the text according to the natural order of the language.
type RecursiveSplitter struct {
	ChunkSize  int
	Overlap    int
	Separators []string
}

// NewRecursiveSplitter creates new recursive splitter
func NewRecursiveSplitter(
	chunkSize int,
	overlap int,
	separators []string,
) *RecursiveSplitter {
	return &RecursiveSplitter{
		ChunkSize:  chunkSize,
		Overlap:    overlap,
		Separators: separators,
	}
}

// forciblySplit splits the text forcibly according to number of characters
func (splitter *RecursiveSplitter) forciblySplit(
	ctx context.Context,
	text string,
) (
	[]string,
	error,
) {
	fixed := fixedsplitter.NewFixedSplitter(
		splitter.ChunkSize,
		splitter.Overlap,
	)
	chunks, err := fixed.SplitText(
		ctx,
		text,
	)
	return chunks, err
}

// recursiveSplit splits single document using recursive splitting method for each step
func (splitter *RecursiveSplitter) recursiveSplit(
	ctx context.Context,
	document string,
	sepIdx int,
) (
	[]string,
	error,
) {
	var result []string = []string{}
	runedDocument := []rune(document)
	// Directly return the document if the length of the document is smaller than ChunkSize
	if len(runedDocument) <= splitter.ChunkSize {
		result = append(result, document)
		return result, nil
	}
	// Forcibly split if all the separators are used.
	if sepIdx >= len(splitter.Separators) {
		return splitter.forciblySplit(ctx, document)
	}
	// Separate the text using the separator
	sep := splitter.Separators[sepIdx]
	parts := strings.Split(document, sep)

	// Start splitting
	currentChunk := ""
	for _, part := range parts {
		candidate := currentChunk
		if candidate != "" {
			candidate += sep
		}
		candidate += part
		if len([]rune(candidate)) < splitter.ChunkSize {
			currentChunk = candidate
		} else {
			if currentChunk != "" {
				result = append(result, currentChunk)
			}
			// If this part is still too long, use recursive method to split using next Splitter
			if len([]rune(part)) > splitter.ChunkSize {
				chunks, err := splitter.recursiveSplit(
					ctx,
					part,
					sepIdx+1,
				)
				if err != nil {
					return chunks, err
				}
				result = append(result, chunks...)
				currentChunk = ""
			} else {
				currentChunk = part
			}
		}
	}
	if currentChunk != "" {
		result = append(result, currentChunk)
	}
	return result, nil
}

// SplitText splits single document using recursive splitting method
func (splitter *RecursiveSplitter) SplitText(
	ctx context.Context,
	document string,
) (
	[]string,
	error,
) {
	result, err := splitter.recursiveSplit(
		ctx,
		document,
		0,
	)
	return result, err
}
