package analysis

import (
	"logos-lsp/lsp"
	"regexp"
	"strconv"
	"strings"
)

// Biblical quote
// Struct to store Bible quote along with its position (line and character number)
type Quote struct {
	Reference
	Uri   string
	Range lsp.Range
}

type Reference struct {
	Book       string
	Chapter    int
	StartVerse int
	EndVerse   int
}

// Function to find Bible quotes and their positions (line number, char position)
func FindBibleQuotesWithPosition(uri string, text string) []Quote {
	// Regex pattern to match different Bible reference formats
	pattern := `(?i)\b([A-Za-z]+)\s+(\d+)[\s:\-](\d+)(?:-(\d+))?\b`
	re := regexp.MustCompile(pattern)

	// Split the text by lines
	lines := strings.Split(text, "\n")
	var results []Quote

	// Iterate through each line, keeping track of the line number
	for lineNumber, line := range lines {
		// Find all matches within the line
		matches := re.FindAllStringIndex(line, -1)

		// Iterate over the matches and collect the quote along with position info
		for _, match := range matches {
			quote := line[match[0]:match[1]]
			charPosition := match[0] // Convert from index (0-based) to char position (1-based)

			start := lsp.Position{Line: lineNumber, Character: charPosition}
			end := lsp.Position{Line: lineNumber, Character: charPosition + len(quote)}
			reference := ParseReference(quote)

			results = append(results, Quote{
				Reference: reference,
				Uri:       uri,
				Range: lsp.Range{
					Start: start,
					End:   end,
				},
			})
		}
	}

	return results
} // Function to find Bible quotes

func ParseReference(quote string) Reference {
	// Regex pattern to match different Bible reference formats
	pattern := `(?i)\b([A-Za-z]+)\s+(\d+)[\s:\-](\d+)(?:-(\d+))?\b`
	re := regexp.MustCompile(pattern)

	// Find the first match
	match := re.FindStringSubmatch(quote)

	// Extract the book, chapter, and verse
	book := match[1]
	chapter := match[2]
	startVerse := match[3]
	var endVerseInt int

	// Convert chapter and verse to integers
	// Handle errors if conversion fails
	chapterInt, _ := strconv.Atoi(chapter)
	startVerseInt, _ := strconv.Atoi(startVerse)

	println("match[4]", match[4])
	if match[4] != "" {
		endVerse := match[4]
		verseInt, _ := strconv.Atoi(endVerse)
		endVerseInt = verseInt
	} else {
		endVerseInt = startVerseInt
	}

	return Reference{
		Book:       NormalizeBookName(book),
		Chapter:    chapterInt,
		StartVerse: startVerseInt,
		EndVerse:   endVerseInt,
	}
}

// Check if the quote is in the range
func (q Quote) IsInRange(line int, character int) bool {
	return q.Range.Start.Line <= line && q.Range.End.Line >= line && q.Range.Start.Character <= character && q.Range.End.Character >= character
}
