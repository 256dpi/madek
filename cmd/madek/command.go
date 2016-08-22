package main

import "github.com/docopt/docopt-go"

type command struct {
	// commands
	cSet bool

	// arguments
	aID string

	// options
	oAddress  string
	oUsername string
	oPassword string
}

func parseCommand() *command {
	usage := `madek.

Usage:
  madek set <id> [options]

Options:
  -h --help                   Show this screen.
  -a --address=<url>          The address of the madek API [default: https://medienarchiv.zhdk.ch/api].
  -u --username=<username>    The username used for authentication.
  -p --password=<password>    The password used for authentication.
`

	a, _ := docopt.Parse(usage, nil, true, "0.1", false)

	return &command{
		// commands
		cSet: getBool(a["set"]),

		// arguments
		aID: getString(a["<id>"]),

		// options
		oAddress:  getString(a["--address"]),
		oUsername: getString(a["--username"]),
		oPassword: getString(a["--password"]),
	}
}

func getBool(field interface{}) bool {
	if val, ok := field.(bool); ok {
		return val
	}

	return false
}

func getString(field interface{}) string {
	if str, ok := field.(string); ok {
		return str
	}

	return ""
}
