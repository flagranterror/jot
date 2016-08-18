PROJECT = jot
VERSION = v1.0.4
RELEASE = $(PROJECT)-$(VERSION)
BUILDDIR = dist/$(RELEASE)
.PHONY: scripts clean

all: win64 win32 linux64 linux32 darwin

install:
	go install

win64:
	GOOS=windows GOARCH=amd64 go build -o $(BUILDDIR)/$(PROJECT).exe

win32:
	GOOS=windows GOARCH=386 go build -o $(BUILDDIR)/$(PROJECT)32.exe

linux64:
	GOOS=linux GOARCH=amd64 go build -o $(BUILDDIR)/$(PROJECT)_linux

linux32:
	GOOS=linux GOARCH=386 go build -o $(BUILDDIR)/$(PROJECT)_linux32

darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILDDIR)/$(PROJECT)_darwin

scripts:
	cp scripts/* $(BUILDDIR) && cp README.md $(BUILDDIR) && cp LICENSE $(BUILDDIR)

zip: scripts
	cd dist/ && zip -r $(RELEASE).zip $(RELEASE)

tar: scripts
	cd dist/ && tar -czf $(RELEASE).tar.gz $(RELEASE)

package: zip tar

clean:
	rm -rf dist/*
