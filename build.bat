@echo off
setlocal enabledelayedexpansion

REM target platform array
set platforms=windows/amd64 darwin/amd64 linux/386 linux/amd64 linux/arm/6 linux/arm/7 linux/arm64 linux/ppc64le linux/s390x

REM output directory
set output_dir=build

REM make sure the output directory exists
if not exist %output_dir% mkdir %output_dir%

REM compile for each target platform
for %%p in (%platforms%) do (
    set "platform=%%p"

    for /F "tokens=1,2,3 delims=/" %%a in ("!platform!") do (
        set "GOOS=%%a"
        set "GOARCH=%%b"
        set "GOARM=%%c"

        set "output_name=h-ui-!GOOS!-!GOARCH!"
        if "!GOARCH!" == "arm" (
            set "output_name=!output_name!v!GOARM!"
        )

        if "!GOOS!" == "windows" (
            set "output_name=!output_name!.exe"
        )

        echo Building for !GOOS!/!GOARCH! !GOARM!...
        set CGO_ENABLED=0
        set GOOS=!GOOS!
        set GOARCH=!GOARCH!
        if defined GOARM (
            set GOARM=!GOARM!
        )

        go build -o %output_dir%/!output_name! -trimpath -ldflags "-s -w"

        if !errorlevel! == 0 (
            echo Build succeeded for !GOOS!/!GOARCH! !GOARM!
        ) else (
            echo Error occurred during building for !GOOS!/!GOARCH! !GOARM!
            echo CGO_ENABLED=!CGO_ENABLED! GOOS=!GOOS! GOARCH=!GOARCH! GOARM=!GOARM!
            exit /b 1
        )
    )
)

echo All builds completed successfully!
endlocal
