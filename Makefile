PACKAGE=github.com/nikhil1raghav/kindle-send
CURRENT_DIR=$(shell pwd)
DIST_DIR=${CURRENT_DIR}/dist

backend:
	CGO=0 go build -v -o ${DIST_DIR}/medastra-backend ./main.go