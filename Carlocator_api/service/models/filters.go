package models

import (
	"math"
)

type GetAllVehiclesFilter struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	//added key,staff filter in GetAllVehiclesFilter in order to get keys,staff or not in GetAll-vehicles API
	Keys                bool        `json:"keys"`
	Staff               bool        `json:"staff"`
	SearchFilterVehicle SearchQuery `json:"search"`
}

func (f GetAllVehiclesFilter) CalculateOffset() int {
	return (f.Page - 1) * f.Limit
}

type Metadata struct {
	CurrentPage  int `json:"page"`
	Limit        int `json:"limit"`
	FirstPage    int `json:"first_page"`
	LastPage     int `json:"last_page"`
	TotalRecords int `json:"total_records"`
}

type GetStaffMembersFilter struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (f GetStaffMembersFilter) CountOffset() int {
	return (f.Page - 1) * f.Limit
}

type GetAssignedKeyFilter struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (f GetAssignedKeyFilter) CalculatedOffset() int {
	return (f.Page - 1) * f.Limit
}
func ComputeMetadata(totalRecords, page, limit int) *Metadata {
	if totalRecords == 0 {
		return &Metadata{}
	}
	return &Metadata{
		CurrentPage:  page,
		Limit:        limit,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(limit))),
		TotalRecords: totalRecords,
	}

}

type SearchQuery struct {
	Model struct {
		Value *string `json:"value"`
	} `json:"model"`
	Color struct {
		Value *string `json:"value"`
	} `json:"veh_color"`
	VehicleAvailable struct {
		Value *bool `json:"value"`
	} `json:"veh_available"`
	VehicleName struct {
		Value *string `json:"value"`
	} `json:"veh_name"`
	Year struct {
		Value *string `json:"value"`
	} `json:"year"`
}
