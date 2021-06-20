go build -v -ldflags="-X 'fiber-101/build.Version=$(./version.sh)' -X 'fiber-101/build.User=$(id -u -n)' -X 'fiber-101/build.Time=$(date)'"
