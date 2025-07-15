package repository_test

import (
	"go-api/model"
	"go-api/repository"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "product_name", "price"}).
		AddRow(1, "Produto A", 9.99).
		AddRow(2, "Produto B", 19.99)

	mock.ExpectQuery("SELECT id, product_name, price FROM product").
		WillReturnRows(rows)

	repo := repository.NewProductRepository(db)
	products, err := repo.GetProducts()

	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Produto A", products[0].Name)
	assert.Equal(t, 19.99, products[1].Price)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewProductRepository(db)

	product := model.Product{
		Name:  "Café",
		Price: 12.5,
	}

	mock.ExpectPrepare("INSERT INTO product").
		ExpectQuery().
		WithArgs(product.Name, product.Price).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := repo.CreateProduct(product)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}