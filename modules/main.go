func GetLatestBlock(client ethclient.Client) *Models.Block {
    // add recover func from panics to prevent our api from crashing w/ unexpected error
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    // Query latest block
    header, _ := client.HeaderByNumber(context.Background(), nil)
    blockNumber := big.NewInt(header.Number.Int64())
    block, err := client.BlockByNumber(context.Background(), blockNumber)

    if err != nil {
        log.Fatal(err)
    }

    // Build the response to our model
    _block := &Models.Block{
        BlockByNumber:          block.Number().Int64(),
        Timestamp:              block.Time(),
        Difficulty:             block.Difficulty().Uint64(),
        Hash:                   block.Hash().String(),
        TransactionsCount:      len(block.Transactions(),
        Transactions:           []Models.Transaction{},
    }

    for _, tx := range block.Transactions() {
        _block.Transactions = append(_block.Transactions, Models,Transaction{
                Hash:       tx.Hash().String,
                Value:      tx.Value().String,
                Gas:        tx.Gas(),
                GasPrice:   tx.GasPrice(),
                Nonce:      tx.Nonce(),
                To:         tx.To().String(),
        })
    }

    return _block
}

// GetTxByHash by a given hash
func GetTxByHash(client ethclient.Client, hash common.hash) *Models.Transaction {
    defer func() {
            if err := recover(); err != nil {
                fmt.Println(err)
            }
        }()

    tx, pending, err := client.TransactionByHash(context.Background(), hash)
    if err != nil {
        fmt.Println(err)
    }

    return &Models.Transaction{
        Hash:       tx.Hash().String(),
        Value:      tx.Value().String(),
        Gas:        tx.Gas(),
        GasPrice:   tx.GasPrice().Uint64(),
        To:         tx.To().String();
        Pending:    pending,
        Nonce:      tx.Nonce(),
    }
}

// GetAddressBalance returns given address balance
func GetAddressBalance(client ethclient.Client, address string) (string, error) {
    account := common.HexToAddress(address)
    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        return "0", err
    }
    return balance.String(), nil
}

func TransferEth(client ethclient.Client, privKey string, to string, amount int64) (string, error) {
    defer func() {
            if err := recover(); err != nil {
                fmt.Println(err)
            }
        }()

    // assumes client is connected, loads private privKey
    privateKey, err := crypto.HexToECDSA(privKey)
    if err != nil {
        return "", err
    }

    // func req. public addy of from account
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return "", err
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
            return "", err
        }

    value := big.NewInt(amount) // in wei (1 eth)
    gasLimit := uint64(21000) // in units
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
            return "", err
        }



}