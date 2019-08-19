module github.com/SpinaNico/go-struct-invoice

go 1.12

replace github.com/SpinaNico/go-struct-invoice/header => ./header

replace github.com/SpinaNico/go-struct-invoice/body => ./body

replace github.com/SpinaNico/go-struct-invoice/share => ./share

require (
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1
)
