package orderbooks

type cancelOrder struct {
	Id                     string
	pair                   string
	price                  string
	side                   string
	status                 string
	timeInForce            string
	orderType              string
	fees                   string
	filled                 string
	total                  string
	initialBaseQuantity    string
	initialQuoteQuantity   string
	remainingBaseQuantity  string
	remainingQuoteQuantity string
	meanExecutionPrice     string
	engineMessage          string
}

type getProOrderFees struct {
	fee                string
	baseCurrencyTotal  string
	quoteCurrencyTotal string
	price              string
}

type postProMarketOrder struct {
	Id                     string
	pair                   string
	price                  string
	side                   string
	status                 string
	timeInForce            string
	orderType              string
	fees                   string
	filled                 string
	total                  string
	initialBaseQuantity    string
	initialQuoteQuantity   string
	remainingBaseQuantity  string
	remainingQuoteQuantity string
	meanExecutionPrice     string
	engineMessage          string
}

type LimitOrder struct {
	Id                     string
	pair                   string
	price                  string
	side                   string
	status                 string
	timeInForce            string
	orderType              string
	fees                   string
	filled                 string
	total                  string
	initialBaseQuantity    string
	initialQuoteQuantity   string
	remainingBaseQuantity  string
	remainingQuoteQuantity string
	meanExecutionPrice     string
	engineMessage          string
}

type getProOrders struct {
	Edges []struct {
		Node struct {
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

type getDepositLink struct {
	Amount      string
	CreatedAt   int64
	Fee         string
	Id          string
	Link        string
	Reference   string
	Status      string
	TotalAmount string
	Type        string
}
