package calculate

import (
	"fmt"
	"gohm/abbrvs"
	"gohm/cli"
	"gohm/utils"
	"math"
)

func cmd_missing_resistance_handler(cmd *cli.Command) string {
	target_resistance := utils.GetValueForRKMElseShorthand(cmd.GetFlagValue("target"), abbrvs.RKM_RESISTOR, abbrvs.RESISTOR)
	if cmd.ArgsLength == 0 {
		panic("too few arguments: [args...]")
	}

	resistance := 1 / ResistanceInParallel(cmd.Args)
	missing := 1 / (1/target_resistance - resistance)

	if missing <= 0 || math.IsInf(missing, 0) || math.IsNaN(missing) {
		panic("invalid resistance")
	}

	switch cmd.GetFlagValue("format") {
	case "json":
		return fmt.Sprintf(`{"resistance":%s,"resistanceAbbreviated":"%sΩ"}`, utils.FormatFloat(missing), utils.GetAbbreviatedValue(missing))
	case "raw":
		return fmt.Sprintf("resistance=%sΩ", utils.FormatFloat(missing))
	default:
		return fmt.Sprintf("resistance=%sΩ", utils.GetAbbreviatedValue(missing))
	}
}
