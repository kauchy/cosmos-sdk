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
    if [ $interval -gt 400 ]; then
        sequence=$($CLI query account $ADDRESS --chain-id=$CHAIN --trust-node=true | jq -r '.value.sequence')
        echo "${password}" | $CLI tx dist withdraw-rewards --is-validator --from $ACCOUNT --chain-id $CHAIN --memo zz --async --sequence $sequence

        sequence=`echo $sequence + 1 | bc`
        sleep 10
        steak=145
        amount_steak=$($CLI query account $ADDRESS --chain-id=$CHAIN --trust-node=true | jq -r '.value.coins[0].amount')
        if [[ $amount_steak -gt 0 && $amount_steak != "null" && $amount_steak -lt 1000000 ]]; then
           echo "About to stake ${amount_steak} steak"
           steak=`echo $amount_steak + $steak | bc`
        fi
        echo "${password}" | $CLI tx stake delegate --amount ${steak}STAKE --from $ACCOUNT --validator $VAL --chain-id $CHAIN --memo zz --async --sequence $sequence
    fi
echo "--------------------------------"
date
sleep $DELAY
done