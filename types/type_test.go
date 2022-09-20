package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestFailed(t *testing.T) {
	t.Log(string(hexutil.MustDecode("0x736a69617364696a616a696f776e64736169646a73616a696f6461736f696a6473616e636f696a736461686f6964687361646f696168646a73616f69686a6461736a696f646a7361696f646a616f69646a61646f69736a61646a616f69646a61")))
	t.Log(hexutil.Bytes(crypto.Keccak256([]byte("notFound(uint256)"))[:4]).String())
}

/*
import Web3 from "web3"
const web3 = new Web3();
  console.log(
    web3.eth.abi.decodeParameter(
      "string",
      "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001e41646d696e3a2063616c6c6572206973206e6f74207468652061646d696e0000"
    )
  );
*/
