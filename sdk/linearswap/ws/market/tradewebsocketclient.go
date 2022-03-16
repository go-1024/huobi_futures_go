package client

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/gostudys/huobi_futures_go/sdk/linearswap/ws"
	"github.com/gostudys/huobi_futures_go/sdk/linearswap/ws/response/market"
)

// Responsible to handle Trade data from WebSocket
type TradeWebSocketClient struct {
	ws.WebSocketClientBase
}

// Initializer
func (p *TradeWebSocketClient) Init(host string) *TradeWebSocketClient {
	p.WebSocketClientBase.Init(host, "/linear-swap-ws")
	return p
}

// Set callback handler
func (p *TradeWebSocketClient) SetHandler(
	connectedHandler ws.ConnectedHandler,
	responseHandler ws.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Request latest 300 trade data
func (p *TradeWebSocketClient) Request(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.trade.detail", symbol)
	req := fmt.Sprintf("{\"req\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(req)

	applogger.Info("WebSocket requested, topic=%s, clientId=%s", topic, clientId)
}

// Subscribe latest completed trade in tick by tick mode
func (p *TradeWebSocketClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.trade.detail", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe trade
func (p *TradeWebSocketClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.trade.detail", symbol)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *TradeWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubTradeDetailResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
