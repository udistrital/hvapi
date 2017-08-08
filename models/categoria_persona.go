package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CategoriaPersona struct {
	Validacion      bool           `orm:"column(validacion);null"`
	PersonaId       int            `orm:"column(persona_id)"`
	IdTipoCategoria *TipoCategoria `orm:"column(id_tipo_categoria);rel(fk)"`
	Vigente         bool           `orm:"column(vigente);null"`
	FechaDato       time.Time      `orm:"column(fecha_dato);type(date);null"`
	Id              int            `orm:"column(id);pk"`
}

func (t *CategoriaPersona) TableName() string {
	return "categoria_persona"
}

func init() {
	orm.RegisterModel(new(CategoriaPersona))
}

// AddCategoriaPersona insert a new CategoriaPersona into database and returns
// last inserted Id on success.
func AddCategoriaPersona(m *CategoriaPersona) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCategoriaPersonaById retrieves CategoriaPersona by Id. Returns error if
// Id doesn't exist
func GetCategoriaPersonaById(id int) (v *CategoriaPersona, err error) {
	o := orm.NewOrm()
	v = &CategoriaPersona{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCategoriaPersona retrieves all CategoriaPersona matches certain condition. Returns empty list if
// no records exist
func GetAllCategoriaPersona(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CategoriaPersona)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []CategoriaPersona
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateCategoriaPersona updates CategoriaPersona by Id and returns error if
// the record to be updated doesn't exist
func UpdateCategoriaPersonaById(m *CategoriaPersona) (err error) {
	o := orm.NewOrm()
	v := CategoriaPersona{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCategoriaPersona deletes CategoriaPersona by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCategoriaPersona(id int) (err error) {
	o := orm.NewOrm()
	v := CategoriaPersona{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CategoriaPersona{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
