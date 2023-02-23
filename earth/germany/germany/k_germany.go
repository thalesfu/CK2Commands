package germany

import (
	"github.com/thalesfu/CK2Commands/earth/germany/germany/alsace"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/baden"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/franconia"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/hesse"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/mainz"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/raetia"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/rhine"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/swabia"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/thurgau"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/thuringia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GermanyKingdom interface {
	feud.Kingdom
	DAlsace阿尔萨斯() alsace.AlsaceDuke
	DBaden巴登() baden.BadenDuke
	DFranconia法兰克尼亚() franconia.FranconiaDuke
	DHesse黑森() hesse.HesseDuke
	DMainz美因茨() mainz.MainzDuke
	DRaetia雷蒂亚() raetia.RaetiaDuke
	DRhine莱茵兰() rhine.RhineDuke
	DSwabia施瓦本() swabia.SwabiaDuke
	DThurgau图尔高() thurgau.ThurgauDuke
	DThuringia图林根() thuringia.ThuringiaDuke
}

type 德意志GermanyKingdom struct {
	feud.BaseKingdom
	阿尔萨斯Alsace     alsace.AlsaceDuke
	巴登Baden        baden.BadenDuke
	法兰克尼亚Franconia franconia.FranconiaDuke
	黑森Hesse        hesse.HesseDuke
	美因茨Mainz       mainz.MainzDuke
	雷蒂亚Raetia      raetia.RaetiaDuke
	莱茵兰Rhine       rhine.RhineDuke
	施瓦本Swabia      swabia.SwabiaDuke
	图尔高Thurgau     thurgau.ThurgauDuke
	图林根Thuringia   thuringia.ThuringiaDuke
}

func (k *德意志GermanyKingdom) DAlsace阿尔萨斯() alsace.AlsaceDuke {
	return k.阿尔萨斯Alsace
}

func (k *德意志GermanyKingdom) DBaden巴登() baden.BadenDuke {
	return k.巴登Baden
}

func (k *德意志GermanyKingdom) DFranconia法兰克尼亚() franconia.FranconiaDuke {
	return k.法兰克尼亚Franconia
}

func (k *德意志GermanyKingdom) DHesse黑森() hesse.HesseDuke {
	return k.黑森Hesse
}

func (k *德意志GermanyKingdom) DMainz美因茨() mainz.MainzDuke {
	return k.美因茨Mainz
}

func (k *德意志GermanyKingdom) DRaetia雷蒂亚() raetia.RaetiaDuke {
	return k.雷蒂亚Raetia
}

func (k *德意志GermanyKingdom) DRhine莱茵兰() rhine.RhineDuke {
	return k.莱茵兰Rhine
}

func (k *德意志GermanyKingdom) DSwabia施瓦本() swabia.SwabiaDuke {
	return k.施瓦本Swabia
}

func (k *德意志GermanyKingdom) DThurgau图尔高() thurgau.ThurgauDuke {
	return k.图尔高Thurgau
}

func (k *德意志GermanyKingdom) DThuringia图林根() thuringia.ThuringiaDuke {
	return k.图林根Thuringia
}

var KGermany德意志 GermanyKingdom = &德意志GermanyKingdom{}

func init() {
	f := KGermany德意志.(*德意志GermanyKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "germany",
		TitleName: "德意志",
		TitleCode: "k_germany",
		Dukes:     map[string]feud.Duke{},
	}

	f.阿尔萨斯Alsace = alsace.DAlsace阿尔萨斯
	f.阿尔萨斯Alsace.SetParent(f)

	f.巴登Baden = baden.DBaden巴登
	f.巴登Baden.SetParent(f)

	f.法兰克尼亚Franconia = franconia.DFranconia法兰克尼亚
	f.法兰克尼亚Franconia.SetParent(f)

	f.黑森Hesse = hesse.DHesse黑森
	f.黑森Hesse.SetParent(f)

	f.美因茨Mainz = mainz.DMainz美因茨
	f.美因茨Mainz.SetParent(f)

	f.雷蒂亚Raetia = raetia.DRaetia雷蒂亚
	f.雷蒂亚Raetia.SetParent(f)

	f.莱茵兰Rhine = rhine.DRhine莱茵兰
	f.莱茵兰Rhine.SetParent(f)

	f.施瓦本Swabia = swabia.DSwabia施瓦本
	f.施瓦本Swabia.SetParent(f)

	f.图尔高Thurgau = thurgau.DThurgau图尔高
	f.图尔高Thurgau.SetParent(f)

	f.图林根Thuringia = thuringia.DThuringia图林根
	f.图林根Thuringia.SetParent(f)

}
