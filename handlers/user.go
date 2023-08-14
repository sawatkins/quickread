package handlers

import (
	"log"
	"time"

	"github.com/sawatkins/upfast.tf-go/database"
	"github.com/sawatkins/upfast.tf-go/models"

	"github.com/gofiber/fiber/v2"
	"github.com/xyproto/randomstring"
)

// CreateUser registers a user
func CreateUser(c *fiber.Ctx) error {
	// later, replace this with param like below
	randomstring.Seed()
	randString := randomstring.HumanFriendlyString(10)

	newUser := models.User{
		Id: randString,
		CreatedOn: time.Now().Format("2006-01-02 15:04:05 UTC-0700"),
	}

	insertedId, err := database.AddUser(newUser)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error": "Failed to add user \"" + newUser.Id + "\" to database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"user":    newUser,
		"insertedID": insertedId,
	})	 
}

// GetUser returns a user
func GetUser(c *fiber.Ctx) error {
	// get user from db
	user, err := database.GetUser(c.Query("id"))
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error": "Failed to get user \"" + c.Query("id") + "\" from database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

// GetAllUsers returns all users
func GetAllUsers(c *fiber.Ctx) error {
	// get all users from db
	return nil
}

// DeleteUser deletes a user
func DeleteUser(c *fiber.Ctx) error {
	// delete user from db
	itemsDeleted, err := database.DeleteUser(c.Query("id"))
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error": "Failed to delete user \"" + c.Query("id") + "\" from database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"itemsDeleted": itemsDeleted,
	})
}