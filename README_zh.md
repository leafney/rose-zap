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

内部采用第三方库 [natefinch/lumberjack](https://github.com/natefinch/lumberjack) 实现日志切割操作，预设默认配置如下。可通过 `SetFileConfig` 及相关方法自定义配置。

```go
zapcore.AddSync(&lumberjack.Logger{
	Filename:   "logs/rzap.log", //日志文件存放目录，如果文件夹不存在会自动创建
	MaxSize:    1024,            //文件大小限制,单位MB
	MaxBackups: 0,               //最大保留日志文件数量
	MaxAge:     1,              //日志文件保留天数
	LocalTime:  true,            //日志文件分割是否采用本地时间，默认为UTC
	Compress:   false,           //是否压缩处理
})
```

#### 单日志文件

`OutSingleFile()` 方法将所有级别日志写入单个日志文件中，同时支持配置是否包含控制台输出。

默认日志文件名为 `logs/rzap.log`，可通过 `SetFileConfig` 方法自定义配置。

```go
    cfg := rzap.NewConfig().
        OutSingleFile(true)
```

#### 分级别输出多文件

通常情况下，为了便于运维人员排查日志，会将低于 `error` 级别的普通日志放到 `info.log` 文件中，`error` 及以上严重级别的日志放到 `error.log` 文件中。可以通过 `OutMultiFile()` 方法实现。

```go
    cfg := rzap.NewConfig().
        OutMultiFile(true)
```

默认日志文件名：普通级别日志为 `logs/info.log`，严重级别日志为 `logs/error.log`。可通过 `SetFileConfig` 及相关方法自定义配置。

#### 分级别输出到控制台和文件

某些情况下，需要将低于 `error` 级别的普通日志直接输出到控制台显示，仅将 `error` 及以上严重级别的日志放到 `error.log` 文件中。可以通过 `OutInfoConsoleErrorFile()` 方法实现。

```go
    cfg := rzap.NewConfig().
        OutInfoConsoleErrorFile()
```

默认日志文件名为 `logs/error.log`，可通过 `SetFileConfig` 及相关方法自定义配置。

#### 更改日志切割配置

内部采用第三方库 `Lumberjack` 实现日志切割操作，如果默认切割配置不满足需求，可以通过 `SetFileConfig` 方法实现自定义配置。

支持以下配置参数：

- `WithFileName()`
- `WithMaxSize()`
- `WithMaxBackups()`
- `WithMaxAge()`
- `WithLocalTime()`
- `WithCompress()`

```go
    cfg := rzap.NewConfig().
        OutSingleFile(true).
        SetFileConfig(WithMaxBackups(2),WithLocalTime(false))
```

#### 分级别多文件日志切割配置

对于分级别的日志文件切割配置，除了采用 `SetFileConfig()` 方法进行 **通用设置** 外，还可以通过 `SetInfoFileConfig()` 或 `SetErrorFileConfig()` 分别单独设置。

```go
    cfg := rzap.NewConfig().
        OutMultiFile(true).
        SetErrorFileConfig(WithFileName("logs/xyz.log"))
```

#### 注意事项

##### WithFileName

`SetFileConfig(WithFileName())` 仅支持 `OutSingleFile()` 方法的日志文件名的修改，对于多级别日志输出 `OutMultiFile()` 和 `OutInfoConsoleErrorFile()` 则需要通过 `SetInfoFileConfig()` 或 `SetErrorFileConfig()` 来修改。

##### WithXXXX

除了 `WithFileName()` 外其他的自定义配置项，按照如下优先级定义：

```
SetFileConfig() < SetInfoFileConfig()/SetErrorFileConfig()
```

----
