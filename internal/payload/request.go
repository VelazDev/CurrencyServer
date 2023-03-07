package payload

type TickerRequest struct {
	Ask          string `json:"ask,omitempty"`
	Bid          string `json:"bid,omitempty"`
	Last         string `json:"last,omitempty"`
	Low          string `json:"low,omitempty"`
	High         string `json:"high,omitempty"`
	Open         string `json:"open,omitempty"`
	Volume       string `json:"volume,omitempty"`
	Volume_quote string `json:"volume_quote,omitempty"`
}

type SymbolRequest struct {
	BaseCurrency string `json:"base_currency,omitempty"`
	FeeCurrency  string `json:"fee_currency,omitempty"`
}

type CurrencyRequest struct {
	FullName string `json:"full_name"`
}

type Request interface {
	TickerRequest | SymbolRequest | CurrencyRequest
}
