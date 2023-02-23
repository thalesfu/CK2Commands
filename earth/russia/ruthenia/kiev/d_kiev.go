package kiev

import (
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/kiev/kiev"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/kiev/korosten"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/kiev/medjybij"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/kiev/vozviahel"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KievDuke interface {
	feud.Duke
	CKiev基辅() kiev.KievCounty
	CKorosten科罗斯坚() korosten.KorostenCounty
	CMedjybij梅德日比日() medjybij.MedjybijCounty
	CVozviahel沃兹维亚格尔() vozviahel.VozviahelCounty
}

type 基辅KievDuke struct {
	feud.BaseDuke
	基辅Kiev          kiev.KievCounty
	科罗斯坚Korosten    korosten.KorostenCounty
	梅德日比日Medjybij   medjybij.MedjybijCounty
	沃兹维亚格尔Vozviahel vozviahel.VozviahelCounty
}

func (d *基辅KievDuke) CKiev基辅() kiev.KievCounty {
	return d.基辅Kiev
}

func (d *基辅KievDuke) CKorosten科罗斯坚() korosten.KorostenCounty {
	return d.科罗斯坚Korosten
}

func (d *基辅KievDuke) CMedjybij梅德日比日() medjybij.MedjybijCounty {
	return d.梅德日比日Medjybij
}

func (d *基辅KievDuke) CVozviahel沃兹维亚格尔() vozviahel.VozviahelCounty {
	return d.沃兹维亚格尔Vozviahel
}

var DKiev基辅 KievDuke = &基辅KievDuke{}

func init() {
	f := DKiev基辅.(*基辅KievDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kiev",
		TitleName: "基辅",
		TitleCode: "d_kiev",
		Counties:  map[string]feud.County{},
	}

	f.基辅Kiev = kiev.CKiev基辅
	f.基辅Kiev.SetParent(f)

	f.科罗斯坚Korosten = korosten.CKorosten科罗斯坚
	f.科罗斯坚Korosten.SetParent(f)

	f.梅德日比日Medjybij = medjybij.CMedjybij梅德日比日
	f.梅德日比日Medjybij.SetParent(f)

	f.沃兹维亚格尔Vozviahel = vozviahel.CVozviahel沃兹维亚格尔
	f.沃兹维亚格尔Vozviahel.SetParent(f)

}
