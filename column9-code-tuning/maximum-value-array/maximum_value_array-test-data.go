package maximumvalue

type testData struct {
	array    []float64
	maxValue float64
}

var tests = []testData{
	{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 11.5},
	{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1, 11.55}, maxValue: 11.55},
	{array: []float64{20, 3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 20},
	{array: []float64{5.1}, maxValue: 5.1},
	{array: []float64{-3.5, -2.00, -11.5, -9.8, -5.1}, maxValue: -2.00},
}

var testsEx = []testData{
	{array: nil, maxValue: 0},
	{array: []float64{}, maxValue: 0},
}

type testDataNarray struct {
	array    []float64
	maxValue float64
	n        int
}

var testsnarray = []testDataNarray{
	{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 11.5},
	{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 4.54, n: 2},
	{array: []float64{5.1}, maxValue: 5.1},
	{array: []float64{-3.5, -2.00, -11.5, 9.8, 5.1}, maxValue: -2.00, n: 3},
}

var testsnarrayEx = []testDataNarray{
	{array: nil, maxValue: 0},
	{array: []float64{}, maxValue: 2},
	{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 4.54},
	{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 4.54, n: -4},
	{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 4.54, n: 6},
}
