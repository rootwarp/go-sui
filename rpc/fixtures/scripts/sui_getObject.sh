#!/bin/bash

. run_env

curl -s -X POST \
    -H "Content-Type: application/json" \
    -d '{
        "jsonrpc": "2.0",
        "id": 1, "method":
        "sui_getObject",
        "params": [
            "0x845d6a0756208f107ebf6d2676641d4c28502e09dbfa9825f15baebc08c0c046",
            {
                "showType": true,
                "showOwner": true,
                "showPreviousTransaction": true,
                "showDisplay": true,
                "showContent": true,
                "showBcs": true,
                "showStorageRebast": true
            }
        ]
    }' \
    $RPC | jq . > ../sui_getObject.json

