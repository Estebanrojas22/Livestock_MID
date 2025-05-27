package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Livestock_MID/MID/services"
)

// CredencialesController operations for Credenciales
type CredencialesController struct {
	beego.Controller
}

// URLMapping ...
func (c *CredencialesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Credenciales
// @Param	body		body 	models.Credenciales	true		"body for Credenciales content"
// @Success 201 {object} models.Credenciales
// @Failure 403 body is empty
// @router / [post]
func (c *CredencialesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Credenciales by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Credenciales
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CredencialesController) GetOne() {
	fmt.Println("Función GetOne")

	id := c.Ctx.Input.Param(":id") // obtiene el ID desde la URL
	if id == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Message": "ID no proporcionado",
		}
		c.ServeJSON()
		return
	}

	// Llamada a servicio para obtener datos
	// endpoint := "credenciales/" + id
	body_response, err := services.Metodo_get_one("host_api1", "credenciales/"+id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Message": "Error al obtener datos",
		}
		c.ServeJSON()
		return
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body_response, &response); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Message": "Error al decodificar respuesta",

		}
		c.ServeJSON()
		return
	}

	// Suponiendo que la contraseña está en el campo "contrasena"
	Contraseña := response["Consulta de id"] // Ajusta según estructura

	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"Message": "Consulta existosa",
		"Data":    Contraseña,
	}
	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description get Credenciales
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Credenciales
// @Failure 403
// @router / [get]
func (c *CredencialesController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Credenciales
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Credenciales	true		"body for Credenciales content"
// @Success 200 {object} models.Credenciales
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CredencialesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Credenciales
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CredencialesController) Delete() {

}
