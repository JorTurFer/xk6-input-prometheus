build:
	xk6 build --with github.com/JorTurFer/xk6-input-prometheus=.

test: build
	go test ./...
	./k6 run -i 1 -u 1 example.js