package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/sifbc/pledge/apiserver/define"
	"github.com/sifbc/pledge/apiserver/sdk"
	"github.com/sifbc/pledge/apiserver/utils"

	"github.com/gin-gonic/gin"
)

var targetOrderAddr string
var keymap = make(map[string]uint)

type QueryRequest struct {
	Key string
}
type QueryResponse struct {
	ResponseCode string
	ResponseMsg  string
	Payload      string
}

func TimeTickerTest(c *gin.Context) {
	var request QueryRequest
	var response QueryResponse
	var err error
	status := http.StatusOK

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	if err = json.Unmarshal(body, &request); err != nil {
		status = http.StatusBadRequest
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = err.Error()
		utils.Response(response, c, status)
		return
	}
	if _, exist := keymap[request.Key]; exist {
		fmt.Println(keymap[request.Key])
		keymap[request.Key] += 1
		fmt.Println(keymap[request.Key])
		response.ResponseCode = strconv.Itoa(status)
		response.ResponseMsg = "success!"
		utils.Response(response, c, status)
		return
	}
	keymap[request.Key] = 1
	go timer(request.Key)
	response.ResponseCode = strconv.Itoa(status)
	response.ResponseMsg = "success!"
	utils.Response(response, c, status)
}

func timer(key string) {
	ticker := time.NewTicker(30 * time.Second)
	<-ticker.C
	fmt.Printf("hello %s timer: %d\n", key, keymap[key])
	delete(keymap, key)
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
	invokeRequest.Key = fmt.Sprintf("%s", GoodsInfo.PledgeGeneralInfo.PledgeNo)

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

	b := request.PledgeNo
	responseData, _, err := sdk.Handler.QueryData("", define.QUERY_DATA, nil, b)
	if err != nil {
		utils.Log.Errorf("Query : %s", err.Error())
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
