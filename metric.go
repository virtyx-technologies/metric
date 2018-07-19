// Virtyx Metrics
//
// Virtyx records metrics in a simple format of name, value, with optional tags
// being added.
package metric

import (
	"encoding/json"
	"fmt"
)

type (
	// The struct returned to the agent by a plugin,
	// containing all metrics and metadata and the
	// text of any error which may have occurred
	Response struct {
		Error     string      `json:"error"`
		Metrics   []*Metric   `json:"metrics"`
		Metadatas []*Metadata `json:"metadata"`
	}

	// Encapsulates a metric captured by a plugin
	Metric struct {
		Name  string                 `json:"name"`  // Metric name
		Value float64                `json:"value"` // Metric value
		Tags  map[string]interface{} `json:"tags"`  // Optional tags that can be used for searching
	}

	// Encapsulates metadata (textual information) captured by a plugin
	Metadata struct {
		Name string `json:"name"` // The metadata name
		Data string `json:"data"` // The data captured
	}
)

// Add a metric to the response
func (r *Response) Metric(m *Metric) {
	r.Metrics = append(r.Metrics, m)
}

// Helper function for recording a value
//  response.Value("system.cpu.percent", 23.32, nil)
func (r *Response) Value(name string, val float64, tags map[string]interface{}) {
	r.Metrics = append(r.Metrics, &Metric{Name: name, Value: val, Tags: tags})
}

// Add metadata to the response
func (r *Response) Metadata(m *Metadata) {
	r.Metadatas = append(r.Metadatas, m)
}

// Add metadata to the response
func (r *Response) Data(name string, data string) {
	r.Metadatas = append(r.Metadatas, &Metadata{Name: name, Data: data})
}

func (r *Response) FindMetric(name string) *Metric {
	for _, m := range r.Metrics {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (r *Response) FindMetadata(name string) *Metadata {
	for _, m := range r.Metadatas {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (m *Metric) String() string {
	tags, err := json.Marshal(m.Tags)
	if err != nil {
		tags = []byte("{}")
	}
	return fmt.Sprintf("%s|%f|%s", m.Name, m.Value, string(tags))
}

func (m *Metadata) String() string {
	return fmt.Sprintf("%s|%s", m.Name, m.Data)
}
