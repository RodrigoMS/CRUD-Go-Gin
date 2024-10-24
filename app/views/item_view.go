package views

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/RodrigoMS/CRUD-Go-Gin/models"
)

func RespondWithItems(c *gin.Context, items []models.Item, err error) {
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, items)
}

func RespondWithItem(c *gin.Context, item models.Item, err error) {
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, item)
}

func RespondWithID(c *gin.Context, id int, err error) {
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"id": id})
}

func RespondWithStatus(c *gin.Context, err error) {
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func RespondWithError(c *gin.Context, err error) {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
