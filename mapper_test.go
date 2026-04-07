package main

import (
	"testing"

	"github.com/go-ldap/ldap/v3"
)

func TestMapLdapEntryToUser_Cases(t *testing.T) {

	cases := []struct {
		name     string
		entry    *ldap.Entry
		expected *User
	}{
		{
			name: "LDAP Attributes",
			entry: &ldap.Entry{
				DN: "cn=John Doe,ou=People,o=Company",
				Attributes: []*ldap.EntryAttribute{
					{Name: "uid", Values: []string{"jdoe"}},
					{Name: "givenName", Values: []string{"John"}},
					{Name: "sn", Values: []string{"Doe"}},
					{Name: "KXI", Values: []string{"12345"}},
					{Name: "mail", Values: []string{"jdoe@example.com"}},
					{Name: "KEntLoc", Values: []string{"LOC1"}},
					{Name: "KDivNo", Values: []string{"DIV1"}},
					{Name: "loginDisabled", Values: []string{"true"}},
					{Name: "initials", Values: []string{"M"}},
					{Name: "fullName", Values: []string{"DOE, JOHN"}},
				},
			},
			expected: &User{
				DN:            "cn=John Doe,ou=People,o=Company",
				UID:           "jdoe",
				GivenName:     "John",
				FamilyName:    "Doe",
				ExternalID:    "12345",
				Email:         "jdoe@example.com",
				Location:      "LOC1",
				Division:      "DIV1",
				LoginDisabled: true,
				MiddleInitial: "M",
				FullName:      "DOE, JOHN",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := MapLdapEntryToUser(tc.entry)

			if tc.expected == nil && got != nil {
				t.Fatalf("expected nil user, got %+v", got)
			}
			if tc.expected == nil && got == nil {
				return
			}
			if tc.expected != nil && got == nil {
				t.Fatalf("expected user %+v, got nil", tc.expected)
			}

			if got.DN != tc.expected.DN {
				t.Errorf("DN: expected %q got %q", tc.expected.DN, got.DN)
			}
			if got.UID != tc.expected.UID {
				t.Errorf("UID: expected %q got %q", tc.expected.UID, got.UID)
			}
			if got.GivenName != tc.expected.GivenName {
				t.Errorf("GivenName: expected %q got %q", tc.expected.GivenName, got.GivenName)
			}
			if got.FamilyName != tc.expected.FamilyName {
				t.Errorf("FamilyName: expected %q got %q", tc.expected.FamilyName, got.FamilyName)
			}
			if got.ExternalID != tc.expected.ExternalID {
				t.Errorf("ExternalID: expected %q got %q", tc.expected.ExternalID, got.ExternalID)
			}
			if got.Email != tc.expected.Email {
				t.Errorf("Email: expected %q got %q", tc.expected.Email, got.Email)
			}
			if got.Location != tc.expected.Location {
				t.Errorf("Location: expected %q got %q", tc.expected.Location, got.Location)
			}
			if got.Division != tc.expected.Division {
				t.Errorf("Division: expected %q got %q", tc.expected.Division, got.Division)
			}
			if got.MiddleInitial != tc.expected.MiddleInitial {
				t.Errorf("MiddleInitial: expected %q got %q", tc.expected.MiddleInitial, got.MiddleInitial)
			}
			if got.FullName != tc.expected.FullName {
				t.Errorf("FullName: expected %q got %q", tc.expected.FullName, got.FullName)
			}
			if got.LoginDisabled != tc.expected.LoginDisabled {
				t.Errorf("LoginDisabled: expected %v got %v", tc.expected.LoginDisabled, got.LoginDisabled)
			}
		})
	}
}
