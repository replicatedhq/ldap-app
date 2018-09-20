package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type handlers struct {
	dbIsReady bool
	storage   *dbClient
}

type User struct {
	UUID           string `json:"uuid"`
	UserID         string `json:"user_id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	FullName       string `json:"full_name"`
	PasswordFormat string `json:"password_format"`
	Password       string `json:"password"`
	Email          string `json:"email"`
}

func (h *handlers) waitForDB() {
	h.storage = getDBClient()
	for {
		logger.Println("Checking if DB is ready")

		h.dbIsReady = h.storage.isReady()
		if h.dbIsReady {
			logger.Println("DB is ready")
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (h *handlers) handlePing(c *gin.Context) {
	if !h.dbIsReady {
		c.AbortWithStatus(http.StatusServiceUnavailable)
	} else {
		c.String(http.StatusNoContent, "")
	}
}

func (h *handlers) handleCreate(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.storage.createUser(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusCreated, "")
}

func (h *handlers) handleModify(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.storage.updateUser(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusNoContent, "")
}

func (h *handlers) handleDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.storage.deleteUser(uuid); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusNoContent, "")
}
