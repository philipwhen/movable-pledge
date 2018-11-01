package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/sifbc/pledge/apiserver/define"
	"github.com/sifbc/pledge/apiserver/sdk"
	"github.com/sifbc/pledge/apiserver/utils"

	"github.com/gin-gonic/gin"
)

var targetOrderAddr string

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

func IfMatrialeRange(dataTime uint64, fromTime uint64, toTime uint64) bool {
	if toTime == 0 {
		if dataTime >= fromTime {
			return true
		}
	} else {
		if dataTime >= fromTime && dataTime <= toTime {
			return true
		}
	}
	return false
}

// 将货物信息上传到区块链上
func UploadGoodsInfo(c *gin.Context) {
	utils.Log.Debug("upload goods info .....")
	var GoodsInfo define.GoodsInfo
	var response define.UploadMaterialsResponse

	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("UploadApplicationMaterials read body : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusNoContent)
		return
	}

	utils.Log.Debugf("Upload header : %v", c.Request.Header)
	utils.Log.Debugf("body : %s", string(body))

	if err = json.Unmarshal(body, &GoodsInfo); err != nil {
		utils.Log.Errorf("Unmarshal : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusBadRequest)
		return
	}
	GoodsInfo.PledgePatrolDetailInfo = make([]define.PledgePatrolDetailInfo, 0)
	txId, nonce, err := sdk.Handler.GetTxId()
	if err != nil {
		utils.Log.Errorf("GetTxId : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	var invokeRequest define.InvokeRequest
	val, _ := json.Marshal(GoodsInfo)
	invokeRequest.Value = string(val)
	invokeRequest.Key = fmt.Sprintf("%s", GoodsInfo.PledgeGeneralInfo.PledgeNoStorage)

	b, _ := json.Marshal(invokeRequest)

	// invoke
	_, err = sdk.Handler.Invoke(txId, nonce, "", define.UPLOAD_GOODS_INFO, nil, b)
	if err != nil {
		utils.Log.Errorf("SaveData Invoke : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "submit success!"
	utils.Response(response, c, status)
}

// 上传巡库信息
func UploadPatrolInfo(c *gin.Context) {
	utils.Log.Debug("upload patrol info .....")
	var info define.PledgePatrolDetailInfo
	var response define.UploadMaterialsResponse

	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("UploadApplicationMaterials read body : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusNoContent)
		return
	}
	utils.Log.Debugf("Upload header : %v", c.Request.Header)
	utils.Log.Debugf("body : %s", string(body))

	if err = json.Unmarshal(body, &info); err != nil {
		utils.Log.Errorf("Unmarshal : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusBadRequest)
		return
	}

	txId, nonce, err := sdk.Handler.GetTxId()
	if err != nil {
		utils.Log.Errorf("GetTxId : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	var invokeRequest define.InvokeRequest
	val, _ := json.Marshal(info)
	invokeRequest.Value = string(val)
	invokeRequest.Key = fmt.Sprintf("%s", info.Realtive_pledge_no)

	utils.Log.Debugf("value : %s", string(val))
	utils.Log.Debugf("key : %s", info.Realtive_pledge_no)

	b, _ := json.Marshal(invokeRequest)

	// invoke
	_, err = sdk.Handler.Invoke(txId, nonce, "", define.UPLOAD_PATROL_INFO, nil, b)
	if err != nil {
		utils.Log.Errorf("SaveData Invoke : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "submit success!"
	utils.Response(response, c, status)
}

// 查询货物信息
func QueryGoodsInfo(c *gin.Context) {
	utils.Log.Debug("QueryApplicationMaterials .....")

	var request define.QueryGoodsRequest
	var response define.QueryGoodsResponse
	var result define.GoodsInfo

	var err error

	// query
	status := http.StatusOK

	utils.Log.Debugf("query header : %v", c.Request.Header)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("read body : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusNoContent)
		return
	}
	utils.Log.Debugf("query body : %s", string(body))

	if err = json.Unmarshal(body, &request); err != nil {
		utils.Log.Errorf("Unmarshal : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusBadRequest)
		return
	}

	b := request.PledgeNoStorage
	responseData, _, err := sdk.Handler.QueryData("", define.QUERY_DATA, nil, b)
	if err != nil {
		utils.Log.Errorf("Query : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusServiceUnavailable)
		return
	}
	if responseData.Payload == nil {
		status = http.StatusNotFound
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = "The specified block was not found"
		utils.Response(response, c, http.StatusNotFound)
		return
	}

	if err = json.Unmarshal(responseData.Payload.([]byte), &result); err != nil {
		utils.Log.Errorf("QueryBlackListPay Unmarshal : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusBadRequest)
		return
	}

	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "query success!"
	response.GoodsInfo = result
	utils.Response(response, c, status)
}

// 同步押品状态
func StatusSync(c *gin.Context) {
	utils.Log.Debug("synchronize goods status .....")
	var request define.StatusSyncRequest
	var response define.UploadMaterialsResponse

	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("UploadApplicationMaterials read body : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusNoContent)
		return
	}

	utils.Log.Debugf("Upload header : %v", c.Request.Header)
	utils.Log.Debugf("body : %s", string(body))

	if err = json.Unmarshal(body, &request); err != nil {
		utils.Log.Errorf("Unmarshal : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusBadRequest)
		return
	}

	txId, nonce, err := sdk.Handler.GetTxId()
	if err != nil {
		utils.Log.Errorf("GetTxId : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	var invokeRequest define.InvokeRequest
	val, _ := json.Marshal(request)
	invokeRequest.Value = string(val)
	invokeRequest.Key = fmt.Sprintf("%s", request.PledgeNoStorage)

	b, _ := json.Marshal(invokeRequest)

	// invoke
	_, err = sdk.Handler.Invoke(txId, nonce, "", define.STATUS_SYNC, nil, b)
	if err != nil {
		utils.Log.Errorf("SaveData Invoke : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "submit success!"
	utils.Response(response, c, status)
}

// 上传预警信息
func UploadWarningInfo(c *gin.Context) {
	utils.Log.Debug("upload warning info .....")
	var msg define.PledgeWarningMsg
	var info define.PledgeWarningInfo
	var response define.UploadMaterialsResponse
	var result define.GoodsInfo

	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("UploadApplicationMaterials read body : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusNoContent)
		return
	}
	utils.Log.Debugf("Upload header : %v", c.Request.Header)
	utils.Log.Debugf("body : %s", string(body))

	if err = json.Unmarshal(body, &msg); err != nil {
		utils.Log.Errorf("Unmarshal : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusBadRequest)
		return
	}

	b := msg.MonitorEquipmentId
	responseData, _, err := sdk.Handler.QueryData("", define.QUERY_DATA, nil, b)
	if err != nil {
		utils.Log.Errorf("Query : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusServiceUnavailable)
		return
	}
	if responseData.Payload == nil {
		status = http.StatusNotFound
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = "The specified block was not found"
		utils.Response(response, c, http.StatusNotFound)
		return
	}

	s := responseData.Payload.([]byte)
	responseData, _, err = sdk.Handler.QueryData("", define.QUERY_DATA, nil, string(s))
	if err != nil {
		utils.Log.Errorf("Query : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusServiceUnavailable)
		return
	}
	if responseData.Payload == nil {
		status = http.StatusNotFound
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = "The specified block was not found"
		utils.Response(response, c, http.StatusNotFound)
		return
	}

	if err = json.Unmarshal(responseData.Payload.([]byte), &result); err != nil {
		utils.Log.Errorf("QueryBlackListPay Unmarshal : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, http.StatusBadRequest)
		return
	}

	generalInfo := result.PledgeGeneralInfo
	info.ChannelSeq = msg.ChannelSeq
	info.PledgeNoStorage = generalInfo.PledgeNoStorage
	info.PledgeName = generalInfo.PledgeName
	info.SocialCreditCode = generalInfo.SocialCreditCode
	info.CoreEnterpriseName = result.PledgeDetailInfo.CoreEnterpriseName
	info.WarningMsg = fmt.Sprintf("货物%s在%s内有%s告警，告警内容为%s，告警来源是设备%s上触发的规则%s。触发时间为%s。", info.PledgeNoStorage, info.CoreEnterpriseName, msg.WarningType, msg.WarningContents, msg.MonitorEquipmentId, msg.RuleId, msg.WarningTime)
	utils.Log.Debugf("WarningMsg of info : %s", info.WarningMsg)

	dbFile := "./pledge.db" //sqlite3数据库名字
	dbFileExist, err := utils.FileOrDirectoryExist(dbFile)
	if err != nil {
		utils.Log.Errorf("check file exist or not error, %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	if !dbFileExist {
		_, err := os.Create(dbFile) // 创建数据库
		if err != nil {
			utils.Log.Errorf("create dbfile error,  %s", err.Error())
			status = http.StatusServiceUnavailable
			response.ResponseCode = strconv.Itoa(status)
			response.ResponseMsg = err.Error()
			utils.Response(response, c, status)
			return
		}
	}
	d, err := utils.ConnectDB("sqlite3", dbFile) // 连接数据库
	if err != nil {
		utils.Log.Errorf("connectdb err, %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	defer d.DisConnectDB()
	if !dbFileExist {
		err := d.CreateTable() // 创建表
		if err != nil {
			utils.Log.Errorf("create table err, %s", err.Error())
			status = http.StatusServiceUnavailable
			response.ResponseCode = strconv.Itoa(status)
			response.ResponseMsg = err.Error()
			utils.Response(response, c, status)
			return
		}
	}
	flag, err := d.QueryAndUpdateTable(info) // 查询并更新数据库
	if err != nil {
		utils.Log.Errorf("query and update table err, %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	if flag {
		var alertPeriod int

		responseData, _, err = sdk.Handler.QueryData("", define.QUERY_DATA, nil, info.ChannelSeq)
		if err != nil {
			utils.Log.Errorf("Query : %s", err.Error())
			status = http.StatusServiceUnavailable
			response.ResponseCode = strconv.Itoa(status)
			response.ResponseMsg = err.Error()
			utils.Response(response, c, http.StatusServiceUnavailable)
			return
		}
		if responseData.Payload == nil {
			alertPeriod = 1
			utils.Log.Debugf("No preset alertPeriod")
		} else {
			str := string(responseData.Payload.([]byte))
			utils.Log.Debugf("alertPeriod: %s", str)
			alertPeriod, err = strconv.Atoi(str)
			if err != nil {
				alertPeriod = 1
				utils.Log.Debugf("No avaliable preset alertPeriod")
			}
		}

		txId, nonce, err := sdk.Handler.GetTxId()
		if err != nil {
			utils.Log.Errorf("GetTxId : %s", err.Error())
			status = http.StatusServiceUnavailable
			response.ResponseCode = strconv.Itoa(status)
			response.ResponseMsg = err.Error()
			utils.Response(response, c, status)
			return
		}

		err = d.InsertTable(info)
		if err != nil {
			utils.Log.Errorf("insert or update table err, %s", err.Error())
			status = http.StatusServiceUnavailable
			response.ResponseCode = strconv.Itoa(status)
			response.ResponseMsg = err.Error()
			utils.Response(response, c, status)
			return
		}

		go timer(info.PledgeNoStorage, alertPeriod, txId, nonce)
	}

	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "submit success!"
	utils.Response(response, c, status)
}

func timer(pledgeNoStorage string, alertPeriod int, txId string, nonce []byte) {
	timer := time.NewTimer(time.Minute * time.Duration(alertPeriod))
	<-timer.C
	utils.Log.Debugf("Timer expired")

	dbFile := "./pledge.db" //sqlite3数据库名字
	dbFileExist, err := utils.FileOrDirectoryExist(dbFile)
	if err != nil {
		utils.Log.Errorf("check file exist or not error, %s", err.Error())
		return
	}
	if !dbFileExist {
		utils.Log.Errorf("dbfile not exist")
		return
	}
	d, err := utils.ConnectDB("sqlite3", dbFile) // 连接数据库
	if err != nil {
		utils.Log.Errorf("connectdb err, %s", err.Error())
		return
	}
	defer d.DisConnectDB()
	flag, info, err := d.QueryTable(pledgeNoStorage) // 查询数据库
	if err != nil {
		utils.Log.Errorf("query table err, %s", err.Error())
		return
	}
	if flag {
		utils.Log.Errorf("pledge information not exist")
		return
	}

	var invokeRequest define.InvokeRequest
	val, _ := json.Marshal(info)
	invokeRequest.Value = string(val)
	invokeRequest.Key = fmt.Sprintf("%s", info.PledgeNoStorage)

	utils.Log.Debugf("value : %s", string(val))
	utils.Log.Debugf("key : %s", info.PledgeNoStorage)

	b, _ := json.Marshal(invokeRequest)

	// invoke
	_, err = sdk.Handler.Invoke(txId, nonce, "", define.UPLOAD_WARNING_INFO, nil, b)
	if err != nil {
		utils.Log.Errorf("SaveData Invoke : %s", err.Error())
		return
	}

	err = d.DeleteTable(pledgeNoStorage) // 更新数据库
	if err != nil {
		utils.Log.Errorf("delete table err, %s", err.Error())
		return
	}
	return
}
