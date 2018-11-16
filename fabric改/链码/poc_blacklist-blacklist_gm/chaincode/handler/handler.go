package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
	"github.com/peersafe/poc_blacklist/chaincode/define"
	api_def "github.com/peersafe/poc_blacklist/apiserver/define"
	"github.com/peersafe/poc_blacklist/chaincode/utils"
	"os"
)

var myLogger = logging.MustGetLogger("hanldler")

var KEEPALIVETEST = "keepAliveTest"

func init() {
	format := logging.MustStringFormatter("%{shortfile} %{time:15:04:05.000} [%{module}] %{level:.4s} : %{message}")
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	logging.SetBackend(backendFormatter).SetLevel(logging.DEBUG, "hanldler")
}

func Uploadblacklist(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    var err error
    var ccSaveBlk api_def.BlacklistKeyData
    request := &define.InvokeRequest{}
    if err = json.Unmarshal([]byte(args[1]), request); err != nil {
        return utils.InvokeResponse(stub, err, function, nil, false)
    }

    blk := &api_def.BlackListInfo{}
    if err = json.Unmarshal([]byte(request.Value), blk); err != nil {
        return utils.InvokeResponse(stub, err, function, nil, false)
    }

	listKey := blk.CommData.ListUniqueKey
	currentMonType := fmt.Sprintf("%s-%s", blk.CreatTime.Year, blk.CreatTime.Month)
    b, _ := stub.GetState(request.Key)
    if b == nil {
        ccSaveBlk.BlkLists = make(map[string]api_def.BlackListInfo)
		ccSaveBlk.BlkLists[listKey] = *blk
		ccSaveBlk.BlackListCntInfo.AddListUniqueKey = make([]string, 0)
		ccSaveBlk.BlackListCntInfo.UpdateListUniqueKey = make([]string, 0)
		ccSaveBlk.BlackListCntInfo.DeleteListUniqueKey = make([]string, 0)
		ccSaveBlk.BlackListCntInfo.BlackListCnt = make(map[string]int)
		ccSaveBlk.BlackListCntInfo.AddListUniqueKey = append(ccSaveBlk.BlackListCntInfo.AddListUniqueKey, listKey)
		ccSaveBlk.BlackListCntInfo.BlackListCnt[blk.CommData.ListType] = 1
		ccSaveBlk.BlackListCntInfo.BlackListCnt[api_def.BLACKLIST_TOTAL_COUNT] = 1
		ccSaveBlk.BlackListCntInfo.BlackListCnt[currentMonType] = 1
    } else {
        if err = json.Unmarshal(b, &ccSaveBlk); err != nil {
            return utils.InvokeResponse(stub, err, function, nil, false)
		}
		ccSaveBlk.BlackListCntInfo.AddListUniqueKey = make([]string, 0)
		ccSaveBlk.BlackListCntInfo.UpdateListUniqueKey = make([]string, 0)
		ccSaveBlk.BlackListCntInfo.DeleteListUniqueKey = make([]string, 0)
		ccSaveBlk.BlackListCntInfo.BlackListCnt = make(map[string]int)
		_, ok := ccSaveBlk.BlkLists[listKey]
		if ok {
			// ccSaveBlk.BlackListCntInfo.UpdateListUniqueKey = append(ccSaveBlk.BlackListCntInfo.UpdateListUniqueKey, listKey)
			// ccSaveBlk.BlackListCntInfo.BlackListCnt[blk.CommData.ListType] = 0
			// ccSaveBlk.BlackListCntInfo.BlackListCnt[api_def.BLACKLIST_TOTAL_COUNT] = 0
			// ccSaveBlk.BlackListCntInfo.BlackListCnt[currentMonType] = 0
			err = fmt.Errorf("%s", "Blacklist already exist!")
			return utils.InvokeResponse(stub, err, function, nil, false)
		} else {
			ccSaveBlk.BlackListCntInfo.AddListUniqueKey = append(ccSaveBlk.BlackListCntInfo.AddListUniqueKey, listKey)
			ccSaveBlk.BlackListCntInfo.BlackListCnt[blk.CommData.ListType] = 1
			ccSaveBlk.BlackListCntInfo.BlackListCnt[api_def.BLACKLIST_TOTAL_COUNT] = 1
			ccSaveBlk.BlackListCntInfo.BlackListCnt[currentMonType] = 1
		}
        ccSaveBlk.BlkLists[listKey] = *blk
    }

	// 不能在这里直接调时间生成，各个chaincode之间这个参数值可能不一样，导致上传失败
	// ccSaveBlk.BlackListCntInfo.OperationTime = time.Now().Unix()
	// 使用在apiserver传入的时间
	operationTime := fmt.Sprintf("%s-%s-%s %s:%s:%s", blk.CreatTime.Year, blk.CreatTime.Month, blk.CreatTime.Day, blk.CreatTime.Hour, blk.CreatTime.Minute, blk.CreatTime.Second)
	ccSaveBlk.BlackListCntInfo.OperationTime = operationTime
    ccSaveBlk.DataType = api_def.DATATYPE_BLACKLIST

    val, _ := json.Marshal(ccSaveBlk)
    err = stub.PutState(request.Key, val)
    if err != nil {
        myLogger.Errorf("Uploadblacklist err: %s", err.Error())
        return utils.InvokeResponse(stub, err, function, nil, false)
	}

    return utils.InvokeResponse(stub, err, function, ccSaveBlk, true)
}

func SaveData(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	request := &define.InvokeRequest{}
	if err = json.Unmarshal([]byte(args[1]), request); err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	err = stub.PutState(request.Key, []byte(request.Value))
	if err != nil {
		myLogger.Errorf("saveData err: %s", err.Error())
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	var list []string
	list = append(list, request.Value)
	return utils.InvokeResponse(stub, nil, function, list, true)
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

func DslQuery(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	request := &define.QueryRequest{}
	if err = json.Unmarshal([]byte(args[1]), request); err != nil {
		err = fmt.Errorf("DslQuery json decode args failed, err = %s", err.Error())
		return utils.QueryResponse(err, nil, request.SplitPage)
	}
	result, err := utils.GetValueByDSL(stub, request)
	if err != nil {
		return utils.QueryResponse(err, nil, request.SplitPage)
	}

	return utils.QueryResponse(nil, result, request.SplitPage)
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

func UploadApplicationMaterials(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	var materialsInfo api_def.MaterialsInfo
	var allMaterialsInfo api_def.AllMaterialsInfo

	request := &define.InvokeRequest{}
	if err = json.Unmarshal([]byte(args[1]), request); err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	//企业资料
	orgInfo := &api_def.OrgInfo{}
	if err = json.Unmarshal([]byte(request.Value), orgInfo); err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	//企业上传资料的次数
	org, _ := stub.GetState(request.Key)
	if org == nil {
		materialsInfo.UploadTimes = api_def.FIRST_TIME_UPLOAD_APPLICATIONS_INFO //申请批次
		materialsInfo.DataType = api_def.DATATYPE_APPLICATIONMATERIALS  //资料固定字段为2
	} else {
		if err = json.Unmarshal(org, &materialsInfo); err != nil {
			return utils.InvokeResponse(stub, err, function, nil, false)
		}
		materialsInfo.UploadTimes = materialsInfo.UploadTimes + 1
        materialsInfo.VerifiedInfo = nil  /*企业再次上传资料清空之前审批信息*/
        materialsInfo.QueryState = ""
        materialsInfo.Suggestion = ""
	}

    materialsInfo.OrgInfo = *orgInfo
    materialsInfo.OrgID = orgInfo.OrgID
    materialsInfo.Orgname = orgInfo.EnterpriseInfo.Orgname
    materialsInfo.CreateTime = orgInfo.CreateTime

	allOrgName, err := stub.GetState(api_def.ALL_MATERIALS_INFO_KEY)
	if err != nil{
		myLogger.Errorf("UploadApplicationMaterials err: %s", err.Error())
		return utils.InvokeResponse(stub, err, function, nil, false)
	}
	if allOrgName != nil {
		if err = json.Unmarshal(allOrgName, &allMaterialsInfo); err != nil {
			return utils.InvokeResponse(stub, err, function, nil, false)
		}
	} else {
		allMaterialsInfo.OrgNameList = []string{}
	}

	//如果已经存在，不需要跟新
	objectNotExist := true
	for _, v := range allMaterialsInfo.OrgNameList {
		if v == materialsInfo.Orgname {
			objectNotExist = false
		}
	}
	if objectNotExist {
		allMaterialsInfo.OrgNameList = append(allMaterialsInfo.OrgNameList, materialsInfo.Orgname)
	}
	
	val, _ := json.Marshal(materialsInfo)
	err = stub.PutState(request.Key, val)
	if err != nil {
		myLogger.Errorf("UploadApplicationMaterials err: %s", err.Error())
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	if objectNotExist {
		allVal, _ := json.Marshal(allMaterialsInfo)
		err = stub.PutState(api_def.ALL_MATERIALS_INFO_KEY, allVal)
		if err != nil {
			myLogger.Errorf("UploadApplicationMaterials err: %s", err.Error())
			return utils.InvokeResponse(stub, err, function, nil, false)
		}
	}

	return utils.InvokeResponse(stub, err, function, materialsInfo, true)
}

func Verifyqualification(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	verifiedInfo := &api_def.VerifiedInfo{}
	mateinfo := &api_def.MaterialsInfo{}
	request := &define.InvokeRequest{}

	if err = json.Unmarshal([]byte(args[1]), request); err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}
	if err = json.Unmarshal([]byte(request.Value), verifiedInfo); err != nil{
		return utils.InvokeResponse(stub, err, function, nil, false)
	}
	b, err := stub.GetState(verifiedInfo.Auditee)
	if err != nil{
		myLogger.Errorf("Verifyqualification err: %s", err.Error())
		return utils.InvokeResponse(stub, err, function, nil, false)
	}
	if b == nil{
		return utils.InvokeResponse(stub, nil, function, nil, false)
	}

	if err = json.Unmarshal(b, mateinfo); err != nil {
		return utils.InvokeResponse(stub, err, function, nil, false)
	}
	if mateinfo.VerifiedInfo == nil{
		mateinfo.VerifiedInfo = make(map[string]api_def.VerifiedInfo)
	}
	mateinfo.VerifiedInfo[request.Key] = *verifiedInfo
    mateinfo.Suggestion = verifiedInfo.Suggestion
    count := 0
	notAgreeFlag := false
	var verifiedInfoAgree string
    for _, verify := range mateinfo.VerifiedInfo {
        if verify.Agree == api_def.MATERIAL_AGREE {
			count += 1
        } else {
			notAgreeFlag = true
			verifiedInfoAgree = verify.Agree
            break
		}
		if count >= 4 {
			notAgreeFlag = true
			verifiedInfoAgree = api_def.MATERIAL_AGREE
			break
		}
    }

    if notAgreeFlag {
        mateinfo.QueryState = verifiedInfoAgree
    } else {
        mateinfo.QueryState = fmt.Sprintf("4-%d",count)
    }
    mateinfo.VerifiedDate = verifiedInfo.VerifiedDate

	val, _ := json.Marshal(mateinfo)
	err = stub.PutState(verifiedInfo.Auditee, val)
	if err != nil {
		myLogger.Errorf("Verifyqualification err: %s", err.Error())
		return utils.InvokeResponse(stub, err, function, nil, false)
	}

	var list []string
	list = append(list, request.Value)
	return utils.InvokeResponse(stub, nil, function, list, true)
}