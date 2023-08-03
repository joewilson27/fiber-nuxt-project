package pkg

import (
	"github.com/joewilson27/go-vue-project/models"
)

var students = []models.Student{
	{
		ID:      1,
		Name:    "John Doe",
		Address: "Brooklyn, NY",
		Grade:   7,
	},
	{
		ID:      2,
		Name:    "Anne Russell",
		Address: "Boston, MA",
		Grade:   8,
	},
}

func process(numsdata models.NumsRequest) models.NumsResponseData {
	var numsres models.NumsResponseData
	numsres.Add = numsdata.Num1 + numsdata.Num2
	numsres.Mul = numsdata.Num1 * numsdata.Num2
	numsres.Sub = numsdata.Num1 - numsdata.Num2
	numsres.Div = numsdata.Num1 / numsdata.Num2

	return numsres
}

func generateAllStudents() []models.Student {
	return students
}

func singleStudent(id int) models.Student {
	var data models.Student
	for _, v := range students {
		if v.ID == id {
			data = v
		}
	}
	return data
}
