package text_splitter

import (
	"context"
)

// Text splitter interface
type TextSplitter interface {
	SplitText(ctx context.Context) ([]string, error)
	SplitMultipleTexts(ctx context.Context) ([]*SplitResult, error)
}
