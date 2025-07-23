
package handlers

import (
	"cars/models"
	"fmt"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var mu sync.Mutex

// CarInventory godoc
// @Summary      Create a new car
// @Description  Add a new car to the inventory
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        cars   body      models.Car  true  "Add car"
// @Success      201  {object}  models.Car
// @Failure      400  {object}  models.Errors
// @Router       /cars [post]
func CreateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&models.Errors{
			Error: "incorrect input body",
			Details: err.Error(),
		})
	}

	car.Insert()

	fmt.Println("Car saved to the inventory with thew id: ", car.ID)

	return c.Status(fiber.StatusCreated).JSON(car)
}

func UpdateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	car.ID = id

	if err := car.Get(); err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "car not found with the given id",
		})
	}

	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    "incorrect input body",
			"details":  err.Error(),
		})
	}

	car.Update()

	fmt.Println("car updated with the id: ", id)

	return c.Status(fiber.StatusCreated).JSON(car)
}

// CarInventory godoc
// @Summary      Get a car
// @Description  Get a car from the inventory
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id   path      string true   "car id"
// @Success      200  {object}  models.Car
// @Failure      404  {object}  models.Errors
// @Failure      400  {object}  models.Errors
// @Router       /cars/{id} [get]
func GetCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}
	
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&models.Errors{
			Error: "invalid id",
			Details: err.Error(),
		})
	}

	car.ID = id

	if err := car.Get(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&models.Errors{
			Error: "car with the given id not found",
			Details: err.Error(),
		})
	}

	// fmt.Println("car found with the id:", id)

	return c.Status(fiber.StatusOK).JSON(car)
}

func DeleteCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

    id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	car.ID = id
    car.Delete()

	fmt.Println("car deleted with the id: ", id)

	return c.SendStatus(fiber.StatusNoContent)
}
