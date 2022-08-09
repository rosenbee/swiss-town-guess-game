package cantondata

// Abbreviations for names of cantons according to
// https://www.bfs.admin.ch/bfs/en/home/basics/symbols-abbreviations.html
var cantonMap = map[string]string{
	"ZH": "Zurich",
	"BE": "Bern",
	"LU": "Lucerne",
	"UR": "Uri",
	"SZ": "Schwyz",
	"OW": "Obwalden",
	"NW": "Nidwalden",
	"GL": "Glarus",
	"ZG": "Zug",
	"FR": "Fribourg",
	"SO": "Solothurn",
	"BS": "Basel-Stadt",
	"BL": "Basel-Landschaft",
	"SH": "Schaffhausen",
	"AR": "Appenzell A. Rh.",
	"AI": "Appenzell I. Rh.",
	"SG": "St. Gallen",
	"GR": "Graubünden",
	"AG": "Aargau",
	"TG": "Thurgau",
	"TI": "Ticino",
	"VD": "Vaud",
	"VS": "Valais",
	"NE": "Neuchâtel",
	"GE": "Geneva",
	"JU": "Jura",
}

// Get random number between 1 and 770.
// BFS numbers for Swiss towns are in that range.
// Returns nil if canton matching the abbreviation was not found.
func GetCantonName(cantonAbbreviation string) *string {
	cantonName, found := cantonMap[cantonAbbreviation]
	if !found {
		return nil
	}
	return &cantonName
}
