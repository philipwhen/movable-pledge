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

// 上传押品信息
func UploadPledgeInfo(c *gin.Context) {
	utils.Log.Debug("UploadPledgeInfo .....")
	var goodsInfo define.GoodsInfo
	var response define.PledgeResponse
	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("UploadPledgeInfo read body error : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	utils.Log.Debugf("UploadPledgeInfo header : %v", c.Request.Header)
	utils.Log.Debugf("UploadPledgeInfo body : %s", string(body))

	if err = json.Unmarshal(body, &goodsInfo); err != nil {
		utils.Log.Errorf("UploadPledgeInfo unmarshal body error : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	txId, nonce, err := sdk.Handler.GetTxId()
	if err != nil {
		utils.Log.Errorf("UploadPledgeInfo GetTxId error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	goodsInfo.PledgePatrolDetailInfo = make([]define.PledgePatrolDetailInfo, 0)
	var invokeRequest define.InvokeRequest
	val, _ := json.Marshal(goodsInfo)
	invokeRequest.Value = string(val)
	invokeRequest.Key = goodsInfo.PledgeGeneralInfo.PledgeNoStorage

	b, err := json.Marshal(invokeRequest)
	if err != nil {
		utils.Log.Errorf("UploadPledgeInfo marshal invokeRequest error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	// invoke
	_, err = sdk.Handler.Invoke(txId, nonce, "", define.UPLOAD_PLEDGE_INFO, nil, b)
	if err != nil {
		utils.Log.Errorf("UploadPledgeInfo invoke error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "押品上链成功"
	utils.Response(response, c, status)
}

// 上传保险公证信息
func UploadInsurNotarInfo(c *gin.Context) {
	utils.Log.Debug("UploadInsurNotarInfo .....")
	var pledgeInsuranceNotarizationInfo define.PledgeInsuranceNotarizationInfo
	var response define.PledgeResponse
	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("UploadInsurNotarInfo read body error : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	utils.Log.Debugf("UploadInsurNotarInfo header : %v", c.Request.Header)
	utils.Log.Debugf("UploadInsurNotarInfo body : %s", string(body))

	if err = json.Unmarshal(body, &pledgeInsuranceNotarizationInfo); err != nil {
		utils.Log.Errorf("UploadInsurNotarInfo unmarshal body error : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	txId, nonce, err := sdk.Handler.GetTxId()
	if err != nil {
		utils.Log.Errorf("UploadInsurNotarInfo GetTxId error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	var invokeRequest define.InvokeRequest
	invokeRequest.Value = string(body)
	invokeRequest.Key = pledgeInsuranceNotarizationInfo.PledgeInsuranceInfo[0].PledgeNoStorage

	b, err := json.Marshal(invokeRequest)
	if err != nil {
		utils.Log.Errorf("UploadInsurNotarInfo marshal invokeRequest error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	// invoke
	_, err = sdk.Handler.Invoke(txId, nonce, "", define.UPLOAD_INSUR_NOTAR_INFO, nil, b)
	if err != nil {
		utils.Log.Errorf("UploadInsurNotarInfo invoke error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "保险和公证信息上链成功"
	utils.Response(response, c, status)
}

// 上传巡库信息
func UploadPatrolInfo(c *gin.Context) {
	utils.Log.Debug("UploadPatrolInfo .....")
	var pledgePatrolDetailInfo define.PledgePatrolDetailInfo
	var response define.PledgeResponse
	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("UploadPatrolInfo read body error : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	utils.Log.Debugf("UploadPatrolInfo header : %v", c.Request.Header)
	utils.Log.Debugf("UploadPatrolInfo body : %s", string(body))

	if err = json.Unmarshal(body, &pledgePatrolDetailInfo); err != nil {
		utils.Log.Errorf("UploadPatrolInfo unmarshal body error : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	txId, nonce, err := sdk.Handler.GetTxId()
	if err != nil {
		utils.Log.Errorf("UploadPatrolInfo GetTxId error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	var invokeRequest define.InvokeRequest
	invokeRequest.Value = string(body)
	invokeRequest.Key = pledgePatrolDetailInfo.MonitorGeneralInfo.PledgeNoStorage

	b, err := json.Marshal(invokeRequest)
	if err != nil {
		utils.Log.Errorf("UploadPatrolInfo marshal invokeRequest error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	// invoke
	_, err = sdk.Handler.Invoke(txId, nonce, "", define.UPLOAD_PATROL_INFO, nil, b)
	if err != nil {
		utils.Log.Errorf("UploadPatrolInfo invoke error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "巡库信息上链成功"
	utils.Response(response, c, status)
}

// 设置预警周期
func SetAlertPeriod(c *gin.Context) {
	utils.Log.Debug("SetAlertPeriod .....")
	var alertPeriodSetting define.AlertPeriodSetting
	var response define.PledgeResponse
	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("SetAlertPeriod read body error : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	utils.Log.Debugf("SetAlertPeriod header : %v", c.Request.Header)
	utils.Log.Debugf("SetAlertPeriod body : %s", string(body))

	if err = json.Unmarshal(body, &alertPeriodSetting); err != nil {
		utils.Log.Errorf("SetAlertPeriod unmarshal body error : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	txId, nonce, err := sdk.Handler.GetTxId()
	if err != nil {
		utils.Log.Errorf("SetAlertPeriod GetTxId error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	var invokeRequest define.InvokeRequest
	invokeRequest.Value = alertPeriodSetting.AlertPeriod
	invokeRequest.Key = alertPeriodSetting.ChannelSeq

	b, err := json.Marshal(invokeRequest)
	if err != nil {
		utils.Log.Errorf("SetAlertPeriod marshal invokeRequest error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	_, err = sdk.Handler.Invoke(txId, nonce, "", define.SET_ALERT_PERIOD, nil, b)
	if err != nil {
		utils.Log.Errorf("SetAlertPeriod invoke error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "预警周期设置成功"
	utils.Response(response, c, status)
}

// 同步押品状态
func StatusSync(c *gin.Context) {
	utils.Log.Debug("synchronize goods status .....")
	var request define.StatusSyncRequest
	var response define.PledgeResponse

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
	var response define.PledgeResponse
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
	var s []string
	err = json.Unmarshal(responseData.Payload.([]byte), &s)
	for _, value := range s {
		responseData, _, err = sdk.Handler.QueryData("", define.QUERY_DATA, nil, value)
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

// 查询押品信息
func QueryGoodsInfo(c *gin.Context) {
	utils.Log.Debug("QueryGoodsInfo .....")

	var request define.QueryGoodsRequest
	var response define.QueryGoodsResponse
	var result define.GoodsInfo
	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("QueryGoodsInfo read body error : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	utils.Log.Debugf("QueryGoodsInfo header : %v", c.Request.Header)
	utils.Log.Debugf("QueryGoodsInfo body : %s", string(body))

	if err = json.Unmarshal(body, &request); err != nil {
		utils.Log.Errorf("QueryGoodsInfo unmarshal body error : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	b := request.PledgeNoStorage
	responseData, _, err := sdk.Handler.QueryData("", define.QUERY_DATA, nil, b)
	if err != nil {
		utils.Log.Errorf("QueryGoodsInfo query error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	if responseData.Payload == nil {
		status = http.StatusNotFound
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = "区块链上没有存储该押品"
		utils.Response(response, c, status)
		return
	}

	if err = json.Unmarshal(responseData.Payload.([]byte), &result); err != nil {
		utils.Log.Errorf("QueryGoodsInfo unmarshal response payload error : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "押品信息查询成功"
	response.GoodsInfo = result
	utils.Response(response, c, status)
}

// 查询预警周期
func QueryPeriod(c *gin.Context) {
	utils.Log.Debug("QueryPeriod .....")

	var request define.QueryPeriodRequest
	var response define.QueryPeriodResponse
	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("QueryPeriod read body error : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	utils.Log.Debugf("QueryPeriod header : %v", c.Request.Header)
	utils.Log.Debugf("QueryPeriod body : %s", string(body))

	if err = json.Unmarshal(body, &request); err != nil {
		utils.Log.Errorf("QueryPeriod unmarshal body error : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	b := request.ChannelSeq
	responseData, _, err := sdk.Handler.QueryData("", define.QUERY_DATA, nil, b)
	if err != nil {
		utils.Log.Errorf("QueryPeriod query error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	if responseData.Payload == nil {
		status = http.StatusNotFound
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = "区块链上没有预警周期信息"
		utils.Response(response, c, status)
		return
	}

	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "预警周期查询成功"
	response.AlertPeriod = string(responseData.Payload.([]byte))
	utils.Response(response, c, status)
}

// 查询监控设备监控的押品信息
func QueryMonitorPledge(c *gin.Context) {
	utils.Log.Debug("QueryMonitorPledge .....")

	var request define.QueryMonitorPledgeRequest
	var response define.QueryMonitorPledgeResponse
	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("QueryMonitorPledge read body error : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	utils.Log.Debugf("QueryMonitorPledge header : %v", c.Request.Header)
	utils.Log.Debugf("QueryMonitorPledge body : %s", string(body))

	if err = json.Unmarshal(body, &request); err != nil {
		utils.Log.Errorf("QueryMonitorPledge unmarshal body error : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	b := request.MonitorEquipmentId
	responseData, _, err := sdk.Handler.QueryData("", define.QUERY_DATA, nil, b)
	if err != nil {
		utils.Log.Errorf("QueryMonitorPledge query error : %s", err.Error())
		status = http.StatusServiceUnavailable
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	if responseData.Payload == nil {
		status = http.StatusNotFound
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = "区块链上没有该监控设备监控的押品信息"
		utils.Response(response, c, status)
		return
	}

	if err = json.Unmarshal(responseData.Payload.([]byte), &response.PledgeNoStorageList); err != nil {
		utils.Log.Errorf("QueryMonitorPledge unmarshal response payload error : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "监控设备监控的押品信息查询成功"
	utils.Response(response, c, status)
}

// 押品信息接收接口
func PledgeReception(c *gin.Context) {
	utils.Log.Debug("PledgeReception .....")
	var goodsInfo define.GoodsInfo
	var response define.PledgeResponse
	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		utils.Log.Errorf("PledgeReception read body error : %s", err.Error())
		status = http.StatusNoContent
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	utils.Log.Debugf("PledgeReception header : %v", c.Request.Header)
	utils.Log.Debugf("PledgeReception body : %s", string(body))

	if err = json.Unmarshal(body, &goodsInfo); err != nil {
		utils.Log.Errorf("PledgeReception unmarshal body error : %s", err.Error())
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}

	utils.Log.Debugf("PledgeReception goodsInfo pledgeNo : %s", goodsInfo.PledgeGeneralInfo.PledgeNoStorage)

	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "押品信息接收成功"
	utils.Response(response, c, status)
}
