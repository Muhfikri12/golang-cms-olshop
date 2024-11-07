package handler

import (
	"net/http"

	"github.com/Muhfikri12/golang-cms-olshop/database"
	"github.com/Muhfikri12/golang-cms-olshop/repository"
	"github.com/Muhfikri12/golang-cms-olshop/service"
)

func Transaction(w http.ResponseWriter, r *http.Request)  {
	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "failed to connect database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repo := repository.NewRepoTransaction(db)
	service := service.NewServiceTransaction(repo)

	trx, err := service.Transaction()
	if err != nil {
		http.Error(w, "data is empty " + err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.ExecuteTemplate(w, "transaction-list", trx); err != nil {
		http.Error(w, "Template Execution Error: "+err.Error(), http.StatusInternalServerError)
	}
}