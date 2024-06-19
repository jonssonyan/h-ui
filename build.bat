::Windows amd64
::SET CGO_ENABLED=1
::SET GOOS=windows
::SET GOARCH=amd64
::go build -o build/h-ui-windows-amd64.exe -trimpath -ldflags "-s -w"
::Mac amd64
::SET CGO_ENABLED=1
::SET GOOS=darwin
::SET GOARCH=amd64
::go build -o build/h-ui-darwin-amd64 -trimpath -ldflags "-s -w"
::Linux 386
SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=386
go build -o build/h-ui-linux-386 -trimpath -ldflags "-s -w"
::Linux amd64
SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=amd64
go build -o build/h-ui-linux-amd64 -trimpath -ldflags "-s -w"
::Linux armv6
SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=arm
SET GOARM=6
go build -o build/h-ui-linux-armv6 -trimpath -ldflags "-s -w"
::Linux armv7
SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=arm
SET GOARM=7
go build -o build/h-ui-linux-armv7 -trimpath -ldflags "-s -w"
::Linux arm64
SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=arm64
go build -o build/h-ui-linux-arm64 -trimpath -ldflags "-s -w"
::Linux ppc64le
SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=ppc64le
go build -o build/h-ui-linux-ppc64le -trimpath -ldflags "-s -w"
::Linux s390x
SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=s390x
go build -o build/h-ui-linux-s390x -trimpath -ldflags "-s -w"