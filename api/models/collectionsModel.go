package models

import ()

type CollectionsModel struct {
	ID         	int    	`json:"id"`
	Status 		string 	`json:"status"`
	Amount 		int 	`json:"amount"`
	Metadata 	[]Metadata
}

type Metadata struct {
	Reason 				string 	`json:"reason"`
	ExternalReference 	string 	`json:"externalReference"`
}