package bible

import (
	"encoding/json"
	"fmt"
	"os"
)

// Package bible provides a simple API to search for quotes in the Bible.
// We load the Bible from
type RawBible map[string]map[int]map[int]Verse

type Bible struct {
	Books map[string]Book
}

type Book struct {
	Chapters map[int]Chapter
}

type Chapter struct {
	Verses map[int]Verse
}

type Verse struct {
	Text string
	Uri  string
	Line int
}

func LoadFromFile() Bible {
	//Open "~/.bible.json"
	dirname, _ := os.UserHomeDir()
	content, err := os.ReadFile(dirname + "/.bible.json")
	if err != nil {
		println("Error reading file", err)
	}
	//Unmarshal the content into a RawBible
	var rawBible RawBible
	err = json.Unmarshal(content, &rawBible)
	if err != nil {
		println("Error unmarshalling raw bible", err)
	}

	//Convert the RawBible into a Bible
	bible := Bible{
		Books: make(map[string]Book),
	}
	for bookName, book := range rawBible {
		bible.Books[bookName] = Book{
			Chapters: make(map[int]Chapter),
		}
		for chapterNumber, chapter := range book {
			bible.Books[bookName].Chapters[chapterNumber] = Chapter{
				Verses: chapter,
			}
			for verseNumber, verse := range chapter {
				bible.Books[bookName].Chapters[chapterNumber].Verses[verseNumber] = verse
			}
		}
	}
	return bible
}

func (b Bible) GetQuoteContent(book string, chapter int, startVerse int, endVerse int) string {
	println("Getting quote content for", book, chapter, startVerse, endVerse)
	content := ""
	for verse := startVerse; verse <= endVerse; verse++ {
		v := b.GetVerse(book, chapter, verse)
		content += fmt.Sprintf("%d. %s \n  ", verse, v.Text)
	}
	return content
}

func (b Bible) GetVerse(book string, chapter int, verse int) Verse {
	if book, ok := b.Books[book]; ok {
		if chapter, ok := book.Chapters[chapter]; ok {
			if verse, ok := chapter.Verses[verse]; ok {
				return verse
			}
		}
	}
	return Verse{}
}
