package models

type PlantsResponse struct {
	Plants 	[]Plants `json:"plants"`
}

type Plants struct {
	ID 	 string `json:"id"`
	Status 	 string `json:"status"`
}