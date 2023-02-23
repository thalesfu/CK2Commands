package sjaelland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/sjaelland/fyn"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/sjaelland/lolland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/sjaelland/sjaelland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SjaellandDuke interface {
	feud.Duke
	CFyn菲英() fyn.FynCounty
	CLolland洛兰() lolland.LollandCounty
	CSjaelland西兰() sjaelland.SjaellandCounty
}

type 西兰SjaellandDuke struct {
	feud.BaseDuke
	菲英Fyn       fyn.FynCounty
	洛兰Lolland   lolland.LollandCounty
	西兰Sjaelland sjaelland.SjaellandCounty
}

func (d *西兰SjaellandDuke) CFyn菲英() fyn.FynCounty {
	return d.菲英Fyn
}

func (d *西兰SjaellandDuke) CLolland洛兰() lolland.LollandCounty {
	return d.洛兰Lolland
}

func (d *西兰SjaellandDuke) CSjaelland西兰() sjaelland.SjaellandCounty {
	return d.西兰Sjaelland
}

var DSjaelland西兰 SjaellandDuke = &西兰SjaellandDuke{}

func init() {
	f := DSjaelland西兰.(*西兰SjaellandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sjaelland",
		TitleName: "西兰",
		TitleCode: "d_sjaelland",
		Counties:  map[string]feud.County{},
	}

	f.菲英Fyn = fyn.CFyn菲英
	f.菲英Fyn.SetParent(f)

	f.洛兰Lolland = lolland.CLolland洛兰
	f.洛兰Lolland.SetParent(f)

	f.西兰Sjaelland = sjaelland.CSjaelland西兰
	f.西兰Sjaelland.SetParent(f)

}
