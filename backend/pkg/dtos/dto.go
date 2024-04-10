package dtos

import "backend/pkg/utils"

type RegistroStruct struct {
	Correo     string `json:"correo" validate:"required"`
	Usuario    string `json:"usuario" validate:"required"`
	Contrasena string `json:"contrasena" validate:"required"`
	Telefono   string `json:"telefono" validate:"required"`
}

type LoginStruct struct {
	UsuarioOCorreo string `json:"usuarioocorreo" validate:"required"`
	Contrasena     string `json:"contrasena" validate:"required"`
}

type ErrorStruct struct {
	Error string `json:"error"`
}

type LoginResponse struct {
	Error string `json:"error"`
	Token string `json:"token"`
}

func (l LoginStruct) Validate() error {
	err := utils.ValidateStruct(l)
	if err != nil {
		return err
	}
	return nil
}

func (r RegistroStruct) Validate() error {
	err := utils.ValidateStruct(r)
	if err != nil {
		return err
	}

	err = utils.ValidateEmail(r.Correo)
	if err != nil {
		return err
	}
	err = utils.ValidatePhone(r.Telefono)
	if err != nil {
		return err
	}

	err = utils.ValidatePassword(r.Contrasena)
	if err != nil {
		return err
	}
	return nil
}
