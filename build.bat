cd src/main/go


go tool compile main.go


go tool link -o ../../../pbrain-gogomoku.exe main.o

cd ../../../
