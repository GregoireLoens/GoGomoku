if ["%GOPATH%"] equ [] goto empty_gopath

set GOPATH_TMP=%GOPATH%
set GOPATH=%GOPATH%;%cd%

goto :execute_build

:empty_gopath
set GOPATH_TMP=
set GOPATH=%cd%

:execute_build
go build -o pbrain-gogomoku.exe .\src\gogomoku\main

set GOPATH=%GOPATH_TMP%
