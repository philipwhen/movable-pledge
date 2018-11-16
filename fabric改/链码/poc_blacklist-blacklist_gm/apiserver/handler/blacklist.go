package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
    "os/exec"
    "os"
    "strconv"

	"github.com/peersafe/poc_blacklist/apiserver/define"
	"github.com/peersafe/poc_blacklist/apiserver/sdk"
	"github.com/peersafe/poc_blacklist/apiserver/utils"

	"github.com/gin-gonic/gin"
)

var targetOrderAddr string

func Uploadblacklist(c *gin.Context) {
    utils.Log.Debug("Uploadblacklist .....")

    var blklist []define.BlackListInfo
    var response define.BlackListResponse
    var err error

    status := http.StatusOK

    body, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        utils.Log.Errorf("SaveData read body : %s", err.Error())
        status = http.StatusNoContent
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = err.Error()
        utils.Response(response, c, http.StatusNoContent)
        return
    }

    utils.Log.Debugf("SaveData header : %v", c.Request.Header)
    utils.Log.Debugf("SaveData body : %s", string(body))

    if err = json.Unmarshal(body, &blklist); err != nil {
        utils.Log.Errorf("SaveData Unmarshal : %s", err.Error())
        status = http.StatusBadRequest
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = err.Error()
        utils.Response(response, c, http.StatusBadRequest)
        return
    }

    for _, blklist := range blklist {
        txId, nonce, err := sdk.Handler.GetTxId()
        if err != nil {
            utils.Log.Errorf("SaveData GetTxId : %s", err.Error())
            status = http.StatusServiceUnavailable
            response.ResponseCode = strconv.Itoa(status)
            response.ResponseMsg = err.Error()
            utils.Response(response, c, status)
            return
        }

        blklist.CommData.FabricTxId = txId
        // b, err := utils.FormatBlackSaveData(blklist)
        b, err := utils.FormatBlackSaveDataByCrpto(blklist)
        if err != nil {
            utils.Log.Errorf("SaveData FormatRequestMessage : %s", err.Error())
            status = http.StatusServiceUnavailable
            response.ResponseCode = strconv.Itoa(status)
            response.ResponseMsg = err.Error()
            utils.Response(response, c, status)
            return
        }

        // invoke
        _, err = sdk.Handler.Invoke(txId, nonce, "", define.UPLOAD_BLACKLIST, nil, b)
        if err != nil {
            utils.Log.Errorf("SaveData Invoke : %s", err.Error())
            status = http.StatusServiceUnavailable
            response.ResponseCode = strconv.Itoa(status)
            response.ResponseMsg = err.Error()
            utils.Response(response, c, status)
            return
        }
        response.FabricID = append(response.FabricID, txId)

        // err = utils.BlackListTransfer(blklist.CommData.PaymentAddr)
        // err = utils.ChiansqlPay(blklist.CommData.PaymentAddr)
        // if err != nil {
        //     utils.Log.Errorf("transfer intergral error : %s", err.Error())
        //     response.ResponseCode = strconv.Itoa(define.BLACK_TRANSFER_ERR)
        //     response.ResponseMsg = err.Error()
        //     utils.Response(response, c, status)
        //     return
        // }
    }
    
    response.ResponseCode = strconv.Itoa(status)
    response.ResponseMsg = "submit success!"
    utils.Response(response, c, status)
}

func QueryBlackListUnpay(c *gin.Context) {
    utils.Log.Debug("QueryBlackListUnpay .....")

    var request define.BlackListQueryUnpayRequest
    var response define.BlackListQueryUnpayResponse
    var blkListKeyData define.BlacklistKeyData
    var err error
    var errMsg string

    // query
    status := http.StatusOK

    utils.Log.Debugf("query header : %v", c.Request.Header)

    body, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        utils.Log.Errorf("QueryBlackListUnpay read body : %s", err.Error())
        status = http.StatusNoContent
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = err.Error()
        utils.Response(errMsg, c, http.StatusNoContent)
        return
    }
    utils.Log.Debugf("query body : %s", string(body))

    if err = json.Unmarshal(body, &request); err != nil {
        utils.Log.Errorf("QueryBlackListUnpay Unmarshal : %s", err.Error())
        status = http.StatusBadRequest
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = err.Error()
        utils.Response(response, c, http.StatusBadRequest)
        return
    }
    b := fmt.Sprintf("%s-%s", request.UserID, request.UserName)
    responseData, _, err := sdk.Handler.QueryData("", define.QUERY_DATA, nil, b)
    if err != nil {
        utils.Log.Errorf("QueryBlackListUnpay Query : %s", err.Error())
        status = http.StatusServiceUnavailable
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = err.Error()
        utils.Response(response, c, http.StatusBadRequest)
        return
    }

    if responseData.Payload == nil {
        status = http.StatusNotFound
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = "The specified block was not found"
        utils.Response(response, c, http.StatusNotFound)
        return
    }
    if err = json.Unmarshal(responseData.Payload.([]byte), &blkListKeyData); err != nil {
        utils.Log.Errorf("QueryBlackListUnpay Unmarshal : %s", err.Error())
        status = http.StatusBadRequest
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = err.Error()
        utils.Response(response, c, http.StatusBadRequest)
        return
    }

    for _, blkList := range blkListKeyData.BlkLists {
        response.Payload = append(response.Payload, blkList)
    }
    
    response.ResponseCode = strconv.Itoa(status)
    response.ResponseMsg = "query success!"
    utils.Response(response, c, status)
}

/* ##############################################
* description: 从本地数据库读取黑名单统计信息
* input:       无
* output:      黑名单统计信息
* ###############################################*/
func QueryBlackListTotal(c *gin.Context) {
    utils.Log.Debug("query blacklist statistic info .....")
    var response define.BlackListQueryTotalCntResponse
    status := http.StatusOK

    // dbFile := "./eventserver/blacklist.db"  //sqlite3数据库名字
    dbFile := utils.Sqlite3DbPath
	dbFileExist, err := utils.FileOrDirectoryExist(dbFile)
	if err != nil {
        utils.Log.Errorf("check file exist or not error, %s", err.Error())
        status = http.StatusServiceUnavailable
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = err.Error()
        utils.Response(response, c, http.StatusBadRequest)
		return
	}
	if !dbFileExist {
		_, err := os.Create(dbFile)
		if err != nil {
            utils.Log.Errorf("create dbfile error,  %s", err.Error())
            status = http.StatusServiceUnavailable
            response.ResponseCode = strconv.Itoa(status)
            response.ResponseMsg = err.Error()
            utils.Response(response, c, http.StatusBadRequest)
            return
		}
	}
	d, err := utils.ConnectDB("sqlite3", dbFile)
	if err != nil {
        utils.Log.Errorf("connectdb err, %s", err.Error())
		status = http.StatusServiceUnavailable
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = err.Error()
        utils.Response(response, c, http.StatusBadRequest)
		return
    }
    defer d.DisConnectDB()
	if !dbFileExist {
        err := d.CreateTable()
        if err != nil {
            utils.Log.Errorf("create table err, %s", err.Error())
            status = http.StatusServiceUnavailable
            response.ResponseCode = strconv.Itoa(status)
            response.ResponseMsg = err.Error()
            utils.Response(response, c, http.StatusBadRequest)
            return
        }
	}
    blacklistCntList, err := d.QueryTable()
    if err != nil {
        utils.Log.Errorf("query table err, %s", err.Error())
        status = http.StatusServiceUnavailable
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = err.Error()
        utils.Response(response, c, http.StatusBadRequest)
        return
    }
    blacklistCntType := map[string]int{"1":0,"2":0,"3":0,"4":0,"5":0,"6":0,"7":0}
    for _, blacklistCnt := range blacklistCntList {
        if blacklistCnt.Type == define.BLACKLIST_TOTAL_COUNT {
            response.Payload.Total = blacklistCnt.ListCnt
            continue
        }
        _, exists := blacklistCntType[blacklistCnt.Type]
        if exists {
            var btc define.BlackTypeCnt
            btc.ListType = blacklistCnt.Type
            btc.ListCnt = blacklistCnt.ListCnt
            response.Payload.TypeCount = append(response.Payload.TypeCount, btc)
            continue
        }
        t := utils.GetTranDateTime()
        currentMonType := fmt.Sprintf("%s-%s", t.Year, t.Month)
        if blacklistCnt.Type == currentMonType {
            response.Payload.CurMonthCount = blacklistCnt.ListCnt
        }
    }

    response.ResponseCode = strconv.Itoa(status)
    response.ResponseMsg = "count success!"
    utils.Response(response, c, status)
}

func QueryBlackListTotal_bak(c *gin.Context) {
    utils.Log.Debug("query target bill .....")

    var request define.QueryRequest
    var response define.BlackListQueryTotalCntResponse
    var err error
    var totalCnt uint64
    var currentMonCnt uint64
    var typeCnt map[string]uint64

    // query
    status := http.StatusOK
    request.DslSyntax = fmt.Sprintf("{\"selector\":{\"DataType\": %d}}", define.DATATYPE_BLACKLIST)
    b, _ := json.Marshal(request)
    responseData, _, err := sdk.Handler.QueryDslData("", define.DSL_QUERY, nil, string(b))
    if err != nil {
        utils.Log.Errorf("DslQuery Query : %s", err.Error())
        status = http.StatusServiceUnavailable
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = err.Error()
        utils.Response(response, c, http.StatusBadRequest)
        return
    }

    t := utils.GetTranDateTime()

    blklists, ok := responseData.Payload.([]string)
    if ok {
        typeCnt = make(map[string]uint64)
        for _, jsonVal := range blklists {
            var blkKeyData define.BlacklistKeyData
            err = json.Unmarshal([]byte(jsonVal), &blkKeyData)
            if err == nil {
                for _, blklst := range blkKeyData.BlkLists {
                    if blklst.CreatTime.Year == t.Year && 
                       blklst.CreatTime.Month == t.Month {
                        currentMonCnt += 1
                    }
                    totalCnt += 1
                    lstType := blklst.CommData.ListType
                    typeCnt[lstType] += 1
                }
            } else {
                utils.Log.Errorf("Unmarshal err: %s", err.Error())
            }
        }

        for lstTp, tCnt := range typeCnt {
            var btc define.BlackTypeCnt
            btc.ListType = lstTp
            btc.ListCnt = tCnt
            response.Payload.TypeCount = append(response.Payload.TypeCount, btc)
        }

        response.Payload.Total = totalCnt
        response.Payload.CurMonthCount = currentMonCnt
        response.ResponseCode = strconv.Itoa(status)
        response.ResponseMsg = "count success!"
    } 

    utils.Response(response, c, status)
}

func KeepaliveQuery(c *gin.Context) {
	status := http.StatusOK

	if !sdk.Handler.PeerKeepalive(define.KEEPALIVE_QUERY) {
		status = define.PEER_FAIL_CODE
		utils.Log.Errorf("peer cann't be reached.")
	} else if !OrderKeepalive() {
		status = define.ORDER_FAIL_CODE
		utils.Log.Errorf("order cann't be reached.")
	}

	utils.Response(nil, c, status)
}

func OrderKeepalive() bool {
	//use nc command to detect whether the order's port is available
	orderCommand := fmt.Sprintf("nc -v %s", targetOrderAddr)
	cmd := exec.Command("/bin/bash", "-c", orderCommand)
	err := cmd.Run()
	if nil != err {
		utils.Log.Errorf("Order(%s) cann't be reached: %s", targetOrderAddr, err.Error())
		return false
	} else {
		return true
	}
}

func SetOrderAddrToProbe(addr string) bool {
	if addr == "" {
		utils.Log.Error("order address to be Probed is null!")
		return false
	}

	targetOrderAddr = addr
	utils.Log.Error("order address to be Probed is", targetOrderAddr)

	return true
}
