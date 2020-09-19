package meta

import "github.com/maengsanha/instacrawler/model/instagram"

// Response represents the body of meta search response.
type Response struct {
	HigherLayer      []instagram.Post `json:"higher_layer"`
	LowerLayer       []instagram.Post `json:"lower_layer"`
	HigherLayerCache []string         `json:"higher_layer_cache"`
	LowerLayerCache  []string         `json:"lower_layer_cache"`
}
