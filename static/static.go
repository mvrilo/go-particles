package static

import (
	"embed"
)

//go:embed particles.js wasm_exec.js particles.wasm index.html
var Files embed.FS
