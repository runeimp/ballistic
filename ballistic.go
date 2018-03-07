package main


//
// IMPORTS
//
import (
	"encoding/json"
	// "errors"
	"fmt"
	// "github.com/chzyer/readline"
	// "github.com/rjeczalik/notify"
	"gopkg.in/urfave/cli.v1" // imports as package "cli"
	"log"
	"math"
	// "menteslibres.net/gosexy/to"
	// "menteslibres.net/gosexy/yaml"
	"os"
	// "os/signal"
	"regexp"
	"sort"
	"strconv"
	"strings"
	// "syscall"
)


//
// CONSTANTS
//
const APP_VERSION = "0.2.0"
const ENERGY_FROM_JOULES_TO_FOOTPOUNDS = 0.737562
const ENERGY_LABEL_FOOTPOUNDS = "foot-pounds"
const ENERGY_LABEL_JOULES = "joules"

const FORCE_FROM_KILOGRAMS_TO_NEWTONS float64 = 9.80665 // kg times meters per second squared
const FORCE_LABEL_NEWTONS string = "newtons"
const FORCE_LABEL_FOOTPOUNDS = "foot-pounds"
const GRAVITY_MPS float64 = 9.80665 // meters per second squared

const MOMENTUM_LABEL_FPS = "foot-pound per second"
const MOMENTUM_LABEL_MKS = "meter kilogram per second"
const MOMENTUM_LABEL_NS = "newton second"

const LENGTH_FROM_CENTIMETERS_TO_METERS float64 = 0.01
const LENGTH_FROM_FEET_TO_METERS float64 = 0.3048
const LENGTH_FROM_INCHES_TO_CM float64 = 2.54
const LENGTH_FROM_INCHES_TO_METERS float64 = 0.0254
const LENGTH_FROM_KILOMETERS_TO_METERS float64 = 1000.0
const LENGTH_FROM_METERS_TO_CENTIMETERS float64 = 100
const LENGTH_FROM_METERS_TO_FEET float64 = 3.28084
const LENGTH_FROM_METERS_TO_INCHES float64 = 39.3701
const LENGTH_FROM_METERS_TO_KILOMETERS float64 = 0.001
const LENGTH_FROM_METERS_TO_MILES float64 = 0.000621371
const LENGTH_FROM_METERS_TO_MILLIMETERS float64 = 1000
const LENGTH_FROM_METERS_TO_NAUTICAL_MILES float64 = 0.000539957
const LENGTH_FROM_MILES_TO_METERS float64 = 1609.34
const LENGTH_FROM_MILLIMETERS_TO_METERS float64 = 0.001
const LENGTH_FROM_NAUTICAL_MILES_TO_METERS float64 = 1852
const LENGTH_FROM_YARDS_TO_METERS float64 = 0.9144
const LENGTH_LABEL_CENTIMETER = "centimeters"
const LENGTH_LABEL_FOOT = "feet"
const LENGTH_LABEL_INCH = "inches"
const LENGTH_LABEL_KILOMETER = "kilometers"
const LENGTH_LABEL_METER = "meters"
const LENGTH_LABEL_MILE = "miles"
const LENGTH_LABEL_MILLIMETER = "millimeters"
const LENGTH_LABEL_NAUTICAL_MILE = "nautical miles"
const LENGTH_LABEL_YARD = "yards"

const MASS_FROM_GRAINS_TO_GRAMS float64 = 0.0647989
const MASS_FROM_GRAINS_TO_KILOGRAMS float64 = 0.0000647989
const MASS_FROM_GRAMS_TO_KILOGRAMS float64 = 0.001
const MASS_FROM_KILOGRAMS_TO_POUNDS float64 = 2.20462
const MASS_FROM_POUNDS_TO_GRAMS float64 = 453.592
const MASS_FROM_POUNDS_TO_KILOGRAMS float64 = 0.453592
const MASS_FROM_STONE_TO_GRAMS float64 = 6350.288
const MASS_FROM_STONE_TO_KILOGRAMS float64 = 6.350288
const MASS_FROM_STONE_TO_POUNDS float64 = 14.0
const MASS_FROM_TONS_LONG_TO_GRAMS float64 = 1016047.203454
const MASS_FROM_TONS_LONG_TO_KILOGRAMS float64 = 1016.047203454
const MASS_FROM_TONS_METRIC_TO_GRAMS float64 = 1000000.0
const MASS_FROM_TONS_METRIC_TO_KILOGRAMS float64 = 1000.0
const MASS_FROM_TONS_SHORT_TO_GRAMS float64 = 907185.0
const MASS_FROM_TONS_SHORT_TO_KILOGRAMS float64 = 907.185
const MASS_LABEL_GRAINS = "grains"
const MASS_LABEL_GRAMS = "grams"
const MASS_LABEL_LONG_TON = "long ton"
const MASS_LABEL_METRIC_TONNE = "metric tonne"
const MASS_LABEL_POUNDS = "pounds"
const MASS_LABEL_SHORT_TON = "short ton"
const MASS_LABEL_STONE = "stone"

const VELOCITY_FROM_FPS_TO_MPS float64 = 0.3048
const VELOCITY_FROM_KMPH_TO_MPS float64 = 0.277778
const VELOCITY_FROM_KNOTS_TO_KMPH float64 = 1.852
const VELOCITY_FROM_KNOTS_TO_MPS float64 = 0.514444
const VELOCITY_FROM_MPH_TO_MPS float64 = 0.44704
const VELOCITY_FROM_MPS_TO_FPS float64 = 3.28084
const VELOCITY_FROM_MPS_TO_KMPH float64 = 3.6
const VELOCITY_FROM_MPS_TO_KNOTS float64 = 1.94384
const VELOCITY_FROM_MPS_TO_MPH float64 = 2.23694
const VELOCITY_LABEL_FPS = "feet per second"
const VELOCITY_LABEL_KMPH = "kilometers per hour"
const VELOCITY_LABEL_KNOTS = "knots"
const VELOCITY_LABEL_MPH = "miles per hour"
const VELOCITY_LABEL_MPS = "meters per second"


var /* const */ VALUE_RE = regexp.MustCompile("([0-9]*[0-9.]?[0-9]*)([a-z#]*)")
// var /* const */ VALUE_RE = regexp.MustCompile("([0-9.]+)([a-z#]*)")
const VALUE_TYPE_LENGTH string = "length"
const VALUE_TYPE_VELOCITY string = "velocity"
const VALUE_TYPE_MASS string = "weight"


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

type InputUnits struct {
	length string
	mass string
	metric bool
	velocity string
	weight string
}

type LabeledValue struct {
	Label string  `json:"label,omitempty"`
	Value float64 `json:"value,omitempty"`
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

type ParsedData struct {
	label string
	value float64
	user_label string
	user_value float64
}


//
// VARIABLES
//
var data BallisticData
var output OutputData
var input_units InputUnits
var output_debug bool = false
var output_json bool = false
var output_pretty bool = false
var output_indent string = "    "


//
// FUNCTIONS
//

/** Build output data */
func buildOutputData(data BallisticData) {

	if data.projectile_velocity.value > 0 {
		output.Velocity = velocity_to_velocity(data)
	}
	
	output.Energy = calc_energy(data)
	output.Momentum = calc_momentum(data)

	if output_debug {
		fmt.Println("")
		fmt.Println("Internal Metric:")
		if data.projectile_velocity.value > 0 {
			fmt.Printf("  Projectile Velocity: %12.6f %s\n", data.projectile_velocity.value, data.projectile_velocity.label)
		}
		if output.Energy.Value > 0 {
			fmt.Printf("    Projectile Energy: %12.6f %s\n", output.Energy.Value, output.Energy.Label)
		}
		if output.Momentum.Value > 0 {
			fmt.Printf("  Projectile Momentum: %12.6f %s\n", output.Momentum.Value, output.Momentum.Label)
		}
		if data.mpbr.value > 0 {
			fmt.Printf("Max Point Blank Range: %12.6f %s\n", data.mpbr.value, data.mpbr.label)
		}
		fmt.Println("") 
	}

	if input_units.metric == false {
		output.Energy.Value *= ENERGY_FROM_JOULES_TO_FOOTPOUNDS
		output.Energy.Label = ENERGY_LABEL_FOOTPOUNDS

		output.Momentum.Value *= MASS_FROM_KILOGRAMS_TO_POUNDS * VELOCITY_FROM_MPS_TO_FPS
		output.Momentum.Label = MOMENTUM_LABEL_FPS
	}

	if data.mpbr.value > 0 {
		if output_debug { fmt.Printf("MPBR %f %s\n", data.mpbr.value, data.mpbr.label) }
		output.Mpbr = mpbr_to_mpbr(data)
		if output_debug { fmt.Printf("MPBR %f %s\n", output.Mpbr.Value, output.Mpbr.Label) }
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
func calc_energy(data BallisticData) (energy LabeledValue) {
	// 1 Joule == 1 N⋅m (Newton meter)
	mass_kg := data.projectile_mass.value
	velocity_mps := data.projectile_velocity.value

	energy.Value = mass_kg * velocity_mps * velocity_mps
	energy.Label = ENERGY_LABEL_JOULES

	if output_debug {
		log.Printf("calc_energy()   <| mass: %f kg", mass_kg)
		log.Printf("calc_energy()   <| velocity: %f mps", velocity_mps)
		log.Printf("calc_energy()    | energy: %f %s", energy.Value, energy.Label)
	}

	return energy
}


/** Calculate force Newtons */ 
func calc_force(avg_draw_weight float64) (draw_force ParsedData) {
	draw_force.value = avg_draw_weight * FORCE_FROM_KILOGRAMS_TO_NEWTONS
	draw_force.label = FORCE_LABEL_NEWTONS

	if output_debug {
		log.Printf("calc_force()    <| avg_draw_weight: %f kg", avg_draw_weight)
		log.Printf("calc_force()     | draw_force: %f %s", draw_force.value, draw_force.label)
	}
	
	return draw_force
}


/** Calculate momentum */
func calc_momentum(data BallisticData) (momentum LabeledValue) {
	// kg⋅m/s (kilogram meters per second)
	mass_kg := data.projectile_mass.value
	velocity_mps := data.projectile_velocity.value

	momentum.Value = mass_kg * velocity_mps
	momentum.Label = MOMENTUM_LABEL_MKS

	if output_debug {
		log.Printf("calc_momentum() <| mass: %f kg", mass_kg)
		log.Printf("calc_momentum() <| velocity: %f mps", velocity_mps)
		log.Printf("calc_momentum()  | momentum: %f %s", momentum.Value, momentum.Label)
	}
		
	return momentum
}


/** Calculate Maximum Point Blank Range */
func calc_mpbr(data BallisticData) (mpbr ParsedData) {
	distance := 0.0
	diameter := data.target_radius.value * 2
	drop := data.target_radius.value
	velocity := data.projectile_velocity.value

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

	mpbr.value = distance
	mpbr.label = "meters"

	// log.Printf("calc_mpbr() <|   target radius: %12.6f m", data.target_radius.value)
	// log.Printf("calc_mpbr() <| target diameter: %12.6f m", diameter)
	// log.Printf("calc_mpbr() <| projectile drop: %12.6f m", drop)
	// log.Printf("calc_mpbr()  |            MPBR: %12.6f m", distance)

	return mpbr
}

/**
 * Calculate velocity
 */
func calc_velocity(data BallisticData) (projectile_velocity ParsedData) {
	projectile_mass := data.projectile_mass.value
	draw_length := data.draw_length.value
	draw_force := data.draw_force.value

	release_time := math.Sqrt(projectile_mass * draw_length / draw_force)

	projectile_velocity.value = draw_length / release_time
	projectile_velocity.label = VELOCITY_LABEL_MPS

	if len(input_units.velocity) == 0 {
		if input_units.metric {
			input_units.velocity = VELOCITY_LABEL_MPS
		} else {
			input_units.velocity = VELOCITY_LABEL_FPS
		}
	}

	if output_debug {
		log.Printf("calc_velocity() <|     projectile mass: %15.6f kg", projectile_mass)
		log.Printf("calc_velocity() <|         draw length: %15.6f m", draw_length)
		log.Printf("calc_velocity() <|          draw force: %15.6f N", draw_force)
		log.Printf("calc_velocity()  | projectile velocity: %15.6f mps", projectile_velocity.value)
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

	if data.Energy.Value != 0 {
		data_obj["energy"] = data.Energy
	}
	if data.Momentum.Value != 0 {
		data_obj["momentum"] = data.Momentum
	}
	if data.Mpbr.Value != 0 {
		data_obj["mpbr"] = data.Mpbr
	}
	if data.Velocity.Value != 0 {
		data_obj["velocity"] = data.Velocity
	}

	return data_obj
}


/** Convert MPBR in meters to input units */
func mpbr_to_mpbr(data BallisticData) (mpbr LabeledValue) {
	mpbr.Label = ""
	mpbr.Value = 0.0

	user_label := data.projectile_velocity.user_label
	if len(user_label) == 0 {
		user_label = input_units.velocity
	}

	switch user_label {
	case VELOCITY_LABEL_FPS:
		mpbr.Label = LENGTH_LABEL_FOOT
		mpbr.Value = data.mpbr.value * LENGTH_FROM_METERS_TO_FEET
	case VELOCITY_LABEL_KMPH:
		mpbr.Label = LENGTH_LABEL_KILOMETER
		mpbr.Value = data.mpbr.value * LENGTH_FROM_METERS_TO_KILOMETERS
	case VELOCITY_LABEL_KNOTS:
		mpbr.Label = LENGTH_LABEL_NAUTICAL_MILE
		mpbr.Value = data.mpbr.value * LENGTH_FROM_METERS_TO_NAUTICAL_MILES
	case VELOCITY_LABEL_MPS:
		mpbr.Label = LENGTH_LABEL_METER
		mpbr.Value = data.mpbr.value
	case VELOCITY_LABEL_MPH:
		mpbr.Label = LENGTH_LABEL_MILE
		mpbr.Value = data.mpbr.value * LENGTH_FROM_METERS_TO_MILES
	}

	if output_debug {
		log.Printf("mpbr_to_mpbr()  <|  projectile velocity: %s", data.projectile_velocity.user_label)
		log.Printf("mpbr_to_mpbr()  <| input_units velocity: %s", input_units.velocity)
		log.Printf("mpbr_to_mpbr()   |  projectile velocity: %15.6f mps", mpbr.Label, mpbr.Label)
	}

	return mpbr
}


/** Print Human Readable Output */
func outputHuman(data OutputData) {
	fmt.Println("")
	if data.Velocity.Value > 0 {
		fmt.Printf("  Projectile Velocity: %12.6f %s\n", data.Velocity.Value, data.Velocity.Label)
	}
	if data.Energy.Value > 0 {
		fmt.Printf("    Projectile Energy: %12.6f %s\n", data.Energy.Value, data.Energy.Label)
	}
	if data.Momentum.Value > 0 {
		fmt.Printf("  Projectile Momentum: %12.6f %s\n", data.Momentum.Value, data.Momentum.Label)
	}
	if data.Mpbr.Value > 0 {
		fmt.Printf("Max Point Blank Range: %12.6f %s\n", data.Mpbr.Value, data.Mpbr.Label)
	}
	fmt.Println("")
}


/** Outputs JSON data */
func outputJSON(data OutputData) {
	if output_debug {
		fmt.Println("JSON data!")
		fmt.Printf("data.Energy: %f %s\n", data.Energy.Value, data.Energy.Label)
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


/** Parse user input value and normalize it for internal use */
func parse_value(value, value_type string) (parsed_data ParsedData) {
	// log.Printf("parse_value()  <| value: %s | value_type: %s", value, value_type)

	if len(value) > 0 {
		value_match := VALUE_RE.FindStringSubmatch(value)

		number, _ := strconv.ParseFloat(value_match[1], 64)
		suffix := strings.ToLower(value_match[2])

		var designation string
		var norm_type string
		var norm_value float64

		switch value_type {
		case VALUE_TYPE_LENGTH:
			norm_type = "meter"

			switch suffix {
			case "feet", "foot", "ft", "f":
				norm_value = number * LENGTH_FROM_FEET_TO_METERS
				designation = LENGTH_LABEL_FOOT
				input_units.metric = false
			case "inches", "inch", "in", "i":
				norm_value = number * LENGTH_FROM_INCHES_TO_METERS
				designation = LENGTH_LABEL_INCH
				input_units.metric = false
			case "nmi", "nm", "M":
				// Actual Values: M, NM, Nm, nm, nmi
				norm_value = number * LENGTH_FROM_NAUTICAL_MILES_TO_METERS
				designation = LENGTH_LABEL_NAUTICAL_MILE
			case "yards", "yard", "yrd", "yd", "y":
				norm_value = number * LENGTH_FROM_YARDS_TO_METERS
				designation = LENGTH_LABEL_YARD
				input_units.metric = false
			case "kilometers", "kilometer", "kilo", "km", "k":
				norm_value = number * LENGTH_FROM_KILOMETERS_TO_METERS
				designation = LENGTH_LABEL_KILOMETER
				input_units.metric = true
			case "meters", "m", "":
				norm_value = number
				designation = LENGTH_LABEL_METER
				input_units.metric = true
			case "centimeters", "centimeter", "centi", "cm", "c":
				norm_value = number * LENGTH_FROM_CENTIMETERS_TO_METERS
				designation = LENGTH_LABEL_CENTIMETER
				input_units.metric = true
			case "millimeters", "millimeter", "milli", "mm":
				norm_value = number * LENGTH_FROM_MILLIMETERS_TO_METERS
				designation = LENGTH_LABEL_MILLIMETER
				input_units.metric = true
			}
			input_units.length = designation
		case VALUE_TYPE_MASS:
			norm_type = "kilogram"

			switch suffix {
			case "grams", "g", "":
				norm_value = number * MASS_FROM_GRAMS_TO_KILOGRAMS
				designation = MASS_LABEL_GRAMS
				input_units.metric = true
			case "grains", "gr":
				norm_value = number * MASS_FROM_GRAINS_TO_KILOGRAMS
				designation = MASS_LABEL_GRAINS
				input_units.metric = false
			case "pounds", "#", "lb", "lbs":
				norm_value = number * MASS_FROM_POUNDS_TO_KILOGRAMS
				designation = MASS_LABEL_POUNDS
				input_units.metric = false
			case "stone", "st":
				norm_value = number * MASS_FROM_STONE_TO_KILOGRAMS
				designation = MASS_LABEL_STONE
				input_units.metric = false
			case "ton":
				norm_value = number * MASS_FROM_TONS_SHORT_TO_KILOGRAMS
				designation = MASS_LABEL_SHORT_TON
				input_units.metric = false
			case "lt":
				norm_value = number * MASS_FROM_TONS_LONG_TO_KILOGRAMS
				designation = MASS_LABEL_LONG_TON
				input_units.metric = false
			case "mt":
				norm_value = number * MASS_FROM_TONS_METRIC_TO_KILOGRAMS
				designation = MASS_LABEL_METRIC_TONNE
				input_units.metric = true
			}
			input_units.mass = designation
		case VALUE_TYPE_VELOCITY:
			norm_type = "meters per second"

			switch suffix {
			case "fps":
				norm_value = number * VELOCITY_FROM_FPS_TO_MPS
				designation = VELOCITY_LABEL_FPS
				input_units.metric = false
			case "knots", "knot", "kn", "kt":
				norm_value = number * VELOCITY_FROM_KNOTS_TO_MPS
				designation = VELOCITY_LABEL_KNOTS
			case "kmph", "k":
				norm_value = number * VELOCITY_FROM_KMPH_TO_MPS
				designation = VELOCITY_LABEL_KMPH
				input_units.metric = true
			case "mph":
				norm_value = number * VELOCITY_FROM_MPH_TO_MPS
				designation = VELOCITY_LABEL_MPH
				input_units.metric = false
			case "mps", "":
				norm_value = number
				designation = VELOCITY_LABEL_MPS
				input_units.metric = true
			}
			
			input_units.velocity = designation
		}

		if output_debug {
			// log.Printf("parse_value()   <| value_match: %s", value_match)
			// log.Printf("parse_value()   <|      number: %f", number)
			// log.Printf("parse_value()   <|      suffix: %s", suffix)
			// log.Printf("parse_value()    | designation: %s", designation)
			// log.Printf("parse_value()    |  norm_value: %f", norm_value)
			// log.Printf("parse_value()    |   norm_type: %s", norm_type)

			// log.Printf("parse_value()   <| %8s value: %17.6f %s (%s)", value_type, number, suffix, designation )
			// log.Printf("parse_value()    | %8s normilazed: %12.6f %s", value_type, norm_value, norm_type)

			suffix_designation := suffix + " (" + designation + ")"
			log.Printf("parse_value()    | %8s: %12.6f %-20s | %12.6f %s", value_type, number, suffix_designation, norm_value, norm_type)
		}

		parsed_data.label = norm_type
		parsed_data.value = norm_value
		parsed_data.user_label = designation
		parsed_data.user_value = number
	}

	return parsed_data
}


/** Convert velocity in mps to input units */
func velocity_to_velocity(data BallisticData) (velocity LabeledValue) {
	velocity.Label = ""

	user_label := input_units.velocity
	if len(user_label) == 0 {
		switch input_units.mass {
		case MASS_LABEL_GRAINS, MASS_LABEL_LONG_TON, MASS_LABEL_POUNDS, MASS_LABEL_SHORT_TON, MASS_LABEL_STONE:
			user_label = VELOCITY_LABEL_FPS
		default:
			user_label = VELOCITY_LABEL_MPS
		}
	}

	switch user_label {
	case VELOCITY_LABEL_FPS:
		velocity.Label = user_label
		velocity.Value = data.projectile_velocity.value * VELOCITY_FROM_MPS_TO_FPS
	case VELOCITY_LABEL_KMPH:
		velocity.Label = user_label
		velocity.Value = data.projectile_velocity.value * VELOCITY_FROM_MPS_TO_KMPH
	case VELOCITY_LABEL_KNOTS:
		velocity.Label = user_label
		velocity.Value = data.projectile_velocity.value * VELOCITY_FROM_MPS_TO_KNOTS
	case VELOCITY_LABEL_MPS:
		velocity.Label = user_label
		velocity.Value = data.projectile_velocity.value
	case VELOCITY_LABEL_MPH:
		velocity.Label = user_label
		velocity.Value = data.projectile_velocity.value * VELOCITY_FROM_MPS_TO_MPH
	}

	// log.Printf("velocity_to_velocity() | user_label: '%s'", user_label)
	// log.Printf("velocity_to_velocity() | projectile_mass user value & label: %f %s", data.projectile_mass.user_value, data.projectile_mass.user_label)
	// log.Printf("velocity_to_velocity() | input_units | velocity: '%s' | mass: '%s' | metric: %t", input_units.velocity, input_units.mass, input_units.metric)
	// log.Printf("velocity_to_velocity() | projectile_velocity user value & label: %f %s", data.projectile_velocity.user_value, data.projectile_velocity.user_label)
	// log.Printf("velocity_to_velocity() | projectile_velocity norm value & label: %f %s", data.projectile_velocity.value, data.projectile_velocity.label)
	// log.Printf("velocity_to_velocity() | projectile_velocity calc value & label: %f %s", velocity.Value, velocity.Label)

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
			Name: "mass, m",
			Usage: "Projectile `MASS` (weight). Used to calculate projectile velocity, energy, etc.",
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

	cli.AppHelpTemplate = fmt.Sprintf(`
NAME:
   {{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

VERSION:
   {{.Version}}{{end}}{{end}}{{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if len .Authors}}

AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
   {{range $index, $author := .Authors}}{{if $index}}
   {{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}

GLOBAL OPTIONS:
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}

COPYRIGHT:
   {{.Copyright}}{{end}}

VALUE SUFFIXES:
  All input values may be suffixed to allow for broader input selection.

  LENGTH
    c, cm, centi, centimeter, centimeters
    f, ft, foot, feet
    i, in, inch, inches
    k, km, kilo, kilometer, kilometers
    m, meters †
    M, NM, Nm, nm, nmi  (Nautical Miles)
    mm, milli, millimeter, millimeters
    y, yd, yrd, yard, yards
  MASS
    #, lb, lbs, pound, pounds
    g, gram, grams
    gr, grain, grains
    kg, kilo, kilogram, kilograms †
    lt, long-ton
    mt, tonne, metric-tonne
    st, stone
    t, ton, short-ton
  VELOCITY
    fps, feet-per-second
    kmph, kilometers-per-hour
    kn, kt, knot, knots
    mph, miles-per-hour
    mps, meters-per-second †

†  This is the default and will be used if no suffix is specified

If most or all of the input values are in imperial units then the output will use imperial units as well.

`)


// ‡  This is the default if you set BALLISTIC_UNITS to imperial instead of metric
// The environment variable BALLISTIC_UNITS can be defined as imperial or metric. If it is defined the output units will always be of that system. If BALLISTIC_UNITS is not defined and most or all of the input values are in imperial units then the output will use imperial units as well. Otherwise ballistic defaults to metric.

	cli.VersionFlag = cli.BoolFlag{
		Name: "version, V",
		Usage: "Output the ballistic app version",
	}

	// app.HideHelp = true
	// app.Commands = []cli.String{}
	// app.Commands = []cli.Command{
	// 	{
	// 		Name: "mpbr",
	// 		Usage: "Calculates the maximum point blank range or a weapon.",
	// 		Flags: []cli.Flag {
	// 			cli.StringFlag{
	// 				Name: "radius, r",
	// 				Usage: "The `RADIUS` of the target area",
	// 			},
	// 			cli.StringFlag{
	// 				Name: "velocity, v",
	// 				Usage: "The projectile `VELOCITY` or SPEED",
	// 			},
	// 		},
	// 		Action:  func(c *cli.Context) error {
	// 			log.Printf("  radius: %f", c.Float64("radius"))
	// 			log.Printf("velocity: %f", c.Float64("velocity"))
	// 			// radius := strconv.ParseFloat(c.String("radius"), 64)
	// 			// velocity := strconv.ParseFloat(c.String("velocity"), 64)
	// 			// calc_mpbr(c.Float64("radius"), c.Float64("velocity"))
	// 			return nil
	// 		},
	// 	},
	// }

	sort.Sort(cli.FlagsByName(app.Flags))
	// sort.Sort(cli.CommandsByName(app.Commands))

	app.Action = func(c *cli.Context) error {
		output_debug = c.Bool("debug")
		output_json = c.Bool("json")
		output_pretty = c.Bool("pretty-print")

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
			log.Printf("        draw length: %9s (%d)", c.String("draw-length"), len(c.String("draw-length")))
			log.Printf("        draw weight: %9s (%d)", c.String("draw-weight"), len(c.String("draw-length")))
			log.Printf("      target radius: %9s (%d)", c.String("radius"), len(c.String("draw-length")))
			log.Printf("projectile velocity: %9s (%d)", c.String("velocity"), len(c.String("draw-length")))
			log.Printf("    projectile mass: %9s (%d)", c.String("mass"), len(c.String("draw-length")))
			// log.Println("")
			// fmt.Println("")
		}


		flags_set := 0
		for _, flag_name := range c.GlobalFlagNames() {
			// fmt.Printf("Flag: %s\n", flag_name)
			flag_value := c.String(flag_name)
			if len(flag_value) > 0 {
				if flag_value != "false" && flag_value != "true" {
					// fmt.Printf("Flag Set: %s\n", flag_value)
					flags_set += 1
				}
			}
		}
		// fmt.Printf("Flags Set: %d\n", flags_set)

		if flags_set == 0 {
			// fmt.Println("No Args!")
			cli.ShowAppHelpAndExit(c, 0)
		}


		if len(c.String("draw-length")) > 0 {
			data.draw_length = parse_value(c.String("draw-length"), VALUE_TYPE_LENGTH)
		}
		if len(c.String("draw-weight")) > 0 {
			data.draw_weight = parse_value(c.String("draw-weight"), VALUE_TYPE_MASS)
			avg_draw_weight := data.draw_weight.value * 0.5
			data.draw_force = calc_force(avg_draw_weight)
		}
		if len(c.String("velocity")) > 0 {
			data.projectile_velocity = parse_value(c.String("velocity"), VALUE_TYPE_VELOCITY)
		}
		if len(c.String("mass")) > 0 {
			data.projectile_mass = parse_value(c.String("mass"), VALUE_TYPE_MASS)
		}

		if data.projectile_velocity.value == 0 {
			if data.projectile_mass.value > 0 && data.draw_length.value > 0 && data.draw_force.value > 0 {
				data.projectile_velocity = calc_velocity(data)
			}
		}

		if len(c.String("radius")) > 0 {
			data.target_radius = parse_value(c.String("radius"), VALUE_TYPE_LENGTH)
		} else {
			data.target_radius = parse_value("225mm", VALUE_TYPE_LENGTH)
		}
		if data.projectile_velocity.value > 0 {
			data.mpbr = calc_mpbr(data)
		}

		buildOutputData(data)

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


