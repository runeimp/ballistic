package main


//
// IMPORTS
//
import (
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
	"strconv"
	"strings"
	// "syscall"
)


//
// CONSTANTS
//
const FORCE_TO_NEWTONS float64 = 9.80665 // kg times meters per second squared
const GRAVITY_MPS float64 = 9.80665 // kg times meters per second squared

const LENGTH_CENTIMETERS_TO_METERS float64 = 0.01
const LENGTH_FEET_TO_METERS float64 = 0.3048
const LENGTH_INCHES_TO_CM float64 = 2.54
const LENGTH_INCHES_TO_METERS float64 = 0.0254
const LENGTH_KILOMETERS_TO_METERS float64 = 1000.0
const LENGTH_METERS_TO_CENTIMETERS float64 = 100
const LENGTH_METERS_TO_FEET float64 = 3.28084
const LENGTH_METERS_TO_INCHES float64 = 39.3701
const LENGTH_METERS_TO_KILOMETERS float64 = 0.001
const LENGTH_METERS_TO_MILES float64 = 0.000621371
const LENGTH_METERS_TO_MILLIMETERS float64 = 1000
const LENGTH_METERS_TO_NAUTICAL_MILES float64 = 0.000539957
const LENGTH_MILES_TO_METERS float64 = 1609.34
const LENGTH_MILLIMETERS_TO_METERS float64 = 0.001
const LENGTH_NAUTICAL_MILES_TO_METERS float64 = 1852
const LENGTH_YARDS_TO_METERS float64 = 0.9144

const MASS_GRAINS_TO_GRAMS float64 = 0.0647989
const MASS_GRAINS_TO_KILOGRAMS float64 = 0.0000647989
const MASS_GRAMS_TO_KILOGRAMS float64 = 0.001
const MASS_POUNDS_TO_GRAMS float64 = 453.592
const MASS_POUNDS_TO_KILOGRAMS float64 = 0.453592
const MASS_STONE_TO_GRAMS float64 = 6350.288
const MASS_STONE_TO_KILOGRAMS float64 = 6.350288
const MASS_STONE_TO_POUNDS float64 = 14.0
const MASS_TONS_LONG_TO_GRAMS float64 = 1016047.203454
const MASS_TONS_LONG_TO_KILOGRAMS float64 = 1016.047203454
const MASS_TONS_METRIC_TO_GRAMS float64 = 1000000.0
const MASS_TONS_METRIC_TO_KILOGRAMS float64 = 1000.0
const MASS_TONS_SHORT_TO_GRAMS float64 = 907185.0
const MASS_TONS_SHORT_TO_KILOGRAMS float64 = 907.185

const VELOCITY_FPS_TO_MPS float64 = 0.3048
const VELOCITY_KMPH_TO_MPS float64 = 0.277778
const VELOCITY_KNOTS_TO_KMPH float64 = 1.852
const VELOCITY_KNOTS_TO_MPS float64 = 0.514444
const VELOCITY_MPH_TO_MPS float64 = 0.44704
const VELOCITY_MPS_TO_FPS float64 = 3.28084
const VELOCITY_MPS_TO_KMPH float64 = 3.6
const VELOCITY_MPS_TO_KNOTS float64 = 1.94384
const VELOCITY_MPS_TO_MPH float64 = 2.23694


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

type LabeledValue struct {
	label string
	value float64
}

type OutputData struct {
	energy LabeledValue
	momentum LabeledValue
	mpbr LabeledValue
	velocity LabeledValue
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


//
// FUNCTIONS
//
/** Calculate drop in flight */
func calc_drop(distance, velocity float64) (drop float64) {
	flight_time := distance / velocity
	drop = GRAVITY_MPS * 0.5 * (flight_time * flight_time)
	return drop
}

/** Calculate energy */
func calc_energy(data BallisticData) (energy LabeledValue) {
	// 1 Joule == 1 N⋅m (Newton meter)
	mass_kg := data.projectile_mass.value
	velocity_mps := data.projectile_velocity.value

	energy.value = mass_kg * velocity_mps * velocity_mps
	energy.label = "joule"

	return energy
}


/** Calculate force Newtons */ 
func calc_force(avg_draw_weight float64) (draw_force ParsedData) {
	draw_force.value = avg_draw_weight * GRAVITY_MPS
	draw_force.label = "newton"
	return draw_force
}


/** Calculate momentum */
func calc_momentum(data BallisticData) (momentum LabeledValue) {
	// kg⋅m/s (kilogram meters per second)
	mass_kg := data.projectile_mass.value
	velocity_mps := data.projectile_velocity.value

	momentum.value = mass_kg * velocity_mps
	momentum.label = "kilogram meters per second"

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
		drop = calc_drop(distance, velocity)
	}
	for drop > diameter {
		distance -= 0.1
		drop = calc_drop(distance, velocity)
	}
	for drop < diameter {
		distance += 0.01
		drop = calc_drop(distance, velocity)
	}
	for drop > diameter {
		distance -= 0.001
		drop = calc_drop(distance, velocity)
	}
	for drop < diameter {
		distance += 0.0001
		drop = calc_drop(distance, velocity)
	}
	for drop > diameter {
		distance -= 0.00001
		drop = calc_drop(distance, velocity)
	}
	for drop < diameter {
		distance += 0.000001
		drop = calc_drop(distance, velocity)
	}
	for drop > diameter {
		distance -= 0.0000001
		drop = calc_drop(distance, velocity)
	}

	mpbr.value = distance
	mpbr.label = "meter"

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
	projectile_velocity.label = "meters per second"

	// log.Printf("calc_velocity() <|     projectile mass: %15.6f kg", projectile_mass)
	// log.Printf("calc_velocity() <|         draw length: %15.6f m", draw_length)
	// log.Printf("calc_velocity() <|          draw force: %15.6f N", draw_force)
	// log.Printf("calc_velocity()  | projectile velocity: %15.6f mps", projectile_velocity.value)
	return projectile_velocity
}


/** Convert MPBR in meters to input units */
func mpbr_to_mpbr(data BallisticData) (mpbr LabeledValue) {
	mpbr.label = ""
	mpbr.value = 0.0

	switch data.projectile_velocity.user_label {
	case "feet per second":
		mpbr.label = "foot"
		mpbr.value = data.mpbr.value * LENGTH_METERS_TO_FEET
	case "kilometers per hour":
		mpbr.label = "kilometer"
		mpbr.value = data.mpbr.value * LENGTH_METERS_TO_KILOMETERS
	case "knots":
		mpbr.label = "nautical mile"
		mpbr.value = data.mpbr.value * LENGTH_METERS_TO_NAUTICAL_MILES
	case "meters per second":
		mpbr.label = "meter"
		mpbr.value = data.mpbr.value
	case "miles per hour":
		mpbr.label = "mile"
		mpbr.value = data.mpbr.value * LENGTH_METERS_TO_MILES
	}

	return mpbr
}


/** Generate output data */
func output_data(data BallisticData) {
	if data.projectile_velocity.value > 0 {
		output.velocity = velocity_to_velocity(data)
	}
	
	output.energy = calc_energy(data)
	output.momentum = calc_momentum(data)

	if data.mpbr.value > 0 {
		// fmt.Printf("MPBR ", data.mpbr.value)
		// output.mpbr.value = data.mpbr.value
		// output.mpbr.label = velocity_to_length(data.projectile_velocity.label)
		output.mpbr = mpbr_to_mpbr(data)
	}

	print_output(output)
}

func print_output(data OutputData) {
	fmt.Println("")
	if data.velocity.value > 0 {
		fmt.Printf("  Projectile Velocity: %12.6f %s\n", data.velocity.value, data.velocity.label)
	}
	if data.energy.value > 0 {
		fmt.Printf("    Projectile Energy: %12.6f %s\n", data.energy.value, data.energy.label)
	}
	if data.momentum.value > 0 {
		fmt.Printf("  Projectile Momentum: %12.6f %s\n", data.momentum.value, data.momentum.label)
	}
	if data.mpbr.value > 0 {
		fmt.Printf("Max Point Blank Range: %12.6f %s\n", data.mpbr.value, data.mpbr.label)
	}
	fmt.Println("")
}


/** Parse user input value and normalize it for internal use */
func parse_value(value, value_type string) (parsed_data ParsedData) {
	// log.Printf("parse_value() <| value: %s | value_type: %s", value, value_type)

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
				norm_value = number * LENGTH_FEET_TO_METERS
				designation = "foot"
			case "inches", "inch", "in", "i":
				norm_value = number * LENGTH_INCHES_TO_METERS
				designation = "inch"
			case "nmi", "nm", "M":
				// Actual Values: M, NM, Nm, nm, nmi
				norm_value = number * LENGTH_NAUTICAL_MILES_TO_METERS
				designation = "nautical mile"
			case "yards", "yard", "yrd", "yd", "y":
				norm_value = number * LENGTH_YARDS_TO_METERS
				designation = "yard"
			case "kilometers", "kilometer", "kilo", "km", "k":
				norm_value = number * LENGTH_KILOMETERS_TO_METERS
				designation = "kilometer"
			case "meters", "m", "":
				norm_value = number
				designation = "meter"
			case "centimeters", "centimeter", "centi", "cm", "c":
				norm_value = number * LENGTH_CENTIMETERS_TO_METERS
				designation = "centimeter"
			case "millimeters", "millimeter", "milli", "mm":
				norm_value = number * LENGTH_MILLIMETERS_TO_METERS
				designation = "millimeter"
			}
		case VALUE_TYPE_VELOCITY:
			norm_type = "meters per second"

			switch suffix {
			case "fps":
				norm_value = number * VELOCITY_FPS_TO_MPS
				designation = "feet per second"
			case "knots", "knot", "kn":
				norm_value = number * VELOCITY_KNOTS_TO_MPS
				designation = "knots"
			case "kmph", "k":
				norm_value = number * VELOCITY_KMPH_TO_MPS
				designation = "kilometers per hour"
			case "mph":
				norm_value = number * VELOCITY_MPH_TO_MPS
				designation = "miles per hour"
			case "mps", "":
				norm_value = number
				designation = "meters per second"
			}
		case VALUE_TYPE_MASS:
			norm_type = "kilogram"

			switch suffix {
			case "grams", "g", "":
				norm_value = number * MASS_GRAMS_TO_KILOGRAMS
				designation = "grams"
			case "grains", "gr":
				norm_value = number * MASS_GRAINS_TO_KILOGRAMS
				designation = "grains"
			case "pounds", "#", "lb", "lbs":
				norm_value = number * MASS_POUNDS_TO_KILOGRAMS
				designation = "pounds"
			case "stone", "st":
				norm_value = number * MASS_STONE_TO_KILOGRAMS
				designation = "stone"
			case "ton":
				norm_value = number * MASS_TONS_SHORT_TO_KILOGRAMS
				designation = "short ton"
			case "lt":
				norm_value = number * MASS_TONS_LONG_TO_KILOGRAMS
				designation = "long ton"
			case "mt":
				norm_value = number * MASS_TONS_METRIC_TO_KILOGRAMS
				designation = "metric ton"
			}
		}

		// log.Printf("parse_value()  | value_match: %s", value_match)
		// log.Printf("parse_value()  |      number: %f", number)
		// log.Printf("parse_value()  |      suffix: %s", suffix)
		// log.Printf("parse_value()  | designation: %s", designation)
		// log.Printf("parse_value()  |  norm_value: %f", norm_value)
		// log.Printf("parse_value()  |   norm_type: %s", norm_type)

		// log.Printf("parse_value() <| %8s value: %17.6f %s (%s)", value_type, number, suffix, designation )
		// log.Printf("parse_value()  | %8s normilazed: %12.6f %s", value_type, norm_value, norm_type)
		suffix_designation := suffix + " (" + designation + ")"
		log.Printf("parse_value() | %8s: %12.6f %-20s | %12.6f %s", value_type, number, suffix_designation, norm_value, norm_type)

		parsed_data.label = norm_type
		parsed_data.value = norm_value
		parsed_data.user_label = designation
		parsed_data.user_value = number
	}

	return parsed_data
}


/** Convert velocity label to length label */
func velocity_to_length(velocity_label string) (length string) {
	length = ""
	switch velocity_label {
	case "fps":
		length = "foot"
	case "knots":
		length = "nautical mile"
	case "kilometers per hour":
		length = "kilometer"
	case "meters per second":
		length = "meter"
	case "miles per hour":
		length = "mile"
	}

	return length
}


/** Convert velocity in mps to input units */
func velocity_to_velocity(data BallisticData) (velocity LabeledValue) {
	velocity.label = ""

	switch data.projectile_velocity.user_label {
	case "feet per second":
		velocity.label = "foot"
		velocity.value = data.projectile_velocity.value * VELOCITY_MPS_TO_FPS
	case "kilometers per hour":
		velocity.label = "kilometer"
		velocity.value = data.projectile_velocity.value * VELOCITY_MPS_TO_KMPH
	case "knots":
		velocity.label = "nautical mile"
		velocity.value = data.projectile_velocity.value * VELOCITY_MPS_TO_KNOTS
	case "meters per second":
		velocity.label = "meter"
		velocity.value = data.projectile_velocity.value
	case "miles per hour":
		velocity.label = "mile"
		velocity.value = data.projectile_velocity.value * VELOCITY_MPS_TO_MPH
	}

	return velocity
}


//
// MAIN ENTRYPOINT
//
func main() {
	// log.Print("Ballistic ...")

	app := cli.NewApp()
	app.Name = "Ballistic"
	app.Usage = "Handles many common and uncommon ballistics calculations"
	app.Version = "0.1.0"
	

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "mass, m",
			Usage: "Projectile `MASS` (weight). Used to calculate projectile velocity, energy, etc.",
		},
		cli.StringFlag{
			Name: "draw-weight, draw, d",
			Usage: "Bow or sling shot draw `WEIGHT`. Used to calculate projectile velocity, energy, etc.",
		},
		cli.StringFlag{
			Name: "draw-length, length, l",
			Usage: "Bow or sling shot draw `LENGTH`. Used to calculate projectile velocity, energy, etc.",
		},
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
		Usage: "Print this help info",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name: "version, V",
		Usage: "Print the ballistic version",
	}


	app.Commands = []cli.Command{
		{
			Name: "mpbr",
			Usage: "Calculates the maximum point blank range or a weapon.",
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "radius, r",
					Usage: "The `RADIUS` of the target area",
				},
				cli.StringFlag{
					Name: "velocity, v",
					Usage: "The projectile `VELOCITY` or SPEED",
				},
			},
			Action:  func(c *cli.Context) error {
				log.Printf("  radius: %f", c.Float64("radius"))
				log.Printf("velocity: %f", c.Float64("velocity"))
				// radius := strconv.ParseFloat(c.String("radius"), 64)
				// velocity := strconv.ParseFloat(c.String("velocity"), 64)
				// calc_mpbr(c.Float64("radius"), c.Float64("velocity"))
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Going Ballistic!")
		log.Printf("        draw length: %9s (str)", c.String("draw-length"))
		log.Printf("        draw weight: %9s (str)", c.String("draw-weight"))
		log.Printf("      target radius: %9s (str)", c.String("radius"))
		log.Printf("projectile velocity: %9s (str)", c.String("velocity"))
		log.Printf("    projectile mass: %9s (str)", c.String("mass"))
		log.Println("")
		// value_match := VALUE_RE.FindStringSubmatch(c.String("weight"))
		// log.Printf("value_match:    %s", value_match)
		// log.Printf("value_match[1]: %s", value_match[1])
		// log.Printf("value_match[2]: %s", value_match[2])


		// var draw_length_meters float64 = 0
		// var draw_weight_kg float64 = 0
		// var force_newtons float64 = 0
		// var projectile_mass_kg float64 = 0
		// var projectile_velocity float64 = 0

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
		// if len(c.String("length")) {}
		// if len(c.String("length")) {}

		if data.projectile_velocity.value == 0 {
			if data.projectile_mass.value > 0 && data.draw_length.value > 0 && data.draw_force.value > 0 {
				data.projectile_velocity = calc_velocity(data)
			}
		}

		if len(c.String("radius")) > 0 {
			data.target_radius = parse_value(c.String("radius"), VALUE_TYPE_LENGTH)
			if data.projectile_velocity.value > 0 {
				data.mpbr = calc_mpbr(data)
			}
		}


		output_data(data)


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


