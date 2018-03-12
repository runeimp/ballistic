//
// PACKAGES
//
package main


//
// IMPORTS
//
import (
	// humanize "github.com/dustin/go-humanize"
	. "github.com/runeimp/ballistic" // Import ballistic package into this namespace for constants, etc.
	"github.com/runeimp/locale"
	// "golang.org/x/text/message" // International formating options. Unlike "github.com/dustin/go-humanize".
	"encoding/json"
	// "errors"
	"fmt"
	// "github.com/rjeczalik/notify"
	"gopkg.in/urfave/cli.v1" // imports as package "cli"
	"log"
	"math"
	// "menteslibres.net/gosexy/to"
	// "menteslibres.net/gosexy/yaml"
	"os"
	// "os/signal"
	"sort"
	// "strconv"
	// "strings"
	// "syscall"
)


//
// CONSTANTS
//
const APP_VERSION = "0.4.2"


//
// Structs
//
type BallisticData struct {
	draw_force ParsedData
	draw_length ParsedData
	draw_weight ParsedData
	mpbr ParsedData // max_point_blank_range ParsedData
	projectile_energy ParsedData
	projectile_mass ParsedData
	projectile_velocity ParsedData
	target_radius ParsedData
}


type LabeledValue struct {
	Label string       `json:"label,omitempty"`
	ValueFloat float64 `json:"value,omitempty"`
	ValueString string `json:"value,omitempty"`
}

// func (t LabeledValue) MarshalJSON() ([]byte, error) {
// 	return []byte{}, nil
// 	// return nil, nil // <- same effect.
// }

type OutputData struct {
	Energy LabeledValue   `json:"energy,omitempty"`
	Momentum LabeledValue `json:"momentum,omitempty"`
	Mpbr LabeledValue     `json:"mpbr,omitempty"`
	Velocity LabeledValue `json:"velocity,omitempty"`
}


//
// VARIABLES
//
var data BallisticData
var decimal_places int = 6
var locale_str string
var output OutputData
var output_debug bool = false
var output_indent string = "    "
var output_json bool = false
var output_pretty bool = false


//
// FUNCTIONS
//

/** Build output data */
func buildOutputData(data BallisticData) {

	if data.projectile_velocity.Value > 0 {
		output.Velocity = velocity_to_velocity(data)
	}
	
	output.Energy = calcEnergy(data)
	output.Momentum = calcMomentum(data)

	if output_debug {
		fmt.Println("")
		fmt.Println("Internal Metric:")
		if data.projectile_velocity.Value > 0 {
			fmt.Printf("  Projectile Velocity: %16.6f %s\n", data.projectile_velocity.Value, data.projectile_velocity.Label)
		}
		if output.Energy.ValueFloat > 0 {
			fmt.Printf("    Projectile Energy: %16.6f %s\n", output.Energy.ValueFloat, output.Energy.Label)
		}
		if output.Momentum.ValueFloat > 0 {
			fmt.Printf("  Projectile Momentum: %16.6f %s\n", output.Momentum.ValueFloat, output.Momentum.Label)
		}
		if data.mpbr.Value > 0 {
			fmt.Printf("Max Point Blank Range: %16.6f %s\n", data.mpbr.Value, data.mpbr.Label)
		}
		fmt.Println("") 
	}

	if InputData.Metric == false {
		output.Energy.ValueFloat *= ENERGY_FROM_JOULES_TO_FOOTPOUNDS
		output.Energy.Label = ENERGY_LABEL_FOOTPOUNDS

		output.Momentum.ValueFloat *= MASS_FROM_KILOGRAMS_TO_POUNDS * VELOCITY_FROM_MPS_TO_FPS
		output.Momentum.Label = MOMENTUM_LABEL_FPS
	}

	if data.mpbr.Value > 0 {
		if output_debug { fmt.Printf("MPBR %f %s\n", data.mpbr.Value, data.mpbr.Label) }
		output.Mpbr = mpbr_to_mpbr(data)
		if output_debug { fmt.Printf("MPBR %f %s\n", output.Mpbr.ValueFloat, output.Mpbr.Label) }
	}
}


/** Calculate drop in flight */
func calcDropAtDistance(distance, velocity float64) (drop float64) {
	flight_time := distance / velocity
	drop = GRAVITY_MPS * 0.5 * (flight_time * flight_time)
	return drop
}

func calcDistanceForDrop(drop, velocity float64) (distance float64) {
	distance = drop * velocity
	// G * 0.5 * Squared(distance / velocity) = drop
	// 9.80665 * 0.5 = 4.903325 * Squared(distance / 500) = 1m
	// 4.903325 / 1m = 4.903325

	//  GRAVITY
	//     Gm1m2
	// F = ------
	//       r2
	return distance
}


/** Calculate energy */
func calcEnergy(data BallisticData) (energy LabeledValue) {
	// 1 Joule == 1 N⋅m (Newton meter)
	mass_kg := data.projectile_mass.Value
	velocity_mps := data.projectile_velocity.Value

	energy.ValueFloat = mass_kg * velocity_mps * velocity_mps
	energy.Label = ENERGY_LABEL_JOULES

	if output_debug {
		log.Printf("calcEnergy()   <| mass: %f kg", mass_kg)
		log.Printf("calcEnergy()   <| velocity: %f mps", velocity_mps)
		log.Printf("calcEnergy()    | energy: %f %s", energy.ValueFloat, energy.Label)
	}

	return energy
}


/** Calculate force Newtons */ 
func calc_force(avg_draw_weight float64) (draw_force ParsedData) {
	draw_force.Value = avg_draw_weight * FORCE_FROM_KILOGRAMS_TO_NEWTONS
	draw_force.Label = FORCE_LABEL_NEWTONS

	if output_debug {
		log.Printf("calc_force()    <| avg_draw_weight: %f kg", avg_draw_weight)
		log.Printf("calc_force()     | draw_force: %f %s", draw_force.Value, draw_force.Label)
	}
	
	return draw_force
}


/** Calculate momentum */
func calcMomentum(data BallisticData) (momentum LabeledValue) {
	// kg⋅m/s (kilogram meters per second)
	mass_kg := data.projectile_mass.Value
	velocity_mps := data.projectile_velocity.Value

	momentum.ValueFloat = mass_kg * velocity_mps
	momentum.Label = MOMENTUM_LABEL_MKS

	if output_debug {
		log.Printf("calcMomentum() <| mass: %f kg", mass_kg)
		log.Printf("calcMomentum() <| velocity: %f mps", velocity_mps)
		log.Printf("calcMomentum()  | momentum: %f %s", momentum.ValueFloat, momentum.Label)
	}
		
	return momentum
}


/** Calculate Maximum Point Blank Range */
func calcMPBR(data BallisticData) (mpbr ParsedData) {
	distance := 0.0
	diameter := data.target_radius.Value * 2
	drop := data.target_radius.Value
	velocity := data.projectile_velocity.Value

	for drop < diameter || drop == 0.0 {
		distance += 1.0
		drop = calcDropAtDistance(distance, velocity)
	}
	for drop > diameter {
		distance -= 0.1
		drop = calcDropAtDistance(distance, velocity)
	}
	for drop < diameter {
		distance += 0.01
		drop = calcDropAtDistance(distance, velocity)
	}
	for drop > diameter {
		distance -= 0.001
		drop = calcDropAtDistance(distance, velocity)
	}
	for drop < diameter {
		distance += 0.0001
		drop = calcDropAtDistance(distance, velocity)
	}
	for drop > diameter {
		distance -= 0.00001
		drop = calcDropAtDistance(distance, velocity)
	}
	for drop < diameter {
		distance += 0.000001
		drop = calcDropAtDistance(distance, velocity)
	}
	for drop > diameter {
		distance -= 0.0000001
		drop = calcDropAtDistance(distance, velocity)
	}

	mpbr.Value = distance
	mpbr.Label = "meters"

	// log.Printf("calcMPBR() <|   target radius: %12.6f m", data.target_radius.Value)
	// log.Printf("calcMPBR() <| target diameter: %12.6f m", diameter)
	// log.Printf("calcMPBR() <| projectile drop: %12.6f m", drop)
	// log.Printf("calcMPBR()  |            MPBR: %12.6f m", distance)

	return mpbr
}

/**
 * Calculate velocity
 */
func calcVelocity(data BallisticData) (projectile_velocity ParsedData) {
	projectile_mass := data.projectile_mass.Value
	draw_length := data.draw_length.Value
	draw_force := data.draw_force.Value

	release_time := math.Sqrt(projectile_mass * draw_length / draw_force)

	projectile_velocity.Value = draw_length / release_time
	projectile_velocity.Label = VELOCITY_LABEL_MPS

	if len(InputData.Velocity) == 0 {
		if InputData.Metric {
			InputData.Velocity = VELOCITY_LABEL_MPS
		} else {
			InputData.Velocity = VELOCITY_LABEL_FPS
		}
	}

	if output_debug {
		log.Printf("calcVelocity() <|     projectile mass: %15.6f kg", projectile_mass)
		log.Printf("calcVelocity() <|         draw length: %15.6f m", draw_length)
		log.Printf("calcVelocity() <|          draw force: %15.6f N", draw_force)
		log.Printf("calcVelocity()  | projectile velocity: %15.6f mps", projectile_velocity.Value)
	}

	return projectile_velocity
}


/**
 * Cleans up OutputData for JSON parsing
 *
 * Removes top level keys from OutputData when the child fields have no value.
 * This is necessary because the standard JSON package does not check child
 * structs if they are empty or not.
 */
func cleanupJSON(data OutputData) (data_obj map[string]interface{}) {
	data_obj = make(map[string]interface{})

	if data.Energy.ValueFloat != 0 {
		data_obj["energy"] = data.Energy
	}
	if data.Momentum.ValueFloat != 0 {
		data_obj["momentum"] = data.Momentum
	}
	if data.Mpbr.ValueFloat != 0 {
		data_obj["mpbr"] = data.Mpbr
	}
	if data.Velocity.ValueFloat != 0 {
		data_obj["velocity"] = data.Velocity
	}

	return data_obj
}


/** Function defined by a call to github.com/runeimp/locale.NumberFormatter() */
var locale_NumberFormatter func(number float64, scale int) string


/** Returns the largest integer in the list of arguments */
func maxInt(nums ...int) (max_int int) {
	max_int = math.MinInt64
	for _, num := range nums {
		if max_int < num {
			max_int = num
		}
	}
	return max_int
}


/** Convert MPBR in meters to input units */
func mpbr_to_mpbr(data BallisticData) (mpbr LabeledValue) {
	mpbr.Label = ""
	mpbr.ValueFloat = 0.0

	user_label := data.projectile_velocity.UserLabel
	if len(user_label) == 0 {
		user_label = InputData.Velocity
	}

	switch user_label {
	case VELOCITY_LABEL_FPS:
		mpbr.Label = LENGTH_LABEL_FOOT
		mpbr.ValueFloat = data.mpbr.Value * LENGTH_FROM_METERS_TO_FEET
	case VELOCITY_LABEL_KMPH:
		mpbr.Label = LENGTH_LABEL_KILOMETER
		mpbr.ValueFloat = data.mpbr.Value * LENGTH_FROM_METERS_TO_KILOMETERS
	case VELOCITY_LABEL_KNOTS:
		mpbr.Label = LENGTH_LABEL_NAUTICAL_MILE
		mpbr.ValueFloat = data.mpbr.Value * LENGTH_FROM_METERS_TO_NAUTICAL_MILES
	case VELOCITY_LABEL_MPS:
		mpbr.Label = LENGTH_LABEL_METER
		mpbr.ValueFloat = data.mpbr.Value
	case VELOCITY_LABEL_MPH:
		mpbr.Label = LENGTH_LABEL_MILE
		mpbr.ValueFloat = data.mpbr.Value * LENGTH_FROM_METERS_TO_MILES
	}

	if output_debug {
		log.Printf("mpbr_to_mpbr()  <|  projectile velocity: %s", data.projectile_velocity.UserLabel)
		log.Printf("mpbr_to_mpbr()  <| InputData velocity: %s", InputData.Velocity)
		log.Printf("mpbr_to_mpbr()   |  projectile velocity: %15.6f mps", mpbr.Label, mpbr.Label)
	}

	return mpbr
}


/** Takes a float64 and returns it's formated number value and it's string width */
func numberFormatter(number float64) (value string, width int) {
	value = locale_NumberFormatter(number, decimal_places)
	width = len(value)

	return value, width
}


/** Print Human Readable Output */
func outputHuman(data OutputData) {
	fmt.Println("")

	var energy_value string
	var energy_width int
	var momentum_value string
	var momentum_width int
	var mpbr_value string
	var mpbr_width int
	var velocity_value string
	var velocity_width int

	if data.Velocity.ValueFloat > 0 {
		velocity_value, velocity_width = numberFormatter(data.Velocity.ValueFloat)
	}
	if data.Energy.ValueFloat > 0 {
		energy_value, energy_width = numberFormatter(data.Energy.ValueFloat)
	}
	if data.Momentum.ValueFloat > 0 {
		momentum_value, momentum_width = numberFormatter(data.Momentum.ValueFloat)
	}
	if data.Mpbr.ValueFloat > 0 {
		mpbr_value, mpbr_width = numberFormatter(data.Mpbr.ValueFloat)
	}

	max_width := fmt.Sprintf("%d", maxInt(
		velocity_width,
		energy_width,
		momentum_width,
		mpbr_width,
	))


	if velocity_width > 0 {
		msg_format := "  Projectile Velocity: %" + max_width + "s %s\n"
		fmt.Printf(msg_format, velocity_value, data.Velocity.Label)
	}
	if energy_width > 0 {
		msg_format := "    Projectile Energy: %" + max_width + "s %s\n"
		fmt.Printf(msg_format, energy_value, data.Energy.Label)
	}
	if momentum_width > 0 {
		msg_format := "  Projectile Momentum: %" + max_width + "s %s\n"
		fmt.Printf(msg_format, momentum_value, data.Momentum.Label)
	}
	if mpbr_width > 0 {
		msg_format := "Max Point Blank Range: %" + max_width + "s %s\n"
		fmt.Printf(msg_format, mpbr_value, data.Mpbr.Label)
	}
	
	fmt.Println("")
}

// func numberFormat(number float64) (result string) {
// 	// numberFormatBase(number)
// 	str_float := fmt.Sprintf("%.6f", number)
// 	str_scale := strings.Split(str_float, ".")[1]
// 	result = humanize.Comma(int64(number)) + "." + str_scale
// 	return result
// }

/** Outputs JSON data */
func outputJSON(data OutputData) {
	if output_debug {
		fmt.Println("JSON data!")
		fmt.Printf("data.Energy: %f %s\n", data.Energy.ValueFloat, data.Energy.Label)
	}
	var err error
	var json_data []byte

	data_obj := cleanupJSON(data)
	// data_obj := data

	if output_pretty {
		json_data, err = json.MarshalIndent(data_obj, "", output_indent)
	} else {
		json_data, err = json.Marshal(data_obj)
	}
	// fmt.Printf("json_data: %#v\n", json_data)
	

	if output_debug {
		log.Printf("err: %v\n", err)
	}
	
	if err == nil {
		fmt.Println(string(json_data))
	} else {
		log.Println("JSON encoding error")
		log.Printf("%s\n", err)
		log.Printf("err: %#v\n", err)
	}
}




/** Convert velocity in mps to input units */
func velocity_to_velocity(data BallisticData) (velocity LabeledValue) {
	velocity.Label = ""

	user_label := InputData.Velocity
	if len(user_label) == 0 {
		switch InputData.Mass {
		case MASS_LABEL_GRAINS, MASS_LABEL_LONG_TON, MASS_LABEL_POUNDS, MASS_LABEL_SHORT_TON, MASS_LABEL_STONE:
			user_label = VELOCITY_LABEL_FPS
		default:
			user_label = VELOCITY_LABEL_MPS
		}
	}

	switch user_label {
	case VELOCITY_LABEL_FPS:
		velocity.Label = user_label
		velocity.ValueFloat = data.projectile_velocity.Value * VELOCITY_FROM_MPS_TO_FPS
	case VELOCITY_LABEL_KMPH:
		velocity.Label = user_label
		velocity.ValueFloat = data.projectile_velocity.Value * VELOCITY_FROM_MPS_TO_KMPH
	case VELOCITY_LABEL_KNOTS:
		velocity.Label = user_label
		velocity.ValueFloat = data.projectile_velocity.Value * VELOCITY_FROM_MPS_TO_KNOTS
	case VELOCITY_LABEL_MPS:
		velocity.Label = user_label
		velocity.ValueFloat = data.projectile_velocity.Value
	case VELOCITY_LABEL_MPH:
		velocity.Label = user_label
		velocity.ValueFloat = data.projectile_velocity.Value * VELOCITY_FROM_MPS_TO_MPH
	}

	// log.Printf("velocity_to_velocity() | user_label: '%s'", user_label)
	// log.Printf("velocity_to_velocity() | projectile_mass user value & label: %f %s", data.projectile_mass.UserValue, data.projectile_mass.UserLabel)
	// log.Printf("velocity_to_velocity() | InputData | velocity: '%s' | mass: '%s' | metric: %t", InputData.Velocity, InputData.Mass, InputData.Metric)
	// log.Printf("velocity_to_velocity() | projectile_velocity user value & label: %f %s", data.projectile_velocity.UserValue, data.projectile_velocity.UserLabel)
	// log.Printf("velocity_to_velocity() | projectile_velocity norm value & label: %f %s", data.projectile_velocity.Value, data.projectile_velocity.Label)
	// log.Printf("velocity_to_velocity() | projectile_velocity calc value & label: %f %s", velocity.ValueFloat, velocity.Label)

	return velocity
}


//
// MAIN ENTRYPOINT
//
func main() {
	// log.Print("Ballistic ...")

	app := cli.NewApp()
	app.Name = "Ballistic"
	app.Usage = "Calculates what it can based on provided input."
	app.Version = APP_VERSION

	app.Flags = []cli.Flag {
		cli.BoolFlag{
			Name: "debug, d",
			Usage: "Output debug info",
		},
		cli.StringFlag{
			Name: "draw-weight, weight, w",
			Usage: "Bow or sling shot draw `WEIGHT`. Used to calculate projectile velocity, energy, etc.",
		},
		cli.StringFlag{
			Name: "draw-length, length, l",
			Usage: "Bow or sling shot draw `LENGTH`. Used to calculate projectile velocity, energy, etc.",
		},
		// cli.BoolFlag{
		// 	Name: "help, h",
		// 	Usage: "Output this help info",
		// },
		cli.BoolFlag{
			Name: "json, j",
			Usage: "Output JSON data",
		},
		cli.StringFlag{
			Name: "locale, local",
			Value: "en_US",
			Usage: "The `LOCALE` to format number output for.",
			EnvVar: "LC_CTYPE,LANG",
		},
		cli.StringFlag{
			Name: "projectile, mass, m",
			Usage: "Projectile `MASS` (weight). Used to calculate projectile velocity, energy, etc.",
		},
		cli.StringFlag{
			Name: "precision, float, f",
			Value: "6",
			Usage: "The output floating point `PRECISION` (numbers after decimal mark).",
		},
		cli.BoolFlag{
			Name: "pretty-print, pretty, p",
			Usage: "Pretty printed JSON output",
		},
		// cli.BoolFlag{
		// 	Name: "pretty, p",
		// 	Usage: "Pretty printed JSON output",
		// },
		// cli.StringFlag{
		// 	Name: "pretty-print",
		// 	Usage: "Pretty printed JSON output specifying `INDENT` string",
		// },
		cli.StringFlag{
			Name: "radius, r",
			Value: "225mm",
			Usage: "The `RADIUS` of the target area. Used to calculate MPBR (Maximum Point Blank Range).",
		},
		cli.StringFlag{
			Name: "velocity, v",
			Usage: "The projectile `VELOCITY` (speed). Used to calculate projectile energy, momentum, etc.",
		},
	}

	cli.HelpFlag = cli.BoolFlag{
		Name: "help, h",
		Usage: "Output this help info",
	}

	cli.AppHelpTemplate = fmt.Sprintf(HELP_TEMPLATE)

	cli.VersionFlag = cli.BoolFlag{
		Name: "version, V",
		Usage: "Output the ballistic app version",
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	// sort.Sort(cli.CommandsByName(app.Commands))

	app.Action = func(c *cli.Context) error {
		output_debug = c.Bool("debug")
		output_json = c.Bool("json")
		output_pretty = c.Bool("pretty-print")
		decimal_places = c.Int("precision")
		locale_str = c.String("locale")

		// output_pretty = c.Bool("pretty")
		// if len(c.String("pretty-print")) > 0 {
		// 	log.Println("pretty-print!!!")
		// 	output_indent = c.String("pretty-print") 
		// }

		if output_pretty && ! output_json {
			output_json = true
		}

		if output_debug {
			fmt.Println("Going Ballistic!")
			log.Printf("             locale: %12s (%d)", locale_str, len(locale_str))
			log.Printf("        draw length: %12s (%d)", c.String("draw-length"), len(c.String("draw-length")))
			log.Printf("        draw weight: %12s (%d)", c.String("draw-weight"), len(c.String("draw-weight")))
			log.Printf("      target radius: %12s (%d)", c.String("radius"), len(c.String("radius")))
			log.Printf("projectile velocity: %12s (%d)", c.String("velocity"), len(c.String("velocity")))
			log.Printf("    projectile mass: %12s (%d)", c.String("projectile"), len(c.String("projectile")))
			log.Printf("          precision: %12s (%d)", c.String("precision"), len(c.String("precision")))
			// log.Println("")
			// fmt.Println("")
		}


		flags_set := 0
		for _, flag_name := range c.GlobalFlagNames() {
			// fmt.Printf("Flag: %s\n", flag_name)
			switch flag_name {
			case "locale", "precision", "radius":
			default:
				flag_value := c.String(flag_name)
				if len(flag_value) > 0 {
					if flag_value != "false" && flag_value != "true" {
						// fmt.Printf("Flag Set: %s\n", flag_value)
						flags_set += 1
					}
				}
			}
		}
		// fmt.Printf("Flags Set: %d\n", flags_set)

		/** If no flags set default to displaying help */
		if flags_set == 0 {
			cli.ShowAppHelpAndExit(c, 0)
		}


		if len(c.String("draw-length")) > 0 {
			data.draw_length = ParseValue(c.String("draw-length"), VALUE_TYPE_LENGTH)
		}
		if len(c.String("draw-weight")) > 0 {
			data.draw_weight = ParseValue(c.String("draw-weight"), VALUE_TYPE_MASS)
			avg_draw_weight := data.draw_weight.Value * 0.5
			data.draw_force = calc_force(avg_draw_weight)
		}
		if len(c.String("velocity")) > 0 {
			data.projectile_velocity = ParseValue(c.String("velocity"), VALUE_TYPE_VELOCITY)
		}
		if len(c.String("mass")) > 0 {
			data.projectile_mass = ParseValue(c.String("mass"), VALUE_TYPE_MASS)
		}

		if data.projectile_velocity.Value == 0 {
			if data.projectile_mass.Value > 0 && data.draw_length.Value > 0 && data.draw_force.Value > 0 {
				data.projectile_velocity = calcVelocity(data)
			}
		}

		data.target_radius = ParseValue(c.String("radius"), VALUE_TYPE_LENGTH)

		if data.projectile_velocity.Value > 0 {
			data.mpbr = calcMPBR(data)
		}

		buildOutputData(data)

		locale_NumberFormatter = locale.NumberFormatter(locale_str)
		// locale_NumberFormatter = locale.NumberFormatter("TESTONE")
		// locale_NumberFormatter(123456789.1234567)

		if output_json {
			outputJSON(output)
		} else {
			outputHuman(output)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	BALLISTIC_WEIGHT, bw_present := os.LookupEnv("BALLISTIC_WEIGHT")

	if bw_present {
		fmt.Printf("BALLISTIC_WEIGHT: %s", BALLISTIC_WEIGHT)
	}

	// x := fmt.Sprintf("Hello, %s!", )
	// x := fmt.Sprintf("Hello, %s!", os.Args[1])
}


