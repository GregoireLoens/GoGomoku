set GOPATH_TMP=%GOPATH%
set GOPATH=%GOPATH%;%cd%

go build -o pbrain-gogomoku.exe src/gogomoku_main/main.go

set GOPATH=%GOPATH_TMP%