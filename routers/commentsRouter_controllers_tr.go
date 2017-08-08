package routers

import (
	"github.com/astaxie/beego"
)

func init() {

		beego.GlobalControllerRouter["kyronApi/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:PersonaEscalafonController"],
			beego.ControllerComments{
				Method: "GetAllPregrado",
				Router: `pregrado`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["kyronApi/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:PersonaEscalafonController"],
			beego.ControllerComments{
				Method: "GetAllPosgrado",
				Router: `posgrado`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		}

		
