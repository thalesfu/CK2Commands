package nubia

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/nubia/atbara"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/nubia/makuria"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/nubia/napata"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NubiaDuke interface {
	feud.Duke
	CAtbara阿特巴拉() atbara.AtbaraCounty
	CMakuria马库里亚() makuria.MakuriaCounty
	CNapata纳帕塔() napata.NapataCounty
}

type 马库里亚NubiaDuke struct {
	feud.BaseDuke
	阿特巴拉Atbara  atbara.AtbaraCounty
	马库里亚Makuria makuria.MakuriaCounty
	纳帕塔Napata   napata.NapataCounty
}

func (d *马库里亚NubiaDuke) CAtbara阿特巴拉() atbara.AtbaraCounty {
	return d.阿特巴拉Atbara
}

func (d *马库里亚NubiaDuke) CMakuria马库里亚() makuria.MakuriaCounty {
	return d.马库里亚Makuria
}

func (d *马库里亚NubiaDuke) CNapata纳帕塔() napata.NapataCounty {
	return d.纳帕塔Napata
}

var DNubia马库里亚 NubiaDuke = &马库里亚NubiaDuke{}

func init() {
	f := DNubia马库里亚.(*马库里亚NubiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nubia",
		TitleName: "马库里亚",
		TitleCode: "d_nubia",
		Counties:  map[string]feud.County{},
	}

	f.阿特巴拉Atbara = atbara.CAtbara阿特巴拉
	f.阿特巴拉Atbara.SetParent(f)

	f.马库里亚Makuria = makuria.CMakuria马库里亚
	f.马库里亚Makuria.SetParent(f)

	f.纳帕塔Napata = napata.CNapata纳帕塔
	f.纳帕塔Napata.SetParent(f)

}
