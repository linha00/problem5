# Wallet Blockchain README

## Introduction

The Wallet Blockchain is a decentralized application built using the Cosmos SDK and Ignite CLI. It allows users to create, manage, and interact with digital wallets. Each wallet stores essential information, including the wallet balance, creator, and unique ID.

## Features

- **Create Wallet**: Users can create new wallets with an initial balance.
- **List Wallets**: Retrieve a list of wallets filtered by balance.
- **Show Wallet Details**: Get detailed information about a specific wallet.
- **Update Wallet Balance**: Modify the balance of an existing wallet.
- **Delete Wallet**: Remove a wallet from the blockchain.

## Commands

- **Create Wallet**:
`walletd tx wallet create-wallet <initial-balance> --from <account> --chain-id wallet`
- **List Wallets**:
`walletd q wallet list-wallet`
- **Show Wallet Details**:
`walletd q wallet show-wallet <wallet-id>`
- **Update Wallet Balance**:
`walletd tx wallet update-wallet <new-balance> <wallet-id> --from <creator> --chain-id wallet`
- **Delete Wallet**:
`walletd tx wallet delete-wallet <wallet-id> --from <creator> --chain-id wallet`