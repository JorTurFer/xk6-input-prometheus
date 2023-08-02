import prometheus from 'k6/x/prometheusread';

export default function () {
  var client = prometheus.newPrometheusClient("http://demo.robustperception.io:9090", "user", "password");
  var response = client.queryScalar("alertmanager_notifications_total");
  
  response.forEach(item => {   
    var jsonBytes = item.marshalJSON()
    var {metric, value} = JSON.parse(String.fromCharCode(...jsonBytes))
    
    console.log(metric.__name__) // show metric name
    console.log(value[1]) // show metric value
  });
}