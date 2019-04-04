// +build wasm

package engine

 func divmod(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}
