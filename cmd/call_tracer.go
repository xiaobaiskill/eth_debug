package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
	"github.com/xiaobaiskill/eth_debug/share"
	"github.com/xiaobaiskill/eth_debug/types"
)

var call = &cli.Command{
	Name:  "call_tracer",
	Usage: "call trace",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  share.Debug,
			Value: false,
		},

		&cli.StringFlag{
			Name:     share.Tx,
			Required: true,
			Value:    "",
		},
	},
	Action: func(c *cli.Context) error {
		client, err := rpc.DialContext(context.Background(), c.String(share.RpcUrl))
		if err != nil {
			return err
		}

		var res = types.CallFrame{}
		err = client.Call(&res, "debug_traceTransaction", c.String(share.Tx), &types.Config{Tracer: "callTracer"})
		if err != nil {
			return err
		}
		if res.Error != "" {
			if strings.HasPrefix(res.Output, "0x08c379a0") {
				res, err := decodeParameter("string", hexutil.MustDecode(fmt.Sprintf("0x%s", res.Output[10:])))
				if err != nil {
					return err
				}
				fmt.Println(res[0].(string))
			} else {
				if res.Output == "" {
					fmt.Println(res.Error)
				} else {
					fmt.Println(res.Output)
				}
			}
		}
		if c.Bool(share.Debug) {
			bytesData, _ := json.Marshal(res)
			fmt.Println(string(bytesData))
		}
		return nil
	},
}

func decodeParameter(typ string, data []byte) ([]interface{}, error) {
	t, err := abi.NewType(typ, "", nil)
	if err != nil {
		return nil, err
	}
	args := abi.Arguments{{
		Type: t,
	},
	}
	return args.UnpackValues(data)
}
