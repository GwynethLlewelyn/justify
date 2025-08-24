package justify

import (
	"testing"
)

const loremipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

func TestInvalidInput_ZeroWidth(t *testing.T) {
	got := Justify(loremipsum, 0)
	if got != "" {
		t.Errorf("expected empty string with line width of zero, got: %q", got)
	}
}

func TestInvalidInput_EmptyText(t *testing.T) {
	got := Justify("", 160)
	if got != "" {
		t.Errorf("expected empty string with empty input, got: %q", got)
	}
}

func TestBasicJustify(t *testing.T) {
	testCases := []struct {
		testName  string
		formatted string
		width     int
	}{
		{"Line width 40",
			`Lorem  ipsum dolor sit amet, consectetur
adipiscing  elit,  sed do eiusmod tempor
incididunt  ut  labore  et  dolore magna
aliqua.  Ut  enim  ad minim veniam, quis
nostrud   exercitation  ullamco  laboris
nisi ut aliquip ex ea commodo consequat.
Duis  aute  irure dolor in reprehenderit
in voluptate velit esse cillum dolore eu
fugiat  nulla  pariatur.  Excepteur sint
occaecat cupidatat non proident, sunt in
culpa  qui  officia deserunt mollit anim
id est laborum.
`, 40},
		{"Line width 7",
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
laborum.
`, 7},
		{"Line width 80",
			`Lorem  ipsum  dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt  ut  labore  et  dolore  magna  aliqua. Ut enim ad minim veniam, quis
nostrud  exercitation  ullamco  laboris nisi ut aliquip ex ea commodo consequat.
Duis  aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu
fugiat  nulla  pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.
`, 80},
		{"single line with huge line size", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.\n", 1000000007},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			expected := tc.formatted
			got := Justify(loremipsum, tc.width)

			if got != expected {
				t.Errorf("mismatch for line=%d, got:\n------\n%s\n------\nexpected:\n------\n%s\n------\n", tc.width, got, expected)
			}
		})
	}
}
