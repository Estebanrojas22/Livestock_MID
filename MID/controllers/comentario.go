package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/fatih/color"
	"github.com/sena_2824182/Livestock_MID/MID/services"
)

// ComentarioController operations for Comentario
type ComentarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *ComentarioController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Comentario
// @Param	body		body 	models.Comentario	true		"body for Comentario content"
// @Success 201 {object} models.Comentario
// @Failure 403 body is empty
// @router / [post]
func (c *ComentarioController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Comentario by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comentario
// @Failure 403 :id is empty
// @router /:id/:id_2 [get]
func (c *ComentarioController) GetOne() {
	var resultado_final []map[string]interface{}
	verde := color.New(color.FgGreen).SprintFunc()
	//fmt.Println(verde("Funcion Get"))

	id_ingreso, id_ingreso_2 := c.Ctx.Input.Param(":id"), c.Ctx.Input.Param(":id_2")

	body, _ := services.Metodo_get_one("host_api2", "comentario?query=IdPublicacion:"+id_ingreso+",IdTPublicacionTipoPublicacion.Id:"+id_ingreso_2)

	body_map, _ := services.ProcessarJson(body)
	//fmt.Println(verde("getOne"), body_map)
	comentarios := body_map["Consulta de id"].([]interface{})
	fmt.Println(verde("comentarios: "), comentarios)
	comentarios_map2, _ := services.ToMapStringInterface(comentarios[0])
	resultado_final = append(resultado_final, map[string]interface{}{
		"id_ Publicacion":    comentarios_map2["IdPublicacion"],
		"nombre_Publicacion": comentarios_map2["IdTPublicacionTipoPublicacion"].(map[string]interface{})["NombrePublicacion"],
	})
	for _, comentario := range comentarios {

		comentario_map, _ := services.ToMapStringInterface(comentario)

		fmt.Println(verde("comentario usuario: "), comentario_map)
		id_usuario := comentario_map["IdUsuario"]

		fmt.Println(verde("comentario id usuario 1:"), id_usuario)
		id_usuario_string := fmt.Sprintf("%v", id_usuario)
		url := "registro_usuario/" + id_usuario_string
		body_usuario, _ := services.Metodo_get_one("host_api", url)
		body_usuario_map, _ := services.ProcessarJson(body_usuario)

		//fmt.Println(verde("usuario: "), body_usuario_map)

		usuario := body_usuario_map["Consulta de id"]
		nombre_usuario := usuario.(map[string]interface{})["Nombre"]
		apellido_usuario := usuario.(map[string]interface{})["Apellido"]
		fmt.Println(verde("Nombre de usuario: "), nombre_usuario, apellido_usuario)
		nombre_completo := fmt.Sprintf("%v %v", nombre_usuario, apellido_usuario)

		resultado_final = append(resultado_final, map[string]interface{}{
			"comentario": comentario_map["Comentarios"],
			"usuario":    nombre_completo,
		})
	}
	c.Data["json"] = resultado_final
	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description get Comentario
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Comentario
// @Failure 403
// @router / [get]
func (c *ComentarioController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Comentario
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Comentario	true		"body for Comentario content"
// @Success 200 {object} models.Comentario
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ComentarioController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Comentario
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ComentarioController) Delete() {

}
