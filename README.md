# pb-client

API to connect with the PocketBase database.

---

## Setup Instructions

Create a .env file at the root of this project with the following structure:

```text
POCKETBASE_ADMIN_EMAIL= //Email of PB superuser
POCKETBASE_ADMIN_PASS= //Password of PB superuser
POCKETBASE_URL= //URL to PocketBase (default local url: http://127.0.0.1:8090)
APP_ENV=dev
```

---

## Usage

Remember to load ENV variables, as this API uses `os.Getenv` under the hood,
to avoid exposing the admin email and pass.

ENV variables can be loaded easily using `godotenv.Load()` for example.

### Constructor

All API calls are exposed through the `PocketBase` struct. It can be created
using the pocketbase URL as follows:

```go
baseUrl := os.Getenv("POCKETBASE_URL")
pb, err := api.NewPocketBase(baseUrl)
if err != nil {
  log.Fatal(err)
}
```

---

### Create

Each struct has a `Create` method for creating a new PocketBase record of
the relevant type.

---

#### CreateKino

Takes in a `cast.Kino` object and returns an error if any occured.

```go
err := pb.CreateKino(&kino)
```

---

#### CreateFilm

Takes in a `cast.Film` object and returns an error if any occured.

```go
err := pb.CreateFilm(&film)
```

---

#### CreateScreening

Takes in a `cast.Screening` object and returns an error if any occured.

```go
err = pb.CreateScreening(&screening)
```

---

### Get

Every struct has a `Get` method that takes in a `model.Filter` for defining the
search criteria, shortcut methods for each relevant search criteria,
and a `GetAll` method for returning all records in the collection.

---

#### GetKino

The `GetKino` endpoint uses fuzzy matching, so you don't need to specify
the exact name.

It takes in a `Filter` struct that allows you to specify which fields you
want to search by. The available options so far are Kino Name and Kino Id.

It returns a single `cast.Kino` object, or `nil` and an error if none is found.

```go
filter := &models.Filter{
  KinoName: "aero"
}
kino, err := pb.GetKino(filter)
```

##### Kino Shortcuts

There are shortcut methods available to fetch a Kino directly by its name or id.

```go
kino, err := pb.GetKinoByName("aero")
kinoById, err := pb.GetKinoById("x92d1vgf5anpuv6")
```

#### GetAllKinos

Returns a list of all available Kinos and an error.

```go
allKinos, err := pb.GetAllKinos()
if err != nil {
  log.Fatal(err)
}
for _, k := range allKinos {
  fmt.Println(k.Name)
}
```

---

#### GetFilm

The `GetFilm` endpoint uses fuzzy matching, so you don't need to specify
the exact titles.

It takes in a `Filter` struct that allows you to specify which fields you want
to search by. The available options so far are Film Title and Film Id.

It returns a single `cast.Film` object, or `nil` and an error if none is found.

```go
filter := &models.Filter{
  FilmTitle: "The Emperor's New Groove"
}
film, err := pb.GetFilm(filter)
```

```go
filter := &models.Filter{Id: "wureceg5i0hpl28"}
film, err := pb.GetFilm(filter)
```

##### Film Shortcuts

It also features shortcut methods to fetch directly by title or id.

```go
filmByTitle, err := pb.GetFilmByTitle("fire walk with me")
filmById, err := pb.GetFilmById("x92d1vgf5anpuv6")
```

---

#### GetAllFilms

This endpoint returns all Films in PocketBase. It takes in no parameters.
It returns a list of `cast.Film`.

```go
allFilms, err := pb.GetAllFilms()
```

---

#### GetScreening

`GetScreening` takes in a `Filter` struct that allows you to specify which fields
you want to search by.

The available options so far are:

- Screening Id
- Film title
- Kino name
- Exact screening date
- Screening date from
- Screening date to

It returns a single `cast.Screening` object, or `nil` and an error if none is found.

```go
screening, err := pb.GetScreening(&models.Filter{
  Id: "epp8eaukme8enad",
})

screening, err := pb.GetScreening(&models.Filter{
  ScreeningKino: "kino aero",
  ScreeningFilm: "pirates of the caribbean",
})
```

#### GetScreenings

`GetScreenings` takes in the same parameters as `GetScreening` and
returns a list of all matching items.

The following example returns all screenings that occured in the
past 7 days at Kino Lucerna.

```go
  screenings, err := pb.GetScreenings(&models.Filter{
    ScreeningDateFrom: time.Now().AddDate(0, 0, -7),
    ScreeningDateTo:   time.Now(),
    ScreeningKino:     "Lucerna",
})
```

#### GetAllScreenings

This endpoint returns all available Screenings. It takes in no parameters.
It returns a list of `cast.Screening` and an error.

```go
allScreenings, err := pb.GetAllScreenings()
if err != nil {
  log.Fatal(err)
}
for _, s := range *allScreenings {
  fmt.Println(s)
}
```

---
