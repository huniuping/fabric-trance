package manufacture_sdk

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"time"
)

var (
	manufacture_sdk     *fabsdk.FabricSDK
	conf_path      string = "../application/conf/sdk_conf/manufactureOrg/manufacture-conf.yaml"
	channel_id     string = "manufacturechannel"
	chaincode_name string = "gongdancc"
	peer_domain    string = "peer0.manufacture.huniuping.com"
	peer_domain1    string = "peer0.supervise.huniuping.com"
	org            string = "ManufactureOrg"
	user           string = "Admin"
	err            error
)

func init() {
	manufacture_sdk, err = fabsdk.New(config.FromFile(conf_path))
	if err != nil {
		ret_err := fmt.Sprintf("读取配置文件出错，错误信息：", err)
		panic(ret_err)
	}
}

func ChannelExecute(chaincode_name, execute_fcn string, args [][]byte) (channel.Response, error) {
	ctx := manufacture_sdk.ChannelContext(channel_id, fabsdk.WithOrg(org), fabsdk.WithUser(user))
	cli, err := channel.New(ctx)
	if err != nil {
		panic("初始化通道出错")
	}
	rsp, err := cli.Execute(channel.Request{
		ChaincodeID: chaincode_name,
		Fcn:         execute_fcn,
		Args:        args,
	}, channel.WithTargetEndpoints(peer_domain), channel.WithTimeout(fab.TimeoutType(3), time.Second*2))
	return rsp, err
}


func ChannelQuery(chaincode_name, query_fcn string, args [][]byte) (channel.Response, error) {
	ctx := manufacture_sdk.ChannelContext(channel_id, fabsdk.WithOrg(org), fabsdk.WithUser(user))
	cli, err := channel.New(ctx)
	if err != nil {
		ret_err := fmt.Sprintf("初始化通道出错:%v", err)
		panic(ret_err)
	}
	rsp, err := cli.Query(channel.Request{
		ChaincodeID: chaincode_name,
		Fcn:         query_fcn,
		Args:        args,
	}, channel.WithTargetEndpoints(peer_domain))
	return rsp, err
}
