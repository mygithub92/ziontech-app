package main

import (
	"encoding/json"
	"fmt"
	"log"
	"errors"
	"math/rand"
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
	EstimatedWeight  float32 `json:"estimatedWeight"`
	ActualWeight  float32 `json:"actualWeight"`
}

type Winery struct {
	Name string `json:"name"`
	Volume  float32 `json:"volume"`
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
	StageId float32 `json:"stageId"`
	Start string `json:"start"`
	End string `json:"end"`
}

type Distribution struct {
	DriverId string `json:driverId`
	PlateNumber string `json:plateNumber`
}

type Product struct {
	Key int `json:"key"`
	CompanyName string `json:"companyName"`
	Grape Grape `json:"grape"`
	Winery Winery `json:"winery"`
	Wine Wine `json:"wine"`
	Distributions [] Distribution `json:"distributions`
	Transactions []Transaction `json:"transactions"`
}

var products = []Product{
	{
		Key: 1,
		CompanyName: "Hoggies Estate",
		Grape: Grape {
			Name: "Gaven",
			Region: "Merbein", 
			Vineyard: "Thompson",
			Block: "2",
			RowRange: "1-3",
			Variety: "Muscat of Alexandria", 
			Vintage: 2018, 
			EstimatedWeight: 20,
			ActualWeight: 18,
		},
		Winery: Winery {
			Name: "Trentham Estate",
			Volume: 5000,
		},
		Wine: Wine {
			Name: "Best Bottlers",
			Label: "Hoggies",
			CorkCap: "",
			Status: "Labeled",
			Seller: "Liquid Shop",
			Brand: "Vintage Reserve Shiraz",
		},
		Distributions: []Distribution {
			{
				DriverId: "KT2456",
				PlateNumber: "DWT345",
			},
		},
		Transactions: []Transaction {
			{
				StageId: 10,
				Start: "2015-03-23",
			},
		},
	},
	{
		Key: 2,
		CompanyName: "Penley",
		Grape: Grape {
			Region: "Coonawarra", 
			Vineyard: "Ladbroke",
			Block: "3",
			RowRange: "4-10",
			Variety: "Shiraz", 
			Vintage: 2016, 
			EstimatedWeight: 35,
			ActualWeight: 33,
		},
		Winery: Winery {
			Name: "Limestone Coast Wines",
			Volume: 5600,
		},
		Wine: Wine {
			Name: "Liquid Goods",
			Label: "Olivias",
			CorkCap: "",
			Status: "Not Labeled",
			Seller: "BWS",
			Brand: "Riesling",
		},
	},
}

var newProduct = Product{
	CompanyName: "Hoggies Estate",
	Grape: Grape {
		Name: "Gaven-new",
		Region: "Merbein-new", 
		Vineyard: "Thompson-new",
		Block: "2-new",
		RowRange: "1-3-new",
		Variety: "Muscat of Alexandria-new", 
		Vintage: 2018, 
		EstimatedWeight: 20,
		ActualWeight: 18,
	},
	Winery: Winery {
		Name: "Trentham Estate-new",
		Volume: 5000,
	},
	Wine: Wine {
		Name: "Best Bottlers-new",
		Label: "Hoggies-new",
		CorkCap: "",
		Status: "Labeled",
		Seller: "Liquid Shop",
		Brand: "Vintage Reserve Shiraz",
	},
	Distributions: []Distribution {
		{
			DriverId: "KT2456",
			PlateNumber: "DWT345",
		},
	},
	Transactions: []Transaction {
		{
			StageId: 10,
			Start: "2015-11-12",
			End: "2015-12-12",
		},
	},
}

func initLedger() {

	json, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}


func queryProduct() []byte {
	json, err := json.MarshalIndent(products, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return json;
}

func addProduct(newProduct Product) {
	products = append(products, newProduct);
}

func transferProduct(key int, product string) {
	updatedProduct := Product{}
	json.Unmarshal([]byte(product), &updatedProduct)

	if key < 0 {
		updatedProduct.Key = rand.Intn(1000)
		addProduct(updatedProduct)
		printProducts(products)
	} else {
		current, err := getProductByKey(key);
		if err != nil {
			log.Fatal(err)
		}
	
		lastTransaction := current.Transactions[len(current.Transactions) -1]
		switch stageId := lastTransaction.StageId; stageId {
			case 10://driver from grower to winery
				current.Distributions = append(current.Distributions, updatedProduct.Distributions[0])
				updatedProduct.Transactions[0].StageId = stageId + 10
				current.Transactions = append(current.Transactions, updatedProduct.Transactions[0])
			case 20://winery making wine
				current.Winery = updatedProduct.Winery
				updatedProduct.Transactions[0].StageId = stageId + 10
				current.Transactions = append(current.Transactions, updatedProduct.Transactions[0])
			case 30://driver from winery to bottler
				current.Distributions = append(current.Distributions, updatedProduct.Distributions[0])
				updatedProduct.Transactions[0].StageId = stageId + 10
				current.Transactions = append(current.Transactions, updatedProduct.Transactions[0])
			case 40://bottler label wine
				current.Wine = updatedProduct.Wine
				updatedProduct.Transactions[0].StageId = stageId + 10
				current.Transactions = append(current.Transactions, updatedProduct.Transactions[0])
			case 50://driver from bottler to warehouse
				current.Distributions = append(current.Distributions, updatedProduct.Distributions[0])
				updatedProduct.Transactions[0].StageId = stageId + 10
				current.Transactions = append(current.Transactions, updatedProduct.Transactions[0])
		}
		products[2] = current
		printProduct(current)
	}
}

func getProductByKey(key int) (product Product, err error) {
	for _, v := range products {
		if v.Key == key {
			product = v
			return
		}
	}
	err = errors.New("No record found")
	return 
}

func printProduct(product Product) {
	json, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
	fmt.Println("--------------------------------------------")
}

func printProducts(products []Product) {
	json, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
	fmt.Println("--------------------------------------------")
}

func main() {
	json, err := json.Marshal(newProduct)
	if err != nil {
		log.Fatal(err)
	}

	transferProduct(-1, string(json));
	transferProduct(81, string(json));
	transferProduct(81, string(json));
	transferProduct(81, string(json));
	transferProduct(81, string(json));
}