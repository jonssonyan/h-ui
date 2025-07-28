#!/usr/bin/env bash

# target platform array
platforms=(
  "windows/amd64"
  "darwin/amd64"
  "linux/386"
  "linux/amd64"
  "linux/arm/6"
  "linux/arm/7"
  "linux/arm64"
  "linux/ppc64le"
  "linux/s390x"
)

# output directory
output_dir="build"

# make sure the output directory exists
mkdir -p "$output_dir"

# compile for each target platform
for platform in "${platforms[@]}"; do
  IFS='/' read -r -a platform_split <<<"$platform"
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}
  GOARM=${platform_split[2]}

  output_name="h-ui-${GOOS}-${GOARCH}"
  if [ "$GOARCH" == "arm" ]; then
    output_name+="v${GOARM}"
  fi

  # add the appropriate extension
  if [ "$GOOS" == "windows" ]; then
    output_name+=".exe"
  fi

  echo "Building for $GOOS/$GOARCH ${GOARM:-}"
  build_env="CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH"
  [ "$GOARCH" == "arm" ] && build_env+=" GOARM=$GOARM"

  env $build_env go build -o "$output_dir/$output_name" -trimpath -ldflags "-s -w"
  if [ $? -ne 0 ]; then
    echo "Error occurred during building for $GOOS/$GOARCH ${GOARM:-}"
    exit 1
  fi
done

echo "All builds completed successfully!"
