package picardie

import (
	"github.com/thalesfu/CK2Commands/earth/france/france/picardie/amiens"
	"github.com/thalesfu/CK2Commands/earth/france/france/picardie/clermont"
	"github.com/thalesfu/CK2Commands/earth/france/france/picardie/ponthieu"
	"github.com/thalesfu/CK2Commands/earth/france/france/picardie/vermandois"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PicardieDuke interface {
	feud.Duke
	CAmiens亚眠() amiens.AmiensCounty
	CClermont克莱蒙() clermont.ClermontCounty
	CPonthieu蓬蒂耶() ponthieu.PonthieuCounty
	CVermandois韦尔芒杜瓦() vermandois.VermandoisCounty
}

type 皮卡第PicardieDuke struct {
	feud.BaseDuke
	亚眠Amiens        amiens.AmiensCounty
	克莱蒙Clermont     clermont.ClermontCounty
	蓬蒂耶Ponthieu     ponthieu.PonthieuCounty
	韦尔芒杜瓦Vermandois vermandois.VermandoisCounty
}

func (d *皮卡第PicardieDuke) CAmiens亚眠() amiens.AmiensCounty {
	return d.亚眠Amiens
}

func (d *皮卡第PicardieDuke) CClermont克莱蒙() clermont.ClermontCounty {
	return d.克莱蒙Clermont
}

func (d *皮卡第PicardieDuke) CPonthieu蓬蒂耶() ponthieu.PonthieuCounty {
	return d.蓬蒂耶Ponthieu
}

func (d *皮卡第PicardieDuke) CVermandois韦尔芒杜瓦() vermandois.VermandoisCounty {
	return d.韦尔芒杜瓦Vermandois
}

var DPicardie皮卡第 PicardieDuke = &皮卡第PicardieDuke{}

func init() {
	f := DPicardie皮卡第.(*皮卡第PicardieDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "picardie",
		TitleName: "皮卡第",
		TitleCode: "d_picardie",
		Counties:  map[string]feud.County{},
	}

	f.亚眠Amiens = amiens.CAmiens亚眠
	f.亚眠Amiens.SetParent(f)

	f.克莱蒙Clermont = clermont.CClermont克莱蒙
	f.克莱蒙Clermont.SetParent(f)

	f.蓬蒂耶Ponthieu = ponthieu.CPonthieu蓬蒂耶
	f.蓬蒂耶Ponthieu.SetParent(f)

	f.韦尔芒杜瓦Vermandois = vermandois.CVermandois韦尔芒杜瓦
	f.韦尔芒杜瓦Vermandois.SetParent(f)

}
