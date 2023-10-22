BIN_DIRS=$(wildcard ./cmd/*)
BIN_NAMES=$(foreach dir, ${BIN_DIRS}, $(shell basename ${dir}))
BIN_PATHS=$(foreach name, ${BIN_NAMES}, ./cmd/${name}/${name}.go)

BIN_DIR=./bin

default: build

# build all files in ./cmd
build: outdir
	@for path in ${BIN_PATHS}; 			\
	do 									\
		echo "building $$path";			\
		go build -o ${BIN_DIR} $$path;	\
	done;

outdir:
	@mkdir -p ${BIN_DIR}

clean:
	@rm ${BIN_DIR}/*

keys: # ${BIN_DIR}/server.key ${BIN_DIR}/server.crt
	openssl req \
		-new \
		-newkey rsa:4096 \
		-days 365 \
		-nodes \
		-x509 \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=Default/CN=https2-rapid-reset-example" \
		-keyout ${BIN_DIR}/server.key \
		-out ${BIN_DIR}/server.crt
