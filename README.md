# eth_debug

### 简介
```
简单的debug 查询工具
1、 tx 交易失败, 错误原因查询(call_tracer)
2、 gas 评估及错误原因(gas_estimate)
```

### debug_traceTransaction
```
./eth_debug --rpcUrl https://bas-archive.ankr.com/rpc call_tracer \
--tx 0x6625d6ecd33fc38dd49feefbe6d2e1da35a7028805d03eb685e79bf397dc5221 \
--debug
```

### eth_estimateGas
```
./eth_debug --rpcUrl https://rpc.ankr.com/bsc gas_estimate \
--data 0x5e583a5a0000000000000000000000005704075803a122fc5afc8b60f07b84b77e065b5e0000000000000000000000000000000000000000000000000000000005eeb12c0000000000000000000000000000000000000000000000000000000000003fde0000000000000000000000005257ad0c2dcffcc0b3a30eaed151e76ff441eb60 \
--to 0x15443be27EfAb7B027051c7c5e2894fE030Db043 \
--from 0x5257ad0c2dCFfCC0B3a30eAed151E76fF441Eb60 \
--value 0x0
```