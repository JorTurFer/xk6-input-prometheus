import prometheus from 'k6/x/prometheusread';

export default function () {
  var client = prometheus.newPrometheusClient("http://demo.robustperception.io:9090", "user", "password")
  var response = client.queryScalar("sum(alertmanager_notifications_total)")
  console.log(response)
}