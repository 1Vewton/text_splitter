package textsplitter

import (
	"context"
)

// TextSplitter interface
type TextSplitter interface {
	// SplitText splits single text
	SplitText(
		ctx context.Context,
		document string,
	) (
		[]string,
		error,
	)
	// SplitMultipleTexts splits multiple texts
	SplitMultipleTexts(
		ctx context.Context,
		documents []string,
	) (
		[]*SplitResult,
		error,
	)
}
