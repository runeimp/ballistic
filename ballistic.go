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
	// "menteslibres.net/gosexy/to"
	// "menteslibres.net/gosexy/yaml"
	"os"
	// "os/signal"
	// "strconv"
	// "strings"
	// "syscall"
)


//
// CONSTANTS
//
const GRAVITY_MPS float64 = 9.80665


//
// FUNCTIONS
//
func calc_drop(distance, velocity float64) (drop float64) {
	flight_time := distance / velocity
	drop = GRAVITY_MPS * 0.5 * (flight_time * flight_time)
	return drop
}


func mpbr(radius, velocity float64) (distance float64) {
	distance = 0.0
	// var diameter, distance, drop float64
	// diameter = radius * 2
	// distance = 0.0
	// drop = radius
	diameter := radius * 2
	drop := radius

	for drop < diameter || drop == 0.0 {
		distance += 1.0
		drop = calc_drop(distance, velocity)
	}

	return distance
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
			Name: "weight, w",
			Usage: "Projectile `WEIGHT`",
		},
		// cli.StringFlag{
		// 	Name: "velocity, v",
		// 	Usage: "Projectile `SPEED`",
		// },
		cli.StringFlag{
			Name: "draw, d",
			Usage: "Bow or sling shot draw `WEIGHT`",
		},
		// cli.StringFlag{
		// 	Name: "radius, r",
		// 	Usage: "The `RADIUS` of the target area",
		// },
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
			Usage: "Calculates the maximum point blank range",
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
				// radius := strconv.ParseFloat(c.String('radius'), 64)
				// velocity := strconv.ParseFloat(c.String('velocity'), 64)
				mpbr(c.Float64("radius"), c.Float64("velocity"))
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Default Action!")
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


