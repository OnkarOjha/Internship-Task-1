@echo off
echo Running go mod tidy...
go mod tidy

echo.
echo Running go run main.go...
go run main.go
