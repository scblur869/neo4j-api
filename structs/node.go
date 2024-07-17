package structs

type Node struct {
	Alias      string         `json:"alias"`
	Label      string         `json:"label"`
	ElementId  string         `json:"elementId,omitempty"`
	Properties map[string]any `json:"properties,omitempty"`
}

type Relationship struct {
	Relation           string         `json:"relation"`
	RelationProperties map[string]any `json:"relationProperties"`
	FromAlias          string         `json:"fromAlias"`
	FromLabel          string         `json:"fromLabel"`
	FromProperties     map[string]any `json:"fromProperties"`
	ToAlias            string         `json:"toAlias"`
	ToLabel            string         `json:"toLabel"`
	ToProperties       map[string]any `json:"toProperties"`
}

type DbActions struct {
	Action string `json:"action"`
	Node   string `json:"node"`
	Key    string `json:"key"`
}
