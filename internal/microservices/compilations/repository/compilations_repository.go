package repository

import (
	"database/sql"
	"fmt"
	"myapp/internal/microservices/compilations"
	"myapp/internal/microservices/compilations/proto"

	"github.com/lib/pq"
)

type movieCompilationsStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) compilations.Storage {
	return &movieCompilationsStorage{db: db}
}

const (
	getAllMoviesSQL = "SELECT m.id, m.name, m.picture FROM movies AS m WHERE m.is_movie=$1 LIMIT $2 OFFSET $3"
	getByGenreSQL   = "SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_genre m_g ON " +
		"m_g.movie_id = m.id WHERE m_g.genre_id=$1"
	getGenreNameSQL = "SELECT name FROM genre WHERE id=$1"
	getByCountrySQL = "SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_countries m_c ON " +
		"m_c.movie_id = m.id WHERE m_c.country_id=$1"
	getCountryNameSQL = "SELECT name FROM country WHERE id=$1"
	getByMovieSQL     = "SELECT DISTINCT m.id, m.name, m.picture FROM movies AS m " +
		"JOIN movies_genre m_g ON m_g.movie_id = m.id " +
		"JOIN movies_genre m_g2 ON m_g2.genre_id = m_g.genre_id " +
		"WHERE m_g2.movie_id=$1"
	getByPersonSQL = "SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_staff m_s ON " +
		"m_s.movie_id = m.id WHERE m_s.person_id=$1"
	getTopSQL       = "SELECT id, name, picture FROM movies ORDER BY kinopoisk_rating DESC LIMIT $1"
	getTopByYearSQL = "SELECT id, name, picture FROM movies WHERE year=$1 ORDER BY kinopoisk_rating DESC"
	getFavorites    = "SELECT id, name, picture FROM movies WHERE id = any ($1) LIMIT $2 OFFSET $3;"
	findMovie       = "SELECT id, name, picture FROM movies where to_tsvector('russian', name) @@ to_tsquery('russian', $1) AND is_movie=$2;"
)

func (ms *movieCompilationsStorage) GetAllMovies(limit, offset int, isMovie bool) (*proto.MovieCompilation, error) {
	var selectedMovieCompilation proto.MovieCompilation

	rows, err := ms.db.Query(getAllMoviesSQL, isMovie, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetByGenre(genreID int) (*proto.MovieCompilation, error) {
	var selectedMovieCompilation proto.MovieCompilation

	err := ms.db.QueryRow(getGenreNameSQL, genreID).Scan(&selectedMovieCompilation.Name)
	if err != nil {
		return nil, err
	}

	rows, err := ms.db.Query(getByGenreSQL, genreID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetByCountry(countryID int) (*proto.MovieCompilation, error) {
	var selectedMovieCompilation proto.MovieCompilation

	err := ms.db.QueryRow(getCountryNameSQL, countryID).Scan(&selectedMovieCompilation.Name)
	if err != nil {
		return nil, err
	}

	rows, err := ms.db.Query(getByCountrySQL, countryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetByMovie(movieID int) (*proto.MovieCompilation, error) {
	var selectedMC proto.MovieCompilation
	selectedMC.Name = "Похожие по жанру"

	rows, err := ms.db.Query(getByMovieSQL, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		// Костыль запроса. Необходимо добавить в запрос исключение исходного фильма
		if selectedMovie.ID != int64(movieID) {
			selectedMC.Movies = append(selectedMC.Movies, &selectedMovie)
		}
	}
	return &selectedMC, nil
}

func (ms *movieCompilationsStorage) GetByPerson(personID int) (*proto.MovieCompilation, error) {

	var selectedMovieCompilation proto.MovieCompilation
	selectedMovieCompilation.Name = "Фильмография"
	rows, err := ms.db.Query(getByPersonSQL, personID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetTop(limit int) (*proto.MovieCompilation, error) {

	var selectedMovieCompilation proto.MovieCompilation
	selectedMovieCompilation.Name = "Топ рейтинга"
	rows, err := ms.db.Query(getTopSQL, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}

	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetTopByYear(year int) (*proto.MovieCompilation, error) {

	var selectedMovieCompilation proto.MovieCompilation
	selectedMovieCompilation.Name = fmt.Sprintf("Лучшее за %d год", year)
	rows, err := ms.db.Query(getTopByYearSQL, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetFavorites(data *proto.GetFavoritesOptions) (*proto.MovieCompilation, error) {
	var selectedMovieCompilation proto.MovieCompilation

	rows, err := ms.db.Query(getFavorites, pq.Array(data.Id), data.Limit, data.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) FindMovie(text string, isMovie bool) (*proto.MovieCompilation, error) {
	var selectedMovieCompilation proto.MovieCompilation

	rows, err := ms.db.Query(findMovie, text, isMovie)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}
