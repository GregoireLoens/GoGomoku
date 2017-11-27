:: Move to sources folder
cd src/main/go

:: Compile all Go sources
go tool compile main.go

:: Link the executable
go tool link -o ../../../pbrain-gogomoku.exe main.o

:: Return to origin
cd ../../../
