package handle

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sifbc/pledge/apiserver/define"
)

var (
	PledgeReceptionUrl string
	WarningUrl         string
	StatusUrl          string
)

func CheckEvent(fromListen chan BlockInfoAll) {
	//Handle the transactions from listen module
	for {
		select {
		case blockInfo := <-fromListen:
			err := ParseBlockInfo(blockInfo)
			if err != nil {
				logger.Errorf("parse block info failed: %s", err.Error())
				continue
			}
		}
	}
}

func ParseBlockInfo(blockInfo BlockInfoAll) error {
	logger.Info("enter ParseBlockInfo function......")
	switch blockInfo.EventName {
	case define.UPLOAD_PLEDGE_INFO:
		logger.Info("upload pledge info")

	//	pledgeReceptionUrl = "http://10.47.137.211:5555/scfs/blockChain/pledgeReception"
		message, err := json.Marshal(blockInfo.MsgInfo.(define.GoodsInfo))
		if err != nil {
			logger.Errorf("marshal request error : %s", err.Error())
			return err
		}
		_, err = SendMessage(message, PledgeReceptionUrl)
		if err != nil {
			logger.Errorf("send message error : %s", err.Error())
			return err
		}
	case define.UPLOAD_WARNING_INFO:
		logger.Info("upload warning info")

		//pledgeReceptionUrl := "http://10.47.137.211:5555/scfs/blockChain/pledgeReception"
		message, err := json.Marshal(blockInfo.MsgInfo.(define.PledgeWarningInfo))
		if err != nil {
			logger.Errorf("marshal request error : %s", err.Error())
			return err
		}
		_, err = SendMessage(message, WarningUrl)
		if err != nil {
			logger.Errorf("send message error : %s", err.Error())
			return err
		}
	case define.STATUS_SYNC:
		logger.Info("status sync")

		//pledgeReceptionUrl = "http://10.37.48.181:5555/scfs/blockChain/pledgeReception"
		message, err := json.Marshal(blockInfo.MsgInfo.(OperateRequest))
		if err != nil {
			logger.Errorf("marshal request error : %s", err.Error())
			return err
		}
		_, err = SendMessage(message, StatusUrl)
		if err != nil {
			logger.Errorf("send message error : %s", err.Error())
			return err
		}
	}
	return nil
}

func SendMessage(data []byte, destUrl string) ([]byte, error) {
	PostFunc := func(data []byte) ([]byte, error) {
		req, err := http.NewRequest("POST", destUrl, bytes.NewBuffer(data))
		if err != nil {
			logger.Errorf("New post request error , %s", err.Error())
			return []byte("empty"), err
		}
		//req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logger.Errorf("client do error , %s", err.Error())
			return []byte("empty"), err
		}
		defer resp.Body.Close()
		res, _ := ioutil.ReadAll(resp.Body)

		return res, nil
	}
	logger.Infof("push request body , %s", string(data))
	res, err := PostFunc(data)
	return res, err
}
