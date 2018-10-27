voll-daneben: voll-daneben.go
	go build voll-daneben.go

voll-daneben.exe: voll-daneben.go
	env GOOS=windows GOARCH=amd64 go build voll-daneben.go
