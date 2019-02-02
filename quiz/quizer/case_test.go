package quizer

var isCorrectCases = []struct {
	description string
	given       string
	real        string
	want        bool
}{
	{
		"Equal string should return true",
		"5",
		"5",
		true,
	},
	{
		"Non Equal string should return false",
		"5",
		"7",
		false,
	},
}
