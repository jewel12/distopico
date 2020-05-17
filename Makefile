.PHONY: build run start

export SLACK_TOKEN
DEVICE := 1
CLASSIFIER := data/classifier.xml

build: distopico

distopico:
	go build -o $@

run: distopico ${CLASSIFIER}
	./$< -d=${DEVICE} -c=${CLASSIFIER} -t=${SLACK_TOKEN}

start:
	watch -n60 $(MAKE) run

clean:
	-rm distopico