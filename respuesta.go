package rhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type respuesta struct {
	Datos    any               `json:"datos,omitempty"`
	Mensajes []*MensajeUsuario `json:"mensajes,omitempty"`
}

func ManejarRespuesta(w http.ResponseWriter, r *http.Request, f func(r *http.Request) Resultado) {
	exitoso, err := f(r).Res()
	if exitoso == nil && err == nil {
		panic("No hay resultado exitoso ni error.")
	} else if exitoso != nil && err != nil {
		panic("No debe haber resultado exitoso y error a la vez.")
	}

	if err != nil {
		manejarRespuestaConError(w, err)
		return
	}

	manejarRespuestaExitosa(w, exitoso)
}

func ManejarRespuestaConError(w http.ResponseWriter, E *Err) {
	manejarRespuestaConError(w, E)
}

func manejarRespuestaConError(w http.ResponseWriter, E *Err) {
	estadoHTTP := asegurarCodigoDeEstadoHTTP(http.StatusInternalServerError, E.cdg)
	if estadoHTTP < 400 || estadoHTTP > 599 {
		panic("No debe haber un código diferente a 4XX o 5XX si hay error.")
	}

	rta := respuesta{
		Mensajes: E.msj,
	}

	for i := range E.msj {
		if E.err == nil {
			fmt.Printf("[%s] %s\n", E.msj[i].Tipo, E.msj[i].Mensaje)
		} else {
			fmt.Printf("[%s] %s: %s\n", E.msj[i].Tipo, E.msj[i].Mensaje, E.err)
		}
	}

	responder(w, estadoHTTP, &rta)
}

func manejarRespuestaExitosa(w http.ResponseWriter, exitoso *ResultadoExitoso) {
	estadoHTTP := asegurarCodigoDeEstadoHTTP(http.StatusOK, exitoso.estadoHTTP)
	if estadoHTTP < 200 || estadoHTTP > 299 {
		panic("No debe haber un código diferente a 2XX si no hay error.")
	}

	rta := respuesta{
		Datos:    exitoso.datos,
		Mensajes: exitoso.mensajes,
	}

	responder(w, estadoHTTP, &rta)
}

func responder(w http.ResponseWriter, estadoHTTP int, rta *respuesta) {
	w.WriteHeader(estadoHTTP)

	if rta.Datos != nil || len(rta.Mensajes) > 0 {
		rtaJSON, er := json.Marshal(rta)
		if er != nil {
			// La respuesta ya se empezó a transmitir en el momento de enviar el código de estado HTTP.
			// Por lo tanto, no hay mucho que hacer, excepto registrar el error.
			fmt.Println("rhttp: no se pudo convertir los datos de respuesta a formato JSON:", er.Error())
			return
		}

		_, er = w.Write(rtaJSON)
		if er != nil {
			// La respuesta ya se empezó a transmitir en el momento de enviar el código de estado HTTP.
			// Por lo tanto, no hay mucho que hacer, excepto registrar el error.
			fmt.Println("rhttp: se detectó un error al enviar los datos al cliente:", er.Error())
			return
		}
	}
}

func asegurarCodigoDeEstadoHTTP(predeterminado int, recibido int) int {
	if recibido == 0 {
		return predeterminado
	}

	return recibido
}
