package container

import "testing"

func Test_IsValid(t *testing.T) {
	tests := []struct {
		title          string
		containerSetup string
		expectedResult bool
	}{
		{
			title:          "Testing valid setup {[()]}",
			containerSetup: "{[()]}",
			expectedResult: true,
		},
		{
			title:          "Testing valid setup ([]{}())",
			containerSetup: "([]{}())",
			expectedResult: true,
		},
		{
			title:          "Testing valid setup (){}()",
			containerSetup: "(){}()",
			expectedResult: true,
		},
		{
			title:          "Testing valid setup ([)]",
			containerSetup: "([)]",
			expectedResult: false,
		},
		{
			title:          "Testing valid setup {[}",
			containerSetup: "{[}",
			expectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result, _ := BracketCheck(test.containerSetup)
			if result != test.expectedResult {
				t.Errorf("expected result: `%t` got: `%t`", test.expectedResult, result)
			}
		})
	}
}
