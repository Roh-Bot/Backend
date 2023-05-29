package DSS

type DSS struct {
	Id1 struct {
		StrategyName string   `json:"strategyName"`
		Stocks       []string `json:"stocks"`
	} `json:"id1"`
	Id2 struct {
		StrategyName string   `json:"strategyName"`
		Stocks       []string `json:"stocks"`
	} `json:"id2"`
}
