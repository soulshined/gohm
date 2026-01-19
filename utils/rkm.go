package utils

import "gohm/abbrvs"

var RKM_MAPPING = map[rune]map[rune]SIPrefix{
	abbrvs.RKM_RESISTOR: { // resistance
		'L': SI_MAPPING[abbrvs.SI_MILLI],
		'K': SI_MAPPING[abbrvs.SI_KILO],
		'M': SI_MAPPING[abbrvs.SI_MEGA],
		'G': SI_MAPPING[abbrvs.SI_GIGA],
		'T': SI_MAPPING[abbrvs.SI_TERA],
	},
	abbrvs.RKM_FARAD: { // capacitor | farad
		'p': SI_MAPPING[abbrvs.SI_PICO],
		'n': SI_MAPPING[abbrvs.SI_NANO],
		'μ': SI_MAPPING[abbrvs.SI_MICRO],
		'L': SI_MAPPING[abbrvs.SI_MILLI],
		'K': SI_MAPPING[abbrvs.SI_KILO],
		'M': SI_MAPPING[abbrvs.SI_MEGA],
		'G': SI_MAPPING[abbrvs.SI_GIGA],
		'T': SI_MAPPING[abbrvs.SI_TERA],
	},
}
