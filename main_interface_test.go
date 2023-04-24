package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:8080"

func TestGetAlbumsInterface(t *testing.T) {
	resp, err := http.Get(baseURL + "/albums")
	assert.NoError(t, err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	var albums []album
	err = json.Unmarshal(body, &albums)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, len(albums) > 0)
}

func TestGetAlbumByIDInterface(t *testing.T) {
	resp, err := http.Get(baseURL + "/albums/2")
	assert.NoError(t, err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	var album album
	err = json.Unmarshal(body, &album)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "2", album.ID)
}

func TestCreateAlbumInterface(t *testing.T) {
	newAlbum := album{
		Title:  "Test Album",
		Artist: "Test Artist",
		Price:  9.99,
	}

	jsonBytes, err := json.Marshal(newAlbum)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", baseURL+"/albums", bytes.NewBuffer(jsonBytes))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	var createdAlbum album
	err = json.Unmarshal(body, &createdAlbum)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, newAlbum.Title, createdAlbum.Title)
	assert.Equal(t, newAlbum.Artist, createdAlbum.Artist)
	assert.Equal(t, newAlbum.Price, createdAlbum.Price)
}
