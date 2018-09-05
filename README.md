# Coffee Maker API

![](https://github.com/todd1794/test/coffeemaker.png)

This is a not complete. This is a first check.


### Requirements

```
go get github.com/gin-gonic/gin
go get github.com/satori/go.uuid
go get gopkg.in/go-playground/validator.v9
go get -u github.com/Depado/ginprom
go get -u github.com/cweill/gotests/...
go get -u github.com/shenwei356/rush/
```

### Building and running

```
go build coffeeMachine.go 
./coffeeMachine
```


### Make a cup

```
curl -s -X POST localhost:8080/BrewCup
```

### Make 1000 cups

```
for a in `seq 1000`;do curl -s -X POST localhost:8080/BrewCup;done
```

### Make 1000 cups at once

```
seq 1000 | rush -j 1000 'curl -s -X POST localhost:8080/BrewCup'
```

### Scrape Metrics

```
curl -s localhost:8080/metrics
```

