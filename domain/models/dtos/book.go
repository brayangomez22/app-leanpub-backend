package dtos

import (
	"leanpub-app/domain/models"
	"time"
)

type BookContentDto struct {
	Chapter  string               `json:"chapter" bson:"chapter"`
	Sections []models.BookSection `json:"sections" bson:"sections"`
}

type BookDto struct {
	Id             string                 `json:"id" bson:"_id"`
	Authors        []models.Author        `json:"authors"`
	AuthorCount    int                    `json:"authorCount" bson:"authorCount"`
	Title          string                 `json:"title" bson:"title"`
	AboutTheBook   string                 `json:"aboutTheBook" bson:"aboutTheBook"`
	Description    string                 `json:"description" bson:"description"`
	Content        []BookContentDto       `json:"content" bson:"content"`
	CoverImage     string                 `json:"coverImage" bson:"coverImage"`
	MinimumPrice   float64                `json:"minimumPrice" bson:"minimumPrice"`
	SuggestedPrice float64                `json:"suggestedPrice" bson:"suggestedPrice"`
	Reviews        int                    `json:"reviews" bson:"reviews"`
	State          models.StateBook       `json:"state" bson:"state"`
	CreatedAt      time.Time              `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time              `json:"updatedAt" bson:"updatedAt"`
	LanguageName   string                 `json:"languageName" bson:"languageName"`
	LanguageCode   string                 `json:"languageCode" bson:"languageCode"`
	Categories     []string               `json:"categories" bson:"categories"`
	ReadingOptions []models.ReadingOption `json:"readingOptions" bson:"readingOptions"`
}
