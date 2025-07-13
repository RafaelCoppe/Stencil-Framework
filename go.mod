module github.com/RafaelCoppe/Stencil-Framework

go 1.24.1

replace github.com/RafaelCoppe/Stencil-Go => ../Stencil-Go

replace github.com/RafaelCoppe/Stencil-Framework/core/framework => ./core/framework

replace github.com/RafaelCoppe/Stencil-Framework/components => ./components

require (
	github.com/RafaelCoppe/Stencil-Framework/components v0.0.0-00010101000000-000000000000
	github.com/RafaelCoppe/Stencil-Framework/core/framework v0.0.0-20250712155831-1876e3cba7fa
	github.com/RafaelCoppe/Stencil-Go v1.1.1
)
