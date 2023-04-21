package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/rootwarp/go-sui/types"
)

type Client interface {
	SuiXGetBalance(address string) (*types.Coin, error)
}

type client struct {
	rpc string
}

// TODO: Should support another cointype.
func (c *client) SuiXGetBalance(address string) (*types.Coin, error) {
	log.Info().Msg("GetBalance")

	reqBody := struct {
		JsonRpc string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Method  string `json:"method"`
		Params  []any  `json:"params"`
	}{
		JsonRpc: "2.0",
		ID:      1,
		Method:  "suix_getBalance",
		Params:  []any{address},
	}

	bodyData, err := json.Marshal(&reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.rpc, bytes.NewReader(bodyData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	cli := http.Client{}

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed with status %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respData := struct {
		JsonRpc string     `json:"jsonrpc"`
		ID      int        `json:"id"`
		Result  types.Coin `json:"result"`
	}{}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		return nil, err
	}

	return &respData.Result, nil
}

// NewClient returns new rpc client.
func NewClient(rpc string) Client {
	return &client{rpc: rpc}
}
