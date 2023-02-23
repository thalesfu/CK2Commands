package munster

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/munster/desmond"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/munster/iarmond"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/munster/ormond"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/munster/thomond"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MunsterDuke interface {
	feud.Duke
	CDesmond德斯蒙德() desmond.DesmondCounty
	CIarmond雅尔蒙德() iarmond.IarmondCounty
	COrmond奥蒙德() ormond.OrmondCounty
	CThomond托蒙德() thomond.ThomondCounty
}

type 芒斯特MunsterDuke struct {
	feud.BaseDuke
	德斯蒙德Desmond desmond.DesmondCounty
	雅尔蒙德Iarmond iarmond.IarmondCounty
	奥蒙德Ormond   ormond.OrmondCounty
	托蒙德Thomond  thomond.ThomondCounty
}

func (d *芒斯特MunsterDuke) CDesmond德斯蒙德() desmond.DesmondCounty {
	return d.德斯蒙德Desmond
}

func (d *芒斯特MunsterDuke) CIarmond雅尔蒙德() iarmond.IarmondCounty {
	return d.雅尔蒙德Iarmond
}

func (d *芒斯特MunsterDuke) COrmond奥蒙德() ormond.OrmondCounty {
	return d.奥蒙德Ormond
}

func (d *芒斯特MunsterDuke) CThomond托蒙德() thomond.ThomondCounty {
	return d.托蒙德Thomond
}

var DMunster芒斯特 MunsterDuke = &芒斯特MunsterDuke{}

func init() {
	f := DMunster芒斯特.(*芒斯特MunsterDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "munster",
		TitleName: "芒斯特",
		TitleCode: "d_munster",
		Counties:  map[string]feud.County{},
	}

	f.德斯蒙德Desmond = desmond.CDesmond德斯蒙德
	f.德斯蒙德Desmond.SetParent(f)

	f.雅尔蒙德Iarmond = iarmond.CIarmond雅尔蒙德
	f.雅尔蒙德Iarmond.SetParent(f)

	f.奥蒙德Ormond = ormond.COrmond奥蒙德
	f.奥蒙德Ormond.SetParent(f)

	f.托蒙德Thomond = thomond.CThomond托蒙德
	f.托蒙德Thomond.SetParent(f)

}
