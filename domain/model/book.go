package model

import "time"

type StateBook string

const (
	StatePublished   StateBook = "PUBLISHED"
	StateUnpublished StateBook = "UNPUBLISHED"
	StateRetired     StateBook = "RETIRED"
	StateClosed      StateBook = "CLOSED"
)

type Author struct {
	AuthorId string `json:"authorId" bson:"authorId"`
}

type ReadingOption struct {
	Option      string `json:"option" bson:"option"`
	Description string `json:"description" bson:"description"`
}

type BookSection struct {
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}

type BookContent struct {
	Chapter  string        `json:"chapter" bson:"chapter"`
	Sections []BookSection `json:"sections" bson:"sections"`
}

type Book struct {
	Id              string          `json:"id" bson:"_id"`
	Authors         []Author        `json:"authors" bson:"authors"`
	AuthorCount     int             `json:"authorCount" bson:"authorCount"`
	Title           string          `json:"title" bson:"title"`
	AboutTheBook    string          `json:"aboutTheBook" bson:"aboutTheBook"`
	Description     string          `json:"description" bson:"description"`
	Content         []BookContent   `json:"content" bson:"content"`
	CoverImage      string          `json:"coverImage" bson:"coverImage"`
	MinimumPrice    float64         `json:"minimumPrice" bson:"minimumPrice"`
	SuggestedPrice  float64         `json:"suggestedPrice" bson:"suggestedPrice"`
	Reviews         int             `json:"reviews" bson:"reviews"`
	State           StateBook       `json:"state" bson:"state"`
	CreatedAt       time.Time       `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt" bson:"updatedAt"`
	LastPublishedAt time.Time       `json:"lastPublishedAt" bson:"lastPublishedAt"`
	LanguageName    string          `json:"languageName" bson:"languageName"`
	LanguageCode    string          `json:"languageCode" bson:"languageCode"`
	Categories      []string        `json:"categories" bson:"categories"`
	ReadingOptions  []ReadingOption `json:"readingOptions" bson:"readingOptions"`
}
