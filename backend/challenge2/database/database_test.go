package database_test

import (
	"testing"

	"github.com/Clarilab/codechallenge/backend/challenge2/database"
	"github.com/Clarilab/codechallenge/backend/challenge2/models"
)

func Test_Repository_GetByName(t *testing.T) {
	repo := database.NewRepository()

	result := repo.GetByName("finux")

	if result == nil {
		t.Error("Expected to find any result")
	}

	if len(result) != 2 {
		t.Error("Expected to find two results")
	}
}

func Test_Repository_Get(t *testing.T) {
	repo := database.NewRepository()

	result := repo.Get(models.SearchRequest{
		CompanyName: "finux",
	})

	if result == nil {
		t.Error("Expected to find any result")
	}

	if len(result) != 2 {
		t.Error("Expected to find two results")
	}
}
