build:
	go build -o gomod-viz .
	mv ./gomod-viz /usr/local/bin

clean:
	rm /usr/local/bin/gomod-viz
