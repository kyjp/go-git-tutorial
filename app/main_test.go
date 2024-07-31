package main

import "testing"

import (
	"web/app/models"
	"web/app/infra"
	"net/http/httptest"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setuoTestData(db *gorm.DB) {
	items := []models.Item {
		{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, UserId: 1},
		{Name: "テストアイテム2", Price: 2000, Description: "テスト2", SoldOut: true, UserId: 2},
		{Name: "テストアイテム3", Price: 3000, Description: "テスト3", SoldOut: false, UserId: 3},
	}

	users := []models.User {
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
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
	db.AutoMigrate(&models.Item{}, $models.User{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	router.ServerHttp(w, req)

	var res map[string][]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

func TestCreate(t *testing.T) {
	router := setup()

	token, err := services.CreateToken(1, "test1@example.com")
	assert.Equal(t, nil, err)

	createItemInput := dto.CreateItemInput{
		Name: "テストアイテム4",
		Price: 4000,
		Description: "Createテスト"
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer " + *token)


	router.ServerHttp(w, req)

	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, uint(4), res["data"].ID)
}

func TestCreateUnauthorized(t *testing.T) {
	router := setup()

	createItemInput := dto.CreateItemInput{
		Name: "テストアイテム4",
		Price: 4000,
		Description: "Createテスト"
	}
	reqBody, _ := json.Marshal(createItemInput)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(reqBody))


	router.ServerHttp(w, req)

	var res map[string]models.Item
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
