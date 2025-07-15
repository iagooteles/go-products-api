package controller_test

import (
	"bytes"
	"encoding/json"
	"go-api/controller"
	"go-api/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockUsecase struct{}

func (m *mockUsecase) GetProducts() ([]model.Product, error) {
	return []model.Product{
		{ID: 1, Name: "Juice", Price: 10.0},
		{ID: 2, Name: "Soda", Price: 5.0},
	}, nil
}

func (m *mockUsecase) CreateProduct(product model.Product) (model.Product, error) {
	product.ID = 1
	return product, nil
}

func (m *mockUsecase) GetProductById(id int) (*model.Product, error) {
	return nil, nil
}

func (m *mockUsecase) UpdateProduct(product model.Product) (model.Product, error) {
	return model.Product{}, nil
}

func (m *mockUsecase) DeleteProduct(id int) (bool, error) {
	return false, nil
}

func TestGetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUC := &mockUsecase{}
	ctrl := controller.NewProductController(mockUC)

	router := gin.Default()
	router.GET("/products", ctrl.GetProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var products []model.Product
	err := json.Unmarshal(resp.Body.Bytes(), &products)
	assert.Nil(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Juice", products[0].Name)
	assert.Equal(t, 10.0, products[0].Price)
	assert.Equal(t, "Soda", products[1].Name)
	assert.Equal(t, 5.0, products[1].Price)
}

func TestCreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUC := &mockUsecase{}
	controller := controller.NewProductController(mockUC)

	router := gin.Default()
	router.POST("/product", controller.CreateProduct)

	product := model.Product{
		Name:  "Juice",
		Price: 10.0,
	}
	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var created model.Product
	err := json.Unmarshal(resp.Body.Bytes(), &created)
	assert.Nil(t, err)
	assert.Equal(t, "Juice", created.Name)
	assert.Equal(t, 10.0, created.Price)
	assert.Equal(t, 1, created.ID)
}
