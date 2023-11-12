package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type DealContract struct {
	ID    		string `json:"dealID"`
	RepAndRes 	string `json:"repAndRes"`
	Etc  		string `json:"etc"`
	Price   string `json:"proce"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	dealContracts := []DealContract{
		{ID : "1", RepAndRes : "판매자는 물품을 인도하기 전, 물품의 하자나 손상 여부를 확인하여야 합니다.", Etc : "책임은 판매자가 부담합니다.",Price :"300,000"},
		{ID : "2", RepAndRes : "구매자는 물품 수령 후, 물품에 대한 하자나 손상 여부를 확인하여야 합니다.", Etc : "책임은 구매자가 부담합니다.",,Price :"530,000"},
	}

	for _, dealContract := range dealContracts {
		dealContractJSON, err := json.Marshal(dealContract)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(dealContract.ID, dealContractJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, repAndRes string, etc string, price string) error {
	exists, err := s.DealContractExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the deal %s already exists", id)
	}

	dealContract := DealContract{
		ID:             id,
		RepAndRes:      repAndRes,
		Etc:			etc,
		Price:    price,
	}
	dealContractJSON, err := json.Marshal(dealContract)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(id, dealContractJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	return  nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*DealContract, error) {
	dealContractJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if dealContractJSON == nil {
		return nil, fmt.Errorf("the deal %s does not exist", id)
	}

	var dealContract DealContract
	err = json.Unmarshal(dealContractJSON, &dealContract)
	if err != nil {
		return nil, err
	}

	return &dealContract, nil
}


// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.DealContractExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}


// DealContractExists returns true when asset with given ID exists in world state
func (s *SmartContract) DealContractExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	dealContractJSON, err := ctx.GetStub().GetState(id)
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
