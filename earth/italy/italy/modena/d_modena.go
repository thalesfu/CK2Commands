package modena

import (
	"github.com/thalesfu/CK2Commands/earth/italy/italy/modena/modena"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/modena/parma"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ModenaDuke interface {
	feud.Duke
	CModena摩德纳() modena.ModenaCounty
	CParma帕尔马() parma.ParmaCounty
}

type 摩德纳ModenaDuke struct {
	feud.BaseDuke
	摩德纳Modena modena.ModenaCounty
	帕尔马Parma  parma.ParmaCounty
}

func (d *摩德纳ModenaDuke) CModena摩德纳() modena.ModenaCounty {
	return d.摩德纳Modena
}

func (d *摩德纳ModenaDuke) CParma帕尔马() parma.ParmaCounty {
	return d.帕尔马Parma
}

var DModena摩德纳 ModenaDuke = &摩德纳ModenaDuke{}

func init() {
	f := DModena摩德纳.(*摩德纳ModenaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "modena",
		TitleName: "摩德纳",
		TitleCode: "d_modena",
		Counties:  map[string]feud.County{},
	}

	f.摩德纳Modena = modena.CModena摩德纳
	f.摩德纳Modena.SetParent(f)

	f.帕尔马Parma = parma.CParma帕尔马
	f.帕尔马Parma.SetParent(f)

}
