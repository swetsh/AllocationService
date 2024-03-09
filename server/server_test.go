package main

// import (
// 	pb "CurrencyConverterService/converter"
// 	"CurrencyConverterService/pkg/models"
// 	"CurrencyConverterService/service"
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/mock"
// 	"gorm.io/gorm"
// )

// type MockDatabase struct {
// 	mock.Mock
// }

// func (m *MockDatabase) DatabaseConnection() error {
// 	return nil
// }

// func (m *MockDatabase) GetDB() *gorm.DB {
// 	args := m.Called()
// 	return args.Get(0).(*gorm.DB)
// }

// func TestConvertCurrency(t *testing.T) {
// 	mockDB := new(MockDatabase)

// 	fromCurrency := models.Currency{
// 		Currency:     "USD",
// 		ExchangeRate: 1.0,
// 	}
// 	toCurrency := models.Currency{
// 		Currency:     "EUR",
// 		ExchangeRate: 0.85,
// 	}

// 	mockDB.On("Where", "currency = ?", "USD").Return(mockDB)
// 	mockDB.On("Find", &models.Currency{}).Return(nil)
// 	mockDB.On("Where", "currency = ?", "EUR").Return(mockDB)
// 	mockDB.On("Find", &models.Currency{}).Return(nil)

// 	srv := &service.Server{
// 		DBManager: mockDB,
// 	}

// 	ctx := context.Background()

// 	req := &pb.ConversionRequest{
// 		FromCurrency: "USD",
// 		ToCurrency:   "EUR",
// 		Amount:       100,
// 	}

// 	resp, err := srv.ConvertCurrency(ctx, req)
// 	if err != nil {
// 		t.Errorf("ConvertCurrency returned error: %v", err)
// 	}

// 	expectedConvertedAmount := (req.Amount / toCurrency.ExchangeRate) * fromCurrency.ExchangeRate
// 	if resp.ConvertedAmount != expectedConvertedAmount {
// 		t.Errorf("Expected converted amount: %f, got: %f", expectedConvertedAmount, resp.ConvertedAmount)
// 	}
// }
