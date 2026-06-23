package textsplitter

import (
	"context"
)

// TextSplitter interface
type TextSplitter interface {
	SplitText(
		ctx context.Context,
		document string,
	) (
		[]string,
		error,
	)
	SplitMultipleTexts(
		ctx context.Context,
		documents []string,
	) (
		[]*SplitResult,
		error,
	)
}
