package main

import (
	"regexp"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Config struct {
	ServerHost   string   `envconfig:"LDAP_SERVER_URI" required:"true"`
	ServerPort   int      `envconfig:"LDAP_SERVER_PORT" default:"389"`
	SearchBase   string   `envconfig:"LDAP_SEARCH_BASE" required:"true"`
	SearchFilter string   `default:"(&(%s=%s))"`
	BindUserDN   string   `envconfig:"LDAP_USER_ACCOUNT" required:"true"`
	BindPassword string   `envconfig:"LDAP_USER_PASSWORD" required:"true"`
	Attributes   []string `default:"isMemberOf,*"`
}

type User struct {
	Config        string
	DN            string
	UID           string
	GivenName     string
	FamilyName    string
	MiddleInitial string
	FullName      string
	Division      string
	Location      string
	Email         string
	ExternalID    string
	LoginDisabled bool
	ManagerDN     string
	Manager       *User
	Title         string
}

var patternCommonName = "(?i)cn=([A-Za-z.][A-Za-z.][A-Za-z0-9.][0-9][0-9][0-9][0-9]),[A-Za-z0-9]+"
var patternEuid = "^[A-Za-z.][A-Za-z.][A-Za-z0-9.][0-9][0-9][0-9][0-9]$"
var patternMail = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,4}$"

var regexCommonName = regexp.MustCompile(patternCommonName)
var regexEuid = regexp.MustCompile(patternEuid)
var regexMail = regexp.MustCompile(patternMail)

var titleCase = cases.Title(language.English, cases.NoLower)
