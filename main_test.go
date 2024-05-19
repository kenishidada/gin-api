package main

import (
	"bytes"
	"encoding/json"
	"gin-api/dto"
	"gin-api/infra"
	"gin-api/models"
	"gin-api/services"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		panic("Error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "test_item1", Price: 1000, Description: "test_description1", SoldOut: false, UserID: 1},
		{Name: "test_item2", Price: 2000, Description: "test_description2", SoldOut: true, UserID: 1},
		{Name: "test_item3", Price: 3000, Description: "test_description3", SoldOut: false, UserID: 1},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test_password1"},
		{Email: "test2@example.com", Password: "test_password2"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	r := setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	r.ServeHTTP(w, req)

	var res map[string][]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

func TestCreate(t *testing.T) {
	r := setup()
	// create token
	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name:        "test_item4",
		Price:       4000,
		Description: "test_description4",
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+*token)

	r.ServeHTTP(w, req)

	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}

func TestCreateUnauthorized(t *testing.T) {
	router := setup()

	createItemInput := dto.CreateItemInput{
		Name:        "テストアイテム4",
		Price:       4000,
		Description: "Createテスト",
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))

	router.ServeHTTP(w, req)

	var res map[string]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
