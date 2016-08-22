package main

import "github.com/docopt/docopt-go"

type command struct {
	// commands
	cFetch  bool
	cServer bool

	// arguments
	aID string

	// options
	oAddress  string
	oUsername string
	oPassword string
	oCache    bool
}

func parseCommand() *command {
	usage := `madek.

Usage:
  madek fetch <id> [options]
  madek server [options]

Options:
  -h --help                   Show this screen.
  -a --address=<url>          The address of the madek instance [default: https://medienarchiv.zhdk.ch].
  -u --username=<username>    The username used for authentication.
  -p --password=<password>    The password used for authentication.
  -c --cache                  Cache requests in server mode.
`

	a, _ := docopt.Parse(usage, nil, true, "0.1", false)

	return &command{
		// commands
		cFetch:  getBool(a["fetch"]),
		cServer: getBool(a["server"]),

		// arguments
		aID: getString(a["<id>"]),

		// options
		oAddress:  getString(a["--address"]),
		oUsername: getString(a["--username"]),
		oPassword: getString(a["--password"]),
		oCache:    getBool(a["--cache"]),
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
