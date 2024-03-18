build :
	@go build -o go_app

release : build 
	@./go_app 

clean :
	@rm -rf ./go_app

run: release clean

watch:
	@npx nodemon --exec "make run" ./main.go --signal SIGTERM