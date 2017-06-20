// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"kyronApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/programa",
			beego.NSInclude(
				&controllers.ProgramaController{},
			),
		),

		beego.NSNamespace("/formacion_academica",
			beego.NSInclude(
				&controllers.FormacionAcademicaController{},
			),
		),

		beego.NSNamespace("/experiencia_docente",
			beego.NSInclude(
				&controllers.ExperienciaDocenteController{},
			),
		),

		beego.NSNamespace("/cursos",
			beego.NSInclude(
				&controllers.CursosController{},
			),
		),

		beego.NSNamespace("/institucion",
			beego.NSInclude(
				&controllers.InstitucionController{},
			),
		),

		beego.NSNamespace("/titulo",
			beego.NSInclude(
				&controllers.TituloController{},
			),
		),

		beego.NSNamespace("/tipo_investigacion",
			beego.NSInclude(
				&controllers.TipoInvestigacionController{},
			),
		),

		beego.NSNamespace("/tipo_dedicacion",
			beego.NSInclude(
				&controllers.TipoDedicacionController{},
			),
		),

		beego.NSNamespace("/investigacion",
			beego.NSInclude(
				&controllers.InvestigacionController{},
			),
		),

		beego.NSNamespace("/nivel_formacion",
			beego.NSInclude(
				&controllers.NivelFormacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
