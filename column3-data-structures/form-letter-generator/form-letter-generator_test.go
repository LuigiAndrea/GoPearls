// +build all column3 formgenerator lettergenerator generator form

package formgenerator

import (
	"os"
	"strings"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
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
		{in: par.firstName, out: 2},
		{in: par.lastName, out: 1},
		{in: par.supervisorName, out: 2},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.out, strings.Count(result, test.in)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestFormLetterGeneratorLessParameters(t *testing.T) {
	par := emailFormalAbsence{firstName: "Luigi", lastName: "Gracco", supervisorName: "Enzo"}

	result, _ := formLetterGenerator(filename, par.firstName, par.lastName, par.supervisorName)

	tests := []testData{
		{in: par.firstName, out: 2},
		{in: par.lastName, out: 1},
		{in: "#4", out: 1},
		{in: par.returningDate, out: len(result) + 1},
		{in: "#1", out: 0},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.out, strings.Count(result, test.in)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestFormLetterGeneratorWrongFilename(t *testing.T) {
	_, e := formLetterGenerator("file?")
	if err := a.AssertException(&os.PathError{}, e); err != nil {
		t.Error(err)
	}
}
