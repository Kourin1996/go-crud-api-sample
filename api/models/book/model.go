package book

import (
	"github.com/Kourin1996/go-crud-api-sample/api/models/common"
	"github.com/Kourin1996/go-crud-api-sample/api/models/user"
)

const MODEL_NAME = "Book"

type Book struct {
	tableName struct{} `pg:"books"`
	common.BaseModel
	Name        string     `json:"name" pg:"name"`
	Description string     `json:"description" pg:"description"`
	Price       int64      `json:"price" pg:"price"`
	UserId      int64      `json:"-" pg:"user_id"`
	User        *user.User `json:"user,omitempty" pg:"rel:has-one"`
}

func NewEmptyBook() *Book {
	b := &Book{}
	b.Model = MODEL_NAME
	return b
}

func NewBook(userId int64, dto *CreateBookDto) (*Book, error) {
	b := NewEmptyBook()
	b.UserId = userId
	b.Name = dto.Name
	b.Description = dto.Description
	b.Price = dto.Price
	return b, nil
}

type UpdateBook struct {
	tableName struct{} `pg:"books"`
	common.BaseModel
	Name        *string `json:"name" pg:"name"`
	Description *string `json:"description" pg:"description"`
	Price       *int64  `json:"price" pg:"price"`
}

func NewEmptyUpdateBook() *UpdateBook {
	b := &UpdateBook{}
	b.Model = MODEL_NAME
	return b
}

func NewUpdateBook(hashId string, dto *UpdateBookDto) (*UpdateBook, error) {
	b := NewEmptyUpdateBook()
	if err := b.SetHashId(hashId); err != nil {
		return nil, err
	}
	b.Name = dto.Name
	b.Description = dto.Description
	b.Price = dto.Price
	return b, nil
}
