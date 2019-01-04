########################################
### Build/Install

build:
ifeq ($(OS),Windows_NT)
	go build -o build/kepler.exe
else
	go build -o build/kepler
endif

install:
	go install
