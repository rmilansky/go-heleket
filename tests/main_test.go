package tests

import (
	"github.com/rmilansky/go-heleket"
	"net/http"
	"os"
	"testing"
)

var TestHeleket *heleket.Heleket

func TestMain(m *testing.M) {
	httpClient := http.Client{}
	merchant := "replace with your merchant id"
	paymentAPIKey := "replace with your payment API key"
	payoutAPIKey := "replace with your payout API key"

	TestHeleket = heleket.New(&httpClient, merchant, paymentAPIKey, payoutAPIKey)

	os.Exit(m.Run())
}
