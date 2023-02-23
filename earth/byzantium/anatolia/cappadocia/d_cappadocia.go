package cappadocia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/cappadocia/nyssa"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/cappadocia/tyana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CappadociaDuke interface {
	feud.Duke
	CNyssa尼撒() nyssa.NyssaCounty
	CTyana堤亚纳() tyana.TyanaCounty
}

type 卡帕多西亚CappadociaDuke struct {
	feud.BaseDuke
	尼撒Nyssa  nyssa.NyssaCounty
	堤亚纳Tyana tyana.TyanaCounty
}

func (d *卡帕多西亚CappadociaDuke) CNyssa尼撒() nyssa.NyssaCounty {
	return d.尼撒Nyssa
}

func (d *卡帕多西亚CappadociaDuke) CTyana堤亚纳() tyana.TyanaCounty {
	return d.堤亚纳Tyana
}

var DCappadocia卡帕多西亚 CappadociaDuke = &卡帕多西亚CappadociaDuke{}

func init() {
	f := DCappadocia卡帕多西亚.(*卡帕多西亚CappadociaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cappadocia",
		TitleName: "卡帕多西亚",
		TitleCode: "d_cappadocia",
		Counties:  map[string]feud.County{},
	}

	f.尼撒Nyssa = nyssa.CNyssa尼撒
	f.尼撒Nyssa.SetParent(f)

	f.堤亚纳Tyana = tyana.CTyana堤亚纳
	f.堤亚纳Tyana.SetParent(f)

}
