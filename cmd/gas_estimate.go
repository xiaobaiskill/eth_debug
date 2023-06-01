package cmd

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
	"github.com/xiaobaiskill/eth_debug/share"
)

var (
	client *ethclient.Client
	err    error
)
var esatmate = &cli.Command{
	Name:  "gas_estimate",
	Usage: "gas estimate",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     share.From,
			Required: false,
			Value:    "0x0000000000000000000000000000000000000000",
		},
		&cli.StringFlag{
			Name:     share.To,
			Required: false,
			Value:    "",
		},
		&cli.StringFlag{
			Name:     share.Value,
			Required: false,
			Value:    "0x0",
		},

		&cli.StringFlag{
			Name:     share.Data,
			Required: true,
			Value:    "",
		},
	},
	Action: func(c *cli.Context) error {
		rpcClient, err := rpc.Dial(c.String(share.RpcUrl))
		if err != nil {
			return err
		}
		client = ethclient.NewClient(rpcClient)

		var to *common.Address
		if c.String(share.To) != "" {
			toAddr := common.HexToAddress(c.String(share.To))
			to = &toAddr
		}

		callMsg := ethereum.CallMsg{
			From:  common.HexToAddress(c.String(share.From)),
			To:    to,
			Value: hexutil.MustDecodeBig(c.String(share.Value)),
			Data:  hexutil.MustDecode(c.String(share.Data)),
		}
		gas, err := client.EstimateGas(context.Background(), callMsg)
		// fmt.Println(callMsg)
		// var hex hexutil.Uint64
		// err = rpcClient.CallContext(context.Background(), &hex, "eth_estimateGas", toCallArg(callMsg))

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("need gas: ", gas)
		}
		return nil
	},
}
