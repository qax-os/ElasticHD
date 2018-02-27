# Binary name
BINARY=elasticHD
VERSION="1.4"
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Builds the project
build:
	cd main && go build -o ../${BINARY} ${LDFLAGS}
# webpack
webpack:
	@sh scripts/package.sh
release:
	# Clean
	make clean
	# Build for mac
	cd main && go build ${LDFLAGS} -o ../${BINARY}
	tar czvf ${BINARY}-mac64-${VERSION}.tar.gz ./${BINARY}
	# Build for linux
	cd main && go clean
	cd main && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ../${BINARY}
	tar czvf ${BINARY}-linux64-${VERSION}.tar.gz ./${BINARY}
	# Build for win
	cd main && go clean
	cd main && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../${BINARY}.exe ${LDFLAGS}
	tar czvf ${BINARY}-win64-${VERSION}.tar.gz ${BINARY}.exe
	cd main && go clean
# Cleans our projects: deletes binaries
clean:
	cd main && go clean
	rm -rf *.gz
	rm -rf *.exe

.PHONY:  clean build

