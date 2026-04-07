package main

import (
	"strconv"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

func MapLdapEntryToUser(ldapEntry *ldap.Entry) *User {
	if ldapEntry == nil {
		return nil
	}
	user := &User{DN: ldapEntry.DN}
	fullName := ""
	displayName := ""

	for _, v := range ldapEntry.Attributes {
		if len(v.Values) == 0 {
			continue
		}
		switch strings.ToLower(v.Name) {
		case "uid":
			user.UID = v.Values[0]
		case "givenname":
			user.GivenName = v.Values[0]
		case "sn":
			user.FamilyName = v.Values[0]
		case "kxi":
			user.ExternalID = v.Values[0]
		case "mail":
			user.Email = v.Values[0]
		case "kentloc":
			user.Location = v.Values[0]
		case "kdivno":
			user.Division = v.Values[0]
		case "logindisabled":
			user.LoginDisabled, _ = strconv.ParseBool(v.Values[0])
		case "initials":
			user.MiddleInitial = v.Values[0]
		case "fullname":
			fullName = v.Values[0]
		case "displayname":
			displayName = v.Values[0]
		case "manager":
			user.ManagerDN = v.Values[0]
		case "title":
			user.Title = v.Values[0]
		}
	}

	if fullName != "" {
		user.FullName = fullName
	} else if displayName != "" {
		user.FullName = displayName
	}
	return user
}
