# logger GO library

### Install
```yaml
import:
- package: github.com/best-expendables-v2/logger
  version: x.x.x
```

### Rules:
Message is truncated at 128 characters

Level table

Priority | #1 | #2 | #3 | #4 | #5 | #6 | #7 | #8 
--- | --- | --- | --- |--- |--- |--- |--- |--- 
Level| emerg | alert | crit | err | warning | notice | info | debug


There are 8 Levels to provide to LoggerProvider
```go
EmergencyLevel
AlertLevel
CriticalLevel
ErrorLevel
WarningLevel
NoticeLevel
InfoLevel
DebugLevel
````
### Examples:

##### With content field
```go
    
	package main
    
    import (
        "github.com/best-expendables-v2/logger"
    	"context"
    )
    
    func main() {
        // Create loggerFactory with level Error above
        loggerFactory := logger.NewLoggerFactory(logger.ErrorLevel)

        // fake context,
        // in real cases, it should be current app's context
        ctx := context.TODO()
    
        // withField will be filled into "content" field
        loggerFactory.Logger(ctx).WithField("test key", "test").Alert("Test")
    
        // simple log, no content field
        loggerFactory.Logger(ctx).Debug("Testing")
        loggerFactory.Logger(ctx).Critical("Testing")
        loggerFactory.Logger(ctx).Error("Testing")
        loggerFactory.Logger(ctx).Warning("Testing")
        loggerFactory.Logger(ctx).Notice("Testing")
    }

```
##### Without content field
```go
     package main
     
     import (
     	"github.com/best-expendables-v2/logger"
     	"context"
     )
     
     func main() {
        // Create loggerFactory with level Error above
        loggerFactory := logger.NewLoggerFactory(logger.ErrorLevel)

        
        // fake context,
        // in real cases, it should be current app's context
        ctx := context.TODO()
    
        // simple log, no content field
        loggerFactory.Logger(ctx).Debug("Testing")
        loggerFactory.Logger(ctx).Critical("Testing")
        loggerFactory.Logger(ctx).Error("Testing")
        loggerFactory.Logger(ctx).Warning("Testing")
        loggerFactory.Logger(ctx).Notice("Testing")
     }

```
##### simple log without any fields
```go
    
	package main
    
    import (
    	"github.com/best-expendables-v2/logger"
    )
    
    func main() {
    	logger.Info("Test log")
    }

```