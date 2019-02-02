package quizer

import "testing"

func TestIsCorrect(t *testing.T) {
	for _, tc := range isCorrectCases {
		given := tc.given
		real := tc.real
		got := isCorrect(given, real)
		expected := tc.want
		if got != expected {
			t.Errorf(`FAIL : %s isCorrect(%s, %s) = %t, want %t`, tc.description, given, real, got, expected)
		}
		t.Log("PASS : ", tc.description)
	}
	t.Log("Tested : ", len(isCorrectCases), " cases")
	// given := "5"
	// real := "5"
	// expected := true
	// actual := isCorrect(given, real)

	// if actual != expected {
	// 	t.Errorf("Expected answer is %t got %t", expected, actual)
	// }
}
