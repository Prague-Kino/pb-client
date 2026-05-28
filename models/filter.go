package models

import (
	"fmt"
	"strings"
	"time"
)

type Filter struct {
	// Exact ID of the record
	Id string
	// Title of a Film record
	FilmTitle string
	// Name of a Kino record
	KinoName string
	// Title of Film in a Screening
	ScreeningFilm string
	// Name of Kino in a Screening
	ScreeningKino string
	// Exact date of a Screening
	ScreeningDate time.Time
	// Filter Screenings from this date forward
	ScreeningDateFrom time.Time
	// Filter Screenings up to this date
	ScreeningDateTo time.Time
}

func (f *Filter) String() string {
	filters := []string{}

	if f == nil {
		return ""
	}

	if f.Id != "" {
		filters = append(filters, fmt.Sprintf(
			"id='%s'",
			f.Id,
		))
	}

	if f.FilmTitle != "" {
		filters = append(
			filters,
			fmt.Sprintf(
				"title~'%s'",
				f.FilmTitle,
			),
		)
	}

	if f.KinoName != "" {
		filters = append(filters, fmt.Sprintf(
			"name~'%s'",
			f.KinoName,
		))
	}

	if f.ScreeningFilm != "" {
		filters = append(filters, fmt.Sprintf(
			"film.title~'%s'",
			f.ScreeningFilm,
		))
	}

	if f.ScreeningKino != "" {
		filters = append(filters, fmt.Sprintf(
			"kino.name~'%s'",
			f.ScreeningKino,
		))
	}

	if !f.ScreeningDate.IsZero() {
		filters = append(filters, fmt.Sprintf(
			"date='%s'",
			DateString(f.ScreeningDate),
		))
	}

	if !f.ScreeningDateFrom.IsZero() && f.ScreeningDate.IsZero() {
		filters = append(filters, fmt.Sprintf(
			"date>='%s'",
			DateString(f.ScreeningDateFrom),
		))
	}

	if !f.ScreeningDateTo.IsZero() && f.ScreeningDate.IsZero() {
		filters = append(filters, fmt.Sprintf(
			"date<='%s'",
			DateString(f.ScreeningDateTo),
		))
	}

	if len(filters) == 0 {
		return ""
	}
	return fmt.Sprintf(
		"(%s)",
		strings.Join(filters, " && "),
	)
}

func DateString(date time.Time) string {
	return date.Format("2006-01-02 15:04:05.000Z")
}
