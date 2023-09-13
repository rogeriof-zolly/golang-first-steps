package addresses_test

import (
	. "addresses/addresses"
	"testing"
)

type testScenario struct {
	addressInput   string
	expectedOutput string
}

func TestValidateAddressType(t *testing.T) {
	// Run tests that have this Parallel function call in parallel
	t.Parallel()

	// Calling in terminal 'go test --cover', the language will output
	// a percentage of coverage of the test

	// You can generate a report using 'go test --coverprofile reportfilename.txt
	// To read the file, use go tool cover --func=reportfilename.txt
	// To generate a html file showing what case you are not covering
	// use go tool cover --html=reportfilename.txt
	testScenarios := []testScenario{
		{"Independence Avenue", "avenue"},
		{"Mable Street", "street"},
		{"ABC Highway", "highway"},
		//{"President Square", "Invalid address!"},
		//{"", "Invalid address!"},
	}

	for _, scenario := range testScenarios {
		receivedAddressType := ValidateAddressType(scenario.addressInput)
		if scenario.expectedOutput != receivedAddressType {
			t.Errorf("Received address type: %s is different from expected: %s", receivedAddressType, scenario.expectedOutput)
		}
	}
}

func TestMocked(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Test failed")
	}
}
