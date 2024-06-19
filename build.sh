# Windows amd64
# export CGO_ENABLED=1
# export GOOS=windows
# export GOARCH=amd64
# go build -o build/h-ui-windows-amd64.exe -trimpath -ldflags "-s -w"
# Mac amd64
# export CGO_ENABLED=1
# export GOOS=darwin
# export GOARCH=amd64
# go build -o build/h-ui-darwin-amd64 -trimpath -ldflags "-s -w"
# Linux 386
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=386
go build -o build/h-ui-linux-386 -trimpath -ldflags "-s -w"
# Linux amd64
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=amd64
go build -o build/h-ui-linux-amd64 -trimpath -ldflags "-s -w"
# Linux armv6
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=arm
export GOARM=6
go build -o build/h-ui-linux-armv6 -trimpath -ldflags "-s -w"
# Linux armv7
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=arm
export GOARM=7
go build -o build/h-ui-linux-armv7 -trimpath -ldflags "-s -w"
# Linux arm64
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=arm64
go build -o build/h-ui-linux-arm64 -trimpath -ldflags "-s -w"
# Linux ppc64le
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=ppc64le
go build -o build/h-ui-linux-ppc64le -trimpath -ldflags "-s -w"
# Linux s390x
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=s390x
go build -o build/h-ui-linux-s390x -trimpath -ldflags "-s -w"