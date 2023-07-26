package main

import "fmt"

var slippiCharactersUrl = "https://slippi.gg/images/characters/stock-icon-%v-0.png"

var SlippiCharactersId = map[string]int{
	"DONKEY_KONG":    0, // Mapped to 255 in database
	"CAPTAIN_FALCON": 1,
	"FOX":            2,
	"GAME_AND_WATCH": 3,
	"KIRBY":          4,
	"BOWSER":         5,
	"LINK":           6,
	"LUIGI":          7,
	"MARIO":          8,
	"MARTH":          9,
	"MEWTWO":         10,
	"NESS":           11,
	"PEACH":          12,
	"PIKACHU":        13,
	"ICE_CLIMBERS":   14,
	"JIGGLYPUFF":     15,
	"SAMUS":          16,
	"YOSHI":          17,
	"ZELDA":          18,
	"SHEIK":          19,
	"FALCO":          20,
	"YOUNG_LINK":     21,
	"DR_MARIO":       22,
	"ROY":            23,
	"PICHU":          24,
	"GANONDORF":      25,
	"None":           256,
}

var SlippiCharacterColors = map[string]string{
	"DONKEY_KONG":    "#2f1003",
	"CAPTAIN_FALCON": "#c51620",
	"FOX":            "#ffb242",
	"GAME_AND_WATCH": "#000000",
	"KIRBY":          "#ffbed8",
	"BOWSER":         "#376218",
	"LINK":           "#073f07",
	"LUIGI":          "#10b91a",
	"MARIO":          "#ff1d1c",
	"MARTH":          "#2f3955",
	"MEWTWO":         "#734c60",
	"NESS":           "#f9ca58",
	"PEACH":          "#ff5488",
	"PIKACHU":        "#ffff00",
	"ICE_CLIMBERS":   "#8a63ff",
	"JIGGLYPUFF":     "#ffd6f0",
	"SAMUS":          "#da490c",
	"YOSHI":          "#008000",
	"ZELDA":          "#ff6ac8",
	"SHEIK":          "#828681",
	"FALCO":          "#494fd6",
	"YOUNG_LINK":     "#009e01",
	"DR_MARIO":       "#d1cfc9",
	"ROY":            "#962000",
	"PICHU":          "#ffff1b",
	"GANONDORF":      "#91763e",
}

func get_key_from_value(value int, lookup map[string]int) string {
	for k, v := range lookup {
		if v == value {
			return k
		}
	}
	return ""
}

func get_character_name(charId int) string {
	if charId == 255 {
		charId = 0
	}
	return get_key_from_value(charId, SlippiCharactersId)
}

func get_character_id(name string, dk_clause bool) int {
	characterId := SlippiCharactersId[name]
	if dk_clause && characterId == 0 {
		return 255
	}
	return characterId
}

func get_character_url(name string) string {
	return fmt.Sprintf(slippiCharactersUrl, get_character_id(name, false))
}
