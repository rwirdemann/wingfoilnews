build-linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/xnews cmd/main.go

clean:
	rm -rf bin
	rm -rf tmp

deploy: build-linux
	ssh wingfoilnews@95.217.180.178 "pkill xnews"
	scp bin/xnews wingfoilnews@95.217.180.178:~/xnews
	ssh wingfoilnews@95.217.180.178 "sh -c 'nohup ~/xnews > xnews.out 2>&1 &'"