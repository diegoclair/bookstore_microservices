package entity

import "github.com/olivere/elastic"

// Item entity
type Item struct {
	ID               string      `json:"id"`
	Seller           int64       `json:"seller"`
	Title            string      `json:"title"`
	Description      Description `json:"description"`
	Pictures         []Picture   `json:"pictures"`
	Video            string      `json:"video"`
	Price            float32     `json:"price"`
	AvailbleQuantity int         `json:"available_quantity"`
	SoldQuantity     int         `json:"sold_quantity"`
	Status           string      `json:"status"`
}

// Description entity
type Description struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
}

// Picture entity
type Picture struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

// EsQuery struct to handlw search elasticsearch documents
type EsQuery struct {
	Equals    []FieldValue `json:"equals"`
	NotEquals []FieldValue `json:"not_equals"`
}

// FieldValue struct to handlw search elasticsearch documents
type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Build the elastic quer with the fields and values
func (q EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()
	equalsQueries := make([]elastic.Query, 0)
	for _, eq := range q.Equals {
		equalsQueries = append(equalsQueries, elastic.NewMatchQuery(eq.Field, eq.Value))
	}
	for _, eq := range q.NotEquals {
		equalsQueries = append(equalsQueries, elastic.NewBoolQuery().MustNot(elastic.NewMatchQuery(eq.Field, eq.Value)))
	}

	query.Must(equalsQueries...)
	return query
}
