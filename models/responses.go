package models

import "github.com/Prague-Kino/cast"

type KinosResponse struct {
	Items      []cast.Kino `json:"items"`
	Page       int         `json:"page"`
	PerPage    int         `json:"perPage"`
	TotalItems int         `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
}

type FilmsResponse struct {
	Items      []cast.Film `json:"items"`
	Page       int         `json:"page"`
	PerPage    int         `json:"perPage"`
	TotalItems int         `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
}

type ScreeningsResponse struct {
	Items      []cast.Screening `json:"items"`
	Page       int              `json:"page"`
	PerPage    int              `json:"perPage"`
	TotalItems int              `json:"totalItems"`
	TotalPages int              `json:"totalPages"`
}
