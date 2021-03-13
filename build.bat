xcopy res build\\res /e /q /y
go build -tags static -ldflags "-s -w" -o build/game.exe