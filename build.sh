export GOPROXY=direct

sudo apt-get update
sudo apt-get install gcc-mingw-w64-i686 gcc-multilib

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o ./linux_amd64_grafanaExp ./cmd
tar -czvf linux_amd64_grafanaExp.tar.gz linux_amd64_grafanaExp

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o ./windows_amd64_grafanaExp ./cmd
tar -czvf windows_amd64_grafanaExp.tar.gz windows_amd64_grafanaExp

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o ./darwin_amd64_grafanaExp ./cmd
tar -czvf darwin_amd64_grafanaExp.tar.gz darwin_amd64_grafanaExp

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o ./darwin_arm64_grafanaExp ./cmd
tar -czvf darwin_arm64_grafanaExp.tar.gz darwin_arm64_grafanaExp