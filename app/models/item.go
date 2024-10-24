package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Item struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}

func GetItems(db *sql.DB) ([]Item, error) {
    rows, err := db.Query("SELECT id, name, price FROM items")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []Item
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}

func GetItem(db *sql.DB, id string) (Item, error) {
    var item Item
    err := db.QueryRow("SELECT id, name, price FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name, &item.Price)
    return item, err
}

func CreateItem(db *sql.DB, item Item) (int, error) {
    var id int
    err := db.QueryRow("INSERT INTO items(name, price) VALUES($1, $2) RETURNING id", item.Name, item.Price).Scan(&id)
    return id, err
}

func UpdateItem(db *sql.DB, id string, item Item) error {
    _, err := db.Exec("UPDATE items SET name = $1, price = $2 WHERE id = $3", item.Name, item.Price, id)
    return err
}

func DeleteItem(db *sql.DB, id string) error {
    _, err := db.Exec("DELETE FROM items WHERE id = $1", id)
    return err
}
