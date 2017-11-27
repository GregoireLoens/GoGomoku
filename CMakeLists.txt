cmake_minimum_required(VERSION 3.2)

execute_process(COMMAND go build -o pbrain-gogomoku.exe src/main/go/main.go)