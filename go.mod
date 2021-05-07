module github.com/torchioalexis/firstgo

go 1.16

require (
	github.com/labstack/echo v3.3.10+incompatible
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf // indirect
	golang.org/x/net v0.0.0-20210505214959-0714010a04ed // indirect
	golang.org/x/sys v0.0.0-20210507014357-30e306a8bba5 // indirect
)

replace github.com/labstack/echo => ./src/echo
