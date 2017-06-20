package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:CursosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ExperienciaDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:InvestigacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:NivelFormacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:ProgramaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoDedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TipoInvestigacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"] = append(beego.GlobalControllerRouter["hojasdevida/controllers:TituloController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
