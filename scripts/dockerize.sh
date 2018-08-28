(cd web && yarn build)

env GOOS=linux GOARCH=arm go build -v cmd/exe/main.go

latesttag=$(git describe --tags)

docker build -t intrinsval:$latesttag .
