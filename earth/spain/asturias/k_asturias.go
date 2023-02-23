package asturias

import (
	"github.com/thalesfu/CK2Commands/earth/spain/asturias/asturias"
	"github.com/thalesfu/CK2Commands/earth/spain/asturias/leon"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsturiasKingdom interface {
	feud.Kingdom
	DAsturias阿斯图里亚斯() asturias.AsturiasDuke
	DLeon莱昂() leon.LeonDuke
}

type 阿斯图里亚斯AsturiasKingdom struct {
	feud.BaseKingdom
	阿斯图里亚斯Asturias asturias.AsturiasDuke
	莱昂Leon         leon.LeonDuke
}

func (k *阿斯图里亚斯AsturiasKingdom) DAsturias阿斯图里亚斯() asturias.AsturiasDuke {
	return k.阿斯图里亚斯Asturias
}

func (k *阿斯图里亚斯AsturiasKingdom) DLeon莱昂() leon.LeonDuke {
	return k.莱昂Leon
}

var KAsturias阿斯图里亚斯 AsturiasKingdom = &阿斯图里亚斯AsturiasKingdom{}

func init() {
	f := KAsturias阿斯图里亚斯.(*阿斯图里亚斯AsturiasKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "asturias",
		TitleName: "阿斯图里亚斯",
		TitleCode: "k_asturias",
		Dukes:     map[string]feud.Duke{},
	}

	f.阿斯图里亚斯Asturias = asturias.DAsturias阿斯图里亚斯
	f.阿斯图里亚斯Asturias.SetParent(f)

	f.莱昂Leon = leon.DLeon莱昂
	f.莱昂Leon.SetParent(f)

}
