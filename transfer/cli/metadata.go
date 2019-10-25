package cli

// Metadata holds docker plugin metadata
type Metadata struct {
    SchemaVersion string `json:"SchemaVersion"`
    Vendor        string `json:"Vendor"`
    Version       string `json:"Version"`
    Experimental  bool   `json:"Experimental"`
}
