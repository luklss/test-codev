test:
	go test -race -coverprofile=profile.out -covermode=atomic .
	cat profile.out >> coverage.txt
	rm profile.out
