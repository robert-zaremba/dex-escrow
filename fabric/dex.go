/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	fmt.Printf("Received request to call funtion %s.", fn)
	switch fn {
	case "init_escrow":
			result, err = init_escrow(stub, args)
	case "acknowledge_eth_transfer":
			result, err = acknowledge_eth_transfer(stub, args)
	case "acknowledge_dollar_transfer":
			result, err = acknowledge_dollar_transfer(stub, args)
  // Will use this function for the demo to shortcut everything
	case "trigger_transaction":
			result, err = trigger_transaction(stub, args)
	case "set":
			result, err = set(stub, args)
	case "get":
			result, err = get(stub, args)
	default:
			fmt.Printf("Unknown function %s.", fn)
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success([]byte(result))
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}

func init_escrow(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 4 {
		return "", fmt.Errorf("Incorrect arguments.")
	}
	set(stub, []string {"buyer", args[0]});
	set(stub, []string {"seller", args[1]});
	set(stub, []string {"dollar_amount", args[2]});
	set(stub, []string {"eth_amount", args[3]});
	set(stub, []string {"buyer_ack", "false"});
	set(stub, []string {"seller_ack", "false"});
	set(stub, []string {"dollar_transfer_ready", "false"});
	set(stub, []string {"eth_transfer_ready", "false"});
	set(stub, []string {"transfer_done", "false"});
	return "", nil
}

func generate_eth_contract(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf("Incorrect arguments.")
	}
	// TODO: Initialize ETH Smart generate_eth_contract using ETH GO Client
	return "", nil
}

func acknowledge_eth_transfer(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf("Incorrect arguments.")
	}
	return set(stub, []string {"seller_ack", "true"});
	// TODO: Trigger transfer if both seller and buyer have acked their transfers
}

func acknowledge_dollar_transfer(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf("Incorrect arguments.")
	}
	return set(stub, []string {"buyer_ack", "true"});
	// TODO: Trigger transfer if both seller and buyer have acked their transfers
}

// Helper function to verify ETH has indeed beed transferred
func verify_eth_transfer(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf("Incorrect arguments.")
	}
	// TODO: Query ETH smart contract to verify ETH have been transferred
	return set(stub, []string {"eth_transfer_ready", "true"});
}

// Helper function to verify dollars  has indeed beed transferred
func verify_dollar_transfer(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf("Incorrect arguments.")
	}
	// TODO: Query banking interface to verify USD have been transferred
	return set(stub, []string {"dollar_transfer_ready", "true"});
}

// Function that releases the escrow
func trigger_transaction(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf("Incorrect arguments.")
	}
	// TODO: Check ETH and Dollar transactions are ready anbd trigger them
	return set(stub, []string {"transfer_done", "true"});
}


// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
