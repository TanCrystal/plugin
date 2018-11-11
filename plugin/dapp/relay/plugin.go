package relay

import (
	"github.com/33cn/chain33/pluginmgr"
	"github.com/33cn/plugin/plugin/dapp/relay/commands"
	"github.com/33cn/plugin/plugin/dapp/relay/executor"
	"github.com/33cn/plugin/plugin/dapp/relay/rpc"
	"github.com/33cn/plugin/plugin/dapp/relay/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.RelayX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.RelayCmd,
		RPC:      rpc.Init,
	})
}
