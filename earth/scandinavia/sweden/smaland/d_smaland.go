package smaland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/smaland/more"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/smaland/oland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/smaland/sevede"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/smaland/tjust"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SmalandDuke interface {
	feud.Duke
	CMore默雷() more.MoreCounty
	COland厄兰() oland.OlandCounty
	CSevede塞韦德() sevede.SevedeCounty
	CTjust修斯特() tjust.TjustCounty
}

type 斯莫兰SmalandDuke struct {
	feud.BaseDuke
	默雷More    more.MoreCounty
	厄兰Oland   oland.OlandCounty
	塞韦德Sevede sevede.SevedeCounty
	修斯特Tjust  tjust.TjustCounty
}

func (d *斯莫兰SmalandDuke) CMore默雷() more.MoreCounty {
	return d.默雷More
}

func (d *斯莫兰SmalandDuke) COland厄兰() oland.OlandCounty {
	return d.厄兰Oland
}

func (d *斯莫兰SmalandDuke) CSevede塞韦德() sevede.SevedeCounty {
	return d.塞韦德Sevede
}

func (d *斯莫兰SmalandDuke) CTjust修斯特() tjust.TjustCounty {
	return d.修斯特Tjust
}

var DSmaland斯莫兰 SmalandDuke = &斯莫兰SmalandDuke{}

func init() {
	f := DSmaland斯莫兰.(*斯莫兰SmalandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "smaland",
		TitleName: "斯莫兰",
		TitleCode: "d_smaland",
		Counties:  map[string]feud.County{},
	}

	f.默雷More = more.CMore默雷
	f.默雷More.SetParent(f)

	f.厄兰Oland = oland.COland厄兰
	f.厄兰Oland.SetParent(f)

	f.塞韦德Sevede = sevede.CSevede塞韦德
	f.塞韦德Sevede.SetParent(f)

	f.修斯特Tjust = tjust.CTjust修斯特
	f.修斯特Tjust.SetParent(f)

}
