package textsplitter

// SplitResult provides the result of splitting text
type SplitResult struct {
	FullText    string
	ChunkResult string
}

// TempSplitResult stores the temporary splitting result
type TempSplitResult struct {
	FullText    string
	ChunkResult []string
}
