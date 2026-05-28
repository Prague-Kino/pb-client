package api

import (
	"encoding/json"
	"fmt"

	"github.com/Prague-Kino/cast"
	"github.com/Prague-Kino/pb-client/internal/errors"
	"github.com/Prague-Kino/pb-client/internal/pocketbase"
	"github.com/Prague-Kino/pb-client/models"
)

// --------------------------------------

func (pb *PocketBase) GetAllKinos() ([]cast.Kino, error) {
	filter := &models.Filter{}
	body, err := pb.get(pocketbase.KinoCollection, filter, nil)
	if err != nil {
		return nil, err
	}

	var res models.KinosResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return res.Items, nil
}

func (pb *PocketBase) GetKino(filter *models.Filter) (*cast.Kino, error) {
	body, err := pb.get(pocketbase.KinoCollection, filter, nil)
	if err != nil {
		return nil, err
	}

	var res models.KinosResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	if len(res.Items) == 0 {
		return nil, &errors.KinoNotFoundError{Filter: *filter}
	}

	return &res.Items[0], nil
}

func (pb *PocketBase) GetKinoByName(name string) (*cast.Kino, error) {
	filter := &models.Filter{KinoName: name}
	return pb.GetKino(filter)
}

func (pb *PocketBase) GetKinoById(id string) (*cast.Kino, error) {
	filter := &models.Filter{Id: id}
	return pb.GetKino(filter)
}

// --------------------------------------

func (pb *PocketBase) GetFilm(filter *models.Filter) (*cast.Film, error) {
	body, err := pb.get(pocketbase.FilmCollection, filter, nil)
	if err != nil {
		return nil, err
	}

	var res models.FilmsResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	if len(res.Items) == 0 {
		return nil, &errors.FilmNotFoundError{Filter: *filter}
	}

	return &res.Items[0], nil
}

func (pb *PocketBase) GetFilmByTitle(title string) (*cast.Film, error) {
	filter := &models.Filter{FilmTitle: title}
	return pb.GetFilm(filter)
}

func (pb *PocketBase) GetFilmById(id string) (*cast.Film, error) {
	filter := &models.Filter{Id: id}
	return pb.GetFilm(filter)
}

func (pb *PocketBase) GetAllFilms() ([]cast.Film, error) {
	filter := &models.Filter{}
	body, err := pb.get(pocketbase.FilmCollection, filter, nil)
	if err != nil {
		return nil, err
	}

	var res models.FilmsResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return res.Items, nil
}

// --------------------------------------

func (pb *PocketBase) GetScreening(filter *models.Filter) (*cast.Screening, error) {
	body, err := pb.get(
		pocketbase.ScreeningCollection,
		filter,
		models.ExpandOpts("kino", "film"),
	)
	if err != nil {
		return nil, err
	}

	var res models.ScreeningsResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	if len(res.Items) == 0 {
		return nil, &errors.ScreeningNotFoundError{Filter: *filter}
	}

	return &res.Items[0], nil
}

func (pb *PocketBase) GetScreenings(filter *models.Filter) (*[]cast.Screening, error) {
	body, err := pb.get(
		pocketbase.ScreeningCollection,
		filter,
		models.ExpandOpts("kino", "film"),
	)
	if err != nil {
		return nil, err
	}

	var res models.ScreeningsResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	if len(res.Items) == 0 {
		return nil, &errors.ScreeningNotFoundError{Filter: *filter}
	}

	return &res.Items, nil
}

func (pb *PocketBase) GetAllScreenings() (*[]cast.Screening, error) {
	body, err := pb.get(pocketbase.ScreeningCollection, nil, models.ExpandOpts("kino", "film"))
	if err != nil {
		return nil, err
	}

	var res models.ScreeningsResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	for _, s := range res.Items {
		fmt.Println(s)
	}

	return &res.Items, nil
}

// --------------------------------------
