package tangiers

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tangiers/cebta"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tangiers/tangiers"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TangiersDuke interface {
	feud.Duke
	CCebta休达() cebta.CebtaCounty
	CTangiers丹吉尔() tangiers.TangiersCounty
}

type 丹吉尔TangiersDuke struct {
	feud.BaseDuke
	休达Cebta     cebta.CebtaCounty
	丹吉尔Tangiers tangiers.TangiersCounty
}

func (d *丹吉尔TangiersDuke) CCebta休达() cebta.CebtaCounty {
	return d.休达Cebta
}

func (d *丹吉尔TangiersDuke) CTangiers丹吉尔() tangiers.TangiersCounty {
	return d.丹吉尔Tangiers
}

var DTangiers丹吉尔 TangiersDuke = &丹吉尔TangiersDuke{}

func init() {
	f := DTangiers丹吉尔.(*丹吉尔TangiersDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tangiers",
		TitleName: "丹吉尔",
		TitleCode: "d_tangiers",
		Counties:  map[string]feud.County{},
	}

	f.休达Cebta = cebta.CCebta休达
	f.休达Cebta.SetParent(f)

	f.丹吉尔Tangiers = tangiers.CTangiers丹吉尔
	f.丹吉尔Tangiers.SetParent(f)

}
