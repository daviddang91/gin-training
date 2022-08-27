run:
	@echo "********** Auto Compile **********"
	CompileDaemon --build="go build -o ./out/app main.go" --command="./out/app"