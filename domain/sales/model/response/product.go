package response

type TierResp struct {
	Id               string `json:"id"`
	TierName         string `json:"tier_name"`
	Feature          string `json:"feature"`
	Limitation       string `json:"limitation"`
	LengthLimitation string `json:"length_limitation"`
}
