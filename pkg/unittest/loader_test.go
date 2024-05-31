package unittest_test

import (
	"reflect"
	"testing"

	"github.com/devbay-io/unittest6502/pkg/unittest"
)

func TestUnitTestFromString(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput unittest.UnitTest
		expectError    bool
	}{
		{
			name: "Valid Input",
			input: `; > platform: simple
; > program_counter: 0x8000
; > assert:
; >   - type: register
; >     location: x
; >     condition:
; >       name: equal
; >       value: value`,
			expectedOutput: unittest.UnitTest{
				Platform:       "simple",
				ProgramCounter: 0x8000,
				Assert: []unittest.UnitTestAssertion{
					{
						Type:     "register",
						Location: "x",
						Condition: unittest.UnitTestAssertionCondition{
							Name:  "equal",
							Value: "value",
						},
					},
				},
			},
			expectError: false,
		},
		{
			name: "Invalid YAML",
			input: `; > platform: simple
; > program_counter: not_a_number`,
			expectedOutput: unittest.UnitTest{},
			expectError:    true,
		},
		{
			name: "Missing assert",
			input: `; > platform: simple
; > program_counter: 0x8000`,
			expectedOutput: unittest.UnitTest{},
			expectError:    true,
		},
		{
			name: "No test info",
			input: `; no yaml info should fail
start:
end:`,
			expectedOutput: unittest.UnitTest{},
			expectError:    true,
		},
		{
			name: "Wrong platform",
			input: `; > platform: wrong
; > program_counter: 0x8000
; > assert:
; >   - type: register
; >     location: x
; >     condition:
; >	      name: equal
; >	      value: 200
start:
end:`,
			expectedOutput: unittest.UnitTest{},
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := unittest.UnitTestFromString(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("UnitTestFromString() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("UnitTestFromString() = %v, expected %v", output, tt.expectedOutput)
			}
		})
	}
}
