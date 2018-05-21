package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Grape struct {
	Name string `json:"name"`
	Region string `json:"region"`
	Vineyard string `json:"vineyard"`
	Block string `json:"block"`
	RowRange string `json:"rowRange"`
	Variety string `json:"variety"`
	Vintage  int `json:"vintage"`
	EstimatedWeight  double `json:"estimatedWeight"`
	ActualWeight  double `json:"actualWeight"`
}

type Winery struct {
	Name string `json:"name"`
	Volume  double `json:"volume"`
}

type Wine struct {
	Name string `json:"name"`
	Label string `json:"label"`
	CorkCap  string `json:"corkCap"`
	Status  string `json:"status"`
	Seller  string `json:"seller"`
	Brand  string `json:"brand"`
}

type Transaction struct {
	StageId double `json:"stageId"`
	Start: time.Time `json:"start"`
	End: time.Time `json:"end"`
}

type Product struct {
	CompanyName string `json:"companyName"`
	Grape Grape `json:"grape"`
	Winery Winery `json:"winery"`
	Wine Wine `json:"wine"`
	Transaction Transaction `json:"transaction"`
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	product := []Product{
		{
			CompanyName: "Hoggies Estate",
			Grape: {
				Name: "Gaven"
				Region: "Merbein", 
				Vineyard: "Thompson",
				Block: "2",
				RowRange: "1-3",
				Variety: "Muscat of Alexandria", 
				Vintage: 2018, 
				EstimatedWeight: "20",
				ActualWeight: "18",
			},
			Winery: {
				Vinery: "Trentham Estate",
				Volume: 5000,
			},
			Wine: {
				Name: "Best Bottlers",
				Label: "Hoggies",
				CorkCap: "",
				Status: "Labeled",
				Seller: "Liquid Shop",
				Brand: "Vintage Reserve Shiraz",
			}
		},
		{
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


func (s *SmartContract) queryProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	wineAsBytes, _ := APIstub.GetState(args[0])
	if wineAsBytes == nil {
		return shim.Error("Could not locate wine")
	}
	return shim.Success(wineAsBytes)
}