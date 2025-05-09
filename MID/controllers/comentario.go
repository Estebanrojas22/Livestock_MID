package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
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
	fmt.Println("Funcion Get")

	id_ingreso, id_ingreso_2 := c.Ctx.Input.Param(":id"), c.Ctx.Input.Param(":id_2")

	body, _ := services.Metodo_get_one("host_api2", "comentario?query=IdPublicacion:"+ id_ingreso +",IdTPublicacionTipoPublicacion.Id:"+id_ingreso_2)

	body_map, _ := services.ProcessarJson(body)
	fmt.Println("getOne", body_map)
	comentarios:= body_map["Consulta de id"].([]interface{})
	for i, comentario := range comentarios {

		fmt.Println("comentario: ", i, comentario)

		comentario_map, _ := services.ToMap(comentario)

		fmt.Println("comentario id usuario: ", comentario_map["IdUsuario"])
		id_usuario := comentario_map["IdUsuario"].(map[string]interface{})["IdUsuario"]
		fmt.Println("comentario id usuario: ", id_usuario)
		// body_usuario, _ := services.Metodo_get_one("host_api","registro_usuario?query=Id:" + id_usuario)
		// usuario_map, _ := services.ProcessarJson(body_usuario)

		// fmt.Println("usuario: ", usuario_map)
	}

	c.Data["json"] = body_map
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
