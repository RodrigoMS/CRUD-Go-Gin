package controllers

import (
    "database/sql"
    "github.com/RodrigoMS/CRUD-Go-Gin/models"
    "github.com/RodrigoMS/CRUD-Go-Gin/views"
    "github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context, db *sql.DB) {
    items, err := models.GetItems(db)
    views.RespondWithItems(c, items, err)
}

func CreateItem(c *gin.Context, db *sql.DB) {
    var item models.Item
    if err := c.ShouldBindJSON(&item); err != nil {
        views.RespondWithError(c, err)
        return
    }

    id, err := models.CreateItem(db, item)
    views.RespondWithID(c, id, err)
}

func GetItem(c *gin.Context, db *sql.DB) {
    id := c.Param("id")
    item, err := models.GetItem(db, id)
    views.RespondWithItem(c, item, err)
}

func UpdateItem(c *gin.Context, db *sql.DB) {
    id := c.Param("id")
    var item models.Item
    if err := c.ShouldBindJSON(&item); err != nil {
        views.RespondWithError(c, err)
        return
    }

    err := models.UpdateItem(db, id, item)
    views.RespondWithStatus(c, err)
}

func DeleteItem(c *gin.Context, db *sql.DB) {
    id := c.Param("id")
    err := models.DeleteItem(db, id)
    views.RespondWithStatus(c, err)
}
