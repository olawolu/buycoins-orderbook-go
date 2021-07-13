package orderbooks

import (
	"context"
	b64 "encoding/base64"
	"fmt"

	"log"

	"github.com/machinebox/graphql"
)

const endpoint = "https://backend.buycoins.tech/api/graphql"

type configCredentials struct {
	basicAuth string
}

func Buycoins(publicKey, secretKey string) configCredentials {
	auth := "Basic " + b64.URLEncoding.EncodeToString([]byte(publicKey+":"+secretKey))
	return configCredentials{
		basicAuth: auth,
	}
}

func (config configCredentials) GetOrders(coinPair, status, side string) (getProOrders, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(`
		query ($pair_: Pair!, $status_: ProOrderStatus!, $side_: OrderSide!) {
			getProOrders (pair:$pair_, status:$status_, side:$side_) {
				edges {
				  node {
					id
					pair
					price
					side
					status
					timeInForce
					orderType
					fees
					filled
					total
					initialBaseQuantity
					initialQuoteQuantity
					remainingBaseQuantity
					remainingQuoteQuantity
					meanExecutionPrice
					engineMessage
				  }
    			}
			}
		}
	`)
	req.Var("pair_", coinPair)
	req.Var("status_", status)
	req.Var("side_", side)
	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()
	res := struct {
		GetProOrders struct {
			Edges []struct {
				Node  struct {
					Id                     string
					Pair                   string
					Price                  string
					Side                   string
					Status                 string
					TimeInForce            string
					OrderType              string
					Fees                   string
					Filled                 string
					Total                  string
					InitialBaseQuantity    string
					InitialQuoteQuantity   string
					RemainingBaseQuantity  string
					RemainingQuoteQuantity string
					MeanExecutionPrice     string
					EngineMessage          string
				}
			}
		}
	}{}
	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected")

	return getProOrders{
		Edges: res.GetProOrders.Edges,
	}, nil
}

func (config configCredentials) CancelOrder(id string) (cancelOrder, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(`
			mutation($id: ID!) {
				cancelOrder(proOrder: $id){
				id
				pair
				price
				side
				status
				timeInForce
				orderType
				fees
				filled
				total
				initialBaseQuantity
				initialQuoteQuantity
				remainingBaseQuantity
				remainingQuoteQuantity
				meanExecutionPrice
				engineMessage
				}
			}
	`)

	req.Var("id", id)
	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()

	res := struct {
		CancelOrder struct {
			Id                     string
			Pair                   string
			Price                  string
			Side                   string
			Status                 string
			TimeInForce            string
			OrderType              string
			Fees                   string
			Filled                 string
			Total                  string
			InitialBaseQuantity    string
			InitialQuoteQuantity   string
			RemainingBaseQuantity  string
			RemainingQuoteQuantity string
			MeanExecutionPrice     string
			EngineMessage          string
		}
	}{}

	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}

	return cancelOrder{
		Id:                     res.CancelOrder.Id,
		pair:                   res.CancelOrder.Pair,
		price:                  res.CancelOrder.Price,
		side:                   res.CancelOrder.Side,
		status:                 res.CancelOrder.Status,
		timeInForce:            res.CancelOrder.TimeInForce,
		orderType:              res.CancelOrder.OrderType,
		fees:                   res.CancelOrder.Fees,
		filled:                 res.CancelOrder.Filled,
		total:                  res.CancelOrder.Total,
		initialBaseQuantity:    res.CancelOrder.InitialBaseQuantity,
		initialQuoteQuantity:   res.CancelOrder.InitialQuoteQuantity,
		remainingBaseQuantity:  res.CancelOrder.RemainingBaseQuantity,
		remainingQuoteQuantity: res.CancelOrder.RemainingQuoteQuantity,
		meanExecutionPrice:     res.CancelOrder.MeanExecutionPrice,
		engineMessage:          res.CancelOrder.EngineMessage,
	}, nil

}

func (config configCredentials) GetProOrderFees(orderType string, pair string, side string, amount float64) (getProOrderFees, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(`
		query($orderType_: OrderMatchingEngineOrder!, $pair_: Pair!, $side_: OrderSide!, $amount_: BigDecimal!) {
			getProOrderFees(orderType: $orderType_, pair: $pair_, side: $side_, amount: $amount_){
			fee
			baseCurrencyTotal
			quoteCurrencyTotal
			price
			}
		}
	`)

	req.Var("orderType_", orderType)
	req.Var("pair_", pair)
	req.Var("side_", side)
	req.Var("amount_", amount)
	ctx := context.Background()

	res := struct {
		GetProOrderFees struct {
			Fees               string
			BaseCurrencyTotal  string
			QuoteCurrencyTotal string
			Price              string
		}
	}{}

	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}

	return getProOrderFees{
		fee:                res.GetProOrderFees.Fees,
		baseCurrencyTotal:  res.GetProOrderFees.BaseCurrencyTotal,
		quoteCurrencyTotal: res.GetProOrderFees.QuoteCurrencyTotal,
		price:              res.GetProOrderFees.Price,
	}, nil
}

func (config configCredentials) PostProMarketOrder(pair string, quantity float64, side string) (postProMarketOrder, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(`
		mutation($pair_: Pair!, $quantity_: BigDecimal!, $side_: OrderSide!) {
			postProMarketOrder(pair: $pair_, quantity: $quantity_, side: $side_){
			id
			pair
			price
			side
			status
			timeInForce
			orderType
			fees
			filled
			total
			initialBaseQuantity
			initialQuoteQuantity
			remainingBaseQuantity
			remainingQuoteQuantity
			meanExecutionPrice
			engineMessage
			}
		}
	`)

	req.Var("pair_", pair)
	req.Var("quantity_", quantity)
	req.Var("side_", side)
	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()

	res := struct {
		PostProMarketOrder struct {
			Id                     string
			Pair                   string
			Price                  string
			Side                   string
			Status                 string
			TimeInForce            string
			OrderType              string
			Fees                   string
			Filled                 string
			Total                  string
			InitialBaseQuantity    string
			InitialQuoteQuantity   string
			RemainingBaseQuantity  string
			RemainingQuoteQuantity string
			MeanExecutionPrice     string
			EngineMessage          string
		}
	}{}

	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}

	return postProMarketOrder{
		Id:                     res.PostProMarketOrder.Id,
		pair:                   res.PostProMarketOrder.Pair,
		price:                  res.PostProMarketOrder.Price,
		side:                   res.PostProMarketOrder.Side,
		status:                 res.PostProMarketOrder.Status,
		timeInForce:            res.PostProMarketOrder.TimeInForce,
		orderType:              res.PostProMarketOrder.OrderType,
		fees:                   res.PostProMarketOrder.Fees,
		filled:                 res.PostProMarketOrder.Filled,
		total:                  res.PostProMarketOrder.Total,
		initialBaseQuantity:    res.PostProMarketOrder.InitialBaseQuantity,
		initialQuoteQuantity:   res.PostProMarketOrder.InitialQuoteQuantity,
		remainingBaseQuantity:  res.PostProMarketOrder.RemainingBaseQuantity,
		remainingQuoteQuantity: res.PostProMarketOrder.RemainingQuoteQuantity,
		meanExecutionPrice:     res.PostProMarketOrder.MeanExecutionPrice,
		engineMessage:          res.PostProMarketOrder.EngineMessage,
	}, nil
}

func (config configCredentials) PostProLimitOrder(pair string, quantity float64, price float64, side string, timeInForce string) (LimitOrder, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(`
		mutation($pair_: Pair!, $quantity_: BigDecimal!, $price_: BigDecimal! $side_: OrderSide!, $timeInForce_: TimeInForce!) {
			postProLimitOrder(pair: $pair_, quantity: $quantity_, price: $price_ side: $side_, timeInForce: $timeInForce_){
			id
			pair
			price
			side
			status
			timeInForce
			orderType
			fees
			filled
			total
			initialBaseQuantity
			initialQuoteQuantity
			remainingBaseQuantity
			remainingQuoteQuantity
			meanExecutionPrice
			engineMessage
			}
		}
	`)

	req.Var("pair_", pair)
	req.Var("quantity_", quantity)
	req.Var("price_", price)
	req.Var("side_", side)
	req.Var("timeInForce_", timeInForce)
	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()

	res := struct {
		PostProLimitOrder struct {
			Id                     string
			Pair                   string
			Price                  string
			Side                   string
			Status                 string
			TimeInForce            string
			OrderType              string
			Fees                   string
			Filled                 string
			Total                  string
			InitialBaseQuantity    string
			InitialQuoteQuantity   string
			RemainingBaseQuantity  string
			RemainingQuoteQuantity string
			MeanExecutionPrice     string
			EngineMessage          string
		}
	}{}

	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}

	return LimitOrder{
		Id:                     res.PostProLimitOrder.Id,
		pair:                   res.PostProLimitOrder.Pair,
		price:                  res.PostProLimitOrder.Price,
		side:                   res.PostProLimitOrder.Side,
		status:                 res.PostProLimitOrder.Status,
		timeInForce:            res.PostProLimitOrder.TimeInForce,
		orderType:              res.PostProLimitOrder.OrderType,
		fees:                   res.PostProLimitOrder.Fees,
		filled:                 res.PostProLimitOrder.Filled,
		total:                  res.PostProLimitOrder.Total,
		initialBaseQuantity:    res.PostProLimitOrder.InitialBaseQuantity,
		initialQuoteQuantity:   res.PostProLimitOrder.InitialQuoteQuantity,
		remainingBaseQuantity:  res.PostProLimitOrder.RemainingBaseQuantity,
		remainingQuoteQuantity: res.PostProLimitOrder.RemainingQuoteQuantity,
		meanExecutionPrice:     res.PostProLimitOrder.MeanExecutionPrice,
		engineMessage:          res.PostProLimitOrder.EngineMessage,
	}, nil
}
