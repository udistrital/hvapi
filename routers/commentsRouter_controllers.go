package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["kyronApi/controllers:CursosController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:CursosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:CursosController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:CursosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:CursosController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:CursosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:CursosController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:CursosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:CursosController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:CursosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TituloController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TituloController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TituloController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TituloController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TituloController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TituloController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["kyronApi/controllers:TituloController"] = append(beego.GlobalControllerRouter["kyronApi/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
