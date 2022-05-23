package client

import (
	"context"
	"fmt"
	"os"
	"tz/models"

	"github.com/joho/godotenv"
	rpc "github.com/ybbus/jsonrpc/v3"
)

const (
	endpoint         = "https://eth.getblock.io/mainnet/"
	lastBlockNumber  = "eth_blockNumber"
	getBlockByNumber = "eth_getBlockByNumber"
	getBlockByHash   = "eth_getBlockByHash"
)

type client struct {
	rpcClient rpc.RPCClient
}

func NewClient() (*client, error) {
	//parse .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	//get api key from env
	apiKey := os.Getenv("apiKey")
	if apiKey == "" {
		return nil, fmt.Errorf("api key is empty")
	}
	return &client{
		rpcClient: rpc.NewClientWithOpts(endpoint, &rpc.RPCClientOpts{
			CustomHeaders: map[string]string{
				"x-api-key": apiKey,
			},
		}),
	}, nil
}

func (c *client) GetLastBlockNumber(ctx context.Context) (string, error) {
	var number string
	err := c.rpcClient.CallFor(ctx, &number, lastBlockNumber)
	if err != nil {
		return "", err
	}
	return number, nil
}

func (c *client) GetBlockByNumber(ctx context.Context, number string) (*models.Block, error) {
	block := models.Block{}
	err := c.rpcClient.CallFor(ctx, &block, getBlockByNumber, number, true)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (c *client) GetBlockByHash(ctx context.Context, hash string) (*models.Block, error) {
	block := models.Block{}
	err := c.rpcClient.CallFor(ctx, &block, getBlockByHash, hash, true)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (c *client) FindBlocks(ctx context.Context, blocksCount int) ([]*models.Block, error) {
	//get last block number
	number, err := c.GetLastBlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	//start fetching last n blocks
	blockCount := 0
	blocks := make([]*models.Block, 0)
	//at start get last block by number
	block, err := c.GetBlockByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	blocks = append(blocks, block)
	blockCount++
	//start fetchung last 99 blocks
	for blockCount < blocksCount {
		fmt.Printf("getting block %d\n", blockCount+1)
		block, err = c.GetBlockByHash(ctx, block.ParentHash)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
		blockCount++
	}
	return blocks, nil
}
