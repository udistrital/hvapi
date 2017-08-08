package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type PersonaEscalafon struct {
    Id int `orm:"column(id);pk"`
    PrimerNombre string `orm:"column(primer_nombre)"`
    SegundoNombre string `orm:"column(segundo_nombre)"`
    PrimerApellido string `orm:"column(primer_apellido)"`
    SegundoApellido string `orm:"column(segundo_apellido)"`
    Escalafon string `orm:"column(escalafon)"`
}

func init() {
	orm.RegisterModel(new(PersonaEscalafon))
}

func GetAllPersonaEscalafonPregrado() (arregloIDs []PersonaEscalafon) {
	o := orm.NewOrm()
	var temp []PersonaEscalafon
	_, err := o.Raw("SELECT p.num_documento_persona id, primer_nombre, segundo_nombre, primer_apellido, segundo_apellido, nombre_categoria escalafon FROM agora.informacion_persona_natural p, kyron.categoria_persona ep, kyron.tipo_categoria e, agora.informacion_proveedor ip WHERE p.num_documento_persona=ep.persona_id AND ep.id_tipo_categoria=e.id AND ip.num_documento=p.num_documento_persona;").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}

func GetAllPersonaEscalafonPosgrado() (arregloIDs []PersonaEscalafon) {
	o := orm.NewOrm()
	var temp []PersonaEscalafon
	_, err := o.Raw("SELECT p.num_documento_persona id, primer_nombre, segundo_nombre, primer_apellido, segundo_apellido, nombre_categoria escalafon FROM agora.informacion_persona_natural p, kyron.categoria_persona ep, kyron.tipo_categoria e WHERE p.num_documento_persona=ep.persona_id AND ep.id_tipo_categoria=e.id AND e.nombre_categoria not like 'Auxiliar';").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}