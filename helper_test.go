package main

import "testing"

func TestCommonNameFromDistinguishedName_Cases(t *testing.T) {

	cases := []struct {
		name     string
		entry    string
		expected string
	}{
		{
			name:     "Valid Common Name",
			entry:    "cn=az12345,ou=people,o=company",
			expected: "az12345",
		},
		{
			name:     "Expanded Common Name",
			entry:    "cn=az12345,ou=organization,ou=people,o=company",
			expected: "az12345",
		},
		{
			name:     "Invalid Common Name",
			entry:    "dn=az12345,ou=organization,ou=people,o=company",
			expected: "dn=az12345,ou=organization,ou=people,o=company",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CommonNameFromDistinguishedName(tc.entry)

			if got != tc.expected {
				t.Errorf("DN: expected %q got %q", tc.expected, got)
			}
		})
	}
}
