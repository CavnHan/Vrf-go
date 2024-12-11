GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

VRF_ABI_ARTIFACT := ./abis/vrf/Vrf.json

vrf:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd

clean:
	rm vrf

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings:
	$(eval temp := $(shell mktemp))

	cat $(VRF_ABI_ARTIFACT) \
    	| jq -r .bytecode.object > $(temp)

	cat $(VRF_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg bindings \
		--abi - \
		--out bindings/vrf.go \
		--type VRF \
		--bin $(temp)

		rm $(temp)
.PHONY: \
	vrf \
	bindings \
	clean \
	test \
	lint