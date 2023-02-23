package latium

import (
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/latium/orbetello"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/latium/orvieto"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/latium/roma"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LatiumDuke interface {
	feud.Duke
	COrbetello奥尔贝泰洛() orbetello.OrbetelloCounty
	COrvieto奥尔维耶托() orvieto.OrvietoCounty
	CRoma罗马() roma.RomaCounty
}

type 拉丁姆LatiumDuke struct {
	feud.BaseDuke
	奥尔贝泰洛Orbetello orbetello.OrbetelloCounty
	奥尔维耶托Orvieto   orvieto.OrvietoCounty
	罗马Roma         roma.RomaCounty
}

func (d *拉丁姆LatiumDuke) COrbetello奥尔贝泰洛() orbetello.OrbetelloCounty {
	return d.奥尔贝泰洛Orbetello
}

func (d *拉丁姆LatiumDuke) COrvieto奥尔维耶托() orvieto.OrvietoCounty {
	return d.奥尔维耶托Orvieto
}

func (d *拉丁姆LatiumDuke) CRoma罗马() roma.RomaCounty {
	return d.罗马Roma
}

var DLatium拉丁姆 LatiumDuke = &拉丁姆LatiumDuke{}

func init() {
	f := DLatium拉丁姆.(*拉丁姆LatiumDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "latium",
		TitleName: "拉丁姆",
		TitleCode: "d_latium",
		Counties:  map[string]feud.County{},
	}

	f.奥尔贝泰洛Orbetello = orbetello.COrbetello奥尔贝泰洛
	f.奥尔贝泰洛Orbetello.SetParent(f)

	f.奥尔维耶托Orvieto = orvieto.COrvieto奥尔维耶托
	f.奥尔维耶托Orvieto.SetParent(f)

	f.罗马Roma = roma.CRoma罗马
	f.罗马Roma.SetParent(f)

}
