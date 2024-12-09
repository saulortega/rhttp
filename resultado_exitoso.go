package rhttp

type ResultadoExitoso struct {
	datos      any
	estadoHTTP int
	mensajes   []*MensajeUsuario
}

func (O *ResultadoExitoso) Res() (*ResultadoExitoso, *Err) {
	return O, nil
}

func Datos(datos any) *ResultadoExitoso {
	if datos == nil {
		panic("Se intentó establecer datos en nil para resultado exitoso.")
	}

	return &ResultadoExitoso{
		datos: datos,
	}
}

func EstadoHTTP(estado int) *ResultadoExitoso {
	if estado < 200 || estado > 299 {
		panic("Se intentó establecer código de estado HTTP inesperado para resultado exitoso.")
	}

	return &ResultadoExitoso{
		estadoHTTP: estado,
	}
}

func Mensaje(tipo MensajeUsuarioTipo, mensaje string) *ResultadoExitoso {
	if tipo == MensajeUsuarioTipoError {
		panic("Se intentó establecer tipo de mensaje de error para resultado exitoso.")
	} else if len(mensaje) == 0 {
		panic("Se intentó establecer mensaje vacío para resultado exitoso.")
	}

	return &ResultadoExitoso{
		mensajes: []*MensajeUsuario{
			{
				Tipo:    tipo,
				Mensaje: mensaje,
			},
		},
	}
}

func (O *ResultadoExitoso) Datos(datos any) *ResultadoExitoso {
	if datos == nil {
		panic("Se intentó establecer datos en nil para resultado exitoso.")
	} else if O.datos != nil {
		panic("Se intentó sobrescribir los datos para resultado exitoso.")
	}

	return &ResultadoExitoso{
		datos: datos,
	}
}

func (O *ResultadoExitoso) EstadoHTTP(estado int) *ResultadoExitoso {
	if estado < 200 || estado > 299 {
		panic("Se intentó establecer código de estado HTTP inesperado para resultado exitoso.")
	} else if O.estadoHTTP > 0 {
		panic("Se intentó sobrescribir el código de estado HTTP para resultado exitoso.")
	}

	return &ResultadoExitoso{
		estadoHTTP: estado,
	}
}

func (O *ResultadoExitoso) Mensaje(tipo MensajeUsuarioTipo, mensaje string) *ResultadoExitoso {
	if tipo == MensajeUsuarioTipoError {
		panic("Se intentó establecer tipo de mensaje de error para resultado exitoso.")
	} else if len(mensaje) == 0 {
		panic("Se intentó establecer mensaje vacío para resultado exitoso.")
	}

	for i := range O.mensajes {
		if O.mensajes[i].Tipo == tipo && O.mensajes[i].Mensaje == mensaje {
			// Evitando duplicados
			return O
		}
	}

	O.mensajes = append(O.mensajes, &MensajeUsuario{
		Tipo:    tipo,
		Mensaje: mensaje,
	})

	return O
}
