package om

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/om/narim"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/om/om"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/om/tara"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OmDuke interface {
	feud.Duke
	CNarim纳林() narim.NarimCounty
	COm鄂木() om.OmCounty
	CTara塔拉() tara.TaraCounty
}

type 鄂木OmDuke struct {
	feud.BaseDuke
	纳林Narim narim.NarimCounty
	鄂木Om    om.OmCounty
	塔拉Tara  tara.TaraCounty
}

func (d *鄂木OmDuke) CNarim纳林() narim.NarimCounty {
	return d.纳林Narim
}

func (d *鄂木OmDuke) COm鄂木() om.OmCounty {
	return d.鄂木Om
}

func (d *鄂木OmDuke) CTara塔拉() tara.TaraCounty {
	return d.塔拉Tara
}

var DOm鄂木 OmDuke = &鄂木OmDuke{}

func init() {
	f := DOm鄂木.(*鄂木OmDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "om",
		TitleName: "鄂木",
		TitleCode: "d_om",
		Counties:  map[string]feud.County{},
	}

	f.纳林Narim = narim.CNarim纳林
	f.纳林Narim.SetParent(f)

	f.鄂木Om = om.COm鄂木
	f.鄂木Om.SetParent(f)

	f.塔拉Tara = tara.CTara塔拉
	f.塔拉Tara.SetParent(f)

}
