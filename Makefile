arm:
	env GOOS=linux GOARCH=arm go build
native:
	go build
clean:
	rm ./remo-numeric-keypad

