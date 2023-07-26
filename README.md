# rose-zap

## Installation

```shell
go get -u github.com/leafney/rose-zap
```
 
Import:

```go
import (
    rzap "github.com/leafney/rose-zap"
)
```

## Quick Start

```go
func main(){
    cfg := rzap.NewConfig()
    
    log := rzap.NewLogger(cfg)
    defer log.Sync()
    
    // Use Logger
    log.Info("Say Hello", zap.String("name", "tom"))
    
    // Use SugaredLogger
    log.SInfof("Fetch url: %s", url)
}
```

## Config
