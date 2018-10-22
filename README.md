Steps I did:
* I copied https://raw.githubusercontent.com/hajimehoshi/ebiten/244e078810bec2e713fc0eac324d148cde81748f/examples/audioinfiniteloop/main.go to main.go
* changed vorbis to mp3
* GO111MODULE=on go mod init
* GO111MODULE=on go mod vendor
* GO111MODULE=on go get ./..
* GO111MODULE=on go run main.go
