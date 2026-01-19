package utils

var EIA_COLOR_MAPPING = map[string]ColorBand{
	"black": {
		SignificantNumeral: 0,
		Multiplier:         0,
		Ansi:               ANSI_BLACK_FG,
	},
	"brown": {
		SignificantNumeral: 1,
		Multiplier:         1,
		Ansi:               ANSI_BROWN_FG,
	},
	"red": {
		SignificantNumeral: 2,
		Multiplier:         2,
		Ansi:               ANSI_RED_BRIGHT_FG,
	},
	"orange": {
		SignificantNumeral: 3,
		Multiplier:         3,
		Ansi:               ANSI_ORANGE_FG,
	},
	"yellow": {
		SignificantNumeral: 4,
		Multiplier:         4,
		Ansi:               ANSI_YELLOW_BRIGHT_FG,
	},
	"green": {
		SignificantNumeral: 5,
		Multiplier:         5,
		Ansi:               ANSI_GREEN_BRIGHT_FG,
	},
	"blue": {
		SignificantNumeral: 6,
		Multiplier:         6,
		Ansi:               ANSI_BLUE_BRIGHT_FG,
	},
	"violet": {
		SignificantNumeral: 7,
		Multiplier:         7,
		Ansi:               ANSI_VIOLET_FG,
	},
	"grey": {
		SignificantNumeral: 8,
		Multiplier:         8,
		Ansi:               ANSI_GRAY_FG,
	},
	"white": {
		SignificantNumeral: 9,
		Multiplier:         9,
		Ansi:               ANSI_WHITE_BRIGHT_FG,
	},
	"gold": {
		SignificantNumeral: -1,
		Multiplier:         -1,
		Ansi:               ANSI_GOLD_FG,
	},
	"silver": {
		SignificantNumeral: -1,
		Multiplier:         -2,
		Ansi:               ANSI_GOLD_FG,
	},
	"pink": {
		SignificantNumeral: -1,
		Multiplier:         -3,
		Ansi:               ANSI_PINK_FG,
	},
}
