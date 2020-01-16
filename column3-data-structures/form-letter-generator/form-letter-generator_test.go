// +build all column3 formgenerator lettergenerator generator

package formgenerator

import (
	"fmt"
	"strings"
	"testing"

	assert "github.com/LuigiAndrea/test-helper/assertions"
)

var filename = "./templateEmailAbsence.data"

type testData struct {
	in  string
	out int
}

type emailFormalAbsence struct {
	firstName, lastName, supervisorName, startDateAway, endDateAway, returningDate string
}

func TestFormLetterGenerator(t *testing.T) {

	par := emailFormalAbsence{firstName: "Luigi", lastName: "Gracco", supervisorName: "Enzo",
		startDateAway: "2 January 2020", endDateAway: "12 January 2020", returningDate: "13 January 2020"}

	result, _ := formLetterGenerator(filename, par.firstName, par.lastName, par.supervisorName, par.startDateAway, par.endDateAway, par.returningDate)

	tests := []testData{
		testData{in: par.firstName, out: 2},
		testData{in: par.lastName, out: 1},
		testData{in: par.supervisorName, out: 2},
	}

	for _, test := range tests {
		if err := assert.AssertDeepEqual(test.out, strings.Count(result, test.in)); err != nil {
			t.Error(err)
		}
	}

	fmt.Println(result)
}

func TestFormLetterGeneratorLessParameters(t *testing.T) {
	par := emailFormalAbsence{firstName: "Luigi", lastName: "Gracco", supervisorName: "Enzo"}

	result, _ := formLetterGenerator(filename, par.firstName, par.lastName, par.supervisorName)

	tests := []testData{
		testData{in: par.firstName, out: 2},
		testData{in: par.lastName, out: 1},
		testData{in: "#4", out: 1},
		testData{in: par.returningDate, out: len(result) + 1},
		testData{in: "#1", out: 0},
	}

	for _, test := range tests {
		if err := assert.AssertDeepEqual(test.out, strings.Count(result, test.in)); err != nil {
			t.Error(err)
		}
	}
}
