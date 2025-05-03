func (c *UsuarioController) Post() {
	fmt.Println("post")

	var body_ingresa map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err == nil {
		fmt.Println("json ingresa", body_ingresa)
		fmt.Println("error", err)
	}

	body_contrasena:= map[string]interface{}{
		"Contrasena":     body_ingresa["Contraseña"],

	}

	bytes_contrasena, err := json.Marshal(body_contrasena)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	body_response_contrasena, _:= services.Metodo_post("hots_crud", "Credenciales", bytes_contrasena)

	var response_json_contrasena map[string]interface{}

	err1 := json.Unmarshal(body_response_contrasena, &response_json_contrasena)
	if err1 != nil {
		fmt.Println("Error al deserializar:", err)
		return
	}

	Id_contrasena:= response_json_contrasena["Data"]
	Id_contrasena= Id_contrasena.(map[string]interface{})["Id"]

	fmt.Println("este es el id", Id_contrasena)

	fmt.Println("tipo dato", reflect.TypeOf(Id_contrasena))

	Id_entero := fmt.Sprintf("%v", Id_contrasena)

	fmt.Println("variable en string", Id_entero)

	Id, _ := strconv.Atoi(Id_entero)

	body_Usuario:= map[string]interface{}{
		"Nombres":     body_ingresa["Nombres"],
		"Apellido": body_ingresa["Apellido"],
		"Email": body_ingresa["Email"],
		"IdCredenciales": map[string]interface{}{"Id": Id},

	}

	bytes_usuario, err := json.Marshal(body_Usuario)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	body_response_usuario, _:= services.Metodo_post("hots_crud", "Usuario", bytes_usuario)

	var response_json_usuario map[string]interface{}

	err2 := json.Unmarshal(body_response_usuario, &response_json_usuario)
	if err2 != nil {
		fmt.Println("Error al deserializar:", err)
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Succes":          true,
		"Status":          200,
		"Message":         "Creación existosa",
		"Data":            body_ingresa,
		
	}
	c.ServeJSON()
	

	// 5555

}