package calculate

import (
	"fmt"
	"gohm/abbrvs"
	"gohm/cli"
	"gohm/utils"
)

func cmd_resistance_handler(cmd *cli.Command) string {
	resistance := 0.
	if cmd.GetFlagValue("circuit") == "parallel" {
		if cmd.ArgsLength < 2 {
			panic("too few arguments: [args...] - requires at least 2")
		}
		resistance = ResistanceInParallel(cmd.Args)
	} else {
		resistance = ResistanceInSeries(cmd.Args)
	}

	switch cmd.GetFlagValue("format") {
	case "json":
		return fmt.Sprintf(`{"resistance":%s,"resistanceAbbreviated":"%sΩ"}`, utils.FormatFloat(resistance), utils.GetAbbreviatedValue(resistance))
	case "raw":
		return fmt.Sprintf("resistance=%sΩ", utils.FormatFloat(resistance))
	default:
		return fmt.Sprintf("resistance=%sΩ", utils.GetAbbreviatedValue(resistance))
	}
}

func ResistanceInSeries(resistor_values []string) float64 {
	resistance := 0.
	for _, v := range resistor_values {
		resistance += utils.GetValueForRKMElseShorthand(v, abbrvs.RKM_RESISTOR, abbrvs.RESISTOR)
	}
	return resistance
}

func ResistanceInParallel(resistor_values []string) float64 {
	resistance := 0.
	for _, v := range resistor_values {
		resistance += 1 / utils.GetValueForRKMElseShorthand(v, abbrvs.RKM_RESISTOR, abbrvs.RESISTOR)
	}
	return 1 / resistance
}
