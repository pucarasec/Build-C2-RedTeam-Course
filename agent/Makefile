build: 
	mkdir -p build
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/x64-linux .
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/x64-windows .
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/x64-darwin .

copy:
	cp build/* ../malon_lp/render_launcher/