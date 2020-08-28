package cli

import (
	"log"
	"os"
	"strconv"
)

const LIMIT = 25

// Command takes a name and datatype for the command
type Command struct {
	Name  string
	Value string
}

// Our interface for Cli...we only need a getter for the data
type ICli interface {
	GetArguments() (subreddit string, limit int)
}

// Our concrete implementation of ICli - merely a container for data (no dependencies)
type Cli struct {
	Subreddit string
	Limit     int
}

// The Cli's getter for it's stored data
func (c Cli) GetArguments() (string, int) {
	return c.Subreddit, c.Limit
}

// Our provider for Cli
// For this one, we are using the init function to parse and store the cli arguments
func InitCli() ICli {
	values := parseArgs(os.Args)

	limit := LIMIT

	if len(values) > 1 {
		l, _ := strconv.ParseInt(values[1].Value, 10, 32)

		if l < 0 || l > 100 {
			log.Println("Limit cannot be less than 0 or greater than 100.Defaulting to 25")
			limit = LIMIT
		} else {
			limit = int(l)
		}
	}

	subr := values[0].Value

	return Cli{Subreddit: subr, Limit: limit}
}

/***************************
     Private Functions
***************************/

// ParseArgs takes in os.Args and parses the system args
func parseArgs(args []string) []Command {
	var commands []Command
	requiredArgs := args[1:]

	if !(len(requiredArgs) > 0) {
		log.Fatal("No arguments were provided")
	}

	for i, v := range requiredArgs {
		if i%2 == 1 {
			continue
		}

		if i == 0 && v != "--subr" {
			log.Fatal("First argument should be --subr _subredditname_")
		}

		commands = append(commands, Command{
			Name:  v,
			Value: requiredArgs[i+1],
		})
	}

	return commands
}
