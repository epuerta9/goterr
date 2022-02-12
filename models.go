package main

type Resource struct {
	Address string `json:"address,omitempty"`
	Change  struct {
		Actions []string `json:"actions,omitempty"`
	} `json:"change,omitempty"`
	Type string `json:"type"`
}

type ResourceChanges struct {
	ResourceChange []Resource `json:"resource_changes,omitempty"`
}
