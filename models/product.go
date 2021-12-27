package models

type Product struct {
	Product_id       string      `json:"product_id,omitempty"`
	Name             string      `json:"name,omitempty"`
	Description      string      `json:"description"`
	Status           string      `json:"status"`
	Creation_date    string      `json:"creation_date"`
	Update_date      string      `json:"update_date"`
	Account_id       string      `json:"account_id"`
	Format_product   interface{} `json:"formtat_product"`
	Value_unit       float64     `json:"value_unit"`
	Unit_name        string      `json:"unit_name"`
	Unit_description string      `json:"unit_description"`
	Stock            int         `json:"stock"`
}
