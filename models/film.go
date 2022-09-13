package models

import "time"

type Film struct {
  ID              int                  `json:"id" gorm:"primary_key:auto_increment"`
  Title           string               `json:"title" form:"title" gorm:"type: varchar(255)"`
  ThumbnailFilm   string               `json:"thumbnailFilm" form:"thumbnailFilm" gorm:"type: varchar(255)"`
  Year            int                  `json:"year" form:"year" gorm:"type: int"`
  Category        []Category           `json:"category" gorm:"many2many:film_categories"`
  Desc            string               `json:"desc" gorm:"type:text" form:"desc"`
  CreatedAt       time.Time            `json:"-"`
  UpdatedAt       time.Time            `json:"-"`
}

type FilmResponse struct {
  ID         int                  `json:"id"`
  Title      string               `json:"title"`
  ThumbnailFilm   string               `json:"thumbnailFilm"`
  Year       int                  `json:"year"`
  Category   []Category           `json:"category" gorm:"many2many:film_categories"`
  Desc       string               `json:"desc"`
}

func (FilmResponse) TableName() string {
  return "films"
}
