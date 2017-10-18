package controllers

import (
    "github.com/udistrital/hvapi/models"
  "github.com/astaxie/beego"
)

type PersonaEscalafonController struct {
	beego.Controller
}

func (c *PersonaEscalafonController) URLMapping() {
	c.Mapping("GetAllPregrado", c.GetAllPregrado)
	c.Mapping("GetAllPosgrado", c.GetAllPosgrado)
}

func (c *PersonaEscalafonController) GetAllPregrado() {
    listaPersonas := models.GetAllPersonaEscalafonPregrado()
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaPersonas
    c.ServeJSON()
}

func (c *PersonaEscalafonController) GetAllPosgrado() {
    listaPersonas := models.GetAllPersonaEscalafonPosgrado()
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaPersonas
    c.ServeJSON()
}