# rose-zap

Wrapper class for the logging library [uber-go/zap](https://github.com/uber-go/zap)

English | [简体中文](README_zh.md)

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

----

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

----

## Configuration

### Default Values

```go
    cfg := rzap.NewConfig()
```

- Default log level is `info`
- Default output is to console window only
- Default line ending is `\n`
- Default output format is `json`
- Default display of file name and line number
- Default output of stack trace for levels above `warning`

### Log Level

Change the default log level using the `SetLevel()` method. The default value is `info`.

Valid log levels are: `debug`, `info`, `warn`, `error`, `panic`, `fatal`

```go
    cfg := rzap.NewConfig().
        SetLevel("debug")
```

### Enable switch

Use the `SetEnable()` method to change whether to enable or disable global log output. The default is `true`.

### Encoder

Change the output encoder using the `UseFmtJson()` method. The default is `Json` format, and it can be changed to `Plain Text` format.

```go
    cfg := rzap.NewConfig().
        UseFmtJson(false)
```

### File Name and Line Number

#### Toggle Display

Control whether to display the file name and line number using the `ShowCaller()` method.

```go
    cfg := rzap.NewConfig().
        ShowCaller(false)
```

#### Change Caller Skip

If you need to wrap functions like `Info`, `SInfow`, etc., you can skip the call to the wrapper function using the `SetCallSkip()` method. The default value is `SetCallSkip(1)`.

```go
    cfg := rzap.NewConfig().
        SetCallSkip(2).
        ShowCaller(true)
```

### Display Stack Trace

Set whether to display the stack trace using the `ShowStacktrace()` method. By default, stack traces are displayed for levels above `warning`.

```go
    cfg := rzap.NewConfig().
        ShowStacktrace(false)
```

### File Output

#### Default Configuration

Internal use of the third-party library [natefinch/lumberjack](https://github.com/natefinch/lumberjack) to implement log rotation. The default configuration is as follows. Custom configurations can be set using `SetFileConfig` and related methods.

```go
zapcore.AddSync(&lumberjack.Logger{
	Filename:   "logs/rzap.log", // Log file directory, the folder will be created automatically if it does not exist
	MaxSize:    1024,            // File size limit, in MB
	MaxBackups: 0,               // Maximum number of log files to retain
	MaxAge:     1,              // Number of days to retain log files
	LocalTime:  true,            // Whether to use local time for log file rotation, default is UTC
	Compress:   false,           // Whether to compress the log file
})
```

#### Single Log File

The `OutSingleFile()` method writes all log levels to a single log file, and the configuration can include console output.

The default log file name is `logs/rzap.log`, which can be customized using the `SetFileConfig` method.

```go
    cfg := rzap.NewConfig().
        OutSingleFile(true)
```

#### Multiple Files by Log Level

Typically, for easier troubleshooting by operations personnel, normal logs below the `error` level are placed in the `info.log` file, and logs at the `error` level and above are placed in the `error.log` file. This can be achieved using the `OutMultiFile()` method.

```go
    cfg := rzap.NewConfig().
        OutMultiFile(true)
```

The default log file names are: `logs/info.log` for normal level logs, and `logs/error.log` for severe level logs. Custom configurations can be set using the `SetFileConfig` and related methods.

#### Separate Console and File Output by Log Level

In some cases, it is necessary to output normal logs below the `error` level directly to the console, and only place logs at the `error` level and above in the `error.log` file. This can be achieved using the `OutInfoConsoleErrorFile()` method.

```go
    cfg := rzap.NewConfig().
        OutInfoConsoleErrorFile()
```

The default log file name is `logs/error.log`, which can be customized using the `SetFileConfig` and related methods.

#### Modify Log Rotation Configuration

The internal log rotation operation is implemented using the third-party library `Lumberjack`. If the default rotation configuration does not meet requirements, custom configurations can be set using the `SetFileConfig` method.

The following configuration parameters are supported:

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

#### Log Rotation Configuration by Log Level

For log file rotation configuration by log level, in addition to using the `SetFileConfig()` method for **general settings**, you can also use the `SetInfoFileConfig()` or `SetErrorFileConfig()` methods to set them separately.

```go
    cfg := rzap.NewConfig().
        OutMultiFile(true).
        SetErrorFileConfig(WithFileName("logs/xyz.log"))
```

#### Notes

##### WithFileName

`SetFileConfig(WithFileName())` only supports modifying the log file name for the `OutSingleFile()` method. For multi-level log output using `OutMultiFile()` and `OutInfoConsoleErrorFile()`, use `SetInfoFileConfig()` or `SetErrorFileConfig()` to modify it.

##### WithXXXX

Except for `WithFileName()`, other custom configuration items are defined based on the following priority:

```
SetFileConfig() < SetInfoFileConfig()/SetErrorFileConfig()
```

----