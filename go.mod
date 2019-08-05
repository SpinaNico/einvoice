module github.com/SpinaNico/go-struct-invoice

go 1.12

replace github.com/SpinaNico/go-struct-invoice/header v0.0.0 => ./header

replace github.com/SpinaNico/go-struct-invoice/body v0.0.0 => ./body

replace github.com/SpinaNico/go-struct-invoice/share v0.0.0 => ./share

require (
	github.com/SpinaNico/go-struct-invoice/body v0.0.0
	github.com/SpinaNico/go-struct-invoice/header v0.0.0
	github.com/SpinaNico/go-struct-invoice/share v0.0.0
)
