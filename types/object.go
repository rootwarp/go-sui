package types

type SuiObject struct {
	ObjectID            string `json:"objectId"`
	Version             string `json:"version"`
	Digest              string `json:"digest"`
	Type                string `json:"type"`
	Owner               Owner  `json:"owner"`
	PreviousTransaction string `json:"previousTransaction"`
	Display             struct {
		Data  interface{} `json:"data"`
		Error interface{} `json:"error"`
	} `json:"display"`
	Content Content `json:"content"`
	BCS     BCS     `json:"bcs"`
}

type Owner struct {
	AddressOwner string `json:"AddressOwner"`
}

type Content struct {
	DataType          string `json:"dataType"`
	Type              string `json:"type"`
	HasPublicTransfer bool   `json:"hasPublicTransfer"`
	Fields            struct {
		Balance string `json:"balance"`
		ID      struct {
			ID string `json:"id"`
		} `json:"id"`
	} `json:"fields"`
}

type BCS struct {
	DataType          string `json:"dataType"`
	Type              string `json:"type"`
	HasPublicTransfer bool   `json:"hasPublicTransfer"`
	Version           int    `json:"version"`
	BCSBytes          string `json:"bcsBytes"`
}
