package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Livestock_MID/MID/services"
)

// Registro_usuariosController operations for Registro_usuarios
type Registro_usuariosController struct {
	beego.Controller
}

// URLMapping ...
func (c *Registro_usuariosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Registro_usuarios
// @Param	body		body 	models.Registro_usuarios	true		"body for Registro_usuarios content"
// @Success 201 {object} models.Registro_usuarios
// @Failure 403 body is empty
// @router / [post]
func (c *Registro_usuariosController) Post() {

	fmt.Println("Funcion Post")

	var body_ingresa map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err == nil {
		fmt.Println("json ingresa", body_ingresa)
	}

	body_contrasena := map[string]interface{}{
		"Contraseña": body_ingresa["contrasena"],
	}

	bytes_contrasena, err := json.Marshal(body_contrasena)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	body_response_contrasena_byte, _ := services.Metodo_post("host_api", "credenciales", bytes_contrasena)

	var response_json_contrasena map[string]interface{}

	err1 := json.Unmarshal(body_response_contrasena_byte, &response_json_contrasena)
	if err1 != nil {
		fmt.Println("Error al deserializar:", err)
		return
	}
	diastring, _ := json.Marshal(body_ingresa["dia"])
	mestring, _ := json.Marshal(body_ingresa["mes"])
	aniotring, _ := json.Marshal(body_ingresa["anio"])

	var fecha_nacimiento = string(diastring) + "/" + string(mestring) + "/" + string(aniotring)

	fmt.Println("fecha", fecha_nacimiento)

	id_contrasena := response_json_contrasena["Datos creados con id"].(map[string]interface{})["Id"]
	id_tipo_documento := body_ingresa["tipoDocumento"]
	id_tipo_documento_string := fmt.Sprintf("%v", id_tipo_documento)
	id_tipo_documento_int, _ := strconv.Atoi(id_tipo_documento_string)
	id_tipo_usuario := body_ingresa["tipoUsuario"]
	id_tipo_usuario_string := fmt.Sprintf("%v", id_tipo_usuario)
	id_tipo_usuario_int, _ := strconv.Atoi(id_tipo_usuario_string)

	fmt.Println("id contraseña", id_contrasena)

	body_Usuario := map[string]interface{}{
		"Nombre":            body_ingresa["nombre"],
		"Apellido":          body_ingresa["apellido"],
		"FNacimiento":       fecha_nacimiento,
		"NDocumento":        body_ingresa["numeroDocumento"],
		"CorreoElectronico": body_ingresa["correo_electronico"],
		"Contrasena":        map[string]interface{}{"id": id_contrasena},
		"IdTipoDocumento":   map[string]interface{}{"id": id_tipo_documento_int},
		"IdTipoUsuario":     map[string]interface{}{"id": id_tipo_usuario_int},
	}
	bytes_usuario, err := json.Marshal(body_Usuario)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	body_response_usuario_byte, _ := services.Metodo_post("host_api", "registro_usuario", bytes_usuario)

	var response_json_usuario map[string]interface{}

	fmt.Println("body_response_usuario_byte", string(body_response_usuario_byte))

	err2 := json.Unmarshal(body_response_usuario_byte, &response_json_usuario)
	if err2 != nil {
		fmt.Println("Error al deserializar:", err)
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"Message": "Creación existosa",
		"Data":    body_ingresa,
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Registro_usuarios by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Registro_usuarios
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Registro_usuariosController) GetOne() {
	
}

// GetAll ...
// @Title GetAll
// @Description get Registro_usuarios
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Registro_usuarios
// @Failure 403
// @router / [get]
func (c *Registro_usuariosController) GetAll() {
	fmt.Println("get all")
	Json_registro, _ := services.Metodo_get_all("host_api", "registro_usuario")
	fmt.Println("Este es el valor de Json registro en byte:", Json_registro)

	Json_procesado_registro, _ := services.ProcessarJson(Json_registro)
	// fmt.Println("Este es el valor de Json registro en Json:", Json_procesado_registro)

	usuario_Json := Json_procesado_registro["Consulta de id"]

	arreglo_map_usuario, _ := services.ConvertInterfaceToSliceMap(usuario_Json)
	// fmt.Println("Estos son los usuarios map:", arreglo_map_usuario)

	var resultado []map[string]interface{}

	for i := range arreglo_map_usuario {

		fmt.Println("Valor de json solo:", arreglo_map_usuario[i])

		resultado_parcial := map[string]interface{}{
			"Nombre":            arreglo_map_usuario[i]["Nombre"],
			"Apellido":          arreglo_map_usuario[i]["Apellido"],
			"FNacimiento":       arreglo_map_usuario[i]["FNacimiento"],
			"TipoDocumento":     arreglo_map_usuario[i]["IdTipoDocumento"].(map[string]interface{})["Nombre"],
			"NDocumento":        arreglo_map_usuario[i]["NDocumento"],
			"Edad":              arreglo_map_usuario[i]["Edad"],
			"CorreoElectronico": arreglo_map_usuario[i]["CorreoElectronico"],
			"Celular":           arreglo_map_usuario[i],
			"TipoUsuario":       arreglo_map_usuario[i]["IdTipoUsuario"].(map[string]interface{})["Nombre"],
			"Contraseña":        arreglo_map_usuario[i]["Contraseña"].(map[string]interface{})["contraseña"],
			"Activo":            arreglo_map_usuario[i]["Activo"],
		}

		resultado = append(resultado, resultado_parcial)
	}

	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"Message": "Consulta existosa",
		"Data":    resultado,
		// "Cantidad paises": len(resultado),
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Registro_usuarios
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Registro_usuarios	true		"body for Registro_usuarios content"
// @Success 200 {object} models.Registro_usuarios
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Registro_usuariosController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Registro_usuarios
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Registro_usuariosController) Delete() {

}
