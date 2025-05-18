package serializers

import (
	"crud-go/models"
	"time"
)

type AuthorResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	YearBorn int    `json:"year_born"`
	UserID   uint   `json:"user_id"`
}

func ToAuthorResponse(author models.Author) AuthorResponse {
	currentYear := time.Now().Year()
	yearBorn := currentYear - author.Age
	return AuthorResponse{
		ID:       author.ID,
		Name:     author.Name,
		YearBorn: yearBorn,
		UserID:   author.UserID,
	}
}

func ToAuthorResponseList(authors []models.Author) []AuthorResponse {
	var responses []AuthorResponse
	for _, author := range authors {
		responses = append(responses, ToAuthorResponse(author))
	}
	return responses
}
