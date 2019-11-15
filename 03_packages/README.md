# Packages

To import more than one packages

```go
import (
    "pkg1"
    "pkg2"
    ...
)
```

To create your own package, make sure you create a folder and some files in there.

For example, let's make a folder ```/strutil``` for string utilities, and make a file ```reverse.go``` to reverse a string.

Once everything is created, we can import our packages into file

```go
import (
    "pkg1"
    "pkg2"
    "github.com/githubUserName/goCrashCourse/03_packages/strutil"
)
```


