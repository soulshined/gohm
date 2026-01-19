package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseShorthand(val string, any_of_targets []string) float64 {
	len_val := len(val)
	digits, digits_end_index, err := get_leading_digits(val, true)

	if err != nil {
		panic(err)
	}

	if digits_end_index >= len_val {
		return digits
	}

	// Get the prefix as a rune - handles multi-byte characters like μ
	remaining := val[digits_end_index:]
	prefix_rune, len_prefix_byte := get_first_rune(remaining)

	// Check if remaining string starts with any target ("Hz", "V", "A")
	if len_target := starts_with_any(remaining, any_of_targets); len_target > 0 && len_target+digits_end_index == len_val {
		return digits
	}

	digits_end_index += len_prefix_byte

	si_prefix, ok := SI_MAPPING[prefix_rune]
	if !ok {
		panic(fmt.Errorf("invalid or unsupported: si prefix %s", string(prefix_rune)))
	}

	if digits_end_index < len_val {
		// Check if remaining string starts with any target
		target_remaining := val[digits_end_index:]
		len_target := starts_with_any(target_remaining, any_of_targets)
		if len_target == 0 {
			panic(fmt.Errorf("invalid: type identifier %s for target", target_remaining))
		}
		digits_end_index += len_target
		if digits_end_index < len_val {
			panic(fmt.Errorf("invalid or unsupported: type identifier %s", val[digits_end_index:]))
		}
	}

	return digits * si_prefix.Pow10
}

func ParseRKMCode(val string, target rune) (float64, error) {
	runes := []rune(val)
	len_rune := len(runes)

	if len_rune < 2 || len_rune > 5 {
		return 0., fmt.Errorf("invalid or unsupported: RKM notation %s", val)
	}

	t, ok := RKM_MAPPING[target]
	if !ok {
		return 0., fmt.Errorf("invalid or unsupported: RKM target %s", string(target))
	}

	_, digits_end_index, err := get_leading_digits(val, false)

	if err != nil {
		return 0., err
	}

	if digits_end_index >= len(val) {
		return 0., fmt.Errorf("invalid or unsupported: RKM notation %s", val)
	}

	remaining := val[digits_end_index:]
	prefix_rune, _ := get_first_rune(remaining)
	pow10 := 1.

	if target != prefix_rune {
		code_letter, ok := t[prefix_rune]
		if !ok {
			return 0., fmt.Errorf("invalid or unsupported: prefix %s for target %s", string(prefix_rune), string(target))
		}
		pow10 = code_letter.Pow10
	}

	replaced := strings.Replace(val, string(prefix_rune), ".", 1)

	as_float, err := strconv.ParseFloat(replaced, 64)
	if err != nil {
		return 0., err
	}

	return as_float * pow10, nil
}

func GetValueForRKMElseShorthand(flag_value string, rkm_target rune, shorthand_targets []string) float64 {
	val, err := ParseRKMCode(flag_value, rkm_target)
	if err != nil {
		val = ParseShorthand(flag_value, shorthand_targets)
	}
	return val
}
