package normandy

import (
	"github.com/thalesfu/CK2Commands/earth/france/france/normandy/arques"
	"github.com/thalesfu/CK2Commands/earth/france/france/normandy/avranches"
	"github.com/thalesfu/CK2Commands/earth/france/france/normandy/caen"
	"github.com/thalesfu/CK2Commands/earth/france/france/normandy/eu"
	"github.com/thalesfu/CK2Commands/earth/france/france/normandy/evreux"
	"github.com/thalesfu/CK2Commands/earth/france/france/normandy/vexin"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NormandyDuke interface {
	feud.Duke
	CArques鲁昂() arques.ArquesCounty
	CAvranches莫尔坦() avranches.AvranchesCounty
	CCaen卡昂() caen.CaenCounty
	CEu厄镇() eu.EuCounty
	CEvreux埃夫勒() evreux.EvreuxCounty
	CVexin韦克桑() vexin.VexinCounty
}

type 诺曼底NormandyDuke struct {
	feud.BaseDuke
	鲁昂Arques     arques.ArquesCounty
	莫尔坦Avranches avranches.AvranchesCounty
	卡昂Caen       caen.CaenCounty
	厄镇Eu         eu.EuCounty
	埃夫勒Evreux    evreux.EvreuxCounty
	韦克桑Vexin     vexin.VexinCounty
}

func (d *诺曼底NormandyDuke) CArques鲁昂() arques.ArquesCounty {
	return d.鲁昂Arques
}

func (d *诺曼底NormandyDuke) CAvranches莫尔坦() avranches.AvranchesCounty {
	return d.莫尔坦Avranches
}

func (d *诺曼底NormandyDuke) CCaen卡昂() caen.CaenCounty {
	return d.卡昂Caen
}

func (d *诺曼底NormandyDuke) CEu厄镇() eu.EuCounty {
	return d.厄镇Eu
}

func (d *诺曼底NormandyDuke) CEvreux埃夫勒() evreux.EvreuxCounty {
	return d.埃夫勒Evreux
}

func (d *诺曼底NormandyDuke) CVexin韦克桑() vexin.VexinCounty {
	return d.韦克桑Vexin
}

var DNormandy诺曼底 NormandyDuke = &诺曼底NormandyDuke{}

func init() {
	f := DNormandy诺曼底.(*诺曼底NormandyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "normandy",
		TitleName: "诺曼底",
		TitleCode: "d_normandy",
		Counties:  map[string]feud.County{},
	}

	f.鲁昂Arques = arques.CArques鲁昂
	f.鲁昂Arques.SetParent(f)

	f.莫尔坦Avranches = avranches.CAvranches莫尔坦
	f.莫尔坦Avranches.SetParent(f)

	f.卡昂Caen = caen.CCaen卡昂
	f.卡昂Caen.SetParent(f)

	f.厄镇Eu = eu.CEu厄镇
	f.厄镇Eu.SetParent(f)

	f.埃夫勒Evreux = evreux.CEvreux埃夫勒
	f.埃夫勒Evreux.SetParent(f)

	f.韦克桑Vexin = vexin.CVexin韦克桑
	f.韦克桑Vexin.SetParent(f)

}
