# Go パラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
OUTPUTDIR=output
BINARY_NAME=ssh-hosts
INSTALL_PATH=/usr/local/bin/


all: build install clean

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

install: build
	mv $(BINARY_NAME) $(INSTALL_PATH)
	@echo "Install $(BINARY_NAME) command"

remove: 
	rm $(INSTALL_PATH)$(BINARY_NAME)
