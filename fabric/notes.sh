# Start network:
# =============

# cd /Users/goutaudi/Blockchain/hyperledger/fabric/fabric-samples/chaincode-docker-devmode
docker-compose -f docker-compose-simple.yaml up
## docker-compose -f docker-compose-simple.yaml down

# On chaincode:
# ============
docker exec -it chaincode bash
go build
CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./start.sh

# DEPLOY DEX:
# ==========
docker exec -it cli bash
peer chaincode install -p chaincodedev/chaincode/dex -n mycc -v 0
peer chaincode instantiate -n mycc -v 0 -c '{"Args":[]}' -C myc

peer chaincode invoke -n mycc -c '{"Args":["init_escrow", "guillaume", "robert", "500", "1"]}' -C myc

peer chaincode invoke -n mycc -c '{"Args":["get", "buyer"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["get", "seller"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["get", "dollar_amount"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["get", "eth_amount"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["get", "buyer_ack"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["get", "seller_ack"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["get", "transfer_done"]}' -C myc

peer chaincode invoke -n mycc -c '{"Args":["acknowledge_eth_transfer"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["get", "seller_ack"]}' -C myc

peer chaincode invoke -n mycc -c '{"Args":["acknowledge_dollar_transfer"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["get", "buyer_ack"]}' -C myc

peer chaincode invoke -n mycc -c '{"Args":["trigger_transaction"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["get", "transfer_done"]}' -C myc
