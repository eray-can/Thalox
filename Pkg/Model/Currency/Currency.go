package Currency

type Currency struct {
	SellPrice   string `json:"satis"`
	BuyingPrice string `json:"alis"`
	Variation   string `json:"degisim"`
	Ratio       string `json:"d_oran"`
	Direction   string `json:"d_yon"`
}

type ExchangeRates struct {
	USD   Currency `json:"USD"`
	EUR   Currency `json:"EUR"`
	GBP   Currency `json:"GBP"`
	GA    Currency `json:"GA"`
	C     Currency `json:"C"`
	GAG   Currency `json:"GAG"`
	BTC   Currency `json:"BTC"`
	ETH   Currency `json:"ETH"`
	XU100 struct {
		Satis   string `json:"satis"`
		Alis    string `json:"alis"`
		Degisim string `json:"degisim"`
	} `json:"XU100"`
}
