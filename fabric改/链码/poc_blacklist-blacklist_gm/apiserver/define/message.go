package define

type QueryContents struct {
	Payload interface{} `json:"payload"`
}

type ParamS  struct {
    Data      string     `json:"Data"`
    PubKey    string     `json:"PubKey"`
}

type EncryptDataRequest struct {
    Method   string     `json:"method"`
    Param    ParamS     `json:"param"`
}

type EncryptDataResponse struct {
    EncryptData   string     `json:"EncryptData"`    
    EncryptKey    string     `json:"EncryptKey"`     
}

type AddrInfo struct {
    DestAddress   string     `json:"destAddress"`
}

type TransferRequest struct {
    Method   string     `json:"method"`
    Param    AddrInfo     `json:"param"`
}

// type TransferResponse struct {
//     ResponseCode        int       `json:"ResponseCode"`    
//     ResponseMsg         string    `json:"ResponseMsg"` 
// }

type Amount struct {
	Currency 	string 		`json:"currency"`
	Issuer 		string 		`json:"issuer"`
	Value 		string 		`json:"value"`
}

type Tx_json struct {
	Account             string    `json:"Account"`
	Amount              Amount    `json:"Amount"`
	Destination         string    `json:"Destination"`
	Fee                 string    `json:"Fee"`
	Flags               int64     `json:"Flags"`
	Sequence            int64     `json:"Sequence"`
	SigningPubKey       string    `json:"SigningPubKey"`
	TransactionType     string    `json:"TransactionType"`
	TxnSignature        string    `json:"TxnSignature"`
	Hash                string    `json:"hash"`
}

type TransferResult struct {
	Engine_result            string     `json:"engine_result"`
	Engine_result_code       int64      `json:"engine_result_code"`
	Engine_result_message    string     `json:"engine_result_message"`
	Status                   string     `json:"status"`
	Tx_blob                  string     `json:"tx_blob"`
	Tx_json                  Tx_json    `json:"tx_json"`
}

type TransferResponse struct {
    Id        int       		`json:"id"`    
    Result    TransferResult    `json:"result"` 
}

type AmountInfo struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
	Issuer   string `json:"issuer"`
}

type Transaction struct {
	TransactionType string     `json:"TransactionType"`
	Account         string     `json:"Account"`
	Destination     string     `json:"Destination"`
	Amount          AmountInfo `json:"Amount"`
}

type PayParams struct {
	Offline bool        `json:"offline"`
	Secret  string      `json:"secret"`
	Tx_json Transaction `json:"tx_json"`
}

type ChainsqlPayRequest struct {
	Method string      `json:"method"`
	Params []PayParams `json:"params"`
	Id     uint32      `json:"id"`
}

type PayResult struct {
	EngineResult     string `json:"engine_result"`
	EngineResultCode int32  `json:"engine_result_code"`
	Status           string `json:"status"`
	TXBlob           string `json:"tx_blob"`
}

type ChainsqlPayResponse struct {
	Id     uint32    `json:"id"`
	Result PayResult `json:"result"`
}
