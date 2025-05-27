package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:ComentarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:ComentarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:ComentarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:ComentarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:ComentarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id/:id_2",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Historial_publicacionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Historial_publicacionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Historial_publicacionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Historial_publicacionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Historial_publicacionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Historial_publicacionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Historial_publicacionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Historial_publicacionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Historial_publicacionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Historial_publicacionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Registro_usuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Registro_usuariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Registro_usuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Registro_usuariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Registro_usuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Registro_usuariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Registro_usuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Registro_usuariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Registro_usuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Livestock_MID/MID/controllers:Registro_usuariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
