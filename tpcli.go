package main

import (
	"fmt"
	"os"

	"time"

	"github.com/bndr/gopencils"
	"github.com/briandowns/spinner"
	"github.com/codegangsta/cli"
	"github.com/kennygrant/sanitize"
	"github.com/mgutz/ansi"
)

// Struct to hold target proces entity data
type entityStruct struct {
	Id           int
	Name         string
	Description  string
	StartDate    string
	EndDate      string
	CreateDate   string
	ResourceType string
	Effort       float32
	Iteration    struct {
		Id   int
		Name string
	}
	Release struct {
		Id   int
		Name string
	}
	Team struct {
		Id   int
		Name string
	}
	EntityState struct {
		Id   int
		Name string
	}
}

// Ansi console decorators
var errorDecorator = ansi.ColorFunc("white:red")
var worningDecorator = ansi.ColorFunc("black:yellow")
var boldDecorator = ansi.ColorFunc("white+b")
var preloader = spinner.New(spinner.CharSets[9], 100*time.Millisecond)

// Get entity data from target process RESTful service
func getEntity(id string, username string, password string, url string, entityType string, resp *entityStruct) (*entityStruct, error) {
	auth := gopencils.BasicAuth{Username: username, Password: password}
	api := gopencils.Api(url, &auth)
	querystring := map[string]string{"format": "json"}
	preloader.Prefix = "Fetching data:"
	preloader.Start()
	_, err := api.Res(entityType, resp).Id(id).Get(querystring)
	preloader.Stop()
	return resp, err
}

// Check if username and password flags are provided or env variables exist
func checkGlobalFlags(c *cli.Context) {
	if c.GlobalString("username") == "" || c.GlobalString("password") == "" || c.GlobalString("url") == "" {
		fmt.Println(errorDecorator("Please provide username, password and url flags!"))
		os.Exit(0)
	}
}

func displayEntity(entity *entityStruct, err error, template string) {

	// If there is error in response display error
	if err != nil {
		fmt.Println(errorDecorator(err.Error()))
		os.Exit(0)
	}

	// If there is no entity data display No entity message
	if entity.Id == 0 {
		fmt.Println(worningDecorator("Entity not found!"))
		os.Exit(0)
	}

	// Display appropriate entity template
	switch template {
	case "s":
		fmt.Println(boldDecorator(entity.Name))
		fmt.Println(entity.ResourceType, " | ", entity.Effort, "p", " | ", entity.EntityState.Name)
	case "m":
		{
			fmt.Println(boldDecorator(entity.Name))
			fmt.Println(entity.ResourceType, " | ", entity.Effort, "p", " | ", entity.EntityState.Name)
		}

	case "l":
		{
			fmt.Println(boldDecorator(entity.Name), " | ", entity.Iteration.Name)
			fmt.Println(entity.ResourceType, " | ", entity.Effort, "p", " | ", entity.EntityState.Name)
			fmt.Println(sanitize.HTML(entity.Description))
		}

	}
}

func main() {

	// Define cli app
	app := cli.NewApp()
	app.Name = "Target Process cli tool"
	app.Author = "Ivan Padavic @ipadavic"
	app.Version = "0.1.0"

	// Define app flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "username, u",
			Usage:  "Username for target process Basic Auth",
			EnvVar: "TPCLI_USERNAME",
		},
		cli.StringFlag{
			Name:   "password, p",
			Usage:  "Password for target process Basic Auth",
			EnvVar: "TPCLI_PASSWORD",
		},
		cli.StringFlag{
			Name:   "url",
			Usage:  "Base url for your target process app rest api (custom or tpondemand)",
			EnvVar: "TPCLI_URL",
		},
	}

	// Define response holder
	response := &entityStruct{}

	commandsTemplateFlag := []cli.Flag{
		cli.StringFlag{
			Name:  "template, t",
			Usage: "Template for displaying entity data [s, m, l]",
			Value: "s",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "bug",
			Aliases: []string{"b"},
			Usage:   "Get bug information",
			Flags:   commandsTemplateFlag,
			Action: func(c *cli.Context) {
				checkGlobalFlags(c)
				resp, err := getEntity(
					c.Args().First(),
					c.GlobalString("username"),
					c.GlobalString("password"),
					c.GlobalString("url"),
					"Bugs",
					response)
				displayEntity(resp, err, c.String("template"))
			},
		},
		{
			Name:    "story",
			Aliases: []string{"s"},
			Usage:   "Get user story information",
			Flags:   commandsTemplateFlag,
			Action: func(c *cli.Context) {
				checkGlobalFlags(c)
				resp, err := getEntity(
					c.Args().First(),
					c.GlobalString("username"),
					c.GlobalString("password"),
					c.GlobalString("url"),
					"UserStories",
					response)
				displayEntity(resp, err, c.String("template"))
			},
		},
		{
			Name:    "task",
			Aliases: []string{"t"},
			Usage:   "Get task information",
			Flags:   commandsTemplateFlag,
			Action: func(c *cli.Context) {
				checkGlobalFlags(c)
				resp, err := getEntity(
					c.Args().First(),
					c.GlobalString("username"),
					c.GlobalString("password"),
					c.GlobalString("url"),
					"Tasks",
					response)
				displayEntity(resp, err, c.String("template"))
			},
		},
	}

	app.Run(os.Args)
}
