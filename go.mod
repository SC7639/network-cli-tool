module github.com/sc7639/network-cli-tool

go 1.13

require (
	github.com/sc7639/network-cli-tool/internal/app/network-cli v0.0.0
	github.com/urfave/cli v1.22.3
)

replace github.com/sc7639/network-cli-tool/internal/app/network-cli => ./internal/app/network-cli/
