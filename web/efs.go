package web

import "embed"

//go:embed "components" "static" "pages" "css"
var Files embed.FS
