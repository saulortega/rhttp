package rhttp

import (
	"errors"
	"net/http"
)

type Err struct {
	msj []*MensajeUsuario
	cdg int
	err error
}

func Error(err error) *Err {
	if err == nil {
		panic("Se intentó establecer error en nil para resultado con error.")
	}

	return transformarError(err)
}

func ErrorDeCliente(mensaje string) *Err {
	return errorDeClienteConEstadoHTTP(http.StatusBadRequest, mensaje)
}

func ErrorDeClienteConEstadoHTTP(estadoHTTP int, mensaje string) *Err {
	return errorDeClienteConEstadoHTTP(estadoHTTP, mensaje)
}

func ErrorDeServidor(err error, mensaje string) *Err {
	if len(mensaje) == 0 {
		panic("Se intentó establecer mensaje vacío para resultado con error.")
	}

	E := &Err{}
	E.cdg = http.StatusInternalServerError
	E.err = err
	E.Mensaje(MensajeUsuarioTipoError, mensaje)
	return E
}

func errorDeClienteConEstadoHTTP(estadoHTTP int, mensaje string) *Err {
	if len(mensaje) == 0 {
		panic("Se intentó establecer mensaje vacío para resultado con error.")
	}

	E := &Err{}
	E.cdg = estadoHTTP
	E.Mensaje(MensajeUsuarioTipoError, mensaje)
	return E
}

func (O *Err) Res() (*ResultadoExitoso, *Err) {
	return nil, O
}

func (O *Err) Error() string {
	if len(O.msj) > 0 {
		return O.msj[0].Mensaje
	}

	if O.err != nil {
		return O.err.Error()
	}

	// Esto no debería ocurrir
	return "error inesperado"
}

func (O *Err) EstadoHTTP(cdg int) *Err {
	if cdg < 400 || cdg > 599 {
		panic("Se intentó establecer código de estado HTTP inesperado para resultado con error.")
	} else if O.cdg > 0 {
		panic("Se intentó sobrescribir el código de estado HTTP para resultado con error.")
	}

	O.cdg = cdg
	return O
}

func (O *Err) ErrorOriginal(err error) *Err {
	if err == nil {
		panic("Se intentó establecer error en nil para resultado con error.")
	} else if O.err != nil {
		panic("Se intentó sobrescribir el error original para resultado con error.")
	}

	O.err = err
	return O
}

func (O *Err) Mensaje(tipo MensajeUsuarioTipo, mensaje string) *Err {
	if len(mensaje) == 0 {
		panic("Se intentó establecer mensaje vacío para resultado con error.")
	}

	for i := range O.msj {
		if O.msj[i].Tipo == tipo && O.msj[i].Mensaje == mensaje {
			// Evitando duplicados
			return O
		}
	}

	O.msj = append(O.msj, &MensajeUsuario{
		Tipo:    tipo,
		Mensaje: mensaje,
	})

	return O
}

func transformarError(err error) *Err {
	var E = &Err{}

	var transformado = errors.As(err, &E)
	if !transformado {
		E.err = err
	}

	return E
}
