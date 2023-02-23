package sevilla

import (
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/sevilla/algeciras"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/sevilla/aracena"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/sevilla/cadiz"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/sevilla/huelva"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/sevilla/niebla"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/sevilla/sevilla"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SevillaDuke interface {
	feud.Duke
	CAlgeciras阿尔赫西拉斯() algeciras.AlgecirasCounty
	CAracena阿拉塞纳() aracena.AracenaCounty
	CCadiz加的斯() cadiz.CadizCounty
	CHuelvac_huelva() huelva.HuelvaCounty
	CNiebla涅夫拉() niebla.NieblaCounty
	CSevilla塞维利亚() sevilla.SevillaCounty
}

type 塞维利亚SevillaDuke struct {
	feud.BaseDuke
	阿尔赫西拉斯Algeciras algeciras.AlgecirasCounty
	阿拉塞纳Aracena     aracena.AracenaCounty
	加的斯Cadiz        cadiz.CadizCounty
	c_huelvaHuelva  huelva.HuelvaCounty
	涅夫拉Niebla       niebla.NieblaCounty
	塞维利亚Sevilla     sevilla.SevillaCounty
}

func (d *塞维利亚SevillaDuke) CAlgeciras阿尔赫西拉斯() algeciras.AlgecirasCounty {
	return d.阿尔赫西拉斯Algeciras
}

func (d *塞维利亚SevillaDuke) CAracena阿拉塞纳() aracena.AracenaCounty {
	return d.阿拉塞纳Aracena
}

func (d *塞维利亚SevillaDuke) CCadiz加的斯() cadiz.CadizCounty {
	return d.加的斯Cadiz
}

func (d *塞维利亚SevillaDuke) CHuelvac_huelva() huelva.HuelvaCounty {
	return d.c_huelvaHuelva
}

func (d *塞维利亚SevillaDuke) CNiebla涅夫拉() niebla.NieblaCounty {
	return d.涅夫拉Niebla
}

func (d *塞维利亚SevillaDuke) CSevilla塞维利亚() sevilla.SevillaCounty {
	return d.塞维利亚Sevilla
}

var DSevilla塞维利亚 SevillaDuke = &塞维利亚SevillaDuke{}

func init() {
	f := DSevilla塞维利亚.(*塞维利亚SevillaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sevilla",
		TitleName: "塞维利亚",
		TitleCode: "d_sevilla",
		Counties:  map[string]feud.County{},
	}

	f.阿尔赫西拉斯Algeciras = algeciras.CAlgeciras阿尔赫西拉斯
	f.阿尔赫西拉斯Algeciras.SetParent(f)

	f.阿拉塞纳Aracena = aracena.CAracena阿拉塞纳
	f.阿拉塞纳Aracena.SetParent(f)

	f.加的斯Cadiz = cadiz.CCadiz加的斯
	f.加的斯Cadiz.SetParent(f)

	f.c_huelvaHuelva = huelva.CHuelvac_huelva
	f.c_huelvaHuelva.SetParent(f)

	f.涅夫拉Niebla = niebla.CNiebla涅夫拉
	f.涅夫拉Niebla.SetParent(f)

	f.塞维利亚Sevilla = sevilla.CSevilla塞维利亚
	f.塞维利亚Sevilla.SetParent(f)

}
