package recursivesplitter

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
