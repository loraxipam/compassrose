// Package compassrose provides data structures and minimal functions to use a 32 point compass
package compassrose

// Number of compass rose points for use with the level parameter
const (
	TwoPoints       = iota // N,S
	FourPoints             // N,E,S,W
	EightPoints            // N, NE, E, SE, S, SW, W, NW
	SixteenPoints          // N,NNE,NE,ENE,E,ESE,SE...
	ThirtyTwoPoints        // N,NbE,NNE,NEbN,NE,NEbE...
)

// CompassRose provides the names for the 32 cardinal compass points
var CompassRose = [32][3]string{
	{"N", "North", "Tramontana"},                             // L0..5
	{"NbE", "North by East", "Qto Tramontana v Greco"},       // L5
	{"NNE", "North Northeast", "Tramontana Greco"},           // L4..5
	{"NEbN", "Northeast by North", "Qto Greco v Tramontana"}, // L5

	{"NE", "Northeast", "Greco"}, // L3..5
	{"NEbE", "Northeast by East", "Qto Greco v Levante"},
	{"ENE", "East Northeast", "Greco Levante"}, // L4..5
	{"EbN", "East by North", "Qto Levante v Greco"},

	{"E", "East", "Levante"}, // L2..5
	{"EbS", "East by South", "Qto Levante v Scirocco"},
	{"ESE", "East Southeast", "Levante Scirocco"}, // L4..5
	{"SEbE", "Southeast by East", "Qto Scirocco v Levante"},

	{"SE", "Southeast", "Scirocco"}, // L3..5
	{"SEbS", "Southeast by South", "Qto Scirocco v Ostro"},
	{"SSE", "South Southeast", "Ostro Scirocco"}, // L4..5
	{"SbE", "South by East", "Qto Ostro v Scirocco"},

	{"S", "South", "Ostro"}, // L1..5
	{"SbW", "South by West", "Qto Ostro v Libeccio"},
	{"SSW", "South Southwest", "Ostro Libeccio"}, // L4..5
	{"SWbS", "Southwest by South", "Qto Libeccio v Ostro"},

	{"SW", "Southwest", "Libeccio"}, // L3..5
	{"SWbW", "Southwest by West", "Qto Libeccio v Ponente"},
	{"WSW", "West Southwest", "Ponente Libeccio"}, // L4..5
	{"WbS", "West by South", "Qto Ponente v Libeccio"},

	{"W", "West", "Ponente"}, // L2..5
	{"WbN", "West by North", "Qto Ponente v Maestro"},
	{"WNW", "West Northwest", "Ponente Maestro"}, // L4..5
	{"NWbW", "Northwest by West", "Qto Maestro v Ponente"},

	{"NW", "Northwest", "Maestro"}, // L3..5
	{"NWbN", "Northwest by North", "Qto Maestro v Tramontana"},
	{"NNW", "North Northwest", "Tramontana Maestro"}, // L4..5
	{"NbW", "North by West", "Qto Tramontana v Maestro"},

	// {"N2","North","Tramontana"}
}

// angleCleaner takes any kind of - or + angle and turns it into an angle 360 and 0
func angleCleaner(inputAngle float32) (cleanAngle float32) {
	var quo int
	if inputAngle > 360.0 || inputAngle < -360.0 {
		quo = int(inputAngle) / 360
		cleanAngle = inputAngle - float32(360*quo)
	} else {
		cleanAngle = inputAngle
	}
	if cleanAngle < 0 {
		cleanAngle = 360.0 + cleanAngle
	}
	return
}

// angleLevel converts the finest (32 points) direction to a wider level of
// direction.
func angleLevel(inputDir int, level int) int {
	if level > 4 || level < 0 {
		// return error{"Bad input"}
		return -99
	}
	// levelLU returns the
	levelLU := [5][32]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 8, 8, 8, 8, 8, 8, 8, 8, 16, 16, 16, 16, 16, 16, 16, 16, 24, 24, 24, 24, 24, 24, 24, 24, 0, 0, 0, 0},
		{0, 0, 4, 4, 4, 4, 8, 8, 8, 8, 12, 12, 12, 12, 16, 16, 16, 16, 20, 20, 20, 20, 24, 24, 24, 24, 28, 28, 28, 28, 0, 0},
		{0, 2, 2, 4, 4, 6, 6, 8, 8, 10, 10, 12, 12, 14, 14, 16, 16, 18, 18, 20, 20, 22, 22, 24, 24, 26, 26, 28, 28, 30, 30, 0},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
	}

	//fmt.Println("angleLevel", inputDir, "level", level)

	outputDir := levelLU[level][inputDir]
	return outputDir
}

// DegreeToHeading shows the compass heading
// @param inDegrees - decimal degrees
// @param level - 4: 32 points, 3: 16 points, 2: 8 points, 1: 4 points, 0: N or S, who knows.
// @param standard - true for English description, false for Ancient Mariners of the Mediterranean to understand
// @return shorthand - short string code for the direction
// @return text - the descriptive string for the direction
func DegreeToHeading(inDegrees float32, level int, standard bool) (shorthand string, text string) {

	// look this up based on the precision
	adjAngle := angleCleaner(inDegrees + 5.625)

	//fmt.Println(inDegrees, adjAngle, int(adjAngle/11.25))

	dir := angleLevel(int(adjAngle/11.25), level)
	if dir == -99 {
		return "ERROR", "BAD LEVEL"
	}
	if standard {
		return CompassRose[dir][0], CompassRose[dir][1]
	}
	return CompassRose[dir][0], CompassRose[dir][2]

}

// DegreeToHeadingSouthfacing shows compass heading based on South as cardinal point, not North
// but the rose is still absolutely North-based. That is, a 0 degree input angle will return "S"
// and -45 will return "SE". Same signature as DegreeToHeading.
func DegreeToHeadingSouthfacing(inDegrees float32, level int, standard bool) (string, string) {
	return DegreeToHeading(inDegrees-180, level, standard)
}
