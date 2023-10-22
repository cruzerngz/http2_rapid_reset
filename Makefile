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
