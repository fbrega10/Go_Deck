hello:
	echo "hello"

#run all the deck tests
test deck:
	cd Cards/src && go test

#run main.go (deck) w/o compiling binaries
run deck:
	cd Cards/src && go run main.go deck.go

#compile deck binaries
comp deck:
	cd Cards/src && go build deck.go main.go && sudo mv deck /bin/ 

#run structs file
run structs:
	cd Structs && go run main.go
