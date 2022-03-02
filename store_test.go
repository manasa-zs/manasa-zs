package car

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"manasa/models"
	"testing"
)

func TestDB_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an unexpected error %v", err)
	}
	defer db.Close()

	car := models.Car{Id: uuid.MustParse("cda6498a-235d-4f7e-ae19-661d41bc154d"), Name: "BMW X1", Year: 2016, Brand: "bmw", Fuel: "diesel"}

	testcases := []struct {
		desc string
		err  error
	}{
		{desc: "success case", err: nil},
		{desc: "Failure case", err: errors.New("query error")},
	}

	mock.ExpectExec("INSERT INTO car").WithArgs(car.Id, car.Name, car.Year, car.Brand, car.Fuel).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO car").WithArgs(car.Id, car.Name, car.Year, car.Brand, car.Fuel).
		WillReturnError(errors.New("query error"))

	s := New(db)

	for _, tc := range testcases {
		_, err := s.Create(car)
		assert.Equal(t, err, tc.err)

	}
}

func TestDB_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	defer db.Close()

	testcases := []struct {
		desc string
		err  error
	}{
		{"success case", nil},
		{"failure case", errors.New("query error")},
	}

	mock.ExpectExec("Select Id,Name,Year,Brand,FuelType from car where Id=?",)

}
