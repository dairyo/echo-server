generate: gen-server main.go
	go build main.go

gen-server: swagger.yml
	rm -rf restapi cmd
	swagger generate server --exclude-main
