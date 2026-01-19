package calculate

import (
	"fmt"
	"gohm/abbrvs"
	"gohm/cli"
	"gohm/utils"
	"strings"
)

func cmd_current_divider_handler(cmd *cli.Command) string {
	if cmd.ArgsLength < 2 {
		panic("too few arguments: [args...] - requires at least 2")
	}

	switch cmd.GetFlagValue("circuit") {
	case "capacitive":
		return cmd_current_divider_handler_capacitance(cmd)
	default:
		return cmd_current_divider_handler_resistance(cmd)
	}
}

func cmd_current_divider_handler_capacitance(cmd *cli.Command) string {
	source_current := utils.ParseShorthand(cmd.GetFlagValue("current"), abbrvs.CURRENT)
	parts, total := cmd_current_divider_handler_accumulate(cmd.Args, abbrvs.RKM_FARAD, abbrvs.FARAD, nil)

	var sb strings.Builder

	switch cmd.GetFlagValue("format") {
	case "json":
		sb.WriteRune('[')

		for i, v := range parts {
			c := (source_current * v) / total

			fmt.Fprintf(&sb, `{"current":%s,"currentAbbreviated":"%sA"}`, utils.FormatFloat(c), utils.GetAbbreviatedValue(c))

			if i != cmd.ArgsLength-1 {
				sb.WriteRune(',')
			}
		}

		sb.WriteRune(']')
		return sb.String()
	case "raw":
		for i, v := range parts {
			c := (source_current * v) / total

			fmt.Fprintf(&sb, `c%d=%sF current=%sA`, i+1, utils.FormatFloat(v), utils.FormatFloat(c))

			if i != cmd.ArgsLength-1 {
				sb.WriteRune('\n')
			}
		}

		return sb.String()
	default:
		for i, v := range parts {
			c := (source_current * v) / total

			fmt.Fprintf(&sb, `c%d=%sF current=%sA`, i+1, utils.GetAbbreviatedValue(v), utils.GetAbbreviatedValue(c))

			if i != cmd.ArgsLength-1 {
				sb.WriteRune('\n')
			}
		}

		return sb.String()
	}
}

func cmd_current_divider_handler_resistance(cmd *cli.Command) string {
	source_current := utils.ParseShorthand(cmd.GetFlagValue("current"), abbrvs.CURRENT)
	parts, total := cmd_current_divider_handler_accumulate(cmd.Args, abbrvs.RKM_RESISTOR, abbrvs.RESISTOR, cmd_current_divider_handler_accumulate_resistance)

	total = 1 / total

	var sb strings.Builder

	switch cmd.GetFlagValue("format") {
	case "json":
		sb.WriteRune('[')

		for i, v := range parts {
			c := source_current * (total / v)

			fmt.Fprintf(&sb, `{"current":%s,"currentAbbreviated":"%sA"}`, utils.FormatFloat(c), utils.GetAbbreviatedValue(c))

			if i != cmd.ArgsLength-1 {
				sb.WriteRune(',')
			}
		}

		sb.WriteRune(']')
		return sb.String()
	case "raw":
		for i, v := range parts {
			c := source_current * (total / v)

			fmt.Fprintf(&sb, `r%d=%sΩ current=%sA`, i+1, utils.FormatFloat(v), utils.FormatFloat(c))

			if i != cmd.ArgsLength-1 {
				sb.WriteRune('\n')
			}
		}

		return sb.String()
	default:
		for i, v := range parts {
			c := source_current * (total / v)

			fmt.Fprintf(&sb, `r%d=%sΩ current=%sA`, i+1, utils.GetAbbreviatedValue(v), utils.GetAbbreviatedValue(c))

			if i != cmd.ArgsLength-1 {
				sb.WriteRune('\n')
			}
		}

		return sb.String()
	}
}

// returns parts/component values, total accumulated value
func cmd_current_divider_handler_accumulate(args []string, rkm_identifer rune, shorthand_identifiers []string, accumulate func(float64) float64) ([]float64, float64) {
	total := 0.
	parts := []float64{}

	for _, arg := range args {
		val := utils.GetValueForRKMElseShorthand(arg, rkm_identifer, shorthand_identifiers)
		parts = append(parts, val)
		if accumulate != nil {
			total += accumulate(val)
		} else {
			total += val
		}
	}

	return parts, total
}

func cmd_current_divider_handler_accumulate_resistance(val float64) float64 {
	return 1 / val
}
