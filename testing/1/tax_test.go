package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.00
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount   float64
		expected float64
	}

	table := []calcTax{
		{1000.00, 5.0},
		{1000.00, 5.0},
		{500.00, 5.0},
		{2000.00, 5.0},
		{6000.00, 5.0},
		{10000.00, 5.0},
		{20000.00, 10.00},
		{22000.00, 10.00},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.00)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.00)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-2.0, -8.0, -5, 2.5, 500.00, 1000.00, 1501.00, 2000.00}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
	})
}

//using testify

func TestCalculateTaxTestify(t *testing.T) {
	tax := CalculateTax(1000.00)
	assert.Equal(t, 5.0, tax)
}

// using mocks
type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("Error saving Tax"))

	err := CalculateTaxAndSave(20000.00, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "Error saving Tax")

	//Spy para identificar quantas vezes foram chamado o método.
	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
