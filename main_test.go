package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAlbumsHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums", getAlbums)
	req, err := http.NewRequest("GET", "/albums", nil)
	if err != nil {
		t.Fatalf("failed to construct request, err: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var albums []album
	if err := json.Unmarshal(w.Body.Bytes(), &albums); err != nil {
		t.Fatalf("failed to unmarshal response, err: %v", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, albums)
}

func TestGetAlbumHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/album/:id", getAlbumByID)
	req, err := http.NewRequest("GET", "/album/2", nil)
	if err != nil {
		t.Fatalf("failed to construct request, err: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var albumData album
	if err := json.Unmarshal(w.Body.Bytes(), &albumData); err != nil {
		t.Fatalf("failed to unmarshal response, err: %v", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Jeru", albumData.Title)
	assert.Equal(t, "2", albumData.ID)
}

func TestNewAlbumHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/albums", postAlbums)
	albumID := xid.New().String()
	newAlbum := album{
		ID:     albumID,
		Title:  "Demo Song",
		Artist: "Demo Singer",
		Price:  22.25,
	}
	data, err := json.Marshal(newAlbum)
	if err != nil {
		t.Fatalf("failed to marshal payload, err: %v", err)
	}
	req, err := http.NewRequest("POST", "/albums", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("failed to construct request, err: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
