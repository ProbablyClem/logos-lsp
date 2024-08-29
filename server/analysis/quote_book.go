package analysis

import "strings"

func NormalizeBookName(book string) string {
	lowerName := strings.ToLower(book)
	mapping := map[string]string{
		"gen":    "Genèse",
		"gn":     "Genèse",
		"exo":    "Exode",
		"ex":     "Exode",
		"lev":    "Lévitique",
		"lv":     "Lévitique",
		"nb":     "Nombres",
		"No":     "Nombres",
		"deut":   "Deutéronome",
		"dt":     "Deutéronome",
		"jos":    "Josué",
		"judg":   "Juges",
		"ruth":   "Ruth",
		"1sam":   "1 Samuel",
		"2sam":   "2 Samuel",
		"1ro":    "1 Rois",
		"2ro":    "2 Rois",
		"1ch":    "1 Chroniques",
		"2ch":    "2 Chroniques",
		"Ez":     "Esdras",
		"Ne":     "Néhémie",
		"Es":     "Esther",
		"job":    "Job",
		"ps":     "Psaumes",
		"prov":   "Proverbes",
		"eccl":   "Ecclésiaste",
		"song":   "Cantique des Cantiques",
		"isa":    "Ésaïe",
		"jer":    "Jérémie",
		"lam":    "Lamentations",
		"ezek":   "Ézéchiel",
		"dan":    "Daniel",
		"hos":    "Osée",
		"joel":   "Joël",
		"amos":   "Amos",
		"obad":   "Abdias",
		"jonah":  "Jonas",
		"mic":    "Michée",
		"nah":    "Nahum",
		"hab":    "Habacuc",
		"zeph":   "Sophonie",
		"hag":    "Aggée",
		"zech":   "Zacharie",
		"mal":    "Malachie",
		"matt":   "Matthieu",
		"mt":     "Matthieu",
		"mr":     "Marc",
		"marc":   "Marc",
		"lc":     "Luc",
		"jn":     "Jean",
		"act":    "Actes",
		"rom":    "Romains",
		"ro":     "Romains",
		"1cor":   "1 Corinthiens",
		"1co":    "1 Corinthiens",
		"2cor":   "2 Corinthiens",
		"2co":    "2 Corinthiens",
		"gal":    "Galates",
		"ga":     "Galates",
		"eph":    "Éphésiens",
		"ep":     "Éphésiens",
		"phil":   "Philippiens",
		"ph":     "Philippiens",
		"col":    "Colossiens",
		"co":     "Colossiens",
		"1thess": "1 Thessaloniciens",
		"1th":    "1 Thessaloniciens",
		"2thess": "2 Thessaloniciens",
		"2th":    "2 Thessaloniciens",
		"1tim":   "1 Timothée",
		"2tim":   "2 Timothée",
		"titus":  "Tite",
		"Tt":     "Tite",
		"phlm":   "Philémon",
		"heb":    "Hébreux",
		"hb":     "Hébreux",
		"jas":    "Jacques",
		"ja":     "Jacques",
		"1pi":    "1 Pierre",
		"2pi":    "2 Pierre",
		"1jn":    "1 Jean",
		"2jn":    "2 Jean",
		"3jn":    "3 Jean",
		"ju":     "Jude",
		"ap":     "Apocalypse",
	}

	if val, ok := mapping[lowerName]; ok {
		return val
	}
	book = capitalizeFirstLetter(book)
	return book
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}
