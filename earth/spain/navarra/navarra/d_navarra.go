package navarra

import (
	"github.com/thalesfu/CK2Commands/earth/spain/navarra/navarra/najera"
	"github.com/thalesfu/CK2Commands/earth/spain/navarra/navarra/navarra"
	"github.com/thalesfu/CK2Commands/earth/spain/navarra/navarra/viscaya"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NavarraDuke interface {
	feud.Duke
	CNajera纳赫拉() najera.NajeraCounty
	CNavarra纳瓦拉() navarra.NavarraCounty
	CViscaya比斯开() viscaya.ViscayaCounty
}

type 纳瓦拉NavarraDuke struct {
	feud.BaseDuke
	纳赫拉Najera  najera.NajeraCounty
	纳瓦拉Navarra navarra.NavarraCounty
	比斯开Viscaya viscaya.ViscayaCounty
}

func (d *纳瓦拉NavarraDuke) CNajera纳赫拉() najera.NajeraCounty {
	return d.纳赫拉Najera
}

func (d *纳瓦拉NavarraDuke) CNavarra纳瓦拉() navarra.NavarraCounty {
	return d.纳瓦拉Navarra
}

func (d *纳瓦拉NavarraDuke) CViscaya比斯开() viscaya.ViscayaCounty {
	return d.比斯开Viscaya
}

var DNavarra纳瓦拉 NavarraDuke = &纳瓦拉NavarraDuke{}

func init() {
	f := DNavarra纳瓦拉.(*纳瓦拉NavarraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "navarra",
		TitleName: "纳瓦拉",
		TitleCode: "d_navarra",
		Counties:  map[string]feud.County{},
	}

	f.纳赫拉Najera = najera.CNajera纳赫拉
	f.纳赫拉Najera.SetParent(f)

	f.纳瓦拉Navarra = navarra.CNavarra纳瓦拉
	f.纳瓦拉Navarra.SetParent(f)

	f.比斯开Viscaya = viscaya.CViscaya比斯开
	f.比斯开Viscaya.SetParent(f)

}
