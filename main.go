//
// PACKAGES
//
package main

/*
 * IMPORTS
 */
import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	"github.com/runeimp/ballistic/app"
	"github.com/runeimp/ballistic/parsing"
	"github.com/runeimp/locale"
	cli "gopkg.in/urfave/cli.v1"
)

/*
 * CONSTANTS
 */
const (
	AppVersion = "0.5.2"
)

/*
 * TYPES
 */

// BallisticData holds all data generated
type BallisticData struct {
	drawForce          parsing.ParsedData
	drawLength         parsing.ParsedData
	drawWeight         parsing.ParsedData
	mpbr               parsing.ParsedData // max_point_blank_range parsing.ParsedData
	projectileEnergy   parsing.ParsedData
	projectileMass     parsing.ParsedData
	projectileRange    parsing.ParsedData
	projectileVelocity parsing.ParsedData
	projectionAngle    parsing.ParsedData
	targetRadius       parsing.ParsedData
}

// LabeledValue provides a label, value, and value as a string
type LabeledValue struct {
	Label       string  `json:"label,omitempty"`
	ValueFloat  float64 `json:"value,omitempty"`
	ValueString string  `json:"value_str,omitempty"`
}

// func (t LabeledValue) MarshalJSON() ([]byte, error) {
// 	return []byte{}, nil
// 	// return nil, nil // <- same effect.
// }

// OutputData is the main data used for output to the screen
type OutputData struct {
	Energy   LabeledValue `json:"energy,omitempty"`
	Momentum LabeledValue `json:"momentum,omitempty"`
	Mpbr     LabeledValue `json:"mpbr,omitempty"`
	Velocity LabeledValue `json:"velocity,omitempty"`
}

//
// VARIABLES
//
var (
	data          BallisticData
	decimalPlaces = 6
	localeStr     string
	output        OutputData
	outputDebug   = false
	outputIndent  = "    "
	outputJSON    = false
	outputPretty  = false
)

//
// FUNCTIONS
//

/** Build output data */
func buildOutputData(data BallisticData) {

	if data.projectileVelocity.Value > 0 {
		output.Velocity = velocityToVelocity(data)
	}

	output.Energy = calcKineticEnergy(data)
	output.Momentum = calcMomentum(data)

	if outputDebug {
		fmt.Println("")
		fmt.Println("Internal Metric:")
		if data.projectileVelocity.Value > 0 {
			fmt.Printf("  Projectile Velocity: %16.6f %s\n", data.projectileVelocity.Value, data.projectileVelocity.Label)
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

	if parsing.InputData.Metric == false {
		output.Energy.ValueFloat *= app.EnergyFromJoulesToFootPounds
		output.Energy.Label = app.EnergyLabelFootPounds

		output.Momentum.ValueFloat *= app.MassFromKilogramsToPounds * app.VelocityFromMPSToFPS
		output.Momentum.Label = app.MomentumLabelFPS
	}

	if data.mpbr.Value > 0 {
		if outputDebug {
			fmt.Printf("MPBR %f %s\n", data.mpbr.Value, data.mpbr.Label)
		}
		output.Mpbr = mpbrToMPBR(data)
		if outputDebug {
			fmt.Printf("MPBR %f %s\n", output.Mpbr.ValueFloat, output.Mpbr.Label)
		}
	}
}

/** Calculate drop in flight */
func calcDropAtDistance(distance, velocity float64) (drop float64) {
	flightTime := distance / velocity
	drop = app.GravityMPS * 0.5 * (flightTime * flightTime)
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

/** Calculate kinetic energy */
func calcKineticEnergy(data BallisticData) (energy LabeledValue) {
	// 1 Joule == 1 N⋅m (Newton meter)
	massKg := data.projectileMass.Value
	velocityMPS := data.projectileVelocity.Value

	energy.ValueFloat = massKg * velocityMPS * velocityMPS * 0.5
	energy.Label = app.EnergyLabelJoules

	if outputDebug {
		log.Printf("calcKineticEnergy()   <| mass: %f kg", massKg)
		log.Printf("calcKineticEnergy()   <| velocity: %f mps", velocityMPS)
		log.Printf("calcKineticEnergy()    | energy: %f %s", energy.ValueFloat, energy.Label)
	}

	return energy
}

/** Calculate force Newtons */
func calcForce(avgDrawWeight float64) (drawForce parsing.ParsedData) {
	drawForce.Value = avgDrawWeight * app.ForceFromKilogramsToNewtons
	drawForce.Label = app.ForceLabelNewtons

	if outputDebug {
		log.Printf("calcForce()    <| avgDrawWeight: %f kg", avgDrawWeight)
		log.Printf("calcForce()     | drawForce: %f %s", drawForce.Value, drawForce.Label)
	}

	return drawForce
}

/** Calculate momentum */
func calcMomentum(data BallisticData) (momentum LabeledValue) {
	// kg⋅m/s (kilogram meters per second)
	massKg := data.projectileMass.Value
	velocityMPS := data.projectileVelocity.Value

	momentum.ValueFloat = massKg * velocityMPS
	momentum.Label = app.MomentumLabelMKS

	if outputDebug {
		log.Printf("calcMomentum() <| mass: %f kg", massKg)
		log.Printf("calcMomentum() <| velocity: %f mps", velocityMPS)
		log.Printf("calcMomentum()  | momentum: %f %s", momentum.ValueFloat, momentum.Label)
	}

	return momentum
}

/** Calculate Maximum Point Blank Range */
func calcMPBR(data BallisticData) (mpbr parsing.ParsedData) {
	distance := 0.0
	diameter := data.targetRadius.Value * 2
	drop := data.targetRadius.Value
	velocity := data.projectileVelocity.Value

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

	// log.Printf("calcMPBR() <|   target radius: %12.6f m", data.targetRadius.Value)
	// log.Printf("calcMPBR() <| target diameter: %12.6f m", diameter)
	// log.Printf("calcMPBR() <| projectile drop: %12.6f m", drop)
	// log.Printf("calcMPBR()  |            MPBR: %12.6f m", distance)

	return mpbr
}

/**
 * Calculate velocity
 */
func calcVelocity(data BallisticData) (projectileVelocity parsing.ParsedData) {
	projectileMass := data.projectileMass.Value
	drawLength := data.drawLength.Value
	drawForce := data.drawForce.Value

	releaseTime := math.Sqrt(projectileMass * drawLength / drawForce)

	projectileVelocity.Value = drawLength / releaseTime
	projectileVelocity.Label = app.VelocityLabelMPS

	if len(parsing.InputData.Velocity) == 0 {
		if parsing.InputData.Metric {
			parsing.InputData.Velocity = app.VelocityLabelMPS
		} else {
			parsing.InputData.Velocity = app.VelocityLabelFPS
		}
	}

	if outputDebug {
		log.Printf("calcVelocity() <|     projectile mass: %15.6f kg", projectileMass)
		log.Printf("calcVelocity() <|         draw length: %15.6f m", drawLength)
		log.Printf("calcVelocity() <|          draw force: %15.6f N", drawForce)
		log.Printf("calcVelocity()  | projectile velocity: %15.6f mps", projectileVelocity.Value)
	}

	return projectileVelocity
}

/** Calculate the initial velocity of a projectile */
func calcVelocityInitial(data BallisticData) (initialVelocity parsing.ParsedData) {
	radians := data.projectionAngle.Value * app.AngleDegreesToRadians
	sin := math.Sin(2 * radians)
	Rg := data.projectileRange.Value * app.GravityMPS
	initialVelocity.Value = math.Sqrt(Rg / sin)
	initialVelocity.Label = app.VelocityLabelMPS

	if len(parsing.InputData.Velocity) == 0 {
		if parsing.InputData.Metric {
			parsing.InputData.Velocity = app.VelocityLabelMPS
		} else {
			parsing.InputData.Velocity = app.VelocityLabelFPS
		}
	}

	if outputDebug {
		log.Printf("calcVelocityInitial()  | initial velocity: %15.6f mps", initialVelocity.Value)
	}

	return initialVelocity
}

/**
 * Cleans up OutputData for JSON parsing
 *
 * Removes top level keys from OutputData when the child fields have no value.
 * This is necessary because the standard JSON package does not check child
 * structs if they are empty or not.
 */
func cleanupJSON(data OutputData) (dataObj map[string]LabeledValue) {
	// dataObj = make(map[string]interface{})
	dataObj = make(map[string]LabeledValue)

	if data.Energy.ValueFloat != 0 {
		dataObj["energy"] = data.Energy
	}
	if data.Momentum.ValueFloat != 0 {
		dataObj["momentum"] = data.Momentum
	}
	if data.Mpbr.ValueFloat != 0 {
		dataObj["mpbr"] = data.Mpbr
	}
	if data.Velocity.ValueFloat != 0 {
		dataObj["velocity"] = data.Velocity
	}

	return dataObj
}

/** Function defined by a call to github.com/runeimp/locale.NumberFormatter() */
var localeNumberFormatter func(number float64, scale int) string

/** Returns the largest integer in the list of arguments */
func maxInt(nums ...int) (maxInt int) {
	// maxInt = math.MinInt64 // ./ballistic.go:350:10: constant -9223372036854775808 overflows int when compiling for Win32
	maxInt = math.MinInt32

	for _, num := range nums {
		if maxInt < num {
			maxInt = num
		}
	}
	return maxInt
}

/** Convert MPBR in meters to input units */
func mpbrToMPBR(data BallisticData) (mpbr LabeledValue) {
	mpbr.Label = ""
	mpbr.ValueFloat = 0.0

	userLabel := data.projectileVelocity.UserLabel
	if len(userLabel) == 0 {
		userLabel = parsing.InputData.Velocity
	}

	switch userLabel {
	case app.VelocityLabelFPS:
		mpbr.Label = app.LengthLabelFoot
		mpbr.ValueFloat = data.mpbr.Value * app.LengthFromMetersToFeet
	case app.VelocityLabelKMPH:
		mpbr.Label = app.LengthLabelKilometer
		mpbr.ValueFloat = data.mpbr.Value * app.LengthFromMetersToKilometers
	case app.VelocityLabelKnots:
		mpbr.Label = app.LengthLabelNauticalMile
		mpbr.ValueFloat = data.mpbr.Value * app.LengthFromMetersToNauticalMiles
	case app.VelocityLabelMPS:
		mpbr.Label = app.LengthLabelMeter
		mpbr.ValueFloat = data.mpbr.Value
	case app.VelocityLabelMPH:
		mpbr.Label = app.LengthLabelMile
		mpbr.ValueFloat = data.mpbr.Value * app.LengthFromMetersToMiles
	}

	if outputDebug {
		log.Printf("mpbrToMPBR()  <|  projectile velocity: %s", data.projectileVelocity.UserLabel)
		log.Printf("mpbrToMPBR()  <| parsing.InputData velocity: %s", parsing.InputData.Velocity)
		log.Printf("mpbrToMPBR()   |  projectile velocity: %15.6f mps", mpbr.Label, mpbr.Label)
	}

	return mpbr
}

/** Takes a float64 and returns it's formated number value and it's string width */
func numberFormatter(number float64) (value string, width int) {
	value = localeNumberFormatter(number, decimalPlaces)
	width = len(value)

	return value, width
}

/** Print Human Readable Output */
func outputHuman(data OutputData) {
	fmt.Println("")

	var energyValue string
	var energyWidth int
	var momentumValue string
	var momentumWidth int
	var mpbrValue string
	var mpbrWidth int
	var velocityValue string
	var velocityWidth int

	if data.Velocity.ValueFloat > 0 {
		velocityValue, velocityWidth = numberFormatter(data.Velocity.ValueFloat)
	}
	if data.Energy.ValueFloat > 0 {
		energyValue, energyWidth = numberFormatter(data.Energy.ValueFloat)
	}
	if data.Momentum.ValueFloat > 0 {
		momentumValue, momentumWidth = numberFormatter(data.Momentum.ValueFloat)
	}
	if data.Mpbr.ValueFloat > 0 {
		mpbrValue, mpbrWidth = numberFormatter(data.Mpbr.ValueFloat)
	}

	maxWidth := fmt.Sprintf("%d", maxInt(
		velocityWidth,
		energyWidth,
		momentumWidth,
		mpbrWidth,
	))

	if velocityWidth > 0 {
		msgFormat := "  Projectile Velocity: %" + maxWidth + "s %s\n"
		fmt.Printf(msgFormat, velocityValue, data.Velocity.Label)
	}
	if energyWidth > 0 {
		msgFormat := "    Projectile Energy: %" + maxWidth + "s %s\n"
		fmt.Printf(msgFormat, energyValue, data.Energy.Label)
	}
	if momentumWidth > 0 {
		msgFormat := "  Projectile Momentum: %" + maxWidth + "s %s\n"
		fmt.Printf(msgFormat, momentumValue, data.Momentum.Label)
	}
	if mpbrWidth > 0 {
		msgFormat := "Max Point Blank Range: %" + maxWidth + "s %s\n"
		fmt.Printf(msgFormat, mpbrValue, data.Mpbr.Label)
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
func outputJSONPrinter(data OutputData) {
	if outputDebug {
		fmt.Println("JSON data!")
		fmt.Printf("data.Energy: %f %s\n", data.Energy.ValueFloat, data.Energy.Label)
		fmt.Printf("data.Momentum: %f %s\n", data.Momentum.ValueFloat, data.Momentum.Label)
		fmt.Printf("data.Mpbr: %f %s\n", data.Mpbr.ValueFloat, data.Mpbr.Label)
		fmt.Printf("data.Velocity: %f %s\n", data.Velocity.ValueFloat, data.Velocity.Label)
	}
	var err error
	var jsonData []byte

	dataObj := cleanupJSON(data)
	// dataObj := data
	if outputDebug {
		fmt.Println("JSON data!")
		fmt.Printf("dataObj[\"energy\"]: %f %s\n", dataObj["energy"].ValueFloat, dataObj["energy"].Label)
		fmt.Printf("dataObj[\"momentum\"]: %f %s\n", dataObj["momentum"].ValueFloat, dataObj["momentum"].Label)
		fmt.Printf("dataObj[\"mpbr\"]: %f %s\n", dataObj["mpbr"].ValueFloat, dataObj["mpbr"].Label)
		fmt.Printf("dataObj[\"velocity\"]: %f %s\n", dataObj["velocity"].ValueFloat, dataObj["velocity"].Label)
	}

	if outputPretty {
		jsonData, err = json.MarshalIndent(dataObj, "", outputIndent)
	} else {
		jsonData, err = json.Marshal(dataObj)
	}
	// fmt.Printf("jsonData: %#v\n", jsonData)

	if outputDebug {
		log.Printf("err: %v\n", err)
	}

	if err == nil {
		fmt.Println(string(jsonData))
	} else {
		log.Println("JSON encoding error")
		log.Printf("%s\n", err)
		log.Printf("err: %#v\n", err)
	}
}

/** Convert velocity in mps to input units */
func velocityToVelocity(data BallisticData) (velocity LabeledValue) {
	velocity.Label = ""

	userLabel := parsing.InputData.Velocity
	if len(userLabel) == 0 {
		switch parsing.InputData.Mass {
		case app.MassLabelGrains, app.MassLabelLongTon, app.MassLabelPounds, app.MassLabelShortTon, app.MassLabelStone:
			userLabel = app.VelocityLabelFPS
		default:
			userLabel = app.VelocityLabelMPS
		}
	}

	switch userLabel {
	case app.VelocityLabelFPS:
		velocity.Label = userLabel
		velocity.ValueFloat = data.projectileVelocity.Value * app.VelocityFromMPSToFPS
	case app.VelocityLabelKMPH:
		velocity.Label = userLabel
		velocity.ValueFloat = data.projectileVelocity.Value * app.VelocityFromMPSToKMPH
	case app.VelocityLabelKnots:
		velocity.Label = userLabel
		velocity.ValueFloat = data.projectileVelocity.Value * app.VelocityFromMPSToKnots
	case app.VelocityLabelMPS:
		velocity.Label = userLabel
		velocity.ValueFloat = data.projectileVelocity.Value
	case app.VelocityLabelMPH:
		velocity.Label = userLabel
		velocity.ValueFloat = data.projectileVelocity.Value * app.VelocityFromMPSToMPH
	}

	// log.Printf("velocityToVelocity() | userLabel: '%s'", userLabel)
	// log.Printf("velocityToVelocity() | projectileMass user value & label: %f %s", data.projectileMass.UserValue, data.projectileMass.UserLabel)
	// log.Printf("velocityToVelocity() | parsing.InputData | velocity: '%s' | mass: '%s' | metric: %t", parsing.InputData.Velocity, parsing.InputData.Mass, parsing.InputData.Metric)
	// log.Printf("velocityToVelocity() | projectileVelocity user value & label: %f %s", data.projectileVelocity.UserValue, data.projectileVelocity.UserLabel)
	// log.Printf("velocityToVelocity() | projectileVelocity norm value & label: %f %s", data.projectileVelocity.Value, data.projectileVelocity.Label)
	// log.Printf("velocityToVelocity() | projectileVelocity calc value & label: %f %s", velocity.ValueFloat, velocity.Label)

	return velocity
}

//
// MAIN ENTRYPOINT
//
func main() {
	// log.Print("Ballistic ...")

	cliApp := cli.NewApp()
	cliApp.Name = "Ballistic"
	cliApp.Usage = "Calculates what it can based on provided input."
	cliApp.Version = AppVersion

	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "projection-angle, angle, a",
			Usage: "The projection angle or trajectory of projectile",
		},
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Output debug info",
		},
		cli.StringFlag{
			Name:  "projectile-range, distance, d",
			Usage: "The distance the projectile traveled",
		},
		cli.StringFlag{
			Name:  "draw-weight, weight, w",
			Usage: "Bow or sling shot draw `WEIGHT`. Used to calculate projectile velocity, energy, etc.",
		},
		cli.StringFlag{
			Name:  "draw-length, length, l",
			Usage: "Bow or sling shot draw `LENGTH`. Used to calculate projectile velocity, energy, etc.",
		},
		// cli.BoolFlag{
		// 	Name: "help, h",
		// 	Usage: "Output this help info",
		// },
		cli.BoolFlag{
			Name:  "json, j",
			Usage: "Output JSON data",
		},
		cli.StringFlag{
			Name:   "locale, local",
			Value:  "en_US",
			Usage:  "The `LOCALE` to format number output for.",
			EnvVar: "LC_CTYPE,LANG",
		},
		cli.StringFlag{
			Name:  "projectile, mass, m",
			Usage: "Projectile `MASS` (weight). Used to calculate projectile velocity, energy, etc.",
		},
		cli.StringFlag{
			Name:  "precision, float, f",
			Value: "6",
			Usage: "The output floating point `PRECISION` (numbers after decimal mark).",
		},
		cli.BoolFlag{
			Name:  "pretty-print, pretty, p",
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
			Name:  "radius, r",
			Value: "225mm",
			Usage: "The `RADIUS` of the target area. Used to calculate MPBR (Maximum Point Blank Range).",
		},
		cli.StringFlag{
			Name:  "velocity, v",
			Usage: "The projectile `VELOCITY` (speed). Used to calculate projectile energy, momentum, etc.",
		},
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Output this help info",
	}

	cli.AppHelpTemplate = fmt.Sprintf(app.HelpTemplate)

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, V",
		Usage: "Output the ballistic app version",
	}

	sort.Sort(cli.FlagsByName(cliApp.Flags))
	// sort.Sort(cli.CommandsByName(app.Commands))

	cliApp.Action = func(c *cli.Context) error {
		outputDebug = c.Bool("debug")
		outputJSON = c.Bool("json")
		outputPretty = c.Bool("pretty-print")
		decimalPlaces = c.Int("precision")
		localeStr = c.String("locale")

		// outputPretty = c.Bool("pretty")
		// if len(c.String("pretty-print")) > 0 {
		// 	log.Println("pretty-print!!!")
		// 	outputIndent = c.String("pretty-print")
		// }

		if outputPretty && !outputJSON {
			outputJSON = true
		}

		if outputDebug {
			fmt.Println("Going Ballistic!")
			log.Printf("             locale: %12s (%d)", localeStr, len(localeStr))
			log.Printf("        draw length: %12s (%d)", c.String("draw-length"), len(c.String("draw-length")))
			log.Printf("        draw weight: %12s (%d)", c.String("draw-weight"), len(c.String("draw-weight")))
			log.Printf("      target radius: %12s (%d)", c.String("radius"), len(c.String("radius")))
			log.Printf("projectile velocity: %12s (%d)", c.String("velocity"), len(c.String("velocity")))
			log.Printf("    projectile mass: %12s (%d)", c.String("projectile"), len(c.String("projectile")))
			log.Printf("          precision: %12s (%d)", c.String("precision"), len(c.String("precision")))
			// log.Println("")
			// fmt.Println("")
		}

		flagsSet := 0
		for _, flagName := range c.GlobalFlagNames() {
			// fmt.Printf("Flag: %s\n", flagName)
			switch flagName {
			case "locale", "precision", "radius":
			default:
				flagValue := c.String(flagName)
				if len(flagValue) > 0 {
					if flagValue != "false" && flagValue != "true" {
						// fmt.Printf("Flag Set: %s\n", flagValue)
						flagsSet++
					}
				}
			}
		}
		// fmt.Printf("Flags Set: %d\n", flagsSet)

		/** If no flags set default to displaying help */
		if flagsSet == 0 {
			cli.ShowAppHelpAndExit(c, 0)
		}

		if len(c.String("draw-length")) > 0 {
			data.drawLength = parsing.ParseValue(c.String("draw-length"), app.ValueTypeLength)
		}
		if len(c.String("draw-weight")) > 0 {
			data.drawWeight = parsing.ParseValue(c.String("draw-weight"), app.ValueTypeMass)
			avgDrawWeight := data.drawWeight.Value * 0.5
			data.drawForce = calcForce(avgDrawWeight)
		}
		if len(c.String("velocity")) > 0 {
			data.projectileVelocity = parsing.ParseValue(c.String("velocity"), app.ValueTypeVelocity)
		}
		if len(c.String("mass")) > 0 {
			data.projectileMass = parsing.ParseValue(c.String("mass"), app.ValueTypeMass)
		}
		if len(c.String("projectile-range")) > 0 {
			data.projectileRange = parsing.ParseValue(c.String("projectile-range"), app.ValueTypeLength)
		}
		if len(c.String("projection-angle")) > 0 {
			data.projectionAngle = parsing.ParseValue(c.String("projection-angle"), app.ValueTypeMass)
		}

		if data.projectileVelocity.Value == 0 {
			if data.projectileMass.Value > 0 && data.drawLength.Value > 0 && data.drawForce.Value > 0 {
				data.projectileVelocity = calcVelocity(data)
			} else if data.projectileRange.Value > 0 && data.projectionAngle.Value > 0 {
				data.projectileVelocity = calcVelocityInitial(data)
			}
		}

		data.targetRadius = parsing.ParseValue(c.String("radius"), app.ValueTypeLength)

		if data.projectileVelocity.Value > 0 {
			data.mpbr = calcMPBR(data)
		}

		buildOutputData(data)

		localeNumberFormatter = locale.NumberFormatter(localeStr)
		// localeNumberFormatter = locale.NumberFormatter("TESTONE")
		// localeNumberFormatter(123456789.1234567)

		if outputJSON {
			outputJSONPrinter(output)
		} else {
			outputHuman(output)
		}

		return nil
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	ballisticWeight, bwPresent := os.LookupEnv("BALLISTIC_WEIGHT")

	if bwPresent {
		fmt.Printf("BALLISTIC_WEIGHT: %s", ballisticWeight)
	}

	// x := fmt.Sprintf("Hello, %s!", )
	// x := fmt.Sprintf("Hello, %s!", os.Args[1])
}
