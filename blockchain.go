package main

import (
  "crypto/sha256"
  "encoding/json"
  "fmt"
)

type Blockchain struct {
  Chain []Block
  CurrentTransactions []Transaction
}

func NewBlockchain() *Blockchain {
  return &Blockchain{}
}

func (b Blockchain) NewBlock(proof string, previousHash string) {
  block := Block{Index: TODO, Timestamp: TODO, Transactions: TODO, Proof: proof, PreviousHash: previousHash}


}

func (b Blockchain) NewTransaction(sender string, recipient string, amount int) {
  transaction := Transaction{Sender: sender, Recipient: recipient, Amount: amount}

  b.CurrentTransactions = append(b.CurrentTransactions, transaction)
}

func (b Blockchain) LastBlock() *Block {
  return nil
}

func Hash() {

}

func main() {
  fmt.Println("Chain")
}
