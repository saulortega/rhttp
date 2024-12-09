package rhttp

type MensajeUsuarioTipo string

const (
	MensajeUsuarioTipoInformacion MensajeUsuarioTipo = "informacion" // Azul
	MensajeUsuarioTipoAviso       MensajeUsuarioTipo = "aviso"       // Amarillo
	MensajeUsuarioTipoAdvertencia MensajeUsuarioTipo = "advertencia" // Naranja
	MensajeUsuarioTipoError       MensajeUsuarioTipo = "error"       // Rojo
)

type MensajeUsuario struct {
	Tipo    MensajeUsuarioTipo `json:"tipo"`
	Mensaje string             `json:"mensaje"`
}
