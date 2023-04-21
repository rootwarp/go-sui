#!/bin/bash

RPC=https://fullnode.testnet.sui.io:443

curl -s -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc": "2.0", "id": 1, "method": "suix_getBalance", "params": ["0xb878abfe4fbd421c70c7c725a3c012bd0e70eb0f42ed0b05f0944b9616f3710d"]}' \
    $RPC | jq . > suix_getBalance.json
