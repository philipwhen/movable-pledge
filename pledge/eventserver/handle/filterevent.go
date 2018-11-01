package handle

import (
	"encoding/json"
	//"fmt"
	//"time"

	"github.com/sifbc/pledge/apiserver/define"

	protos_peer "github.com/hyperledger/fabric/protos/peer"
	"github.com/op/go-logging"
)

var (
	logger = logging.MustGetLogger("filter-event")
)

func FilterEvent(event *protos_peer.ChaincodeEvent) (interface{}, bool) {
	logger.Info("enter filterevent function......")
	logger.Info(event.EventName)
	switch event.EventName {
	case define.UPLOAD_PLEDGE_INFO:
		responseData := define.InvokeResponse{}
		goodsInfo := define.GoodsInfo{}
		responseData.Payload = &goodsInfo

		logger.Infof("unmarshal payload is %s.", string(event.Payload))
		err := json.Unmarshal(event.Payload, &responseData)
		if err != nil {
			logger.Errorf("unmarshal payload failed: %s.", err)
			return nil, false
		}
		return goodsInfo, true
	case define.UPLOAD_WARNING_INFO:
		responseData := define.InvokeResponse{}
		warningInfo := define.PledgeWarningInfo{}
		responseData.Payload = &warningInfo

		logger.Infof("unmarshal payload is %s.", string(event.Payload))
		err := json.Unmarshal(event.Payload, &responseData)
		if err != nil {
			logger.Errorf("unmarshal payload failed: %s.", err)
			return nil, false
		}
		return warningInfo, true
	case define.STATUS_SYNC:
		responseData := define.InvokeResponse{}
		OperateRequest := OperateRequest{}
		responseData.Payload = &OperateRequest

		logger.Infof("unmarshal payload is %s.", string(event.Payload))
		err := json.Unmarshal(event.Payload, &responseData)
		if err != nil {
			logger.Errorf("unmarshal payload failed: %s.", err)
			return nil, false
		}
		return OperateRequest, true
	}
	return nil, false
}
