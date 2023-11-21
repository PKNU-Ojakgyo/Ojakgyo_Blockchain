package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provIDes functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type DealContract struct {
	DealID    		string `json:"dealID"`
	RepAndRes 	string `json:"repAndRes"`
	Note  		string `json:"note"`
	Price   string `json:"price"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	dealContracts := []DealContract{
		{DealID: "1", RepAndRes : "판매자는 물품을 인도하기 전, 물품의 하자나 손상 여부를 확인하여야 합니다.", Note : "책임은 판매자가 부담합니다.",Price :"300000"},
		{DealID: "2", RepAndRes : "구매자는 물품 수령 후, 물품에 대한 하자나 손상 여부를 확인하여야 합니다.", Note : "책임은 구매자가 부담합니다.",Price :"530000"},
	}

	for _, dealContract := range dealContracts {
		dealContractJSON, err := json.Marshal(dealContract)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(dealContract.DealID, dealContractJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealID string, repAndRes string, note string, price string) error {
	exists, err := s.DealContractExists(ctx, dealID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the deal %s already exists", dealID)
	}

	dealContract := DealContract{
		DealID:             dealID,
		RepAndRes:      repAndRes,
		Note:			note,
		Price:    price,
	}
	dealContractJSON, err := json.Marshal(dealContract)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(dealID, dealContractJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	return  nil
}

// ReadAsset returns the asset stored in the world state with given ID.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, dealID string) (*DealContract, error) {
	dealContractJSON, err := ctx.GetStub().GetState(dealID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if dealContractJSON == nil {
		return nil, fmt.Errorf("the deal %s does not exist", dealID)
	}

	var dealContract DealContract
	err = json.Unmarshal(dealContractJSON, &dealContract)
	if err != nil {
		return nil, err
	}

	return &dealContract, nil
}


// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, dealID string) error {
	exists, err := s.DealContractExists(ctx, dealID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", dealID)
	}

	return ctx.GetStub().DelState(dealID)
}


// DealContractExists returns true when asset with given ID exists in world state
func (s *SmartContract) DealContractExists(ctx contractapi.TransactionContextInterface, dealID string) (bool, error) {
	dealContractJSON, err := ctx.GetStub().GetState(dealID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return dealContractJSON != nil, nil
}

func main() {
	cc, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting ABstore chaincode: %s", err)
	}
}