package handlers

import pr "github.com/souviks72/flipcart-api/params"

type APIService struct {
	pr.Database
}

func InitAPIService(db pr.Database) *APIService {
	ps := &APIService{
		Database: db,
	}

	return ps
}
