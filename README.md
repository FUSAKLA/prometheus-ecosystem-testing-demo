# Prometheus testing
Demo to show what can be tested in the Prometheus ecosystem.

Commands are meant to be run from root of this repository.

### Used tools:
 - [promtool](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/#syntax-checking-rules)
 - [amtool](https://github.com/prometheus/alertmanager#amtool)
 - [promruval](https://github.com/FUSAKLA/promruval)
 - [Go prometheus client library](https://pkg.go.dev/github.com/prometheus/client_golang@v1.12.2/prometheus/testutil)


## Validate Prometheus configuration
```bash
promtool check config prometheus.yaml
```

## Validate metrics

### Locally
```bash
cat metrics.txt | promtool check metrics
```

### Live app
Terminal 1
```bash
pushd app/
go run main.go 
popd
```
Terminal 2
```bash
curl http://localhost:8080/metrics | promtool check metrics
```


## Test instrumentation
```bash
pushd app/
go test -v ./...
popd
```

## Validate rules
```bash
promtool check rules rules.yaml
```

## Unit test rules
```bash
promtool test rules rules_test.yaml
```

## validate rules semantics
```bash
promruval validate --config-file validation.yaml rules.yaml
```

## Check Alertmanager configuration
```bash
amtool check-config alertmanager.yaml
```

## Test Alertmanager templating
```bash
amtool template render --template.glob notification_template.tmpl --template.data alert.json --template.text='{{ template "test" . }}'
```

## Test Alertmanager routing
```bash
amtool config routes test --config.file alertmanager.yaml --verify.receivers team-A --tree team=team-A
```