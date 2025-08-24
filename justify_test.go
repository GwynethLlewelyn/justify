package justify

import (
	"testing"
)

const loremipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func TestBasicJustify(t *testing.T) {
	testCases := []struct {
		testName  string
		formatted string
		length    int
	}{
		{"empty", "", 0},
		{"empty but with long line length", "", 160},
		{"zero line length", loremipsum, 0},
		{"Line length 40",
			`Lorem  ipsum dolor sit amet, consectetur
adipiscing elit, sed   do eiusmod tempor
incididunt  ut labore  et  dolore  magna
aliqua. Ut enim  ad minim   veniam, quis
nostrud    exercitation  ullamco laboris
nisi ut aliquip ex ea commodo consequat.
Duis aute irure   dolor in reprehenderit
in voluptate velit esse cillum dolore eu
fugiat nulla pariatur.   Excepteur  sint
occaecat cupidatat non proident, sunt in
culpa  qui officia deserunt  mollit anim
id est laborum.`, 40},
		{"Line length 7",
			`Lorem
ipsum
dolor
sit
amet,
consectetur
adipiscing
elit,
sed  do
eiusmod
tempor
incididunt
ut
labore
et
dolore
magna
aliqua.
Ut enim
ad
minim
veniam,
quis
nostrud
exercitation
ullamco
laboris
nisi ut
aliquip
ex   ea
commodo
consequat.
Duis
aute
irure
dolor
in
reprehenderit
in
voluptate
velit
esse
cillum
dolore
eu
fugiat
nulla
pariatur.
Excepteur
sint
occaecat
cupidatat
non
proident,
sunt in
culpa
qui
officia
deserunt
mollit
anim id
est
laborum.`, 7},
		{"Line length 80",
			`Lorem ipsum  dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna  aliqua.   Ut  enim ad minim  veniam, quis
nostrud  exercitation ullamco laboris nisi ut aliquip ex  ea commodo consequat.
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu
fugiat nulla pariatur. Excepteur  sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.`, 80},
		{"single line with huge line size", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", 10 ^ 9 + 7},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			expected := tc.formatted
			got := Justify(loremipsum, tc.length)

			if got != expected {
				t.Errorf("mismatch for line=%d, got:\n%s\nexpected:\n%s\n", tc.length, got, expected)
			}
		})
	}
}
