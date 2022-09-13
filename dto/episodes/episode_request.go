package episodesdto

type CreateEpisodeRequest struct {
	Title         string `json:"title" validate:"required"`
	ThumbnailFilm string `json:"thumbnailFilm" validate:"required"`
	LinkFilm      string `json:"linkFilm" validate:"required"`
	FilmID        int    `json:"film_id" form:"film_id" gorm:"type: int" validate:"required"`
}

type UpdateEpisodeRequest struct {
	Title         string `json:"title"`
	ThumbnailFilm string `json:"thumbnailFilm"`
	LinkFilm      string `json:"linkFilm"`
	FilmID        int    `json:"film_id" form:"film_id" gorm:"type: int"`
}