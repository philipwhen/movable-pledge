package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
	api_def "github.com/sifbc/pledge/apiserver/define"
	"github.com/sifbc/pledge/chaincode/define"
	"github.com/sifbc/pledge/chaincode/utils"
)

var myLogger = logging.MustGetLogger("hanldler")

var KEEPALIVETEST = "keepAliveTest"

func init() {
	format := logging.MustStringFormatter("%{shortfile} %{time:15:04:05.000} [%{module}] %{level:.4s} : %{message}")
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	logging.SetBackend(backendFormatter).SetLevel(logging.DEBUG, "hanldler")
}

func QueryData(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	key := args[1]
	page := define.Page{}
	result, err := stub.GetState(key)
	if err != nil {
		return utils.QueryResponse(err, nil, page)
	}

	return utils.QueryResponse(nil, result, page)
}

func KeepaliveQuery(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	targetValue, err := stub.GetState(KEEPALIVETEST)
	if err != nil {
		err = fmt.Errorf("ERROR! KeepaliveQuery get failed, err = %s", err.Error())
		return []byte("UnReached"), err
	}

	if string(targetValue) != KEEPALIVETEST {
		err = fmt.Errorf("ERROR! KeepaliveQuery get result is %s", string(targetValue))
		return []byte("UnReached"), err
	}

	return []byte("Reached"), nil
}

// new function: upload goods info
func UploadGoodsInfo(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	//	var result api_def.GoodsInfo

	request := &define.InvokeRequest{}
	if err = json.Unmarshal([]byte(args[1]), request); err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	//货物资料
	GoodsInfo := &api_def.GoodsInfo{}
	if err = json.Unmarshal([]byte(request.Value), GoodsInfo); err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	info, err := stub.GetState(request.Key)
	if info != nil || err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}
	//	result = GoodsInfo
	val, _ := json.Marshal(GoodsInfo)
	err = stub.PutState(request.Key, val)
	if err != nil {
		myLogger.Errorf("UploadApplicationMaterials err: %s", err.Error())
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	return utils.InvokeResponse(stub, err, function, nil, true)
}

// new function: upload patrol info
func UploadPatrolInfo(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	var GoodsInfo api_def.GoodsInfo
	var PatrolInfo api_def.PledgePatrolDetailInfo

	request := &define.InvokeRequest{}
	if err = json.Unmarshal([]byte(args[1]), request); err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	//根据key值获取链上信息
	v, err := stub.GetState(request.Key)
	if v == nil || err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}
	if err = json.Unmarshal(v, &GoodsInfo); err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	//将巡库信息添加到链上
	if err = json.Unmarshal([]byte(request.Value), &PatrolInfo); err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	GoodsInfo.PledgePatrolDetailInfo = append(GoodsInfo.PledgePatrolDetailInfo, PatrolInfo)

	val, _ := json.Marshal(GoodsInfo)
	err = stub.PutState(request.Key, val)
	if err != nil {
		myLogger.Errorf("UploadApplicationMaterials err: %s", err.Error())
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	return utils.InvokeResponse(stub, err, function, nil, true)
}
