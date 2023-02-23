package anjou

import (
	"github.com/thalesfu/CK2Commands/earth/france/france/anjou/anjou"
	"github.com/thalesfu/CK2Commands/earth/france/france/anjou/maine"
	"github.com/thalesfu/CK2Commands/earth/france/france/anjou/perche"
	"github.com/thalesfu/CK2Commands/earth/france/france/anjou/saumur"
	"github.com/thalesfu/CK2Commands/earth/france/france/anjou/vendome"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnjouDuke interface {
	feud.Duke
	CAnjou安茹() anjou.AnjouCounty
	CMaine曼恩() maine.MaineCounty
	CPerche佩尔什() perche.PercheCounty
	CSaumur索米尔() saumur.SaumurCounty
	CVendome旺多姆() vendome.VendomeCounty
}

type 安茹AnjouDuke struct {
	feud.BaseDuke
	安茹Anjou    anjou.AnjouCounty
	曼恩Maine    maine.MaineCounty
	佩尔什Perche  perche.PercheCounty
	索米尔Saumur  saumur.SaumurCounty
	旺多姆Vendome vendome.VendomeCounty
}

func (d *安茹AnjouDuke) CAnjou安茹() anjou.AnjouCounty {
	return d.安茹Anjou
}

func (d *安茹AnjouDuke) CMaine曼恩() maine.MaineCounty {
	return d.曼恩Maine
}

func (d *安茹AnjouDuke) CPerche佩尔什() perche.PercheCounty {
	return d.佩尔什Perche
}

func (d *安茹AnjouDuke) CSaumur索米尔() saumur.SaumurCounty {
	return d.索米尔Saumur
}

func (d *安茹AnjouDuke) CVendome旺多姆() vendome.VendomeCounty {
	return d.旺多姆Vendome
}

var DAnjou安茹 AnjouDuke = &安茹AnjouDuke{}

func init() {
	f := DAnjou安茹.(*安茹AnjouDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "anjou",
		TitleName: "安茹",
		TitleCode: "d_anjou",
		Counties:  map[string]feud.County{},
	}

	f.安茹Anjou = anjou.CAnjou安茹
	f.安茹Anjou.SetParent(f)

	f.曼恩Maine = maine.CMaine曼恩
	f.曼恩Maine.SetParent(f)

	f.佩尔什Perche = perche.CPerche佩尔什
	f.佩尔什Perche.SetParent(f)

	f.索米尔Saumur = saumur.CSaumur索米尔
	f.索米尔Saumur.SetParent(f)

	f.旺多姆Vendome = vendome.CVendome旺多姆
	f.旺多姆Vendome.SetParent(f)

}
