package metric

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetricString(t *testing.T) {
	metric := &Metric{Name: "system.cpu.percent", Value: 1, Tags: nil}
	assert.Equal(t, "system.cpu.percent|1.000000|null", metric.String())
}

func TestMetricJson(t *testing.T) {
	metric := &Metric{Name: "system.cpu.percent", Value: 1, Tags: map[string]interface{}{"hostname": "my-great-hostname"}}
	json, err := json.Marshal(metric)
	assert.Nil(t, err)
	assert.Contains(t, string(json), `"name":"system.cpu.percent"`)
	assert.Contains(t, string(json), `"value":1`)
	assert.Contains(t, string(json), `"tags":{"hostname":"my-great-hostname"}`)
}

func TestMetadataJson(t *testing.T) {
	meta := &Metadata{Name: "metadata", Data: "cool"}
	json, err := json.Marshal(meta)
	assert.Nil(t, err)
	assert.Contains(t, string(json), `"name":"metadata"`)
	assert.Contains(t, string(json), `"data":"cool"`)
}

func TestMetadataString(t *testing.T) {
	meta := &Metadata{Name: "metadata", Data: "data"}
	assert.Equal(t, "metadata|data", meta.String())
}

func TestAddMetric(t *testing.T) {
	a := assert.New(t)
	r := &Response{}

	a.Equal(0, len(r.Metrics))
	metric := &Metric{Name: "system.cpu.percent", Value: 1, Tags: nil}
	r.Metric(metric)
	a.Equal(1, len(r.Metrics))
	a.Equal(metric, r.Metrics[0])
}

func TestValue(t *testing.T) {
	a := assert.New(t)
	r := &Response{}
	r.Value("system.cpu.percent", 1, nil)
	a.Equal(1, len(r.Metrics))
	a.Equal("system.cpu.percent", r.Metrics[0].Name)
	a.Equal(float64(1), r.Metrics[0].Value)
	a.Nil(r.Metrics[0].Tags)
}

func TestAddMetadata(t *testing.T) {
	a := assert.New(t)
	r := &Response{}

	a.Equal(0, len(r.Metadatas))
	metadata := &Metadata{Name: "metadata", Data: "unqualified"}
	r.Metadata(metadata)
	a.Equal(1, len(r.Metadatas))
	a.Equal(metadata, r.Metadatas[0])
}

func TestData(t *testing.T) {
	a := assert.New(t)
	r := &Response{}

	r.Data("metadata", "data")
	a.Equal(1, len(r.Metadatas))
	a.Equal("metadata", r.Metadatas[0].Name)
	a.Equal("data", r.Metadatas[0].Data)
}

func TestFindMetric(t *testing.T) {
	a := assert.New(t)
	r := &Response{}

	a.Equal(0, len(r.Metrics))
	m1 := &Metric{Name: "metric.one", Value: 0, Tags: nil}
	r.Metric(m1)
	m2 := &Metric{Name: "metric.two", Value: 1, Tags: nil}
	r.Metric(m2)
	found := r.FindMetric("metric.one")
	a.Equal(m1, found)
	notFound := r.FindMetric("metric.three")
	a.Nil(notFound)
}

func TestFindMetadata(t *testing.T) {
	a := assert.New(t)
	r := &Response{}

	m1 := &Metadata{Name: "metadata.one", Data: "321"}
	r.Metadata(m1)
	m2 := &Metadata{Name: "metadata.two", Data: "xyz"}
	r.Metadata(m2)
	found := r.FindMetadata("metadata.one")
	a.Equal(m1, found)
	notFound := r.FindMetadata("metadata.three")
	a.Nil(notFound)
}
