package rpc

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	rpc = "https://fullnode.testnet.sui.io:443"
)

func Test_sui_getObject(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockResp := loadFixture("./fixtures/sui_getObject.json")

	httpmock.RegisterResponder(
		http.MethodPost,
		rpc,
		httpmock.NewStringResponder(http.StatusOK, mockResp))

	c := NewClient(rpc)

	obj, err := c.SuiGetObject("0x845d6a0756208f107ebf6d2676641d4c28502e09dbfa9825f15baebc08c0c046")

	assert.Nil(t, err)
	assert.Equal(t, "0x845d6a0756208f107ebf6d2676641d4c28502e09dbfa9825f15baebc08c0c046", obj.ObjectID)
	assert.Equal(t, "3601626", obj.Version)
	assert.Equal(t, "8s2GGEMNkwdB9eHWATnzbDpeMSKgcGH467RmAHaQZbD8", obj.Digest)
	assert.Equal(t, "0x2::coin::Coin<0x2::sui::SUI>", obj.Type)
	assert.Equal(t, "0xfa0a634449311cec35e77298b710a3b4f45c111844636d4ff406a070e4e71443", obj.Owner.AddressOwner)
	assert.Equal(t, "HPtMBJm4xEUGExHrwoh1nfxjTX3HiG6da698N2YjrQzS", obj.PreviousTransaction)
	assert.Equal(t, "0x2::coin::Coin<0x2::sui::SUI>", obj.Content.Type)
	assert.Equal(t, "moveObject", obj.Content.DataType)

	assert.Equal(t, "0x845d6a0756208f107ebf6d2676641d4c28502e09dbfa9825f15baebc08c0c046", obj.Content.Fields.ID.ID)
	assert.Equal(t, "5000000000", obj.Content.Fields.Balance)
}

func Test_suix_getBalance(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockResp := loadFixture("./fixtures/suix_getBalance.json")

	httpmock.RegisterResponder(
		http.MethodPost,
		rpc,
		httpmock.NewStringResponder(http.StatusOK, string(mockResp)))

	c := NewClient(rpc)

	coin, err := c.SuiXGetBalance("0xb878abfe4fbd421c70c7c725a3c012bd0e70eb0f42ed0b05f0944b9616f3710d")

	assert.Nil(t, err)
	assert.Equal(t, "0x2::sui::SUI", coin.CoinType)
}

func Test_suix_getOwnedObjects(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockResp := loadFixture("./fixtures/suix_getOwnedObjects.json")

	httpmock.RegisterResponder(
		http.MethodPost,
		rpc,
		httpmock.NewStringResponder(http.StatusOK, string(mockResp)))

	c := NewClient(rpc)

	address := "0xb878abfe4fbd421c70c7c725a3c012bd0e70eb0f42ed0b05f0944b9616f3710d"
	pagedObjects, err := c.SuiXGetOwnedObjects(address)
	assert.Nil(t, err)

	assert.Equal(t, len(pagedObjects.Data), 4)

	assert.Equal(t, "0x00e4b300b42d5ef9b2b8c140f3fdf4215c0879243cadf077cf95532880eacb5f", pagedObjects.Data[0].Data.ObjectID)
	assert.Equal(t, "1814281", pagedObjects.Data[0].Data.Version)
	assert.Equal(t, "HzZrCMvFp9gmpbTchniYPa3zwVJupC1fngqCzVfb8fTC", pagedObjects.Data[0].Data.Digest)
}

func loadFixture(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(content)
}
