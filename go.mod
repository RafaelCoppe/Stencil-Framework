module github.com/RafaelCoppe/Stencil-Framework

go 1.24.1

replace github.com/RafaelCoppe/Stencil-Go => ../Stencil-Go

replace github.com/RafaelCoppe/Stencil-Framework/framework => ./framework

require (
	github.com/RafaelCoppe/Stencil-Framework/framework v0.0.0
	github.com/RafaelCoppe/Stencil-Go v1.0.0
)
