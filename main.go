package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/kelseyhightower/envconfig"
)

var matchEuid bool
var matchMail bool

var users []*User

func processUsers(u *User) {
	if len(users) > 0 {
		users[0].Manager = u
	}
	users = append([]*User{u}, users...)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("*** Error: Please provide a valid Associate EUID")
		os.Exit(1)
	}

	args := os.Args[1:]
	userArg := args[0]

	matchEuid = regexEuid.MatchString(userArg)
	matchMail = regexMail.MatchString(userArg)

	var filterKey string

	if !matchEuid && !matchMail {
		fmt.Println("*** Error: Please provide a valid EUID or Email")
		os.Exit(1)
	} else if matchEuid {
		filterKey = "cn"
	} else if matchMail {
		filterKey = "mail"
	}

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	l, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%d", cfg.ServerHost, cfg.ServerPort))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind(cfg.BindUserDN, cfg.BindPassword)
	if err != nil {
		log.Fatal(err)
	}

	findManagerChain(l, cfg, filterKey, userArg, processUsers)

	if len(users) != 0 {
		users[len(users)-1].PrintUser()

		for index, user := range users {
			var prefix string
			switch index {
			case 0:
				prefix = ""
			default:
				prefix = strings.Repeat("  ", index) + "⌙"

			}
			fmt.Printf("%v%+v\n", prefix, user)
		}
	} else {
		fmt.Println("*** Error: Associate not found")
		os.Exit(1)
	}
}
