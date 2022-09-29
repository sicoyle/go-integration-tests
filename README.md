# go-integration-tests
This repo is meant to contain code for sharing insights on integration testing solutions using Go.

## To run

```bash
make run
```

### Check endpoint
```bash
curl localhost:8081/
```

## To test
```bash
make integration-test
```

## To test and receive a test report
```bash
make integration-test-report
```