# gohm

&nbsp;&nbsp;georg<br>
golang<br>
&nbsp;&nbsp;h<br>
&nbsp;&nbsp;m<br>

A zero dependency, very small footprint, electrical engineering CLI utility for calculations and component identification


## Installation

> [!CAUTION]
> This was not intended or created in mind to be a programmatic tool. This is additionally not published to pkg.go.dev - I just wanted it to be open source for feedback, discussion and contributions
>

1. Download the binary of your choosing from the latest pre-compiled [releases](https://github.com/soulshined/gohm/releases)
2. Run `./<filename> -v`

Optionally add it to your PATH:

1. Rename the downloaded executable to `gohm`
2. Add the renamed executable to a bin folder of your preference or add it to your PATH
3. Restart terminal
4. Test with `gohm -v`


## Input Formats

gohm supports multiple input formats for convenience (support is specified on individual flags):

### Shorthand Notation
Values can be specified with SI prefixes:
- `1k` = 1000
- `4.7k` = 4700
- `100μ` = 0.0001
- `10n` = 0.00000001
- `47p` = 0.000000000047

### RKM Code Notation
Component values using letter as decimal point:
- `4K7` = 4700 (4.7k)
- `10K` = 10000
- `4R7` = 4.7
- `47R` = 47
- `4n7` = 0.0000000047

### Unit Suffixes
Optional unit identifier suffixes:
- `12V`, `5mA`, `1kHz`, `100μF`, `10kΩ`

## Commands

### calculate

Perform electrical calculations

#### Subcommands

##### calculate 555

Calculate 555 timer

**Flags:**
| Flag | Alias | Enum | Description |
|---|---|---|---|
| `-capacitance` <sup style="color:red">required<sup> | `-c` | | Capacitance value (F) - supports RKM & shorthand |
| `-resistance` <sup style="color:red">required<sup> | `-r` | | Resistance value (R) - when specified 2 times - circuit is assumed astable - supports RKM & shorthand |
| `-format` | | `abbr` (default), `raw`, `json` | Output format |

**Examples:**
```
> gohm calculate 555 -capacitance 10μF -resistance 100k
  → time=1.1s
```
_astable circuit example - providing 2 resistors_
```
> gohm calculate 555 -capacitance 1μF -resistance 1k -resistance 1k
  → time_low=693μs time_high=1.386ms frequency=480Hz
```

##### calculate capacitance

Calculate total capacitance of capacitors - values are n args passed in - args supports RKM & shorthand

**Flags:**
| Flag | Alias | Enum | Description |
|---|---|---|---|
| `-circuit` | | `series` (default), `parallel` | Type of circuit |
| `-format` | | `abbr` (default), `raw`, `json` | Output format  |

**Examples:**
```
> gohm calculate capacitance 10μF 22μF
  → capacitance=6.875μF
```
```
> gohm calculate capacitance 10μF 22μF 18μF -circuit parallel
  → 50μF
```

##### calculate current-divider

Calculate current for parallel components - values are n args passed in

**Flags:**
| Flag | Alias | Enum | Description |
|---|---|---|---|
| `-current` <sup style="color:red">required<sup> | `-c`, `i` | | Input current - RKM & shorthand supported |
| `-circuit` | | `resistive` (default), `capacitive` | Circuit divider type - resistive & capacitive args support RKM & shorthand |
| `-format` | | `abbr` (default), `raw`, `json` | Output format |

**Examples:**
```
> gohm calculate current-divider -current 6A 10 20 22
  → r1=10Ω current=3.0697674418604644A
    r2=20Ω current=1.5348837209302322A
    r3=22Ω current=1.3953488372093021A
```
```
> gohm calculate current-divider -current 6A 10pF 22pF 18pF
  → c1=10pF current=1.2000000000000002A
    c2=22pF current=2.64A
    c3=18pF current=2.16A
```

##### calculate missing-resistance

Calculate a resistance value needed to complete a parallel circuit - resistors are n args passed in - RKM & shorthand supported

**Flags:**
| Flag | Alias | Enum | Description |
|---|---|---|---|
| `-target` <sup style="color:red">required<sup> | `-t` | | Desired total/target resistance - RKM & shorthand supported |
| `-format` | | `abbr` (default), `raw`, `json` | Output format |

**Examples:**
```
> gohm calculate missing-resistance -target 150 250
  → resistance=374.99999999999994Ω
```
```
> gohm calculate missing-resistance -target 1k 2.5k 2k
  → resistance=9.999999999999996kΩ
```

##### calculate ohmslaw

Calculate Ohm's Law values (V=IR, P=IV, etc) based on 2 input values

**Flags:**
| Flag | Alias | Enum | Description |
|---|---|---|---|
| `-current` | `-c`, `-i` | | Current value (I) - shorthand supported |
| `-power` | `-p` | | Power value (W) - shorthand supported |
| `-resistance` | `-r` | | Resistance value (R) - can be specified multiple times for series - RKM & shorthand supported |
| `-voltage` | `-v` | | Voltage value (V) - shorthand supported |
| `-format` | | `abbr` (default), `raw`, `json` | Output format |

**Examples:**
```
> gohm calculate ohmslaw -voltage 5V -current 5A
  → voltage=5V current=5A resistance=1Ω power=25W
```
```
> gohm calculate ohmslaw -voltage 5V -resistance 20
  → voltage=5V current=250mA resistance=20Ω power=1.25W
```
_series total resistance 40Ω_
```
> gohm calculate ohmslaw -voltage 5V -resistance 20 -resistance 10 -resistance 10
  → voltage=5V current=125mA resistance=40Ω power=625mW
```

##### calculate resistance

Calculate total resistance of resistors - values are n args passed in - args supports RKM & shorthand

**Flags:**
| Flag | Alias | Enum | Description |
|---|---|---|---|
| `-circuit` | | `series` (default), `parallel` | Type of circuit |
| `-format` | | `abbr` (default), `raw`, `json` | Output format |

**Examples:**
```
> gohm calculate resistance 5k 2.5k 5k
  → resistance=12.5kΩ
```
```
> gohm calculate resistance 5k 2.5k 5k -circuit parallel
  → resistance=1.25kΩ
```

##### calculate voltage-divider

Calculate output voltage for series components

**Flags:**
| Flag | Alias | Enum | Description |
|---|---|---|---|
| `-voltage` <sup style="color:red">required<sup> | `-v` | | Input voltage - shorthand supported |
| `-capacitance` | `-c` | | RKM & shorthand supported |
| `-resistance` | `-r` | | RKM & shorthand supported |
| `-frequency` | `-f` | | used only with a resistor <-> capacitor divider type - shorthand supported |
| `-format` | | `abbr` (default), `raw`, `json` | Output format |

**Examples:**
```
> gohm calculate voltage-divider -voltage 9v -resistance 3k -resistance 3k
  → voltage=4.5V
```

---

### identify

Identify electrical components by visual indicators

#### Subcommands

##### identify capacitor

Identify capacitor value

**Flags:**
| Flag | Alias | Enum | Description |
|---|---|---|---|
| `-eiac` | `-code` | | The eia code or identifier - supports 2-4 digit codes including SMD & EIA-198 |
| `-format` | | `abbr` (default), `raw`, `json` | Output format |

**Examples:**

_2 digit code_
```
> gohm identify capacitor -eiac 21
  → nominal=21pF min=nil max=nil
```
_2 digit eia-198 code_
```
> gohm identify capacitor -eiac A8
  → nominal=100μF min=nil max=nil
```
_3 digit code - SMD or ceramic_
```
> gohm identify capacitor -eiac 100
  → nominal=10pF min=nil max=nil
```
_3 digit decimal_
```
> gohm identify capacitor -eiac 9R4
  → nominal=9.4pF min=nil max=nil
```
_4 digit code_
```
> gohm identify capacitor -eiac 123B
  → nominal=12nF min=11.988zF max=11.988zF
```
_4 digit decimal with tolerance_
```
> gohm identify capacitor -eiac 6R7K
  → nominal=6.7pF min=6.029999999999999yF max=6.029999999999999yF
```

##### identify resistor

Identify resistor value from color bands - color bands are n args passed in

**Flags:**
| Flag | Alias | Enum | Description |
|---|---|---|---|
| `-format` | | `abbr` (default), `raw`, `json` | Output format |

**Examples:**

```
> gohm identify resistor red red brown
  → nominal=220Ω min=176Ω max=264Ω temp_coefficient=nil
```

_eia shorthand_
```
> gohm identify resistor rd rd bn rd
  → nominal=220Ω min=215.6Ω max=224.4Ω temp_coefficient=nil
```