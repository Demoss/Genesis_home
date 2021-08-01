package resources

type Ticker struct{
	Base string `json:"base"`
	Target string `json:"target"`
	Price string `json:"price"`
	Volume string `json:"volume"`
	Change string `json:"change"`
}
type ResponceBTC struct {
	Ticker Ticker `json:"ticker"`
	Timestamp uint64 `json:"timestamp"`
	Success bool `json:"success"`
	Err string `json:"err"`
}