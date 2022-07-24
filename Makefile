generate: gen-server main.go
	go build

gen-server: swagger.yml
	rm -rf restapi cmd
	swagger generate server --exclude-main
