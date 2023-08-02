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