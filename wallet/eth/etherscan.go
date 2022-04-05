package eth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BaseParamData struct {
	Url             string
	Module          string
	Action          string
	Address         string
	Contractaddress string
	Startblock      string
	Endblock        string
	Page            string
	Offset          string
	Sort            string
	Apikey          string
}

type ResultData struct {
	BlockNumber 	string 	`json:"blockNumber"`
	TimeStamp		string 	`json:"timeStamp"`
	Hash 			string 	`json:"hash"`
	From 			string 	`json:"from"`
	To 				string 	`json:"to"`
	Value 			string 	`json:"value"`
	ContractAddress string 	`json:"contractAddress"`
	Gas 			string 	`json:"gas"`
	GasUsed 		string 	`json:"gasUsed"`
	GasPrice        string  `json:"gasPrice"`
	IsError         string  `json:"isError"`
	TxreceiptStatus string  `json:"txreceipt_status"`
	Input           string  `json:"input"`
}

type TxResponse struct {
	Status   string       `json:"status"`
	Message  string       `json:"message"`
	Result   []ResultData `json:"result"`
}

func (this BaseParamData)GetTxByAddress() (*TxResponse, error) {
	reqUrl := this.Url + "?module=" + this.Module + "&action=" + this.Action + "&address=" +
		this.Address + "&startblock=" + this.Startblock + "&endblock=" + this.Endblock +
		"&page=" + this.Page + "&offset=" + this.Offset + "&sort=" + this.Sort + "&apikey=" + this.Apikey
	resp, _ := http.Get(reqUrl)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var tx_rep TxResponse
	if err := json.Unmarshal(body, &tx_rep); err != nil {
		return nil, err
	}
	return &tx_rep, nil
}

func (this BaseParamData)GetErc20TxByAddress() (*TxResponse, error) {
	reqUrl := this.Url + "?module=" +
		this.Module + "&action=" + this.Action +
		"&contractaddress=" + this.Contractaddress +
		"&address=" + this.Address +
		"&page=" + this.Page + "&offset=" +
		this.Offset + "&sort=" + this.Sort +
		"&apikey=" + this.Apikey
	fmt.Println(reqUrl)
	resp, _ := http.Get(reqUrl)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var tx_rep TxResponse
	if err := json.Unmarshal(body, &tx_rep); err != nil {
		return nil, err
	}
	return &tx_rep, nil
}