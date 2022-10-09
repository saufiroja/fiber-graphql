package controllers

import (
	"fiber/fiber-graphql/domain"
	"fiber/fiber-graphql/entity"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	CustomerService domain.CustomerService
}

func NewCustomerController(customerService domain.CustomerService) domain.CustomerController {
	return &CustomerController{
		CustomerService: customerService,
	}
}

func (h *CustomerController) GetAll(c *fiber.Ctx) error {
	customer, err := h.CustomerService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    customer,
	})
}

func (h *CustomerController) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	// convert string to int
	customerId, err := strconv.Atoi(id)

	customer, err := h.CustomerService.GetById(uint(customerId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    customer,
	})
}

func (h *CustomerController) CreateCustomer(c *fiber.Ctx) error {
	customer := entity.Customer{}

	err := c.BodyParser(&customer)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error",
		})
	}

	newCustomer, err := h.CustomerService.CreateCustomer(&customer)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success",
		"data":    newCustomer,
	})
}
