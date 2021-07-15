# Orderbooks

## Introduction

A library for the buycoins orderbooks API.

## Usage

### Install Package

```bash
go get github.com/buycoinsresearch/buycoins-orderbook-go
```

### Buycoins Orderbooks

* Initialise 

```go
package main

import (
	"github.com/buycoinsresearch/buycoins-orderbook-go"
)

var authorize = orderbooks.Buycoins(blahburibdblahbeubblah, blahblahblah)
...
```

>**NOTE**<br/>
>Pleasee ensure you pass both the public key and the secret key

* Get orders

```go
...

func main () {
    getOrders, err := authorize.GetOrders("btc_usdt", "successful", "buy")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", getOrders)
}
```

* Cancel order

```go
...

func main () {
    cancelOrder, err := authorize.CancelOrder("UHJvT3JkZXItYzExYTY2ZmEtMzdiYy00ZDcyLWJmODgtNzg2NzRkNzdhZTlj")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", cancelOrder)
}
```

* Estimate order fees

```go
...

func main () {
    getProOrderFee,err := authorize.GetProOrderFees("market_order", "btc_usdt", "sell", 0.001)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", getProOrderFee)
}
```

* Post Market Order 

```go
...

func main () {
    postProMarketOrder, err := authorize.PostProMarketOrder("btc_usdt", 5.00, "buy")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", postProMarketOrder)
}
```

* Post Liimit Order 

```go
...

func main () {
    postProLimitOrder, err := authorize.PostProLimitOrder("btc_usdt", 5.00, 26000000.00, "buy", "good_til_cancelled")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", postProLimitOrder)
}
```


>**NOTE**<br/>
>Check the `example` directory to see a sample implementation.
