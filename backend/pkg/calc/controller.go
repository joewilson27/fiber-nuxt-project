package pkg

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joewilson27/go-vue-project/models"
)

func MainCalc(c *fiber.Ctx) error {
	// prepare model request to contain data
	var numsRequest models.NumsRequest
	// prepare response model
	var resData models.RespCalc

	if err := c.BodyParser(&numsRequest); err != nil {
		fmt.Println("Errrr", err.Error())
		resData.Data = nil

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
			"data":    resData,
		})
	}

	// process data
	returnData := process(numsRequest)
	resData.Data = returnData
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Ok",
		"data":    resData,
	})

}

func GetStudents(c *fiber.Ctx) error {

	var resData models.RespCalc

	students := generateAllStudents()

	resData.Data = students
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Ok",
		"data":    resData,
	})
}

func GetSingleStudent(c *fiber.Ctx) error {
	var resData models.RespCalc

	id, _ := strconv.Atoi(c.Params("id"))
	student := singleStudent(id)

	resData.Data = student
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Ok",
		"data":    resData,
	})
}
