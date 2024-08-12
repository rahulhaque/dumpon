#!/usr/bin/env bash

version=$1
if [[ -z "$version" ]]; then
  echo "usage: $0 <version>"
  exit 1
fi

rm -rf bin/
mkdir -p bin

platforms=(
	"darwin/amd64"
	"darwin/arm64"
	"linux/amd64"
	"linux/arm64"
	"windows/amd64"
)

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}

	os=$GOOS
	if [ $os = "darwin" ]; then
        os="macOS"
    fi

	output_name="dumpon-${version}-${os}-${GOARCH}"
    zip_name=$output_name
    if [ $os = "windows" ]; then
        output_name+='.exe'
    fi

	echo "Building bin/$output_name..."
    env GOOS=$GOOS GOARCH=$GOARCH go build \
      -ldflags "-X main.version=$version" \
      -o bin/$output_name
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting.'
        exit 1
    fi

	pushd bin > /dev/null
    if [ $os = "windows" ]; then
        zip $zip_name.zip $output_name
        rm $output_name
    else
        chmod a+x $output_name
        gzip $output_name
    fi
    popd > /dev/null
done