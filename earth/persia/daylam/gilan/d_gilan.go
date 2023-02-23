package gilan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/gilan/daylam"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/gilan/gilan"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/gilan/qazwin"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GilanDuke interface {
	feud.Duke
	CDaylam德莱木() daylam.DaylamCounty
	CGilan吉兰() gilan.GilanCounty
	CQazwin加兹温() qazwin.QazwinCounty
}

type 吉兰GilanDuke struct {
	feud.BaseDuke
	德莱木Daylam daylam.DaylamCounty
	吉兰Gilan   gilan.GilanCounty
	加兹温Qazwin qazwin.QazwinCounty
}

func (d *吉兰GilanDuke) CDaylam德莱木() daylam.DaylamCounty {
	return d.德莱木Daylam
}

func (d *吉兰GilanDuke) CGilan吉兰() gilan.GilanCounty {
	return d.吉兰Gilan
}

func (d *吉兰GilanDuke) CQazwin加兹温() qazwin.QazwinCounty {
	return d.加兹温Qazwin
}

var DGilan吉兰 GilanDuke = &吉兰GilanDuke{}

func init() {
	f := DGilan吉兰.(*吉兰GilanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gilan",
		TitleName: "吉兰",
		TitleCode: "d_gilan",
		Counties:  map[string]feud.County{},
	}

	f.德莱木Daylam = daylam.CDaylam德莱木
	f.德莱木Daylam.SetParent(f)

	f.吉兰Gilan = gilan.CGilan吉兰
	f.吉兰Gilan.SetParent(f)

	f.加兹温Qazwin = qazwin.CQazwin加兹温
	f.加兹温Qazwin.SetParent(f)

}
