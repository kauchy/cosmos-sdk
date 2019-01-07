#!/bin/bash

set -x

CLI=gaiacli
CHAIN=game_of_stakes_3
VAL=cosmosvaloper16vgxx4302kcpw2z98a555w5lj5etnwsjr9cj2j
ADDRESS=cosmos16vgxx4302kcpw2z98a555w5lj5etnwsjx3v8xp
ACCOUNT=qos
DELAY=300

echo "Enter your key password:"
read -s password

while true
do
    last_withdraw_height=$($CLI query dist vdi $VAL --chain-id=$CHAIN --trust-node=true | jq -r '.fee_pool_withdrawal_height')
    latest_height=`gaiacli status|jq '.sync_info.latest_block_height'|bc`
    interval=`echo $latest_height - $last_withdraw_height | bc`
    if [ $interval -gt 200 ]; then
        sequence=$($CLI query account $ADDRESS --chain-id=$CHAIN --trust-node=true | jq -r '.value.sequence')
        steak=105
        amount_steak=$($CLI query account $ADDRESS --chain-id=$CHAIN --trust-node=true | jq -r '.value.coins[0].amount')
        if [[ $amount_steak -gt 0 && $amount_steak != "null" && $amount_steak -lt 1000000 ]]; then
           echo "About to stake ${amount_steak} steak"
           steak=`echo $amount_steak + $steak | bc`
        fi
        echo "${password}" | $CLI tx stake withdraw-delegate --from $ACCOUNT --amount ${steak}STAKE --chain-id $CHAIN --memo zz --async --sequence $sequence
    fi
echo "--------------------------------"
date
sleep $DELAY
done