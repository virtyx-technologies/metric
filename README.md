# Virtyx Metrics

Virtyx consumes metrics and tracks series over time for monitoring data. Metrics can be sent in one of two formats. The first is a well formed JSON object that contains the metrics and any metadata.

```
{
  "metrics": [...]
}
```

The second is a string format, which can be used by scripts to easily send data to Virtyx:

```
great.metric.name|123|{"tag":"name"}
```

With tags being optional. Lastly, Virtyx can also injest [Prometheus metrics](https://prometheus.io/docs/instrumenting/exposition_formats/) as well:

```
cpu_usage{service="nginx",host="machine1"} 34.6 1494595898000
```

# Usage

```
r := &metric.Response{}
r.Value("system.cpu.percent", 33.33, map[string]interface{}{"hostname": "my-great-hostname"})
```
