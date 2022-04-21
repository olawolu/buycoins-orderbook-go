package orderbooks

var getProOrdersQuery = `query ($pair_: Pair!, $status_: ProOrderStatus!, $side_: OrderSide!) {
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
}`

var cancelOrderQuery = `mutation($id: ID!) {
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
}`

var getProOrderFeesQuery = `query($orderType_: OrderMatchingEngineOrder!, $pair_: Pair!, $side_: OrderSide!, $amount_: BigDecimal!) {
	getProOrderFees(orderType: $orderType_, pair: $pair_, side: $side_, amount: $amount_){
		fee
		baseCurrencyTotal
		quoteCurrencyTotal
		price
	}
}`

var postProMarketOrderQuery = `mutation($pair_: Pair!, $quantity_: BigDecimal!, $side_: OrderSide!) {
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
}`

var postProLimitOrderQuery = `mutation($pair_: Pair!, $quantity_: BigDecimal!, $price_: BigDecimal! $side_: OrderSide!, $timeInForce_: TimeInForce!) {
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
}`

var getDepositLinkQuery = `mutation($amount: BigDecimal!) {
	createSendCashPayDeposit(amount: $amount){
		amount
		createdAt
		fee
		id
		link
		reference
		status
		totalAmount
		type
	}
}`

var getBalancesQuery = `query($crypto: Cryptocurrency) {
	getBalances(cryptocurrency: $crypto) {
		id
		cryptocurrency
		confirmedBalance
	}
}`