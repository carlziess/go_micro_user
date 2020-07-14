@echo off

start "prod1" go run main.go --server_address :8001&
start "prod2" go run main.go --server_address :8002&
start "prod3" go run main.go --server_address :8003
pause

//cmd同时运行这些命令，打开三个窗口