package models

import (
	"database/sql"
	"fmt"
	"nadia/bookstore/database"
)

type Book struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func GetBooks() ([]Book, error) {
	rows, err := database.DB.Query("SELECT * FROM Books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Books []Book
	for rows.Next() {
		var Book Book
		if err := rows.Scan(&Book.ID, &Book.Title, &Book.Author, &Book.Price); err != nil {
			return nil, err
		}
		Books = append(Books, Book)
	}
	return Books, nil
}

func GetBooksById(id string) (Book, error) {
	var Book Book
	query := "SELECT * FROM Books WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&Book.ID, &Book.Title, &Book.Author, &Book.Price)
	if err == sql.ErrNoRows {
		return Book, fmt.Errorf("Book not found")
	}
	return Book, err
}

func (a *Book) AddBook() error {
	query := "INSERT INTO Books (title, Author, price) VALUES (?,?,?)"
	_, err := database.DB.Exec(query, a.Title, a.Author, a.Price)
	return err
}
