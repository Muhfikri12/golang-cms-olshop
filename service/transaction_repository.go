package service

import (
	"github.com/Muhfikri12/golang-cms-olshop/model"
	"github.com/Muhfikri12/golang-cms-olshop/repository"
)

type ServiceTransaction struct {
	RepoTransaction repository.RepoTransactionDB
}

func NewServiceTransaction(repo repository.RepoTransactionDB) *ServiceTransaction {
	return &ServiceTransaction{RepoTransaction: repo}
}

func (st *ServiceTransaction) Transaction() (*[]model.Transaction, error) {
	trx, err := st.RepoTransaction.Transaction()
	if err != nil {
		return nil, err
	}

	return trx, err
}