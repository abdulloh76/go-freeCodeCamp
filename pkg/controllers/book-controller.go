package controllers

import (
	"book-store/pkg/models"
	"book-store/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooks(response http.ResponseWriter, request *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func GetBookById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func CreateBook(response http.ResponseWriter, request *http.Request) {
	book := &models.Book{}
	utils.ParseBody(request, book)
	newBook := book.CreateBook()
	res, _ := json.Marshal(newBook)

	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(request, updateBook)

	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(id)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)

}

func DeleteBook(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails := models.DeleteBook(id)
	res, _ := json.Marshal(bookDetails)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}
