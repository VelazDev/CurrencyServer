package payload

type DataResponse struct {
	Id          string  `json:"id"`
	FullName    string  `json:"fullName"`
	Ask         float64 `json:"ask"`
	Bid         float64 `json:"bid"`
	Last        float64 `json:"last"`
	Open        float64 `json:"open"`
	Low         float64 `json:"low"`
	High        float64 `json:"high"`
	FeeCurrency string  `json:"feeCurrency"`
}
