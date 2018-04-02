// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define wine structure, with 4 properties.  
Structure tags are used by encoding/json library
*/
type Wine struct {
	CompanyName string `json:"companyName"`
	Region string `json:"region"`
	Vineyard string `json:"vineyard"`
	Block string `json:"block"`
	RowRange string `json:"rowRange"`
	Variety string `json:"variety"`
	Vintage  string `json:"vintage"`
	DateDelivered  string `json:"dateDelivered"`
	Vinery  string `json:"vinery"`
	EstimatedWeight  string `json:"estimatedWeight"`

	ActualWeight  string `json:"actualWeight"`
	Volume  string `json:"volume"`
	ReceivedFrom  string `json:"receivedFrom"`
	TransferredTo  string `json:"transferredTo"`
	BottlingCompany  string `json:"bottlingCompany"`
	
	Label  string `json:"label"`
	CorkCap  string `json:"corkCap"`
	Status  string `json:"status"`
	Seller  string `json:"seller"`
	Brand  string `json:"brand"`
}

/*
 * The Init method *
 called when the Smart Contract "wine-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "wine-chaincode"
 The app also specifies the specific smart contract function to call with args
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryWine" {
		return s.queryWine(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordWine" {
		return s.recordWine(APIstub, args)
	} else if function == "queryAllWine" {
		return s.queryAllWine(APIstub)
	} else if function == "changedByVinery" {
		return s.changedByVinery(APIstub, args)
	}else if function == "changedByBottler" {
		return s.changedByBottler(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The querywine method *
Used to view the records of one particular wine
It takes one argument -- the key for the wine in question
 */
func (s *SmartContract) queryWine(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	wineAsBytes, _ := APIstub.GetState(args[0])
	if wineAsBytes == nil {
		return shim.Error("Could not locate wine")
	}
	return shim.Success(wineAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 wine catches)to our network
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	wine := []Wine{
		Wine{
			CompanyName: "Hoggies Estate",
			Region: "Merbein", 
			Vineyard: "Thompson",
			Block: "2",
			RowRange: "1-3",
			Variety: "Muscat of Alexandria", 
			Vintage: "2018", 
			DateDelivered: "2018-04-12",
			Vinery: "Trentham Estate",
			EstimatedWeight: "20",
			ActualWeight: "18",
			Volume: "5000",
			ReceivedFrom: "",
			TransferredTo: "",
			BottlingCompany: "Best Bottlers",
			Label: "Hoggies",
			CorkCap: "",
			Status: "Labeled",
			Seller: "Liquid Shop",
			Brand: "Vintage Reserve Shiraz",
		},
		Wine{
			CompanyName: "Penley",
			Region: "Coonawarra", 
			Vineyard: "Ladbroke",
			Block: "3",
			RowRange: "4-10",
			Variety: "Shiraz", 
			Vintage: "2016", 
			DateDelivered: "2017-12-23",
			Vinery: "Limestone Coast Wines",
			EstimatedWeight: "35",
			ActualWeight: "33",
			Volume: "5600",
			ReceivedFrom: "",
			TransferredTo: "",
			BottlingCompany: "Liquid Goods",
			Label: "Olivias",
			CorkCap: "",
			Status: "Not Labeled",
			Seller: "BWS",
			Brand: "Riesling",
		},
	}

	i := 0
	for i < len(wine) {
		fmt.Println("i is ", i)
		wineAsBytes, _ := json.Marshal(wine[i])
		APIstub.PutState(strconv.Itoa(i+1), wineAsBytes)
		fmt.Println("Added", wine[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordwine method *
Fisherman like Sarah would use to record each of her wine catches. 
This method takes in five arguments (attributes to be saved in the ledger). 
 */
func (s *SmartContract) recordWine(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	var wine = Wine{ 
		CompanyName:     args[1],
		Region:          args[2], 
		Vineyard:        args[3],
		Block:           args[4],
		RowRange:        args[5],
		Variety:         args[6], 
		Vintage:         args[7], 
		DateDelivered:   args[8], 
		Vinery:          args[9],
		EstimatedWeight: args[10],
	}

	wineAsBytes, _ := json.Marshal(wine)
	err := APIstub.PutState(args[0], wineAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record wine catch: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllwine method *
allows for assessing all the records added to the ledger(all wine catches)
This method does not take any arguments. Returns JSON string containing results. 
 */
func (s *SmartContract) queryAllWine(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllwine:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * The changewineHolder method *
The data in the world state can be updated with who has possession. 
This function takes in 2 arguments, wine id and new holder name. 
 */
func (s *SmartContract) changedByVinery(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	// if len(args) != 6 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 6")
	// }

	wineAsBytes, _ := APIstub.GetState(args[0])
	if wineAsBytes == nil {
		return shim.Error("Could not locate wine")
	}
	wine := Wine{}

	json.Unmarshal(wineAsBytes, &wine)
	// Normally check that the specified argument is a valid holder of wine
	// we are skipping this check for this example
	wine.ActualWeight = args[1]
	wine.Volume = args[2]
	wine.ReceivedFrom = args[3]
	wine.TransferredTo = args[4]
	wine.BottlingCompany = args[5]

	wineAsBytes, _ = json.Marshal(wine)
	err := APIstub.PutState(args[0], wineAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change wine holder: %s", args[0]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) changedByBottler(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	// if len(args) != 6 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 6")
	// }

	wineAsBytes, _ := APIstub.GetState(args[0])
	if wineAsBytes == nil {
		return shim.Error("Could not locate wine")
	}
	wine := Wine{}

	json.Unmarshal(wineAsBytes, &wine)
	// Normally check that the specified argument is a valid holder of wine
	// we are skipping this check for this example
	wine.Label = args[1]
	wine.CorkCap = args[2]
	wine.Status = args[3]
	wine.Seller = args[4]
	wine.Brand = args[5]

	wineAsBytes, _ = json.Marshal(wine)
	err := APIstub.PutState(args[0], wineAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change wine holder: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}