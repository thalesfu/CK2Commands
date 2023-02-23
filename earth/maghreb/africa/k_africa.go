package africa

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/cyrenaica"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/jerid"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/kabylia"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/mzab"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/syrte"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tripolitania"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tunis"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AfricaKingdom interface {
	feud.Kingdom
	DCyrenaica昔兰尼加() cyrenaica.CyrenaicaDuke
	DJerid杰里德() jerid.JeridDuke
	DKabylia卡比利亚() kabylia.KabyliaDuke
	DMzab姆扎卜() mzab.MzabDuke
	DSyrte苏尔特() syrte.SyrteDuke
	DTripolitania的黎波里塔尼亚() tripolitania.TripolitaniaDuke
	DTunis突尼斯() tunis.TunisDuke
}

type 阿非利加AfricaKingdom struct {
	feud.BaseKingdom
	昔兰尼加Cyrenaica       cyrenaica.CyrenaicaDuke
	杰里德Jerid            jerid.JeridDuke
	卡比利亚Kabylia         kabylia.KabyliaDuke
	姆扎卜Mzab             mzab.MzabDuke
	苏尔特Syrte            syrte.SyrteDuke
	的黎波里塔尼亚Tripolitania tripolitania.TripolitaniaDuke
	突尼斯Tunis            tunis.TunisDuke
}

func (k *阿非利加AfricaKingdom) DCyrenaica昔兰尼加() cyrenaica.CyrenaicaDuke {
	return k.昔兰尼加Cyrenaica
}

func (k *阿非利加AfricaKingdom) DJerid杰里德() jerid.JeridDuke {
	return k.杰里德Jerid
}

func (k *阿非利加AfricaKingdom) DKabylia卡比利亚() kabylia.KabyliaDuke {
	return k.卡比利亚Kabylia
}

func (k *阿非利加AfricaKingdom) DMzab姆扎卜() mzab.MzabDuke {
	return k.姆扎卜Mzab
}

func (k *阿非利加AfricaKingdom) DSyrte苏尔特() syrte.SyrteDuke {
	return k.苏尔特Syrte
}

func (k *阿非利加AfricaKingdom) DTripolitania的黎波里塔尼亚() tripolitania.TripolitaniaDuke {
	return k.的黎波里塔尼亚Tripolitania
}

func (k *阿非利加AfricaKingdom) DTunis突尼斯() tunis.TunisDuke {
	return k.突尼斯Tunis
}

var KAfrica阿非利加 AfricaKingdom = &阿非利加AfricaKingdom{}

func init() {
	f := KAfrica阿非利加.(*阿非利加AfricaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "africa",
		TitleName: "阿非利加",
		TitleCode: "k_africa",
		Dukes:     map[string]feud.Duke{},
	}

	f.昔兰尼加Cyrenaica = cyrenaica.DCyrenaica昔兰尼加
	f.昔兰尼加Cyrenaica.SetParent(f)

	f.杰里德Jerid = jerid.DJerid杰里德
	f.杰里德Jerid.SetParent(f)

	f.卡比利亚Kabylia = kabylia.DKabylia卡比利亚
	f.卡比利亚Kabylia.SetParent(f)

	f.姆扎卜Mzab = mzab.DMzab姆扎卜
	f.姆扎卜Mzab.SetParent(f)

	f.苏尔特Syrte = syrte.DSyrte苏尔特
	f.苏尔特Syrte.SetParent(f)

	f.的黎波里塔尼亚Tripolitania = tripolitania.DTripolitania的黎波里塔尼亚
	f.的黎波里塔尼亚Tripolitania.SetParent(f)

	f.突尼斯Tunis = tunis.DTunis突尼斯
	f.突尼斯Tunis.SetParent(f)

}
