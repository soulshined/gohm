package cli

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	stateInit = iota
	stateFlags
	stateArgs
	stateFlagsAfterArgs
)

type Flag struct {
	Name           string
	Aliases        []string
	Description    string
	Default        string
	PossibleValues []string
	IsMulti        bool
	Required       bool
	Value          string
	Values         []string // all values when IsMulti is true
	IsSet          bool
}

type Example struct {
	Command     string
	Description string
	Output      string
}

type Command struct {
	Name        string
	Aliases     []string
	Description string
	Flags       []*Flag
	Subcommands []*Command
	Examples    []Example
	Args        []string
	ArgsLength  int
	Handler     func(*Command) string
	parent      *Command
}

// CLI is the root command handler
type CLI struct {
	Name        string
	Version     string
	Description string
	Root        *Command
}

func NewCLI(name, version, description string) *CLI {
	return &CLI{
		Name:        name,
		Version:     version,
		Description: description,
		Root: &Command{
			Name:        name,
			Description: description,
		},
	}
}

func (c *CLI) AddCommand(cmd *Command) {
	cmd.parent = c.Root
	c.Root.Subcommands = append(c.Root.Subcommands, cmd)
}

func (cmd *Command) AddSubcommand(sub *Command) {
	sub.parent = cmd
	cmd.Subcommands = append(cmd.Subcommands, sub)
}

func (cmd *Command) AddFlag(f *Flag) {
	f.Value = f.Default
	cmd.Flags = append(cmd.Flags, f)
}

func (cmd *Command) GetFlag(name string) *Flag {
	for _, f := range cmd.Flags {
		if f.Name == name {
			return f
		}
		if slices.Contains(f.Aliases, name) {
			return f
		}
	}
	return nil
}

func (cmd *Command) GetFlagValue(name string) string {
	if f := cmd.GetFlag(name); f != nil {
		return f.Value
	}
	return ""
}

func (cmd *Command) GetFlagFloat(name string) (float64, error) {
	val := cmd.GetFlagValue(name)
	if val == "" {
		return 0, fmt.Errorf("flag %s not set", name)
	}
	return strconv.ParseFloat(val, 64)
}

func (cmd *Command) GetFlagValues(name string) []string {
	if f := cmd.GetFlag(name); f != nil {
		return f.Values
	}
	return nil
}

func (cmd *Command) GetFlagFloats(name string) ([]float64, error) {
	values := cmd.GetFlagValues(name)
	if len(values) == 0 {
		return nil, fmt.Errorf("flag %s not set", name)
	}
	result := make([]float64, len(values))
	for i, v := range values {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float value for flag %s: %s", name, v)
		}
		result[i] = f
	}
	return result, nil
}

func (cmd *Command) IsFlagSet(name string) bool {
	if f := cmd.GetFlag(name); f != nil {
		return f.IsSet
	}
	return false
}

func (c *CLI) Run(args []string) string {
	if len(args) < 2 {
		c.Root.print_help()
		os.Exit(1)
	}

	args = args[1:]

	if args[0] == "-help" || args[0] == "--help" || args[0] == "-h" {
		c.Root.print_help()
		os.Exit(0)
	}

	if len(args) == 1 && (args[0] == "-v" || args[0] == "-version" || args[0] == "--version") {
		fmt.Printf("%s %s\n", c.Name, c.Version)
		os.Exit(0)
	}

	return c.execute_command(c.Root, args)
}

// parse_flags parses flags from args and returns remaining positional args
// Flags must come either all before or all after positional args - no mixing allowed
func (cmd *Command) parse_flags(args []string) []string {
	var positional []string
	i := 0
	state := stateInit

	for i < len(args) {
		arg := args[i]

		if arg == "-help" || arg == "--help" || arg == "-h" {
			cmd.print_help()
			os.Exit(0)
		}

		if strings.HasPrefix(arg, "-") {
			flagName := strings.TrimLeft(arg, "-")

			if state == stateFlagsAfterArgs {
				panic("invalid: flags and arguments cannot be mixed - place all flags before or after arguments")
			}

			switch state {
			case stateArgs:
				state = stateFlagsAfterArgs
			case stateInit:
				state = stateFlags
			}

			// -flag=value format
			if flag, value, ok := strings.Cut(flagName, "="); ok {
				if f := cmd.GetFlag(flag); f != nil {
					if f.IsSet && !f.IsMulti {
						panic(fmt.Errorf("invalid: flag -%s specified multiple times but does not support multiple values", flag))
					}
					f.Value = value
					f.Values = append(f.Values, value)
					f.IsSet = true
				} else {
					panic(fmt.Errorf("invalid: unknown flag -%s", flag))
				}
				i++
				continue
			}

			// -flag value format
			if f := cmd.GetFlag(flagName); f != nil {
				if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
					if f.IsSet && !f.IsMulti {
						panic(fmt.Errorf("invalid: flag -%s specified multiple times but does not support multiple values", flagName))
					}
					f.Value = args[i+1]
					f.Values = append(f.Values, args[i+1])
					f.IsSet = true
					i += 2
					continue
				}

				// flag without value (boolean-style-like)
				if f.IsSet && !f.IsMulti {
					panic(fmt.Errorf("invalid: flag -%s specified multiple times but does not support multiple values", flagName))
				}
				f.IsSet = true
				i++
				continue
			}

			panic(fmt.Errorf("invalid: flag -%s", flagName))
		}

		// positional argument
		switch state {
		case stateFlags:
			state = stateArgs
		case stateInit:
			state = stateArgs
		case stateFlagsAfterArgs:
			fmt.Fprintf(os.Stderr, "Error: flags and arguments cannot be mixed. Place all flags before or after arguments.\n")
			os.Exit(1)
		}

		positional = append(positional, arg)
		i++
	}

	return positional
}

func (cmd *Command) print_help() {
	path := cmd.Name
	if cmd.parent != nil && cmd.parent.Name != "" {
		path = cmd.get_full_path()
	}

	fmt.Printf("Usage: %s", path)

	if len(cmd.Subcommands) > 0 {
		fmt.Print(" <command>")
	}
	if len(cmd.Flags) > 0 {
		fmt.Print(" [flags]")
	}
	if cmd.Handler != nil && len(cmd.Subcommands) == 0 {
		fmt.Print(" [args...]")
	}
	fmt.Println()

	if cmd.Description != "" {
		fmt.Printf("\n%s\n", cmd.Description)
	}

	if len(cmd.Aliases) > 0 {
		fmt.Printf("\nAliases: %s\n", strings.Join(cmd.Aliases, ", "))
	}

	if len(cmd.Subcommands) > 0 {
		fmt.Println("\nCommands:")
		maxLen := 0
		for _, sub := range cmd.Subcommands {
			nameWithAliases := sub.Name
			if len(sub.Aliases) > 0 {
				nameWithAliases += " (" + strings.Join(sub.Aliases, ", ") + ")"
			}
			if len(nameWithAliases) > maxLen {
				maxLen = len(nameWithAliases)
			}
		}
		for _, sub := range cmd.Subcommands {
			nameWithAliases := sub.Name
			if len(sub.Aliases) > 0 {
				nameWithAliases += " (" + strings.Join(sub.Aliases, ", ") + ")"
			}
			fmt.Printf("  %-*s  %s\n", maxLen, nameWithAliases, sub.Description)
		}
	}

	if len(cmd.Flags) > 0 {
		fmt.Println("\nFlags:")
		maxLen := 0
		for _, f := range cmd.Flags {
			flagStr := build_flag_string(f)
			if len(flagStr) > maxLen {
				maxLen = len(flagStr)
			}
		}
		for _, f := range cmd.Flags {
			flagStr := build_flag_string(f)
			extras := []string{}
			if f.Required {
				extras = append(extras, "required")
			}
			if len(f.PossibleValues) > 0 {
				extras = append(extras, "values: "+strings.Join(f.PossibleValues, "|"))
			}
			if f.Default != "" {
				extras = append(extras, "default: "+f.Default)
			}
			extraStr := ""
			if len(extras) > 0 {
				extraStr = " (" + strings.Join(extras, ", ") + ")"
			}
			fmt.Printf("  %-*s  %s%s\n", maxLen, flagStr, f.Description, extraStr)
		}
	}

	if len(cmd.Examples) > 0 {
		fmt.Println("\nExamples:")
		for i, ex := range cmd.Examples {
			if i != 0 {
				fmt.Println()
			}

			if ex.Description != "" {
				fmt.Printf("  // %s\n", ex.Description)
			}
			fmt.Printf("  > %s\n", ex.Command)
			if ex.Output != "" {
				fmt.Printf("    → %s\n", ex.Output)
			}
		}
	}

	fmt.Println("\nUse -help with any command for more information.")
}

func build_flag_string(f *Flag) string {
	parts := []string{"-" + f.Name}
	for _, alias := range f.Aliases {
		parts = append(parts, "-"+alias)
	}
	return strings.Join(parts, ", ")
}

func (cmd *Command) get_full_path() string {
	if cmd.parent == nil || cmd.parent.Name == "" {
		return cmd.Name
	}
	return cmd.parent.get_full_path() + " " + cmd.Name
}

func (cmd *Command) validate_required_flags() {
	var missing []string
	for _, f := range cmd.Flags {
		if f.Required && !f.IsSet {
			missing = append(missing, f.Name)
		}
	}
	if len(missing) > 0 {
		fmt.Fprintf(os.Stderr, "Error: missing required flag(s): -%s\n", strings.Join(missing, ", -"))
		os.Exit(1)
	}
}

func (cmd *Command) find_subcommand(name string) *Command {
	name = strings.ToLower(name)
	for _, sub := range cmd.Subcommands {
		if strings.ToLower(sub.Name) == name {
			return sub
		}
		for _, alias := range sub.Aliases {
			if strings.ToLower(alias) == name {
				return sub
			}
		}
	}
	return nil
}

func (c *CLI) find_target_command(cmd *Command, args []string) (*Command, []string) {
	if len(args) == 0 {
		return cmd, args
	}

	if !strings.HasPrefix(args[0], "-") {
		if sub := cmd.find_subcommand(args[0]); sub != nil {
			return c.find_target_command(sub, args[1:])
		}
	}

	return cmd, args
}

func (c *CLI) execute_command(cmd *Command, args []string) string {
	if len(args) == 0 {
		if cmd.Handler != nil {
			return cmd.Handler(cmd)
		}
		cmd.print_help()
		os.Exit(1)
	}

	targetCmd, remainingArgs := c.find_target_command(cmd, args)

	remaining := targetCmd.parse_flags(remainingArgs)
	targetCmd.validate_required_flags()
	targetCmd.Args = remaining
	targetCmd.ArgsLength = len(targetCmd.Args)

	if targetCmd.Handler != nil {
		return targetCmd.Handler(targetCmd)
	}

	targetCmd.print_help()
	os.Exit(1)
	return ""
}
