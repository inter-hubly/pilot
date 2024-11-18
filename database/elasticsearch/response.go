package elasticsearch

type Response struct {
	Source       map[string]interface{} `json:"_source,omitempty"`
	Shard        Shard                  `json:"_shards,omitempty"`
	Index        string                 `json:"_index,omitempty"`
	Type         string                 `json:"_type,omitempty"`
	ID           string                 `json:"_id"`
	Result       string                 `json:"result,omitempty"`
	Version      int                    `json:"_version,omitempty"`
	SecNo        int                    `json:"_seq_no,omitempty"`
	PrimaryTerm  int                    `json:"_primary_term,omitempty"`
	ForceRefresh bool                   `json:"force_refresh,omitempty"`
}

type Shard struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}
