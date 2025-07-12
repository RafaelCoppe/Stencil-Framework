module github.com/RafaelCoppe/Stencil-Framework

go 1.24.1

replace github.com/RafaelCoppe/Stencil-Go => ../Stencil-Go

replace github.com/RafaelCoppe/Stencil-Framework/framework => ./framework

require (
	github.com/RafaelCoppe/Stencil-Framework/framework v0.0.0
	github.com/RafaelCoppe/Stencil-Go v1.0.3
	github.com/spf13/cobra v1.8.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
