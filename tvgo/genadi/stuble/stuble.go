package stuble

import _ "embed"

//go:embed router.tpl
var Router string

//go:embed create.tpl
var Create string

//go:embed edit.tpl
var Edit string

//go:embed page.tpl
var Page string

//go:embed view.tpl
var View string

//go:embed service.tpl
var Service string
