package model

import (
	"github.com/NattapornTee22816/binance-connector-golang/lib"
	"github.com/buger/jsonparser"
	"time"
)

// NewOrder

type OrderParam struct {
	Symbol           string            `json:"symbol" param:"symbol" validate:"required"`
	Side             string            `json:"side" param:"side" validate:"required"`
	OrderType        OrderType         `json:"type" param:"type" validate:"required"`
	TimeInForce      TimeInForce       `json:"timeInForce" param:"timeInForce"`
	Quantity         float64           `json:"quantity" param:"quantity"`
	QuoteOrderQty    float64           `json:"quoteOrderQty" param:"quoteOrderQty"`
	Price            float64           `json:"price"`
	NewClientOrderId string            `json:"newClientOrderId" param:"newClientOrderId"`
	StopPrice        float64           `json:"stopPrice" param:"stopPrice"`
	IcebergQty       float64           `json:"icebergQty" param:"icebergQty"`
	NewOrderRespType OrderResponseType `json:"newOrderRespType" param:"newOrderRespType"`
	RecvWindow       int64             `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

type Order struct {
	Symbol              string       `json:"symbol"`
	OrderId             int64        `json:"orderId"`
	OrderListId         int64        `json:"orderListId"`
	ClientOrderId       string       `json:"clientOrderId"`
	TransactTime        time.Time    `json:"transactTime"`
	Price               string       `json:"price"`
	OrigQty             string       `json:"origQty"`
	ExecutedQty         string       `json:"executedQty"`
	CummulativeQuoteQty string       `json:"cummulativeQuoteQty"`
	Status              string       `json:"status"`
	TimeInForce         string       `json:"timeInForce"`
	Type                string       `json:"type"`
	Side                string       `json:"side"`
	Fills               []*OrderFill `json:"fills,omitempty"`
}

type OrderFill struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	TradeId         int64  `json:"tradeId"`
}

func (r *Parser) ParseOrder(b []byte) (*Order, error) {
	result := new(Order)

	if v, err := jsonparser.GetString(b, "symbol"); err == nil {
		result.Symbol = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "orderId"); err == nil {
		result.OrderId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "orderListId"); err == nil {
		result.OrderListId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "clientOrderId"); err == nil {
		result.ClientOrderId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "transactTime"); err == nil {
		result.TransactTime = lib.ConvertIntToTime(v, 0)
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "price"); err == nil {
		result.Price = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "origQty"); err == nil {
		result.OrigQty = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "executedQty"); err == nil {
		result.ExecutedQty = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "cummulativeQuoteQty"); err == nil {
		result.CummulativeQuoteQty = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "status"); err == nil {
		result.Status = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "timeInForce"); err == nil {
		result.TimeInForce = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "type"); err == nil {
		result.Type = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "side"); err == nil {
		result.Side = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}

	result.Fills = make([]*OrderFill, 0)
	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		fill := new(OrderFill)

		if v, err := jsonparser.GetString(value, "price"); err != nil {
			fill.Price = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "qty"); err != nil {
			fill.Qty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "commission"); err != nil {
			fill.Commission = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "commissionAsset"); err != nil {
			fill.CommissionAsset = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "tradeId"); err != nil {
			fill.TradeId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		result.Fills = append(result.Fills, fill)
	}, "fills")
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return result, nil
}

// CancelOrder

type CancelOrderParam struct {
	Symbol            string `json:"symbol" param:"symbol" validate:"required"`
	OrderId           int64  `json:"orderId" param:"orderId"`
	OrigClientOrderId string `json:"origClientOrderId" param:"origClientOrderId"`
	NewClientOrderId  string `json:"newClientOrderId" param:"newClientOrderId"`
	RecvWindow        int64  `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

type CancelOrder struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderId   string `json:"origClientOrderId"`
	OrderId             int64  `json:"orderId"`
	OrderListId         int64  `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}

func (r *Parser) ParseCancelOrder(b []byte) (*CancelOrder, error) {
	result := new(CancelOrder)

	if v, err := jsonparser.GetString(b, "symbol"); err == nil {
		result.Symbol = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "origClientOrderId"); err == nil {
		result.OrigClientOrderId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "orderId"); err == nil {
		result.OrderId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "orderId"); err == nil {
		result.OrderId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "orderListId"); err == nil {
		result.OrderListId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "clientOrderId"); err == nil {
		result.ClientOrderId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "price"); err == nil {
		result.Price = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "origQty"); err == nil {
		result.OrigQty = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "executedQty"); err == nil {
		result.ExecutedQty = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "cummulativeQuoteQty"); err == nil {
		result.CummulativeQuoteQty = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "status"); err == nil {
		result.Status = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "timeInForce"); err == nil {
		result.TimeInForce = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "type"); err == nil {
		result.Type = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "side"); err == nil {
		result.Side = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// CancelOpenOrder

type CancelOpenOrderParam struct {
	Symbol     string `json:"symbol" param:"symbol" validate:"required"`
	RecvWindow int64  `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

type CancelOpenOrder struct {
	Symbol              string               `json:"symbol,omitempty"`
	OrigClientOrderId   string               `json:"origClientOrderId,omitempty"`
	OrderId             int64                `json:"orderId,omitempty"`
	OrderListId         int64                `json:"orderListId,omitempty"`
	ClientOrderId       string               `json:"clientOrderId,omitempty"`
	Price               string               `json:"price,omitempty"`
	OrigQty             string               `json:"origQty,omitempty"`
	ExecutedQty         string               `json:"executedQty,omitempty"`
	CummulativeQuoteQty string               `json:"cummulativeQuoteQty,omitempty"`
	Status              string               `json:"status,omitempty"`
	TimeInForce         string               `json:"timeInForce,omitempty"`
	Type                string               `json:"type,omitempty"`
	Side                string               `json:"side,omitempty"`
	ContingencyType     string               `json:"contingencyType,omitempty"`
	ListStatusType      string               `json:"listStatusType,omitempty"`
	ListOrderStatus     string               `json:"listOrderStatus,omitempty"`
	ListClientOrderId   string               `json:"listClientOrderId,omitempty"`
	TransactionTime     time.Time            `json:"transactionTime,omitempty"`
	Orders              []*CancelOrderId     `json:"orders,omitempty"`
	OrderReports        []*CancelOrderReport `json:"orderReports,omitempty"`
}

type CancelOrderId struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type CancelOrderReport struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderId   string `json:"origClientOrderId"`
	OrderId             int64  `json:"orderId"`
	OrderListId         int64  `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
}

func (r *Parser) ParseCancelOpenOrder(b []byte) ([]*CancelOpenOrder, error) {
	results := make([]*CancelOpenOrder, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(CancelOpenOrder)

		if v, err := jsonparser.GetString(value, "symbol"); err == nil {
			item.Symbol = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "origClientOrderId"); err == nil {
			item.OrigClientOrderId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "orderId"); err == nil {
			item.OrderId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "orderListId"); err == nil {
			item.OrderListId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "clientOrderId"); err == nil {
			item.ClientOrderId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "price"); err == nil {
			item.Price = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "origQty"); err == nil {
			item.OrigQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "executedQty"); err == nil {
			item.ExecutedQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "cummulativeQuoteQty"); err == nil {
			item.CummulativeQuoteQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "status"); err == nil {
			item.Status = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "timeInForce"); err == nil {
			item.TimeInForce = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "type"); err == nil {
			item.Type = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "side"); err == nil {
			item.Side = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "contingencyType"); err == nil {
			item.ContingencyType = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "listStatusType"); err == nil {
			item.ListStatusType = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "listOrderStatus"); err == nil {
			item.ListOrderStatus = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "listClientOrderId"); err == nil {
			item.ListClientOrderId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "transactionTime"); err == nil {
			item.TransactionTime = lib.ConvertIntToTime(v, 0)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		item.Orders = make([]*CancelOrderId, 0)
		_, err := jsonparser.ArrayEach(value, func(order []byte, dataType jsonparser.ValueType, __offset int, __err error) {
			orderId := new(CancelOrderId)

			if v, err := jsonparser.GetString(value, "symbol"); err == nil {
				orderId.Symbol = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetInt(value, "orderId"); err == nil {
				orderId.OrderId = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "clientOrderId"); err == nil {
				orderId.ClientOrderId = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}

			item.Orders = append(item.Orders, orderId)
		}, "orders")
		if _err = r.errorParser(err); _err != nil {
			return
		}
		item.OrderReports = make([]*CancelOrderReport, 0)
		_, err = jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, __err error) {
			report := new(CancelOrderReport)

			if v, err := jsonparser.GetString(value, "symbol"); err == nil {
				report.ClientOrderId = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "origClientOrderId"); err == nil {
				report.OrigClientOrderId = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetInt(value, "orderId"); err == nil {
				report.OrderId = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetInt(value, "orderListId"); err == nil {
				report.OrderListId = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "clientOrderId"); err == nil {
				report.ClientOrderId = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "price"); err == nil {
				report.Price = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "origQty"); err == nil {
				report.OrigQty = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "executedQty"); err == nil {
				report.ExecutedQty = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "cummulativeQuoteQty"); err == nil {
				report.CummulativeQuoteQty = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "status"); err == nil {
				report.Status = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "timeInForce"); err == nil {
				report.TimeInForce = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "type"); err == nil {
				report.Type = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "side"); err == nil {
				report.Side = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "stopPrice"); err == nil {
				report.StopPrice = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}
			if v, err := jsonparser.GetString(value, "icebergQty"); err == nil {
				report.IcebergQty = v
			} else {
				if __err = r.errorParser(err); __err != nil {
					return
				}
			}

			item.OrderReports = append(item.OrderReports, report)
		}, "orderReports")
		if _err = r.errorParser(err); _err != nil {
			return
		}

		results = append(results, item)
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}

// GetOrder

type GetOrderParam struct {
	Symbol            string `json:"symbol" param:"symbol" validate:"required"`
	OrderId           int64  `json:"orderId" param:"orderId"`
	OrigClientOrderId string `json:"origClientOrderId" param:"origClientOrderId"`
	RecvWindow        int64  `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

type GetOrder struct {
	Symbol              string    `json:"symbol"`
	OrderId             int64     `json:"orderId"`
	OrderListId         int64     `json:"orderListId"`
	ClientOrderId       string    `json:"clientOrderId"`
	Price               string    `json:"price"`
	OrigQty             string    `json:"origQty"`
	ExecutedQty         string    `json:"executedQty"`
	CummulativeQuoteQty string    `json:"cummulativeQuoteQty"`
	Status              string    `json:"status"`
	TimeInForce         string    `json:"timeInForce"`
	Type                string    `json:"type"`
	Side                string    `json:"side"`
	StopPrice           string    `json:"stopPrice"`
	IcebergQty          string    `json:"icebergQty"`
	Time                time.Time `json:"time"`
	UpdateTime          time.Time `json:"updateTime"`
	IsWorking           bool      `json:"isWorking"`
	OrigQuoteOrderQty   string    `json:"origQuoteOrderQty"`
}

func (r *Parser) ParseGetOrder(b []byte) (*GetOrder, error) {
	result := new(GetOrder)

	if v, err := jsonparser.GetString(b, "symbol"); err == nil {
		result.Symbol = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "orderId"); err == nil {
		result.OrderId = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "orderListId"); err == nil {
		result.OrderListId = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "clientOrderId"); err == nil {
		result.ClientOrderId = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "price"); err == nil {
		result.Price = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "origQty"); err == nil {
		result.OrigQty = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "executedQty"); err == nil {
		result.ExecutedQty = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "cummulativeQuoteQty"); err == nil {
		result.CummulativeQuoteQty = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "status"); err == nil {
		result.Status = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "timeInForce"); err == nil {
		result.TimeInForce = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "type"); err == nil {
		result.Type = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "side"); err == nil {
		result.Side = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "stopPrice"); err == nil {
		result.StopPrice = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "icebergQty"); err == nil {
		result.IcebergQty = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "time"); err == nil {
		result.Time = lib.ConvertIntToTime(v, 0)
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "updateTime"); err == nil {
		result.UpdateTime = lib.ConvertIntToTime(v, 0)
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetBoolean(b, "isWorking"); err == nil {
		result.IsWorking = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "origQuoteOrderQty"); err == nil {
		result.OrigQuoteOrderQty = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// GetOpenOrders

type GetOpenOrdersParam struct {
	Symbol     string `json:"symbol" param:"symbol"`
	RecvWindow int64  `json:"recvWindow" param:"recvWindow"`
}

func (r *Parser) ParseGetOpenOrder(b []byte) ([]*GetOrder, error) {
	results := make([]*GetOrder, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		if item, err := r.ParseGetOrder(value); err == nil {
			results = append(results, item)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}

// GetOrders

type GetOrdersParam struct {
	Symbol     string    `json:"symbol" param:"symbol" validate:"required"`
	OrderId    int64     `json:"orderId" param:"orderId"`
	StartTime  time.Time `json:"startTime" param:"startTime"`
	EndTime    time.Time `json:"endTime" param:"endTime"`
	Limit      int64     `json:"limit" param:"limit" validate:"max=1000"`
	RecvWindow int64     `json:"recvWindow" param:"recvWindow" validate:"max=1000"`
}

func (r *Parser) ParseGetOrders(b []byte) ([]*GetOrder, error) {
	return r.ParseGetOpenOrder(b)
}

// NewOcoOrder

type NewOcoOrderParam struct {
	Symbol               string            `json:"symbol" param:"symbol" validate:"required"`
	ListClientOrderId    string            `json:"listClientOrderId" param:"listClientOrderId"`
	Side                 OrderSide         `json:"side" param:"side" validate:"required"`
	Quantity             float64           `json:"quantity" param:"quantity" validate:"required"`
	LimitClientOrderId   string            `json:"limitClientOrderId" param:"limitClientOrderId"`
	Price                float64           `json:"price" param:"price" validate:"required"`
	LimitIcebergQty      float64           `json:"limitIcebergQty" param:"limitIcebergQty"`
	StopClientOrderId    string            `json:"stopClientOrderId" param:"stopClientOrderId"`
	StopPrice            float64           `json:"stopPrice" param:"stopPrice" validate:"required"`
	StopLimitPrice       float64           `json:"stopLimitPrice" param:"stopLimitPrice"`
	StopIcebergQty       float64           `json:"stopIcebergQty" param:"stopIcebergQty"`
	StopLimitTimeInForce TimeInForce       `json:"stopLimitTimeInForce" param:"stopLimitTimeInForce"`
	NewOrderRespType     OrderResponseType `json:"newOrderRespType" param:"newOrderRespType"`
	RecvWindow           int64             `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

type OcoOrder struct {
	OrderListId       int64                `json:"orderListId"`
	ContingencyType   string               `json:"contingencyType"`
	ListStatusType    string               `json:"listStatusType"`
	ListOrderStatus   string               `json:"listOrderStatus"`
	ListClientOrderId string               `json:"listClientOrderId"`
	TransactionTime   time.Time            `json:"transactionTime"`
	Symbol            string               `json:"symbol"`
	Orders            []*OcoOrderBasicInfo `json:"orders"`
	OrderReports      []*OcoOrderReport    `json:"orderReports"`
}

type OcoOrderBasicInfo struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type OcoOrderReport struct {
	Symbol              string    `json:"symbol"`
	OrderId             int64     `json:"orderId"`
	OrderListId         int64     `json:"orderListId"`
	ClientOrderId       string    `json:"clientOrderId"`
	TransactionTime     time.Time `json:"transactionTime"`
	Price               string    `json:"price"`
	OrigQty             string    `json:"origQty"`
	ExecutedQty         string    `json:"executedQty"`
	CummulativeQuoteQty string    `json:"cummulativeQuoteQty"`
	Status              string    `json:"status"`
	TimeInForce         string    `json:"timeInForce"`
	Type                string    `json:"type"`
	Side                string    `json:"side"`
	StopPrice           string    `json:"stopPrice"`
}

func (r *Parser) ParseNewOcoOrder(b []byte) (*OcoOrder, error) {
	result := new(OcoOrder)

	if v, err := jsonparser.GetInt(b, "orderListId"); err == nil {
		result.OrderListId = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "contingencyType"); err == nil {
		result.ContingencyType = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "listStatusType"); err == nil {
		result.ListStatusType = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "listOrderStatus"); err == nil {
		result.ListOrderStatus = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "listClientOrderId"); err == nil {
		result.ListClientOrderId = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "transactionTime"); err == nil {
		result.TransactionTime = lib.ConvertIntToTime(v, 0)
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "symbol"); err == nil {
		result.Symbol = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	result.Orders = make([]*OcoOrderBasicInfo, 0)
	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		if item, err := r.parseOcoOrderBasicInfo(value); err == nil {
			result.Orders = append(result.Orders, item)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
	}, "orders")
	if err = r.errorParser(err); err != nil {
		return nil, err
	}
	result.OrderReports = make([]*OcoOrderReport, 0)
	_, err = jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(OcoOrderReport)

		if v, err := jsonparser.GetString(b, "symbol"); err == nil {
			item.Symbol = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(b, "orderId"); err == nil {
			item.OrderId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(b, "orderListId"); err == nil {
			item.OrderListId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "clientOrderId"); err == nil {
			item.ClientOrderId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(b, "transactionTime"); err == nil {
			item.TransactionTime = lib.ConvertIntToTime(v, 0)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "price"); err == nil {
			item.Price = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "origQty"); err == nil {
			item.OrigQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "executedQty"); err == nil {
			item.ExecutedQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "cummulativeQuoteQty"); err == nil {
			item.CummulativeQuoteQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "status"); err == nil {
			item.Status = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "timeInForce"); err == nil {
			item.TimeInForce = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "type"); err == nil {
			item.Type = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "side"); err == nil {
			item.Side = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(b, "stopPrice"); err == nil {
			item.StopPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		result.OrderReports = append(result.OrderReports, item)
	}, "orderReports")
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Parser) parseOcoOrderBasicInfo(b []byte) (*OcoOrderBasicInfo, error) {
	result := new(OcoOrderBasicInfo)

	if v, err := jsonparser.GetString(b, "symbol"); err == nil {
		result.Symbol = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "orderId"); err == nil {
		result.OrderId = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "clientOrderId"); err == nil {
		result.ClientOrderId = v
	} else {
		if err := r.errorParser(err); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// CancelOcoOrder

type CancelOcoOrderParam struct {
	Symbol            string `json:"symbol" param:"symbol" validate:"required"`
	OrderListId       string `json:"orderListId" param:"orderListId"`
	ListClientOrderId string `json:"listClientOrderId" param:"listClientOrderId"`
	NewClientOrderId  string `json:"newClientOrderId" param:"newClientOrderId"`
	RecvWindow        int64  `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

type CancelOcoOrder struct {
	OrderListId       int64                   `json:"orderListId"`
	ContingencyType   string                  `json:"contingencyType"`
	ListStatusType    string                  `json:"listStatusType"`
	ListOrderStatus   string                  `json:"listOrderStatus"`
	ListClientOrderId string                  `json:"listClientOrderId"`
	TransactionTime   time.Time               `json:"transactionTime"`
	Symbol            string                  `json:"symbol"`
	Orders            []*OcoOrderBasicInfo    `json:"orders"`
	OrderReports      []*CancelOcoOrderReport `json:"orderReports"`
}

type CancelOcoOrderReport struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderId   string `json:"origClientOrderId"`
	OrderId             int64  `json:"orderId"`
	OrderListId         int64  `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
}

func (r *Parser) ParseCancelOcoOrder(b []byte) (*CancelOcoOrder, error) {
	result := new(CancelOcoOrder)

	if v, err := jsonparser.GetInt(b, "orderListId"); err == nil {
		result.OrderListId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "contingencyType"); err == nil {
		result.ContingencyType = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "listStatusType"); err == nil {
		result.ListStatusType = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "listOrderStatus"); err == nil {
		result.ListOrderStatus = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "listClientOrderId"); err == nil {
		result.ListClientOrderId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "transactionTime"); err == nil {
		result.TransactionTime = lib.ConvertIntToTime(v, 0)
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "symbol"); err == nil {
		result.Symbol = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	result.Orders = make([]*OcoOrderBasicInfo, 0)
	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		if item, err := r.parseOcoOrderBasicInfo(value); err != nil {
			result.Orders = append(result.Orders, item)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
	}, "orders")
	if err = r.errorParser(err); err != nil {
		return nil, err
	}
	result.OrderReports = make([]*CancelOcoOrderReport, 0)
	_, err = jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(CancelOcoOrderReport)

		if v, err := jsonparser.GetString(value, "symbol"); err == nil {
			item.Symbol = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "origClientOrderId"); err == nil {
			item.OrigClientOrderId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "orderId"); err == nil {
			item.OrderId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "orderListId"); err == nil {
			item.OrderListId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "clientOrderId"); err == nil {
			item.ClientOrderId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "price"); err == nil {
			item.Price = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "origQty"); err == nil {
			item.OrigQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "executedQty"); err == nil {
			item.ExecutedQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "cummulativeQuoteQty"); err == nil {
			item.CummulativeQuoteQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "status"); err == nil {
			item.Status = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "timeInForce"); err == nil {
			item.TimeInForce = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "type"); err == nil {
			item.Type = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "side"); err == nil {
			item.Side = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "stopPrice"); err == nil {
			item.StopPrice = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		result.OrderReports = append(result.OrderReports, item)
	}, "orderReports")
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return result, nil
}

// GetOcoOrder

type GetOcoOrderParam struct {
	OrderListId       int64  `json:"orderListId" param:"orderListId"`
	OrigClientOrderId string `json:"origClientOrderId" param:"origClientOrderId"`
	RecvWindow        int64  `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

type GetOcoOrder struct {
	OrderListId       int64                `json:"orderListId"`
	ContingencyType   string               `json:"contingencyType"`
	ListStatusType    string               `json:"listStatusType"`
	ListOrderStatus   string               `json:"listOrderStatus"`
	ListClientOrderId string               `json:"listClientOrderId"`
	TransactionTime   time.Time            `json:"transactionTime"`
	Symbol            string               `json:"symbol"`
	Orders            []*OcoOrderBasicInfo `json:"orders"`
}

func (r *Parser) ParseGetOcoOrder(b []byte) (*GetOcoOrder, error) {
	result := new(GetOcoOrder)

	if v, err := jsonparser.GetInt(b, "orderListId"); err == nil {
		result.OrderListId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "contingencyType"); err == nil {
		result.ContingencyType = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "listStatusType"); err == nil {
		result.ListStatusType = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "listOrderStatus"); err == nil {
		result.ListOrderStatus = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "listClientOrderId"); err == nil {
		result.ListClientOrderId = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetInt(b, "transactionTime"); err == nil {
		result.TransactionTime = lib.ConvertIntToTime(v, 0)
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	if v, err := jsonparser.GetString(b, "symbol"); err == nil {
		result.Symbol = v
	} else {
		if err = r.errorParser(err); err != nil {
			return nil, err
		}
	}
	result.Orders = make([]*OcoOrderBasicInfo, 0)
	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		if item, err := r.parseOcoOrderBasicInfo(value); err == nil {
			result.Orders = append(result.Orders, item)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
	}, "orders")
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return result, nil
}

// GetOcoOrders

type GetOcoOrdersParam struct {
	FromId     int64     `json:"fromId" param:"fromId"`
	StartTime  time.Time `json:"startTime" param:"startTime"`
	EndTime    time.Time `json:"endTime" param:"endTime"`
	Limit      int64     `json:"limit" param:"limit" validate:"max=1000"`
	RecvWindow int64     `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

func (r *Parser) ParseGetOcoOrders(b []byte) ([]*GetOcoOrder, error) {
	results := make([]*GetOcoOrder, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		if item, err := r.ParseGetOcoOrder(value); err == nil {
			results = append(results, item)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}

// GetOcoOpenOrders

type GetOcoOpenOrdersParam struct {
	RecvWindow int64 `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

func (r *Parser) ParseGetOcoOpenOrders(b []byte) ([]*GetOcoOrder, error) {
	return r.ParseGetOcoOrders(b)
}

// Account

type AccountParam struct {
	RecvWindow int64 `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

type Account struct {
	MakerCommission  float64           `json:"makerCommission"`
	TakerCommission  float64           `json:"takerCommission"`
	BuyerCommission  float64           `json:"buyerCommission"`
	SellerCommission float64           `json:"sellerCommission"`
	CanTrade         bool              `json:"canTrade"`
	CanWithdraw      bool              `json:"canWithdraw"`
	CanDeposit       bool              `json:"canDeposit"`
	UpdateTime       time.Time         `json:"updateTime"`
	AccountType      string            `json:"accountType"`
	Balances         []*AccountBalance `json:"balances"`
	Permissions      []string          `json:"permissions"`
}

type AccountBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

func (r *Parser) ParseAccount(b []byte) (*Account, error) {
	result := new(Account)

	if v, err := jsonparser.GetFloat(b, "makerCommission"); err == nil {
		result.MakerCommission = v
	} else {
		if wrapper := r.errorParser(err); wrapper != nil {
			return nil, wrapper
		}
	}
	if v, err := jsonparser.GetFloat(b, "takerCommission"); err == nil {
		result.TakerCommission = v
	} else {
		if wrapper := r.errorParser(err); wrapper != nil {
			return nil, wrapper
		}
	}
	if v, err := jsonparser.GetFloat(b, "buyerCommission"); err == nil {
		result.BuyerCommission = v
	} else {
		if wrapper := r.errorParser(err); wrapper != nil {
			return nil, wrapper
		}
	}
	if v, err := jsonparser.GetFloat(b, "sellerCommission"); err == nil {
		result.SellerCommission = v
	} else {
		if wrapper := r.errorParser(err); wrapper != nil {
			return nil, wrapper
		}
	}
	if v, err := jsonparser.GetBoolean(b, "canTrade"); err == nil {
		result.CanTrade = v
	} else {
		if wrapper := r.errorParser(err); wrapper != nil {
			return nil, wrapper
		}
	}
	if v, err := jsonparser.GetBoolean(b, "canWithdraw"); err == nil {
		result.CanWithdraw = v
	} else {
		if wrapper := r.errorParser(err); wrapper != nil {
			return nil, wrapper
		}
	}
	if v, err := jsonparser.GetBoolean(b, "canDeposit"); err == nil {
		result.CanDeposit = v
	} else {
		if wrapper := r.errorParser(err); wrapper != nil {
			return nil, wrapper
		}
	}
	if v, err := jsonparser.GetInt(b, "updateTime"); err == nil {
		result.UpdateTime = lib.ConvertIntToTime(v, 0)
	} else {
		if wrapper := r.errorParser(err); wrapper != nil {
			return nil, wrapper
		}
	}
	if v, err := jsonparser.GetString(b, "accountType"); err == nil {
		result.AccountType = v
	} else {
		if wrapper := r.errorParser(err); wrapper != nil {
			return nil, wrapper
		}
	}
	result.Balances = make([]*AccountBalance, 0)
	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(AccountBalance)

		if v, err := jsonparser.GetString(value, "asset"); err == nil {
			item.Asset = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "free"); err == nil {
			item.Free = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "locked"); err == nil {
			item.Locked = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		result.Balances = append(result.Balances, item)
	}, "balances")
	if err != nil {
		return nil, err
	}
	result.Permissions = make([]string, 0)
	_, err = jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if dataType == jsonparser.String {
			result.Permissions = append(result.Permissions, string(value))
		}
	}, "permissions")
	if err != nil {
		return nil, err
	}

	return result, err
}

// MyTrades

type MyTradesParam struct {
	Symbol     string    `json:"symbol" param:"symbol" validate:"required"`
	OrderId    int64     `json:"orderId" param:"orderId"`
	StartTime  time.Time `json:"startTime" param:"startTime"`
	EndTime    time.Time `json:"endTime" param:"endTime"`
	FromId     int64     `json:"fromId" param:"fromId"`
	Limit      int64     `json:"limit" param:"limit" validate:"max=1000"`
	RecvWindow int64     `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

type MyTrade struct {
	Symbol          string    `json:"symbol"`
	Id              int64     `json:"id"`
	OrderId         int64     `json:"orderId"`
	OrderListId     int64     `json:"orderListId"`
	Price           string    `json:"price"`
	Qty             string    `json:"qty"`
	QuoteQty        string    `json:"quoteQty"`
	Commission      string    `json:"commission"`
	CommissionAsset string    `json:"commissionAsset"`
	Time            time.Time `json:"time"`
	IsBuyer         bool      `json:"isBuyer"`
	IsMaker         bool      `json:"isMaker"`
	IsBestMatch     bool      `json:"isBestMatch"`
}

func (r *Parser) ParseMyTrades(b []byte) ([]*MyTrade, error) {
	results := make([]*MyTrade, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(MyTrade)

		if v, err := jsonparser.GetString(value, "symbol"); err == nil {
			item.Symbol = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "id"); err == nil {
			item.Id = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "orderId"); err == nil {
			item.OrderId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "orderListId"); err == nil {
			item.OrderListId = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "price"); err == nil {
			item.Price = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "qty"); err == nil {
			item.Qty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "quoteQty"); err == nil {
			item.QuoteQty = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "commission"); err == nil {
			item.Commission = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "commissionAsset"); err == nil {
			item.CommissionAsset = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "time"); err == nil {
			item.Time = lib.ConvertIntToTime(v, 0)
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetBoolean(value, "isBuyer"); err == nil {
			item.IsBuyer = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetBoolean(value, "isMaker"); err == nil {
			item.IsMaker = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetBoolean(value, "isBestMatch"); err == nil {
			item.IsBestMatch = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		results = append(results, item)
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}

// GetOrderRateLimit

type GetOrderRateLimitParam struct {
	RecvWindow int64 `json:"recvWindow" param:"recvWindow" validate:"max=60000"`
}

type GetOrderRateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int64  `json:"intervalNum"`
	Limit         int64  `json:"limit"`
	Count         int64  `json:"count"`
}

func (r *Parser) ParseGetOrderRateLimit(b []byte) ([]*GetOrderRateLimit, error) {
	results := make([]*GetOrderRateLimit, 0)

	_, err := jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, _err error) {
		item := new(GetOrderRateLimit)

		if v, err := jsonparser.GetString(value, "rateLimitType"); err == nil {
			item.RateLimitType = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetString(value, "interval"); err == nil {
			item.Interval = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "intervalNum"); err == nil {
			item.IntervalNum = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "limit"); err == nil {
			item.Limit = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}
		if v, err := jsonparser.GetInt(value, "count"); err == nil {
			item.Count = v
		} else {
			if _err = r.errorParser(err); _err != nil {
				return
			}
		}

		results = append(results, item)
	})
	if err = r.errorParser(err); err != nil {
		return nil, err
	}

	return results, nil
}
