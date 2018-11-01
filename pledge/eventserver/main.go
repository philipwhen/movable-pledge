package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/hyperledger/fabric-sdk-go-peersafe/pkg/common"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/sifbc/pledge/eventserver/handle"
	le "github.com/sifbc/pledge/eventserver/listenevent"

	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	logOutput  = os.Stderr
	configPath = flag.String("configPath", "./", "config path")
	configName = flag.String("configName", "client_sdk", "config file name")
)

const CHANNELBUFFER = 1000

func main() {
	flag.Parse()
	err := common.InitSDK(*configPath, *configName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	runtime.GOMAXPROCS(viper.GetInt("peer.gomaxprocs"))

	//setup system-wide logging backend based on settings from core.yaml
	flogging.InitBackend(flogging.SetFormat(viper.GetString("logging.format")), logOutput)
	logging.SetLevel(logging.DEBUG, "client_sdk")

	eventAddress := viper.GetString("peer.listenAddress")
	chainID := viper.GetString("chaincode.id.chainID")

	handle.PledgeReceptionUrl = viper.GetString("bankUrl.pledgeReceptionUrl")
	handle.WarningUrl = viper.GetString("bankUrl.pledgeWarningUrl")

	listenToHandle := make(chan handle.BlockInfoAll, CHANNELBUFFER)

	//check and recover the message
	go handle.CheckEvent(listenToHandle)

	//listen the block event and parse the message
	le.ListenEvent(eventAddress, chainID, handle.FilterEvent, listenToHandle)
}
