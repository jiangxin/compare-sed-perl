TARGET=compare-sed-perl

all: build
	@./$(TARGET)

build: $(TARGET)
$(TARGET): main.go
	go build -o $(TARGET)

clean:
	rm -f $(TARGET)


.PHONY: all build clean
