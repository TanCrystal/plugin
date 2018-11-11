package rpc_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	commonlog "github.com/33cn/chain33/common/log"
	"github.com/33cn/chain33/rpc/jsonclient"
	"github.com/33cn/chain33/types"
	"github.com/33cn/chain33/util/testnode"
	"github.com/33cn/plugin/plugin/dapp/retrieve/rpc"
	pty "github.com/33cn/plugin/plugin/dapp/retrieve/types"

	_ "github.com/33cn/chain33/system"
	_ "github.com/33cn/plugin/plugin"
)

func init() {
	commonlog.SetLogLevel("error")
}

func TestJRPCChannel(t *testing.T) {
	// 启动RPCmocker
	mocker := testnode.New("--notset--", nil)
	defer func() {
		mocker.Close()
	}()
	mocker.Listen()

	jrpcClient := mocker.GetJsonC()

	testCases := []struct {
		fn func(*testing.T, *jsonclient.JSONClient) error
	}{
		{fn: testBackupCmd},
		{fn: testPrepareCmd},
		{fn: testPerformCmd},
		{fn: testCancelCmd},
		{fn: testRetrieveQueryCmd},
	}
	for index, testCase := range testCases {
		err := testCase.fn(t, jrpcClient)
		if err == nil {
			continue
		}
		assert.NotEqualf(t, err, types.ErrActionNotSupport, "test index %d", index)
		if strings.Contains(err.Error(), "rpc: can't find") {
			assert.FailNowf(t, err.Error(), "test index %d", index)
		}
	}
}

func testBackupCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := rpc.RetrieveBackupTx{}
	return jrpc.Call("retrieve.CreateRawRetrieveBackupTx", params, nil)
}

func testPrepareCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := rpc.RetrievePrepareTx{}
	return jrpc.Call("retrieve.CreateRawRetrievePrepareTx", params, nil)
}

func testPerformCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := rpc.RetrievePerformTx{}
	return jrpc.Call("retrieve.CreateRawRetrievePerformTx", params, nil)
}

func testCancelCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := rpc.RetrieveCancelTx{}
	return jrpc.Call("retrieve.CreateRawRetrieveCancelTx", params, nil)
}

func testRetrieveQueryCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var rep interface{}
	var params types.Query4Cli
	req := &pty.ReqRetrieveInfo{}
	params.Execer = "retrieve"
	params.FuncName = "GetRetrieveInfo"
	params.Payload = req
	rep = &pty.RetrieveQuery{}
	return jrpc.Call("Chain33.Query", params, rep)
}
