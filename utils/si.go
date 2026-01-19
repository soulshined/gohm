package utils

import "gohm/abbrvs"

type SIPrefix struct {
	Name  string
	Pow10 float64
}

var SI_MAPPING = map[rune]SIPrefix{
	abbrvs.SI_QUECTO: {
		Name:  "quecto",
		Pow10: MATH_POW_QUECTO,
	},
	abbrvs.SI_RONTO: {
		Name:  "ronto",
		Pow10: MATH_POW_RONTO,
	},
	abbrvs.SI_YOCTO: {
		Name:  "yocto",
		Pow10: MATH_POW_YOCTO,
	},
	abbrvs.SI_ZEPTO: {
		Name:  "zepto",
		Pow10: MATH_POW_ZEPTO,
	},
	abbrvs.SI_ATTO: {
		Name:  "atto",
		Pow10: MATH_POW_ATTO,
	},
	abbrvs.SI_FEMTO: {
		Name:  "femto",
		Pow10: MATH_POW_FEMTO,
	},
	abbrvs.SI_PICO: {
		Name:  "pico",
		Pow10: MATH_POW_PICO,
	},
	abbrvs.SI_NANO: {
		Name:  "nano",
		Pow10: MATH_POW_NANO,
	},
	abbrvs.SI_MICRO: {
		Name:  "micro",
		Pow10: MATH_POW_MICRO,
	},
	abbrvs.SI_MILLI: {
		Name:  "milli",
		Pow10: MATH_POW_MILLI,
	},
	abbrvs.SI_KILO: {
		Name:  "kilo",
		Pow10: MATH_POW_KILO,
	},
	abbrvs.SI_MEGA: {
		Name:  "mega",
		Pow10: MATH_POW_MEGA,
	},
	abbrvs.SI_GIGA: {
		Name:  "giga",
		Pow10: MATH_POW_GIGA,
	},
	abbrvs.SI_TERA: {
		Name:  "tera",
		Pow10: MATH_POW_TERA,
	},
	abbrvs.SI_PETA: {
		Name:  "peta",
		Pow10: MATH_POW_PETA,
	},
	abbrvs.SI_EXA: {
		Name:  "exa",
		Pow10: MATH_POW_EXA,
	},
	abbrvs.SI_ZETTA: {
		Name:  "zetta",
		Pow10: MATH_POW_ZETTA,
	},
	abbrvs.SI_YOTTA: {
		Name:  "yotta",
		Pow10: MATH_POW_YOTTA,
	},
	abbrvs.SI_RONNA: {
		Name:  "ronna",
		Pow10: MATH_POW_RONNA,
	},
	abbrvs.SI_QUETTA: {
		Name:  "quetta",
		Pow10: MATH_POW_QUETTA,
	},
}

var si_prefixes_positive_base10 = []rune{
	abbrvs.SI_QUETTA,
	abbrvs.SI_RONNA,
	abbrvs.SI_YOTTA,
	abbrvs.SI_ZETTA,
	abbrvs.SI_EXA,
	abbrvs.SI_PETA,
	abbrvs.SI_TERA,
	abbrvs.SI_GIGA,
	abbrvs.SI_MEGA,
	abbrvs.SI_KILO,
}

var si_prefixes_negative_base10 = []rune{
	abbrvs.SI_MILLI,
	abbrvs.SI_MICRO,
	abbrvs.SI_NANO,
	abbrvs.SI_PICO,
	abbrvs.SI_FEMTO,
	abbrvs.SI_ATTO,
	abbrvs.SI_ZEPTO,
	abbrvs.SI_YOCTO,
	abbrvs.SI_RONTO,
	abbrvs.SI_QUECTO,
}
