package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/Muhfikri12/golang-cms-olshop/database"
	"github.com/Muhfikri12/golang-cms-olshop/model"
	"github.com/Muhfikri12/golang-cms-olshop/repository"
	"github.com/Muhfikri12/golang-cms-olshop/service"
)

var templates = template.Must(template.ParseGlob("view/*.html"))

func FormCreateBook(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "add-book", nil)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseMultipartForm(10 << 20) // Limit file upload to 10MB
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	category := r.FormValue("category")
	author := r.FormValue("author")

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	discount, err := strconv.ParseFloat(r.FormValue("discount"), 64)
	if err != nil {
		http.Error(w, "Invalid discount", http.StatusBadRequest)
		return
	}

	// Handle file upload for cover image
	coverFile, coverHandler, err := r.FormFile("cover")
	if err != nil {
		http.Error(w, "Error uploading cover image", http.StatusBadRequest)
		return
	}
	defer coverFile.Close()
	coverFilePath := fmt.Sprintf("uploads/images/%s", coverHandler.Filename)
	coverOut, err := os.Create(coverFilePath)
	if err != nil {
		http.Error(w, "Error saving cover image" + err.Error(), http.StatusInternalServerError)
		return
	}
	defer coverOut.Close()
	_, err = io.Copy(coverOut, coverFile)
	if err != nil {
		http.Error(w, "Error copying cover image", http.StatusInternalServerError)
		return
	}

	// Handle file upload for book PDF
	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error uploading PDF file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	filePath := fmt.Sprintf("uploads/files/%s", fileHandler.Filename)
	fileOut, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error saving PDF file", http.StatusInternalServerError)
		return
	}
	defer fileOut.Close()
	_, err = io.Copy(fileOut, file)
	if err != nil {
		http.Error(w, "Error copying PDF file", http.StatusInternalServerError)
		return
	}

	// Connect to the database
	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "Database connection error" + err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Create book object
	book := model.Books{
		Name:     name,
		Category: category,
		Author:   author,
		Cover:    coverFilePath,
		Pdf:      filePath,
		Price:    price,
		Discount: discount,
	}

	// Insert the book into the database
	repo := repository.NewBookRepo(db)
	service := service.NewBookService(repo)
	if err := service.CreateBookService(&book); err != nil {
		http.Error(w, "Error creating book", http.StatusInternalServerError)
		fmt.Println("Error:", err)
		return
	}

	// Redirect to the books list page
	http.Redirect(w, r, "/books-list", http.StatusSeeOther)
}

func ItemsList(w http.ResponseWriter, r *http.Request)  {
	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "Database Connection Error" + err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	repo := repository.NewBookRepo(db)
	service := service.NewBookService(repo)

	items, err := service.ShowALlList()
	if err != nil {
		http.Error(w, "Not Found Items" + err.Error(), http.StatusInternalServerError)
	}

	if err := templates.ExecuteTemplate(w, "book-list", items); err != nil {
		http.Error(w, "Template Execution Error: "+err.Error(), http.StatusInternalServerError)
	}
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	id_int, _ := strconv.Atoi(id)

	err := r.ParseMultipartForm(10 << 20) 
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	category := r.FormValue("category")
	author := r.FormValue("author")

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	discount, err := strconv.ParseFloat(r.FormValue("discount"), 64)
	if err != nil {
		http.Error(w, "Invalid discount", http.StatusBadRequest)
		return
	}

	coverFile, coverHandler, err := r.FormFile("cover")
	if err != nil {
		http.Error(w, "Error uploading cover image", http.StatusBadRequest)
		return
	}
	defer coverFile.Close()
	coverFilePath := fmt.Sprintf("uploads/images/%s", coverHandler.Filename)
	coverOut, err := os.Create(coverFilePath)
	if err != nil {
		http.Error(w, "Error saving cover image" + err.Error(), http.StatusInternalServerError)
		return
	}
	defer coverOut.Close()
	_, err = io.Copy(coverOut, coverFile)
	if err != nil {
		http.Error(w, "Error copying cover image", http.StatusInternalServerError)
		return
	}

	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error uploading PDF file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	filePath := fmt.Sprintf("uploads/files/%s", fileHandler.Filename)
	fileOut, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error saving PDF file", http.StatusInternalServerError)
		return
	}
	defer fileOut.Close()
	_, err = io.Copy(fileOut, file)
	if err != nil {
		http.Error(w, "Error copying PDF file", http.StatusInternalServerError)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "Database connection error: "+ err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	book := model.Books{
		Name:     name,
		Category: category,
		Author:   author,
		Cover:    coverFilePath,
		Pdf:      filePath,
		Price:    price,
		Discount: discount,
	}

	repo := repository.NewBookRepo(db)
	service := service.NewBookService(repo)
	if err := service.UpdateBook(id_int, &book); err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/books-list", http.StatusSeeOther)
}
	