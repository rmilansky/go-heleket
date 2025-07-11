package tests

import (
	"github.com/rmilansky/go-heleket"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateStaticWallet(t *testing.T) {
	staticWalletReq := &heleket.StaticWalletRequest{
		Currency: "TRX",
		Network:  "tron",
		OrderId:  "xxx",
		StaticWalletRequestOptions: &heleket.StaticWalletRequestOptions{
			UrlCallback: "https://example.com/heleket/callback",
		},
	}

	staticWallet, err := TestHeleket.CreateStaticWallet(staticWalletReq)
	require.NoError(t, err)
	require.NotEmpty(t, staticWallet)
}
