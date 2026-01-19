package identify

import (
	"fmt"
	"gohm/cli"
	"gohm/utils"
	"math"
)

func cmd_capacitor_handler(cmd *cli.Command) string {
	format := cmd.GetFlagValue("format")

	if cmd.IsFlagSet("eiac") {
		return get_capacitance_from_eia(cmd.GetFlagValue("eiac"), format)
	}

	panic("unsupported: identify capacitor flags")
}

func get_capacitance_from_eia(val string, format string) string {
	len_val := len(val)

	result_pf := 0.
	var result_tolerance_min *utils.Tolerance = nil
	var result_tolerance_max *utils.Tolerance = nil

	switch len_val {
	case 2:
		v1 := val[0]
		v2 := val[1]

		if utils.IsLetter(v1) && utils.IsDigit(v2) {
			// 2 digit EIA-198
			switch v1 {
			case 'A':
				result_pf = 1
			case 'B':
				result_pf = 1.1
			case 'C':
				result_pf = 1.2
			case 'D':
				result_pf = 1.3
			case 'E':
				result_pf = 1.5
			case 'F':
				result_pf = 1.6
			case 'G':
				result_pf = 1.8
			case 'H':
				result_pf = 2
			case 'J':
				result_pf = 2.2
			case 'K':
				result_pf = 2.4
			case 'L':
				result_pf = 2.7
			case 'M':
				result_pf = 3
			case 'N':
				result_pf = 3.3
			case 'P':
				result_pf = 3.6
			case 'Q':
				result_pf = 3.9
			case 'R':
				result_pf = 4.3
			case 'S':
				result_pf = 4.7
			case 'T':
				result_pf = 5.1
			case 'U':
				result_pf = 5.6
			case 'V':
				result_pf = 6.2
			case 'W':
				result_pf = 6.8
			case 'X':
				result_pf = 7.5
			case 'Y':
				result_pf = 8.2
			case 'Z':
				result_pf = 9.1
			case 'a':
				result_pf = 2.6
			case 'b':
				result_pf = 3.5
			case 'd':
				result_pf = 4
			case 'e':
				result_pf = 4.5
			case 'f':
				result_pf = 5
			case 'm':
				result_pf = 6
			case 'n':
				result_pf = 7
			case 't':
				result_pf = 8
			case 'y':
				result_pf = 9
			default:
				panic(fmt.Errorf("invalid or unsupported: EIA-198 identifier %s", string(v1)))
			}

			result_pf = result_pf * math.Pow10(int(v2-'0'))
		} else if utils.IsDigit(v1) && utils.IsDigit(v2) {
			result_pf = float64(int(v1-'0')*10 + int(v2-'0'))
		} else {
			panic(fmt.Errorf("invalid: capacitor value %s", val))
		}
	case 3, 4:
		v1 := val[0]
		v2 := val[1]
		v3 := val[2]

		if !utils.IsDigit(v1) || !utils.IsDigit(v3) {
			panic(fmt.Errorf("invalid: capacitor value %s", val))
		}

		if utils.IsLetter(v2) {
			if v2 != 'R' {
				panic(fmt.Errorf("invalid or unsupported: capacitor decimal identifier %s", string(v2)))
			}

			result_pf = float64(int(v1-'0')) + float64(int(v3-'0'))/10
		} else if !utils.IsDigit(v2) {
			panic(fmt.Errorf("invalid: capacitor identifier %s", string(v2)))
		} else {
			result_pf = float64(int(v1-'0')*10+int(v2-'0')) * math.Pow10(int(v3-'0'))
		}

		if len_val == 3 {
			break
		}

		v4 := val[3]

		switch v4 {
		case 'B':
			result_tolerance_min = &utils.Tolerance{
				Value:    .1,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
			result_tolerance_max = &utils.Tolerance{
				Value:    .1,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
		case 'C':
			result_tolerance_min = &utils.Tolerance{
				Value:    .25,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
			result_tolerance_max = &utils.Tolerance{
				Value:    .25,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
		case 'D':
			result_tolerance_min = &utils.Tolerance{
				Value:    .5,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
			result_tolerance_max = &utils.Tolerance{
				Value:    .5,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
		case 'F':
			result_tolerance_min = &utils.Tolerance{
				Value:    1,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
			result_tolerance_max = &utils.Tolerance{
				Value:    1,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
		case 'G':
			result_tolerance_min = &utils.Tolerance{
				Value:    2,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
			result_tolerance_max = &utils.Tolerance{
				Value:    2,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
		case 'J':
			result_tolerance_min = &utils.Tolerance{
				Value:    5,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
			result_tolerance_max = &utils.Tolerance{
				Value:    5,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
		case 'K':
			result_tolerance_min = &utils.Tolerance{
				Value:    10,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
			result_tolerance_max = &utils.Tolerance{
				Value:    10,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
		case 'M':
			result_tolerance_min = &utils.Tolerance{
				Value:    20,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
			result_tolerance_max = &utils.Tolerance{
				Value:    20,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
		case 'Z':
			result_tolerance_min = &utils.Tolerance{
				Value:    20,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
			result_tolerance_max = &utils.Tolerance{
				Value:    80,
				UnitType: utils.UNIT_TYPE_PERCENT,
			}
		default:
			panic(fmt.Errorf("invalid or unsupported: capacitor tolerance identifier: %s", string(v4)))
		}
	default:
		panic(fmt.Errorf("unsupported: %d digit codes", len_val))
	}

	result_pf = result_pf * utils.MATH_POW_PICO
	result_tolerance_min_pf, result_tolerance_max_pf := 0., 0.

	if result_tolerance_min != nil {
		if result_tolerance_min.UnitType == utils.UNIT_TYPE_PERCENT {
			result_tolerance_min_pf = result_pf * (1 - result_tolerance_min.Value/100)
		} else {
			result_tolerance_min_pf = result_pf - result_tolerance_min.Value
		}
	}

	if result_tolerance_max != nil {
		if result_tolerance_max.UnitType == utils.UNIT_TYPE_PERCENT {
			result_tolerance_max_pf = result_pf * (1 - result_tolerance_max.Value/100)
		} else {
			result_tolerance_max_pf = result_pf - result_tolerance_max.Value
		}
	}

	result_tolerance_min_pf *= utils.MATH_POW_PICO
	result_tolerance_max_pf *= utils.MATH_POW_PICO

	switch format {
	case "json":
		return fmt.Sprintf(`{"nominal":%s,"nominalAbbreviated":"%sF","actualMin":%s,"actualMinAbbreviated":"%s","actualMax":%s,"actualMaxAbbreviated":"%s"}`,
			utils.FormatFloat(result_pf),
			utils.GetAbbreviatedValue(result_pf),
			utils.If(result_tolerance_min != nil, utils.FormatFloat(result_tolerance_min_pf), "null"),
			utils.If(result_tolerance_min != nil, utils.GetAbbreviatedValue(result_tolerance_min_pf)+"F", "null"),
			utils.If(result_tolerance_min != nil, utils.FormatFloat(result_tolerance_max_pf), "null"),
			utils.If(result_tolerance_min != nil, utils.GetAbbreviatedValue(result_tolerance_max_pf)+"F", "null"),
		)
	case "raw":
		return fmt.Sprintf("nominal=%sF min=%s max=%s",
			utils.FormatFloat(result_pf),
			utils.If(result_tolerance_min != nil, utils.FormatFloat(result_tolerance_min_pf)+"F", "nil"),
			utils.If(result_tolerance_min != nil, utils.FormatFloat(result_tolerance_max_pf)+"F", "nil"),
		)
	default:
		return fmt.Sprintf("nominal=%sF min=%s max=%s",
			utils.GetAbbreviatedValue(result_pf),
			utils.If(result_tolerance_min != nil, utils.GetAbbreviatedValue(result_tolerance_min_pf)+"F", "nil"),
			utils.If(result_tolerance_min != nil, utils.GetAbbreviatedValue(result_tolerance_max_pf)+"F", "nil"),
		)
	}
}
