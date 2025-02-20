BINARY_NAME=task-tracker
SRC=cmd/task-tracker/main.go

build:
	go build -o $(BINARY_NAME) $(SRC)

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

.PHONY: build run clean
