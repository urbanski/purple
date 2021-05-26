
build:
	go build -o purplecli .
	chmod +x purplecli

test:
	go test -v ./...

vet:
	go vet .

image:
	docker build . -t wurb/purple:latest

clean:
	rm -f purplecli

coverage:
	go test -cover ./...