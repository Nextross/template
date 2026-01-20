package wedos

import "encoding/json"

type Request struct {
	User    string          `json:"user"`
	Auth    string          `json:"auth"`
	Command string          `json:"command"`
	ClTRID  string          `json:"clTRID"`
	Test    string          `json:"test,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

type RequestEnvelope struct {
	Request Request `json:"request"`
}

type Response struct {
	Code      int             `json:"code"`
	Result    string          `json:"result"`
	Timestamp uint            `json:"timestamp"`
	ClTRID    string          `json:"clTRID"`
	SvTRID    string          `json:"svTRID"`
	Command   string          `json:"command"`
	Data      json.RawMessage `json:"data,omitempty"`
}

type ResponseEnvelope struct {
	Response Response `json:"response"`
}

type DNSRowsListResponse struct {
	Row Row `json:"row"`
}

type Row []RowItem

type RowItem struct {
	ID            string `json:"ID"`
	Name          string `json:"name"`
	TTL           string `json:"ttl"`
	Rdtype        string `json:"rdtype"`
	Rdata         string `json:"rdata"`
	ChangedDate   string `json:"changed_date"`
	AuthorComment string `json:"author_comment"`
}

type DNSAppendResponse struct {
	Domain string `json:"domain"`
	ID     string `json:"row_id"`
}

type recordKey struct {
	Name string
	Type string
	Data string
}
