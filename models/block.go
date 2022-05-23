package models

type Block struct {
	Difficulty       string        `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Hash             *string       `json:"hash,omitempty"`
	LogsBloom        *string       `json:"logsBloom,omitempty"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            *string       `json:"nonce,omitempty"`
	Number           *string       `json:"number,omitempty"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        string        `json:"timestamp"`
	TotalDifficulty  string        `json:"totalDifficulty"`
	Transactions     []Transaction `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []string      `json:"uncles"`
}

type Transaction struct {
	BlockHash            *string  `json:"blockHash,omitempty"`
	BlockNumber          *string  `json:"blockNumber,omitempty"`
	From                 string   `json:"from"`
	Gas                  string   `json:"gas"`
	GasPrice             string   `json:"gasPrice"`
	MaxFeePerGas         *string  `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *string  `json:"maxPriorityFeePerGas,omitempty"`
	Hash                 string   `json:"hash"`
	Input                string   `json:"input"`
	Nonce                string   `json:"nonce"`
	To                   *string  `json:"to,omitempty"`
	TransactionIndex     *string  `json:"transactionIndex,omitempty"`
	Value                string   `json:"value"`
	Type                 string   `json:"type"`
	AccessList           []Access `json:"accessList,omitempty"`
	ChainId              *string  `json:"chainId,omitempty"`
	V                    string   `json:"v"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
}

type Access struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}
