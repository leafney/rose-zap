# rose-zap
 
日志库 [uber-go/zap](https://github.com/uber-go/zap) 的封装类

----

## 安装

```shell
go get -u github.com/leafney/rose-zap
```

导入:

```go
import (
    rzap "github.com/leafney/rose-zap"
)
```

----

## 快速入门

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

----

## 配置

### 默认值

```go
    cfg := rzap.NewConfig()
```

- 默认输出为 `info` 级别
- 默认仅输出到控制台窗口
- 默认行尾结束符 `\n`
- 默认输出为 `json` 格式
- 默认显示文件名和行号
- 默认输出 `warning` 级别以上堆栈信息

### 日志级别

通过 `SetLevel()` 方法更改默认日志显示级别，默认值为 `info`。

可输入的日志级别为：`debug`、`info`、`warn`、`error`、`panic`、`fatal`

```go
    cfg := rzap.NewConfig().
        SetLevel("debug")
```

### 编码器

通过 `UseFmtJson()` 方法来更改输出的编码器，默认为 `Json` 格式，可更改为 `普通文本` 格式

```go
    cfg := rzap.NewConfig().
        UseFmtJson(false)
```

### 文件名和行号

#### 切换显示

通过 `ShowCaller()` 方法来控制是否显示文件名和行号

```go
    cfg := rzap.NewConfig().
        ShowCaller(false)
```

#### 更改调用函数

如果需要对 `Info`、`SInfow` 等函数再次进行包装，可以通过 `SetCallSkip()` 方法跳过对封装函数的调用。默认值为 `SetCallSkip(1)`

```go
    cfg := rzap.NewConfig().
        SetCallSkip(2).
        ShowCaller(true)
```

### 显示堆栈信息

通过 `ShowStacktrace()` 方法设置是否显示堆栈信息。默认情况下显示 `warning` 级别及以上的堆栈信息。

```go
    cfg := rzap.NewConfig().
        ShowStacktrace(false)
```

### 文件输出

#### 默认配置

采用 `lumberjack` 库实现文件切割。默认情况下，

```go
zapcore.AddSync(&lumberjack.Logger{
	Filename:   "logs/rzap.log", //日志文件存放目录，如果文件夹不存在会自动创建
	MaxSize:    1024,            //文件大小限制,单位MB
	MaxBackups: 0,               //最大保留日志文件数量
	MaxAge:     1,              //日志文件保留天数
	LocalTime:  true,            //日志文件分割采用本地时间
	Compress:   false,           //是否压缩处理
})
```

