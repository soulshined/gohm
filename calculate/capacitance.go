package calculate

import (
	"fmt"
	"gohm/abbrvs"
	"gohm/cli"
	"gohm/utils"
)

func cmd_capacitance_handler(cmd *cli.Command) string {
	capacitance := 0.
	if cmd.GetFlagValue("circuit") == "parallel" {
		capacitance = CapacitanceInParallel(cmd.Args)
	} else {
		capacitance = CapacitanceInSeries(cmd.Args)
	}

	switch cmd.GetFlagValue("format") {
	case "json":
		return fmt.Sprintf(`{"capacitance":%s,"capacitanceAbbreviated":"%sF"}`, utils.FormatFloat(capacitance), utils.GetAbbreviatedValue(capacitance))
	case "raw":
		return fmt.Sprintf("capacitance=%sF", utils.FormatFloat(capacitance))
	default:
		return fmt.Sprintf("capacitance=%sF", utils.GetAbbreviatedValue(capacitance))
	}
}

func CapacitanceInSeries(capacitance_values []string) float64 {
	capacitance := 0.
	for _, v := range capacitance_values {
		capacitance += 1 / utils.GetValueForRKMElseShorthand(v, abbrvs.RKM_FARAD, abbrvs.FARAD)
	}
	return 1 / capacitance
}

func CapacitanceInParallel(capacitance_values []string) float64 {
	capacitance := 0.
	for _, v := range capacitance_values {
		capacitance += utils.GetValueForRKMElseShorthand(v, abbrvs.RKM_FARAD, abbrvs.FARAD)
	}
	return capacitance
}
