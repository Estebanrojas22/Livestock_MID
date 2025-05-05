package controllers

import (
	"fmt"

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


}

// GetOne ...
// @Title GetOne
// @Description get Registro_usuarios by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Registro_usuarios
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Registro_usuariosController) GetOne() {
	fmt.Println("Funcion Get")
	// 	id_contraseña := c.Ctx.Input.Param(":id")
	// 	fmt.Println("EL id de ingreso es:", id_contraseña)

	// 	body, _ := services.Metodo_get("servicio_registro", id_contraseña)
	// 	fmt.Println("EL id de registro es:", body)
	// 	var result map[string]interface{}
	// 	err := json.Unmarshal(body, &result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	id_contraseña = result["Contraseña"].(map[string]interface{})["Id"].(string)

	// fmt.Println("EL id de ingreso es:", result)
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

	for i := range arreglo_map_usuario{

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
			"Activo":            arreglo_map_usuario[i]["Celular"],
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
