package utils

import (
	"math"
	"strconv"
	"strings"
)

const (
	MATH_POW_QUECTO = .000_000_000_000_000_000_000_000_000_001
	MATH_POW_RONTO  = .000_000_000_000_000_000_000_000_001
	MATH_POW_YOCTO  = .000_000_000_000_000_000_000_001
	MATH_POW_ZEPTO  = .000_000_000_000_000_000_001
	MATH_POW_ATTO   = .000_000_000_000_000_001
	MATH_POW_FEMTO  = .000_000_000_000_001
	MATH_POW_PICO   = .000_000_000_001
	MATH_POW_NANO   = .000_000_001
	MATH_POW_MICRO  = .000_001
	MATH_POW_MILLI  = .001

	MATH_POW_KILO   = 1_000.
	MATH_POW_MEGA   = 1_000_000.
	MATH_POW_GIGA   = 1_000_000_000.
	MATH_POW_TERA   = 1_000_000_000_000.
	MATH_POW_PETA   = 1_000_000_000_000_000.
	MATH_POW_EXA    = 1_000_000_000_000_000_000.
	MATH_POW_ZETTA  = 1_000_000_000_000_000_000_000.
	MATH_POW_YOTTA  = 1_000_000_000_000_000_000_000_000.
	MATH_POW_RONNA  = 1_000_000_000_000_000_000_000_000_000.
	MATH_POW_QUETTA = 1_000_000_000_000_000_000_000_000_000_000.
)

const (
	UNIT_TYPE_PERCENT = iota
	UNIT_TYPE_EXACT
)

type Tolerance struct {
	Value    float64
	UnitType int
}

type ColorBand struct {
	SignificantNumeral int
	Multiplier         int
	Ansi               string
}

func GetAbbreviatedValue(val float64) string {
	for i := range si_prefixes_positive_base10 {
		pow10 := math.Pow10(30 - (i * 3))
		if val >= pow10 {
			return FormatFloat(val/pow10) + string(si_prefixes_positive_base10[i])
		}
	}

	if val < 1 && val > 0 {
		for i := range si_prefixes_negative_base10 {
			pow10 := math.Pow10(-3 - (i * 3))
			if val >= pow10 {
				return FormatFloat(val/pow10) + string(si_prefixes_negative_base10[i])
			}
		}
	}

	return FormatFloat(val)
}

func FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func IsDigit[T ~byte | ~rune](t T) bool {
	return t >= 48 && t <= 57
}

func IsLetter[T ~byte | ~rune](t T) bool {
	return IsLowerLetter(t) || IsUpperLetter(t)
}

func IsLowerLetter[T ~byte | ~rune](t T) bool {
	return 97 <= t && t <= 122
}

func IsUpperLetter[T ~byte | ~rune](t T) bool {
	return 65 <= t && t <= 90
}

func If[T any](condition bool, tru T, fals T) T {
	if !condition {
		return fals
	}
	return tru
}

func get_first_rune(s string) (rune, int) {
	for _, r := range s {
		return r, len(string(r))
	}
	return 0, 0
}

func get_leading_digits(val string, supports_decimals bool) (float64, int, error) {
	var sb strings.Builder
	i := 0
	for index, r := range val {
		if (supports_decimals && 46 == r) || 48 <= r && r <= 57 {
			sb.WriteRune(r)
			i = index
		} else {
			break
		}
	}

	digits := sb.String()
	if len(digits) == 0 {
		return 0., 0, nil
	}

	f, err := strconv.ParseFloat(digits, 64)
	if err != nil {
		return 0., -1, err
	}

	return f, i + 1, nil
}

// starts_with_any checks if s starts with any of the given prefixes.
// Returns the length of the matching prefix, or 0 if no match.
func starts_with_any(s string, prefixes []string) int {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return len(prefix)
		}
	}
	return 0
}
