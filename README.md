# Prometheus reader k6 extension

K6 extension to read data from prometheus

## Build

To build a custom `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

1. Download [xk6](https://github.com/grafana/xk6):

    ```bash
    go install go.k6.io/xk6/cmd/xk6@latest
    ```

2. [Build the k6 binary](https://github.com/grafana/xk6#command-usage):

    ```bash
    xk6 build --with  github.com/JorTurFer/xk6-input-prometheus
    ```

   This will create a k6 binary that includes the xk6-input-prometheus extension in your local folder. This k6 binary can now run a k6 test.

### Development
To make development a little smoother, use the `Makefile` in the root folder. The default target will format your code, run tests, and create a `k6` binary with your local code rather than from GitHub.

```shell
git clone git@github.com/JorTurFer/xk6-input-prometheus.git
cd xk6-input-prometheus
make build
```

Using the `k6` binary with `xk6-input-prometheus`, run the k6 test as usual:

```bash
./k6 run example.js

```

## Examples: 

### Document Insertion Test
```js
import prometheus from 'k6/x/prometheusread';

export default function () {
    query()
    queryRange()
}

function query (){
    console.log("Query Sample")
    var client = prometheus.newPrometheusClient("http://demo.robustperception.io:9090", "user", "password")
    var response = client.query("alertmanager_notifications_total")
    
    response.forEach(item => {   
      var jsonBytes = item.marshalJSON()
      var {metric, value} = JSON.parse(String.fromCharCode(...jsonBytes))
  
      console.log(metric.__name__) // show metric name
      console.log(value[1]) // show metric value
    });
}

function queryRange (){
    console.log("QueryRange Sample")
    var client = prometheus.newPrometheusClient("http://demo.robustperception.io:9090", "user", "password")
    var end = new Date()
    var start = new Date(end.getTime() - 5 * 60000)
    var period = "minute"

    var response = client.queryRange("rate(prometheus_tsdb_head_samples_appended_total[5m])", start.toISOString(), end.toISOString(), period)

    response.forEach(item => {   
        var jsonBytes = item.marshalJSON()
        var {metric, values} = JSON.parse(String.fromCharCode(...jsonBytes))
    
        console.log(metric) // show metric
        console.log(values) // show metric values
    });
}

```
