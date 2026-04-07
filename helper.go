package main

import (
	"fmt"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

type CallbackFunc func(*User)

func (u User) String() string {
	return fmt.Sprintf("%v (%v)", strings.ToUpper(u.FullName), strings.ToLower(u.Email))
}

func (u User) PrintUser() {
	fmt.Println("Associate")
	fmt.Println(strings.Repeat("-", 32))
	fmt.Printf("Name: %v (%v)\n", strings.ToUpper(u.FullName), strings.ToLower(u.UID))
	fmt.Printf("Title: %v\n", FormatTitle(titleCase.String(strings.ToLower(u.Title))))
	fmt.Printf("Mail: %v\n", strings.ToLower(u.Email))
	if u.Manager != nil {
		fmt.Printf("Manager: %v\n", strings.ToLower(u.Manager.UID))
	}
	fmt.Println()
}

func CommonNameFromDistinguishedName(dn string) string {
	matches := regexCommonName.FindStringSubmatch(dn)
	if len(matches) >= 2 {
		return matches[1]
	} else {
		return dn
	}
}

func FormatTitle(title string) string {
	replacementStrings := map[string]string{
		"Cio":  "CIO",
		"Ciso": "CISO",
		"Gvp":  "GVP",
		"Ktd":  "KTD",
		"Svp":  "SVP",
		"Vp":   "VP",
	}
	for key, value := range replacementStrings {
		title = strings.Replace(title, key, value, 1)
	}
	return title
}

func findManagerChain(l *ldap.Conn, cfg Config, filterKey string, filterValue string, callback CallbackFunc) {
	searchRequest := ldap.NewSearchRequest(
		cfg.SearchBase,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf(cfg.SearchFilter, filterKey, filterValue),
		cfg.Attributes,
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil || len(sr.Entries) == 0 {
		return
	}

	entry := sr.Entries[0]
	user := MapLdapEntryToUser(entry)

	callback(user)

	manager := entry.GetEqualFoldAttributeValue("manager")

	if manager != "" {
		findManagerChain(l, cfg, "cn", CommonNameFromDistinguishedName(manager), callback)
	}
}
