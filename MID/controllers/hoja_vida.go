package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Livestock_MID/MID/services"
)

// Hoja_vidaController operations for Hoja_vida
type Hoja_vidaController struct {
	beego.Controller
}

// URLMapping ...
func (c *Hoja_vidaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Hoja_vida
// @Param	body		body 	models.Hoja_vida	true		"body for Hoja_vida content"
// @Success 201 {object} models.Hoja_vida
// @Failure 403 body is empty
// @router / [post]
func (c *Hoja_vidaController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Hoja_vida by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Hoja_vida
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Hoja_vidaController) GetOne() {
	fmt.Println("Funcion Get")
	id_contraseña := c.Ctx.Input.Param(":id")
	fmt.Println("EL id de ingreso es:", id_contraseña)

	body, _ := services.Metodo_get("servicio_registro", id_contraseña)
	fmt.Println("EL id de registro es:", body)
	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}
	id_contraseña = result["Contraseña"].(map[string]interface{})["Id"].(string)

	fmt.Println("EL id de ingreso es:", result)

}

// GetAll ...
// @Title GetAll
// @Description get Hoja_vida
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Hoja_vida
// @Failure 403
// @router / [get]
func (c *Hoja_vidaController) GetAll() {
	fmt.Println("get registro")

}

// Put ...
// @Title Put
// @Description update the Hoja_vida
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Hoja_vida	true		"body for Hoja_vida content"
// @Success 200 {object} models.Hoja_vida
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Hoja_vidaController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Hoja_vida
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Hoja_vidaController) Delete() {

}
