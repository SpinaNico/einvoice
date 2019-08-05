module invoice

go 1.12

replace invoice/header v0.0.0 => ./header

replace invoice/body v0.0.0 => ./body

replace invoice/share v0.0.0 => ./share

require (
	invoice/body v0.0.0
	invoice/header v0.0.0
	invoice/share v0.0.0
)
