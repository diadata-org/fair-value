# fair-value feeder

<p align="center">
    <img src="./assets/DIA_logo.png" alt="Dia logo" width="200" height="auto" style="padding: 20px;">
</p>

The node setup instructions are available in our [Wiki](https://github.com/diadata-org/fair-value/wiki) page!

This repository hosts a self-contained application for running a fair-value feeder in the Lumina oracle network. The main purpose of a fair-value feeder is to determine the price of reserve backed assets. It consists of a class of scrapers which collect supply and price data from various on-chain sources. As there is a variety of implementations of such reserve backed asset products, the class consists of different ways on how to come up with a price for the considered asset. For the same reason, the meaning of the values determined by the feeder vary depending on the considered asset class.
In general, the executable feeds a smart contract by writing the following 5 values:
```
fairValue:   The dimensionless price of an asset backed by a single reserve asset given by the amount of the reserve divided by the amount of the asset.
valueUsd:    The USD price of the asset.
numerator:   The amount of the reserve asset in case of a single reserve asset. The USD price of the total of reserve assets in case of multiple reserve assets.
denominator: The amount of the asset in case of a single reserve asset. The USD price of the total amount of the asset in case of multiple reserve assets.
timestamp:   Unix timestamp of the oracle update.
```
It does so by fetching the total supply and the associated reserve tokens along with their amounts for various reserve-backed assets.\
It is worth noting that `fairValue` is only non-zero in case the considered asset has a single reserve asset, as in case of multiple reserve assets, we cannot assign a meaningful number to the considered quotient.

## Requirements

- Docker or Docker Compose installed

- Container has minimal resource requirements, making it suitable for most machines, including Windows, macOS, Linux, and Raspberry Pi, as long as Docker is installed.

- A private key from MetaMask or any other EVM wallet. Alternatively to generate a private key effortlesly, [eth-keys](https://github.com/ethereum/eth-keys) tool can be used for this.

- DIA tokens in your wallet. For running a node in testnet, you can request tokens from the [faucet](https://faucet.diadata.org).

## Quick Start

> **NOTE:** This guide is using docker compose for running a feeder node. You can explore alternative deployment methods [here](https://github.com/diadata-org/decentral-feeder/wiki/Node-deployment#alternative-deployment-methods).

### Navigate to the Docker Compose Folder

- Locate the `docker-compose` folder in this repository.
- Inside, you will find a file named `docker-compose.yaml`.

### Configure Environment Variables

- Create a `.env` file in the same directory as `docker-compose.yaml`. This file should contain the following variables:

```
SOURCE=TwelveData

# Node operator configuration (MUST be filled by node operators)
PRIVATE_KEY=
NODE_OPERATOR_NAME=""
DEPLOYED_CONTRACT=""
BLOCKCHAIN_NODE=""
BACKUP_NODE=""
IMAGE_TAG=""

# Monitoring
PUSHGATEWAY_USER=
PUSHGATEWAY_PASSWORD=
PUSHGATEWAY_URL=

# Chain configuration
# https://github.com/diadata-org/decentral-data-feeder/wiki/Chain-Info
CHAIN_ID=
UPDATE_SECONDS=80
CONFIG_UPDATE_SECONDS=86400
WRITE_TICKER_SECONDS=3600
```

Please refer to the [.env.example](./docker-compose/.env.example) file for the exact environment variables to set.

- Open a terminal in the `docker-compose` folder and start the deployment by running:
  ```bash
  docker-compose up
  ```

### Retrieve Deployed Contract

- Once the container is deployed with `DEPLOYED_CONTRACT` env variable empty the logs will display the deployed contract address in the following format:
  ```plaintext
  │ time="2024-11-25T11:30:08Z" level=info msg="Contract pending deploy: 0xxxxxxxxxxxxxxxxxxxxxxxxxx."
  ```
- Copy the displayed contract address (e.g., `0xxxxxxxxxxxxxxxxxxxxxxxxxx`) and stop the container with `docker rm -f <container_name>`.

- Update your `.env` file with `DEPLOYED_CONTRACT` variable mentioned above. Redeployed the container with `docker-compose up -d`

  ```plaintext
  DEPLOYED_CONTRACT=0xxxxxxxxxxxxxxxxxxxxxxxxxx
  ```

- Check if the container is running correctly by viewing the logs. Run the following command:

  ```bash
  docker-compose logs -f
  ```


- You can optionally cleanup the deployment once you're done by running:

  ```
  docker rm -f <container_name>
  ```

- Verify the container has been removed:
  ```
  docker ps -a
  ```

## Error Handling

> **NOTE:** This guide is specific to docker compose. You can check how to handle errors based on your deployment method [here](https://github.com/diadata-org/decentral-feeder/wiki/Error-Handling).

If any issues arise during deployment, follow these steps:

#### Check Logs:

- Run `docker-compose logs -f`

#### Verify Environment Variables:

- Ensure all required variables (`PRIVATE_KEY`, `DEPLOYED_CONTRACT`,...etc) are correctly set by checking the `.env` file.

#### Restart Deployment:

- ```bash
  docker-compose down && docker-compose up -d
  ```

#### Update or Rebuild:

- Ensure you're using the correct image version:
  ```bash
  docker pull diadata/decentralized-data-feeder:<VERSION>
  ```
- Apply fixes and redeploy.

## Documentation

For the full node deployment instructions, you can visit our [Wiki](https://github.com/diadata-org/fair-value/wiki) page.

To learn about DIA's oracle stacks, you can visit our documentation [here](https://docs.diadata.org/).

## Issues

To report bugs or suggest enhancements, you can create a [Github Issue](https://docs.github.com/en/issues/tracking-your-work-with-issues/using-issues/creating-an-issue) in the repository.

## Contribution Guidelines

Coming soon...

## Community

You can find our team on the following channels:

- [Discord](https://discord.com/invite/RjHBcZ9mEH)
- [Telegram](https://t.me/diadata_org)
- [X](https://x.com/DIAdata_org)
