module github.com/AndreVeiga/go-first

go 1.19

replace view => ./view

replace repository => ./repository

replace integration => ./integration

replace core => ./core

require (
	core v0.0.0-00010101000000-000000000000
	view v0.0.0-00010101000000-000000000000
)

require (
	integration v0.0.0-00010101000000-000000000000 // indirect
	repository v0.0.0-00010101000000-000000000000 // indirect
)
