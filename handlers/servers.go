package handlers

import (
	"time"

	"github.com/sawatkins/quickread/database"
	"github.com/sawatkins/quickread/models"

	"github.com/gofiber/fiber/v2"
	"github.com/xyproto/randomstring"
)

// CreateServer registers a server
func CreateServer(c *fiber.Ctx) error {
	randomstring.Seed()
	randString := randomstring.HumanFriendlyString(10)

	// Validation?
	server := models.Server{
		Id:          randString,
		UserId:      c.Query("user_id"),
		StartingMap: c.Query("starting_map"),
		Region:      c.Query("region"),
		Status:      c.Query("status"),
		Public:      parseBool(c.Query("public")),
		CreatedOn:   time.Now().Format("2006-01-02 15:04:05 UTC-0700"),
	}

	insertedId, err := database.AddServer(server)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to add server \"" + server.Id + "\" to database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success":    true,
		"server":     server,
		"insertedID": insertedId,
	})
}

// GetServer returns a server
func GetServer(c *fiber.Ctx) error {
	// get server from db
	server, err := database.GetServer(c.Query("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to get server \"" + c.Query("id") + "\" from database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"server":  server,
	})
}

// GetAllServers returns all servers
func GetAllServers(c *fiber.Ctx) error {
	// get all servers from db
	servers, err := database.GetAllServers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to get all servers from database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"servers": servers,
	})
}

// UpdateServer updates a server
func UpdateServer(c *fiber.Ctx) error {
	// update server in db
	// querey validation? todo later
	// get current server
	server, err := database.GetServer(c.Query("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Could not find server \"" + c.Query("id") + "\" in database",
		})
	}
	// determine which values to update
	if c.Query("user_id") != "" {
		server.UserId = c.Query("user_id")
	}
	if c.Query("starting_map") != "" {
		server.StartingMap = c.Query("starting_map")
	}
	if c.Query("region") != "" {
		server.Region = c.Query("region")
	}
	if c.Query("url") != "" {
		server.Url = c.Query("url")
	}
	if c.Query("status") != "" {
		server.Status = c.Query("status")
	}
	if c.Query("password") != "" {
		server.Password = c.Query("password")
	}
	// update server in db
	itemsUpdated, err := database.UpdateServer(c.Query("id"), server)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success":       false,
			"error":         "Failed to update server \"" + c.Query("id") + "\" in database",
			"error_message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success":      true,
		"server":       server,
		"itemsUpdated": itemsUpdated,
	})
}

// DeleteServer deletes a server
func DeleteServer(c *fiber.Ctx) error {
	// delete server from db
	itemsDeleted, err := database.DeleteServer(c.Query("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to delete server \"" + c.Query("id") + "\" from database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success":      true,
		"itemsDeleted": itemsDeleted,
	})
}

// GetActivePublicServers returns all active public servers
func GetActivePublicServers(c *fiber.Ctx) error {
	// get all servers from db
	servers, err := database.GetActivePublicServers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success":       false,
			"error":         "Failed to get all active public servers from database",
			"error_message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"servers": servers,
	})
}
