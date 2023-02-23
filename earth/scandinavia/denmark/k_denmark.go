package denmark

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/holstein"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/sjaelland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/skane"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark/slesvig"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DenmarkKingdom interface {
	feud.Kingdom
	DHolstein荷尔斯泰因() holstein.HolsteinDuke
	DSjaelland西兰() sjaelland.SjaellandDuke
	DSkane斯堪尼亚() skane.SkaneDuke
	DSlesvig日德兰() slesvig.SlesvigDuke
}

type 丹麦DenmarkKingdom struct {
	feud.BaseKingdom
	荷尔斯泰因Holstein holstein.HolsteinDuke
	西兰Sjaelland   sjaelland.SjaellandDuke
	斯堪尼亚Skane     skane.SkaneDuke
	日德兰Slesvig    slesvig.SlesvigDuke
}

func (k *丹麦DenmarkKingdom) DHolstein荷尔斯泰因() holstein.HolsteinDuke {
	return k.荷尔斯泰因Holstein
}

func (k *丹麦DenmarkKingdom) DSjaelland西兰() sjaelland.SjaellandDuke {
	return k.西兰Sjaelland
}

func (k *丹麦DenmarkKingdom) DSkane斯堪尼亚() skane.SkaneDuke {
	return k.斯堪尼亚Skane
}

func (k *丹麦DenmarkKingdom) DSlesvig日德兰() slesvig.SlesvigDuke {
	return k.日德兰Slesvig
}

var KDenmark丹麦 DenmarkKingdom = &丹麦DenmarkKingdom{}

func init() {
	f := KDenmark丹麦.(*丹麦DenmarkKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "denmark",
		TitleName: "丹麦",
		TitleCode: "k_denmark",
		Dukes:     map[string]feud.Duke{},
	}

	f.荷尔斯泰因Holstein = holstein.DHolstein荷尔斯泰因
	f.荷尔斯泰因Holstein.SetParent(f)

	f.西兰Sjaelland = sjaelland.DSjaelland西兰
	f.西兰Sjaelland.SetParent(f)

	f.斯堪尼亚Skane = skane.DSkane斯堪尼亚
	f.斯堪尼亚Skane.SetParent(f)

	f.日德兰Slesvig = slesvig.DSlesvig日德兰
	f.日德兰Slesvig.SetParent(f)

}
