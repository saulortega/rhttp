package rhttp

type Resultado interface {
	Res() (*ResultadoExitoso, *Err)
}
