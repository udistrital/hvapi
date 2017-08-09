package routers

import (
	"github.com/astaxie/beego"
)

func init() {

		beego.GlobalControllerRouter["hojasdevida/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:PersonaEscalafonController"],
			beego.ControllerComments{
				Method: "GetAllPregrado",
				Router: `pregrado`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["hojasdevida/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:PersonaEscalafonController"],
			beego.ControllerComments{
				Method: "GetAllPosgrado",
				Router: `posgrado`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		}

		
