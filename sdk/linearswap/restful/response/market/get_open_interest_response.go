﻿package market

type GetOpenInterestResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Amount float32 `json:"amount"`

		Volume float32 `json:"volume"`

		Value float32 `json:"value"`

		TradeAmount float32 `json:"trade_amount"`

		TradeVolume float32 `json:"trade_volume"`

		TradeTurnover float32 `json:"trade_turnover"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
