package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/peersafe/poc_blacklist/apiserver/define"
	"io/ioutil"
	"net/http"
	"time"
	//"encoding/base64"
)

var (
	PushUrl string
	ChainslqUrl string
	Sqlite3DbPath string
)

//var DefaultPubKey string = "cB4Yp2i2Qv1t8o3YdXzySgJAgD41sUGE6aQWi15CRNkxA3jKj6mU"

func GetTranDateTime() define.CreatTime {
	var t define.CreatTime
	currentTime := time.Now()

	t.Year = fmt.Sprintf("%d", currentTime.Year())
	t.Month = fmt.Sprintf("%d", currentTime.Month())
	t.Day = fmt.Sprintf("%d", currentTime.Day())
	t.Hour = fmt.Sprintf("%d", currentTime.Hour())
	t.Minute = fmt.Sprintf("%d", currentTime.Minute())
	t.Second = fmt.Sprintf("%d", currentTime.Second())

	return t
}

func SendChainSqlMessage(data []byte, destUrl string) ([]byte, error) {
	PostFunc := func(data []byte) ([]byte, error) {
		req, err := http.NewRequest("POST", destUrl, bytes.NewBuffer(data))
		if err != nil {
			Log.Errorf("New post request error , %s", err.Error())
			return []byte("empty"), err
		}
		//req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			Log.Errorf("client do error , %s", err.Error())
			return []byte("empty"), err
		}
		defer resp.Body.Close()
		res, _ := ioutil.ReadAll(resp.Body)

		return res, nil
	}
	Log.Infof("push request body , %s", string(data))
	res, err := PostFunc(data)
	return res, err
}

func FormatBlackSaveData(blkList define.BlackListInfo) ([]byte, error) {
	var invokeRequest define.InvokeRequest
	t := GetTranDateTime()
	blkList.CreatTime = t
	val, _ := json.Marshal(blkList)
	invokeRequest.Value = string(val)
	invokeRequest.Key = fmt.Sprintf("%s-%s", blkList.CommData.UserId, blkList.CommData.UserName)
	return json.Marshal(invokeRequest)
}

func BlackListEncrypt(data string, key string) ([]byte, error) {
	var edReq define.EncryptDataRequest
	edReq.Method = define.CHAINSQL_ENCRYPT
	edReq.Param.Data = data
	edReq.Param.PubKey = key
	message, _ := json.Marshal(edReq)

	res, err := SendChainSqlMessage(message, PushUrl)
	if err != nil {
		return nil, fmt.Errorf("chainSql encrypt error : %s", err.Error())
	}

	return res, nil
}

func ChiansqlPay(addr string) error {
	var chainsqlPayReq define.ChainsqlPayRequest
	var payParams define.PayParams
	var transactionInfo define.Transaction
	var amountInfo define.AmountInfo

	var response define.TransferResponse

	amountInfo.Currency = "SND"
	amountInfo.Value = "10"
	amountInfo.Issuer = "zHb9CJAWyB4zj91VRWn96DkukG4bwdtyTh"
	transactionInfo.TransactionType = "Payment"
	transactionInfo.Account = "zHb9CJAWyB4zj91VRWn96DkukG4bwdtyTh"
	transactionInfo.Destination = addr
	transactionInfo.Amount = amountInfo

	payParams.Offline = false
	payParams.Secret = "xnoPBzXtMeMyMHUVTgbuqAfg1SUTb"
	payParams.Tx_json = transactionInfo
	chainsqlPayReq.Method = "submit"
	// chainsqlPayReq.Params[0] = payParams
	chainsqlPayReq.Params = append(chainsqlPayReq.Params, payParams)
	chainsqlPayReq.Id = 1
	message, err := json.Marshal(chainsqlPayReq)
	if err != nil {
		Log.Errorf("marshal request error , %s", err.Error())
		return err
	}
	res, err := SendChainSqlMessage(message, ChainslqUrl)
	if err != nil {
		Log.Errorf("sendsql message error , %s", err.Error())
		return fmt.Errorf("chainSql encrypt error : %s", err.Error())
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		Log.Errorf("unmarshal response error , %s", err.Error())
		return fmt.Errorf("chainSql unmarshal response error : %s", err.Error())
	}
	// if response.ResponseCode != http.StatusOK {
	// 	Log.Errorf("chainsql response error , errcode %d", response.ResponseCode)
	// 	return fmt.Errorf("chainsql response error , errcode %d", response.ResponseCode)
	// }

	return nil
}

// func BlackListTransfer(addr string) error {
// 	var req define.TransferRequest
// 	var response define.TransferResponse
// 	req.Method = define.CHAINSQL_TRANSFERACCOUNT
// 	req.Param.DestAddress = addr
// 	message, _ := json.Marshal(req)

// 	res, err := SendChainSqlMessage(message, PushUrl)
// 	if err != nil {
// 		Log.Errorf("sendsql message error , %s", err.Error())
// 		return fmt.Errorf("chainSql encrypt error : %s", err.Error())
// 	}

// 	err = json.Unmarshal(res, &response)
// 	if err != nil {
// 		Log.Errorf("unmarshal response error , %s", err.Error())
// 		return fmt.Errorf("chainSql unmarshal response error : %s", err.Error())
// 	}

// 	if response.ResponseCode != http.StatusOK {
// 		Log.Errorf("chainsql response error , errcode %d", response.ResponseCode)
// 		return fmt.Errorf("chainsql response error , errcode %d", response.ResponseCode)
// 	}

// 	return nil
// }

func FormatBlackSaveDataByCrpto(blkList define.BlackListInfo) ([]byte, error) {
	var invokeRequest define.InvokeRequest
	t := GetTranDateTime()
	blkList.CreatTime = t
	/*
	   key := GenerateKey(32)
	   cryptoData, err := AesEncrypt([]byte(blkList.SpecialData), key)
	   if err != nil{
	       return nil, fmt.Errorf("AesEncrypt data error : %s", err.Error())
	   }
	   blkList.SpecialData = base64.StdEncoding.EncodeToString(cryptoData)
	   path := define.CRYPTO_PATH + "enrollment.cert"
	   cert, err := ReadFile(path)
	   if err != nil{
	       return nil, fmt.Errorf("AesEncrypt data error : %s", err.Error())
	   }

	   cryptoKey, err := EciesEncrypt(key, cert)
	   if err != nil {
	       return nil, fmt.Errorf("EciesEncrypt  key error : %s", err.Error())
	   }
	   blkList.EncryKey = base64.StdEncoding.EncodeToString(cryptoKey)*/
	chanRes, err := BlackListEncrypt(blkList.SpecialData, blkList.CommData.PaymentPubKey)
	if err == nil {
		var edRes define.EncryptDataResponse
		err1 := json.Unmarshal(chanRes, &edRes)
		if err1 != nil {
			return nil, fmt.Errorf("chainSql encrypt error : %s", err1.Error())
		}
		blkList.SpecialData = edRes.EncryptData
		blkList.EncryKey = edRes.EncryptKey
	}

	val, _ := json.Marshal(blkList)
	invokeRequest.Value = string(val)
	invokeRequest.Key = fmt.Sprintf("%s-%s", blkList.CommData.UserId, blkList.CommData.UserName)
	return json.Marshal(invokeRequest)
}

func FormatMaterialsVerify(m define.MaterialsInfo) define.VerifiedBase {
	var vb define.VerifiedBase
	vb.OrgID = m.OrgInfo.OrgID
	vb.Orgname = m.OrgInfo.EnterpriseInfo.Orgname
	vb.CreateTime = m.OrgInfo.CreateTime
	vb.VerifiedDate = m.VerifiedDate
	vb.UploadTimes = m.UploadTimes
	vb.QueryState = m.QueryState
	vb.Suggestion = m.Suggestion
	return vb
}
