install:
	go install ./cmd/...

run_tradeclient:
	`go env GOPATH`/bin/tradeclient

run_webui:
	`go env GOPATH`/bin/webui
