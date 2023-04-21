package rpc

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_suix_getBalance(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	f, err := os.Open("./fixtures/suix_getBalance.json")
	assert.Nil(t, err)

	mockResp, err := io.ReadAll(f)
	assert.Nil(t, err)

	rpc := "https://fullnode.testnet.sui.io:443"
	httpmock.RegisterResponder(
		http.MethodPost,
		rpc,
		httpmock.NewStringResponder(http.StatusOK, string(mockResp)))

	c := NewClient(rpc)

	coin, err := c.SuiXGetBalance("0xb878abfe4fbd421c70c7c725a3c012bd0e70eb0f42ed0b05f0944b9616f3710d")

	assert.Nil(t, err)
	assert.Equal(t, "0x2::sui::SUI", coin.CoinType)
}
