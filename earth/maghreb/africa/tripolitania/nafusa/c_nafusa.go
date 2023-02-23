package nafusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NafusaCounty interface {
	feud.County
	BErruda埃鲁达() feud.Barony
	BKroussia克鲁西亚() feud.Barony
	BNafusa奈富塞() feud.Barony
	BSejenane塞杰南() feud.Barony
}

type 奈富塞NafusaCounty struct {
	feud.BaseCounty
	埃鲁达Erruda    feud.Barony
	克鲁西亚Kroussia feud.Barony
	奈富塞Nafusa    feud.Barony
	塞杰南Sejenane  feud.Barony
}

func (c *奈富塞NafusaCounty) BErruda埃鲁达() feud.Barony {
	return c.埃鲁达Erruda
}

func (c *奈富塞NafusaCounty) BKroussia克鲁西亚() feud.Barony {
	return c.克鲁西亚Kroussia
}

func (c *奈富塞NafusaCounty) BNafusa奈富塞() feud.Barony {
	return c.奈富塞Nafusa
}

func (c *奈富塞NafusaCounty) BSejenane塞杰南() feud.Barony {
	return c.塞杰南Sejenane
}

var CNafusa奈富塞 NafusaCounty = &奈富塞NafusaCounty{}

func init() {
	f := CNafusa奈富塞.(*奈富塞NafusaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1723",
		Title:     "nafusa",
		TitleName: "奈富塞",
		TitleCode: "c_nafusa",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃鲁达Erruda = BErruda埃鲁达
	f.埃鲁达Erruda.SetParent(f)

	f.克鲁西亚Kroussia = BKroussia克鲁西亚
	f.克鲁西亚Kroussia.SetParent(f)

	f.奈富塞Nafusa = BNafusa奈富塞
	f.奈富塞Nafusa.SetParent(f)

	f.塞杰南Sejenane = BSejenane塞杰南
	f.塞杰南Sejenane.SetParent(f)

}
