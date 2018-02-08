# Binary name
BINARY=elasticHD
VERSION="1.4"
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Builds the project
build:
		go build -o ${BINARY} ${LDFLAGS}
release:
		# Clean
		go clean
		rm -rf *.gz
		# Build for mac
		go build ${LDFLAGS}
		tar czvf ${BINARY}-mac64-${VERSION}.tar.gz ./${BINARY}
		go clean
		# Build for linux
		go clean
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS}
		tar czvf ${BINARY}-linux64-${VERSION}.tar.gz ./${BINARY}
		# Build for win
		go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe ${LDFLAGS}"
		tar czvf ${BINARY}-win64-${VERSION}.tar.gz ./${BINARY}.exe
		go clean
# Cleans our projects: deletes binaries
clean:
		go clean
		rm -rf *.gz

.PHONY:  clean build
