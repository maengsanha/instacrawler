package meta

// Request represents the body of meta search request.
type Request struct {
	HigherLayer      []string `json:"higher_layer"`
	LowerLayer       []string `json:"lower_layer"`
	HigherLayerCache []string `json:"higher_layer_cache"`
	LowerLayerCache  []string `json:"lower_layer_cache"`
}
