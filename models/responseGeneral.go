package models

/*ResponsePositive da una respuesta de true si todo va bien*/
type ResponsePositive struct {
	Status  bool   `json:"status" example:"true"`
	Message string `json:"message,omitempty"`
}

/*ResponseFailure da una respuesta de false si ocurrio un error*/
type ResponseFailure struct {
	Status  bool   `json:"status" example:"false"`
	Message string `json:"message" example:"ERROR"`
}
